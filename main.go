package main

import (
	"bytes"
	"fmt"
	"log"
	"net/url"
	"os"
	"path"
	"regexp"
	"strings"
	"text/template"

	"github.com/go-git/go-git/v5"
	"github.com/go-git/go-git/v5/plumbing/object"
	"github.com/yuin/goldmark"
	meta "github.com/yuin/goldmark-meta"
	"github.com/yuin/goldmark/ast"
	"github.com/yuin/goldmark/extension"
	"github.com/yuin/goldmark/parser"
	"github.com/yuin/goldmark/text"
	"github.com/yuin/goldmark/util"
)

type Site struct {
	AllPages []*Page
	Sections map[string][]*Page
}

func NewSite() *Site {
	return &Site{
		AllPages: []*Page{},
		Sections: map[string][]*Page{},
	}
}

// Page represents a goliki file.
// It contains the file data and metadata at multiple stages in the pipeline.
// This is called Page
type Page struct {
	// File written in markdown
	File     *object.File
	Metadata map[string]interface{}
	// HTML rendered from File
	MarkdownHTML *bytes.Buffer
	LastCommit   *object.Commit

	Section string
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

	site := NewSite()

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

	// file name check for *.md
	pageRegex, err := regexp.Compile(`[\w]+\.md`)
	if err != nil {
		log.Fatal(err)
	}

	// layoutRegex, err := regexp.Compile(`.goliki/layouts/[\w]+\.html`)
	// if err != nil {
	// 	log.Fatal(err)
	// }

	// traverse the tree!
	// add qualifying files to list
	tree.Files().ForEach(func(f *object.File) error {

		if !pageRegex.MatchString(f.Name) {
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

		doc := &Page{
			File:       f,
			LastCommit: lastCommit,
		}

		// fmt.Println(doc.File.Name)
		// fmt.Println(doc.LastCommit.Message)

		site.AllPages = append(site.AllPages, doc)
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
	for _, page := range site.AllPages {

		filepath := page.File.Name
		// dir, filename := path.Split(filepath)
		dirpaths := strings.SplitN(filepath, "/", 2)
		if len(dirpaths) > 1 {
			page.Section = dirpaths[0]
		}

		// get file content
		contents, err := page.File.Contents()
		if err != nil {
			log.Fatalf("could not get contents: %s", filepath)
			continue
		}

		// Why do I have to do this?
		page.MarkdownHTML = new(bytes.Buffer)

		// TODO: Add https://github.com/yuin/goldmark-highlighting
		context := parser.NewContext()
		if err := markdown.Convert([]byte(contents), page.MarkdownHTML, parser.WithContext(context)); err != nil {
			panic(err)
		}

		// TODO: Metadata is too raw
		page.Metadata = meta.Get(context)
		// fmt.Println(d.Metadata)

		// Collections
		site.Sections[page.Section] = append(site.Sections[page.Section], page)
	}

	t := template.Must(template.ParseFiles(
		".goliki/layouts/default.html",
		".goliki/layouts/index.html",
	))
	fmt.Printf("tmpl check!\n%s", t.DefinedTemplates())

	// for each doc we want to render the specified layout with metadata and generated markdown,
	// and write the final content to file.
	for _, page := range site.AllPages {
		// TODO: get YAML front matter and read layout (template)

		//
		// write HTML file to file path
		//

		// e.g. f/a.md > build/f/a.html
		relpath := strings.Replace(page.File.Name, ".md", ".html", 1)
		tmplName := "default.html"

		// special case: check if README and rewrite to index
		if strings.Contains(page.File.Name, "README.md") {
			relpath = strings.Replace(page.File.Name, "README.md", "index.html", 1)
			tmplName = "index.html"
		}

		// FIXME: default var or parameter
		buildpath := path.Join(".goliki/public/", relpath)

		fmt.Println(buildpath)

		// mkdir -p
		os.MkdirAll(path.Dir(buildpath), os.ModePerm)

		f, err := os.Create(buildpath)
		if err != nil {
			panic(err)
		}
		defer f.Close()

		// render template with doc as context and write to file
		err = t.ExecuteTemplate(f, tmplName, page)
		if err != nil {
			panic(err)
		}

		// f.Write(d.MarkdownHTML.Bytes())
	}

}
