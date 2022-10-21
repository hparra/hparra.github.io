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

// Site is collection of all site data: pages, variables, etc.
// It is passed into templates via Page.
type Site struct {
	AllPages []*Page
	Sections map[string][]*Page
}

// NewSite returns a Site with instantiated fields.
func NewSite() *Site {
	return &Site{
		AllPages: []*Page{},
		Sections: map[string][]*Page{},
	}
}

// Page represents a goliki file.
// It contains the file data and metadata at multiple stages in the build pipeline.
// It is passed into templates.
type Page struct {
	// GitFile is the original content file from git.
	GitFile *object.File
	// LastCommit is the latest commit that modified this file.
	LastCommit *object.Commit
	// MarkdownHTML is rendered from GitFile content.
	MarkdownHTML *bytes.Buffer
	// Metadata is read from YAML Front Matter.
	Metadata map[string]interface{}

	Site    *Site
	File    *File
	OutFile *File

	TemplateName string

	RelPermalink string
	Section      string
	Title        string
}

// File represents common file metadata.
// It tries to follow schema for Hugo's .Page.File
// https://gohugo.io/variables/files/
type File struct {
	Path            string
	LogicalName     string
	ContentBaseName string
	BaseFileName    string
	Ext             string
	Dir             string
}

// LinkTransformer is a custom Goldmark ASTTransformer.
// It changes links to properly point to generated files.
// https://pkg.go.dev/github.com/yuin/goldmark@v1.5.2/parser?utm_source=gopls#ASTTransformer
type LinkTransformer struct{}

// Transform changes markdown links referencing relative markdown files to similarly-named html files.
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

//
func main() {

	site := NewSite()

	// save time and check templates immediately
	t := template.Must(template.ParseFiles(
		".goliki/layouts/default.html",
		".goliki/layouts/index.html",
	))
	log.Printf("found templates: %s\n", t.DefinedTemplates())

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
	log.Printf("HEAD %s\n", ref.Hash())

	// retrieve commit object using hash
	commit, err := r.CommitObject(ref.Hash())
	if err != nil {
		log.Fatal(err)
	}

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

	// traverse the tree!
	// add qualifying files to list
	// TODO: don't be lazy; read from file system
	tree.Files().ForEach(func(f *object.File) error {

		if !pageRegex.MatchString(f.Name) {
			return nil
		}
		log.Printf("processing %s\n", f.Name)

		// git log -- FILENAME
		commitIter, err := r.Log(&git.LogOptions{
			From:     ref.Hash(),
			FileName: &f.Name,
		})
		if err != nil {
			log.Fatal(err)
		}

		lastCommit, err := commitIter.Next()
		if err != nil {
			log.Fatal(err)
		}

		page := &Page{
			GitFile:    f,
			LastCommit: lastCommit,
		}

		site.AllPages = append(site.AllPages, page)

		return nil
	})
	log.Printf("processed %d files\n", len(site.AllPages))

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

	// for each doc parse it's markdown and store rendered HTML
	log.Println("starting process")
	for _, page := range site.AllPages {

		// get file content
		contents, err := page.GitFile.Contents()
		if err != nil {
			log.Printf("could not get contents: %s\n", page.GitFile.Name)
			continue
		}

		// Why do I have to do this?
		page.MarkdownHTML = new(bytes.Buffer)

		// TODO: Add https://github.com/yuin/goldmark-highlighting
		context := parser.NewContext()
		if err := markdown.Convert([]byte(contents), page.MarkdownHTML, parser.WithContext(context)); err != nil {
			log.Printf("%s", err)
			continue
		}

		// TODO: Metadata is too raw
		page.Metadata = meta.Get(context)
		// fmt.Println(d.Metadata)

		// cf is our content file, i.e. original markdown files
		cf := &File{}
		cf.Path = page.GitFile.Name
		cf.Dir, cf.LogicalName = path.Split(cf.Path)
		cf.Ext = path.Ext(cf.LogicalName)
		cf.ContentBaseName = strings.Replace(cf.LogicalName, cf.Ext, "", 1)
		page.File = cf

		// will this error?
		page.Section = strings.Split(cf.Dir, "/")[0]

		// sf is our site file, i.e. final rendered files
		sf := &File{}
		sf.Ext = "html"
		sf.ContentBaseName = cf.ContentBaseName
		if cf.ContentBaseName == "README" {
			sf.ContentBaseName = "index"
		}
		sf.LogicalName = fmt.Sprintf("%s.%s", sf.ContentBaseName, sf.Ext)
		sf.Dir = cf.Dir
		sf.Path = path.Join(sf.Dir, sf.LogicalName)
		page.OutFile = sf

		// select our template
		page.TemplateName = "default.html"
		if sf.ContentBaseName == "index" {
			page.TemplateName = "index.html"
		}

		page.RelPermalink = sf.Path
		page.Title = sf.ContentBaseName
	}
	log.Println("finished process")

	// Post-processing for collections, etc.
	log.Println("starting post-process")
	for _, page := range site.AllPages {

		if page.Section != "" {
			site.Sections[page.Section] = append(site.Sections[page.Section], page)
		}

		// embed global Site var into each page
		page.Site = site
	}
	log.Println("finished post-process")

	publishDir := ".goliki/public/"

	// traverse all pages and render to final location
	log.Println("starting file writing")
	for _, page := range site.AllPages {

		// create destination directory
		filedir := path.Join(publishDir, page.OutFile.Dir)
		err := os.MkdirAll(filedir, os.ModePerm)
		if err != nil {
			log.Fatal(err)
		}

		// create file at path
		filepath := path.Join(publishDir, page.OutFile.Path)
		f, err := os.Create(filepath)
		if err != nil {
			log.Fatal(err)
		}
		defer f.Close()

		// write file with rendered template
		err = t.ExecuteTemplate(f, page.TemplateName, page)
		if err != nil {
			log.Fatal(err)
		}

		log.Printf("wrote %s\n", filepath)
	}
	log.Println("ending file writing")

}
