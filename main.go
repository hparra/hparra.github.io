package main

import (
	"bytes"
	"fmt"
	"log"
	"net/url"
	"os"
	"path"
	"strings"

	"github.com/go-git/go-git/v5"
	"github.com/go-git/go-git/v5/plumbing/object"
	"github.com/yuin/goldmark"
	"github.com/yuin/goldmark-meta"
	"github.com/yuin/goldmark/ast"
	"github.com/yuin/goldmark/extension"
	"github.com/yuin/goldmark/parser"
	"github.com/yuin/goldmark/text"
	"github.com/yuin/goldmark/util"
)

// Document represents a goliki file.
// It contains the file data and metadata at multiple stages in the pipeline.
type Document struct {
	// File written in markdown
	File *object.File
	// HTML rendered from File
	MarkdownHTML *bytes.Buffer
	LastCommit   *object.Commit
}

// LinkTransformer implements ASTTransformer.
// https://pkg.go.dev/github.com/yuin/goldmark@v1.5.2/parser?utm_source=gopls#ASTTransformer
type LinkTransformer struct{}

// Transform transforms the provided Markdown AST.
func (*LinkTransformer) Transform(doc *ast.Document, reader text.Reader, pctx parser.Context) {

	ast.Walk(doc, func(node ast.Node, enter bool) (ast.WalkStatus, error) {
		if !enter {
			return ast.WalkContinue, nil
		}

		link, ok := node.(*ast.Link)
		if !ok {
			return ast.WalkContinue, nil
		}

		dest := string(link.Destination)

		u, err := url.Parse(dest)
		if err != nil {
			log.Fatal("bad url")
		}

		// We only change relative-URLs
		if u.IsAbs() {
			return ast.WalkContinue, nil
		}

		// replace html with md
		newDest := strings.Replace(dest, ".md", ".html", 1)
		link.Destination = []byte(newDest)
		// fmt.Printf("%s\n", newDest)

		return ast.WalkContinue, nil
	})
}

func main() {

	gitpath := "."

	r, err := git.PlainOpen(gitpath)
	if err != nil {
		log.Fatal(err)
	}

	// get reference to last commit (HEAD)
	ref, err := r.Head()
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(ref.Hash())

	// retrieve commit object using hasg
	commit, err := r.CommitObject(ref.Hash())
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(commit)

	// get list of all files at that commit
	tree, err := commit.Tree()
	if err != nil {
		log.Fatal(err)
	}

	// traverse the tree and add qualifying files to list
	docs := []*Document{}
	tree.Files().ForEach(func(f *object.File) error {

		// NOTE: first param is a shell pattern, not regex
		matched, err := path.Match("*/*.md", f.Name)
		if err != nil || !matched {
			return nil
		}

		// git log -- FILENAME
		commitIter, err := r.Log(&git.LogOptions{
			From:     ref.Hash(),
			FileName: &f.Name,
		})
		if err != nil {
			return nil
		}

		lastCommit, err := commitIter.Next()
		if err != nil {
			return nil
		}

		doc := &Document{
			File:       f,
			LastCommit: lastCommit,
		}

		// fmt.Println(doc.File.Name)
		// fmt.Println(doc.LastCommit.Message)

		docs = append(docs, doc)
		return nil
	})

	// create configured goldmark
	markdown := goldmark.New(
		goldmark.WithParserOptions(
			parser.WithASTTransformers(
				util.Prioritized(&LinkTransformer{}, 100),
			),
		),
		goldmark.WithExtensions(
			extension.GFM,
			meta.Meta,
		),
	)

	// for each doc we need to calculate the resulting path

	// for each doc parse it's markdown and store rendered HTML
	// TODO: store YAML front matter for each
	for _, d := range docs {

		// get file content
		contents, err := d.File.Contents()
		if err != nil {
			log.Fatalf("could not get contents: %s", d.File.Name)
			continue
		}

		// Why do I have to do this?
		d.MarkdownHTML = new(bytes.Buffer)

		// TODO: Custom Link renderer
		// TODO: Add https://github.com/yuin/goldmark-highlighting
		context := parser.NewContext()
		if err := markdown.Convert([]byte(contents), d.MarkdownHTML, parser.WithContext(context)); err != nil {
			panic(err)
		}
	}

	// TODO: process site-wide metadata from all front matter

	// for each doc we want to render the specified layout with metadata and generated markdown,
	// and write the final content to file.
	for _, d := range docs {
		// TODO: get YAML front matter and read layout (template)
		// TODO: render layout with metadata and markdown

		//
		// write HTML file to file path
		//

		// e.g. f/a.md > build/f/a.html
		relpath := strings.Replace(d.File.Name, ".md", ".html", 1)
		buildpath := path.Join("build", relpath)
		fmt.Println(buildpath)

		// mkdir -p
		os.MkdirAll(path.Dir(buildpath), os.ModePerm)

		f, err := os.Create(buildpath)
		if err != nil {
			panic(err)
		}
		defer f.Close()

		f.Write(d.MarkdownHTML.Bytes())
	}

}
