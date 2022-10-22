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

// Page represents a goliki file.
// It contains the file data and metadata at multiple stages in the build pipeline.
// It is passed into templates.
type Page struct {
	// LastCommit is the latest commit that modified this file.
	LastCommit *object.Commit
	// Markdown is the raw markdown file.
	Markdown string
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

// NewFile returns ref to new F from parsed path to file.
func NewFile(fpath string) *File {
	f := &File{}
	f.Path = fpath
	f.Dir, f.LogicalName = path.Split(f.Path)
	f.Ext = path.Ext(f.LogicalName)
	f.ContentBaseName = strings.Replace(f.LogicalName, f.Ext, "", 1)
	return f
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

	// check for local git repo
	gitpath := "."
	r, err := git.PlainOpen(gitpath)
	if err != nil {
		log.Fatal(err)
	}

	pagesTask(r, `[\w]+\.md`, ".goliki/public/")

}

// pagesTask runt the Page pipeline
func pagesTask(r *git.Repository, filePattern string, publishDir string) {

	// load and parse templates
	t := template.Must(template.ParseFiles(
		".goliki/layouts/default.html",
		".goliki/layouts/index.html",
	))
	log.Printf("found templates: %s\n", t.DefinedTemplates())

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

	// actual pipeline
	pages := readPages(r, filePattern)
	for _, page := range pages {
		addGit(page, r)
		renderMarkdown(page, markdown)
	}
	reduceSite(pages)
	for _, page := range pages {
		renderWritePage(page, publishDir, t)
	}
}

// readPage needs to be refactored when we read from fs instead.
func readPages(r *git.Repository, re string) []*Page {
	pageRegex, err := regexp.Compile(re)
	if err != nil {
		log.Fatal(err)
		return nil
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

	pages := []*Page{}
	tree.Files().ForEach(func(f *object.File) error {

		ff := NewFile(f.Name)
		if !pageRegex.MatchString(ff.Path) {
			return nil
		}
		log.Printf("processing %s\n", ff.Path)

		// get file content
		contents, err := f.Contents()
		if err != nil {
			log.Fatalf("could not get contents: %s\n", f.Name)
		}

		page := &Page{
			File:     ff,
			Markdown: contents,
		}

		pages = append(pages, page)

		return nil
	})
	log.Printf("processed %d files\n", len(pages))

	return pages
}

// addGit adds latest commit info to Page.
func addGit(page *Page, r *git.Repository) *Page {

	// git log -- FILENAME
	commitIter, err := r.Log(&git.LogOptions{
		// it's HEAD by default
		// From:     ref.Hash(),
		FileName: &page.File.Path,
	})
	if err != nil {
		log.Fatal(err)
		return page
	}

	lastCommit, err := commitIter.Next()
	if err != nil {
		log.Fatal(err)
		return page
	}

	page.LastCommit = lastCommit
	return page
}

// renderMarkdown renders and stores HTML from Page's markdown.
// It also parses metadata into appropriate Page variables.
func renderMarkdown(page *Page, markdown goldmark.Markdown) *Page {

	// Why do I have to do this?
	page.MarkdownHTML = new(bytes.Buffer)

	// TODO: Add https://github.com/yuin/goldmark-highlighting
	context := parser.NewContext()
	if err := markdown.Convert([]byte(page.Markdown), page.MarkdownHTML, parser.WithContext(context)); err != nil {
		log.Printf("%s", err)
		return page
	}

	// TODO: Metadata is too raw
	page.Metadata = meta.Get(context)

	// will this error?
	page.Section = strings.Split(page.File.Dir, "/")[0]

	// Outfile represents how the final file is to be written.
	// FIXME: this probably needs to go elsewhere
	sf := &File{}
	sf.Ext = "html"
	sf.ContentBaseName = page.File.ContentBaseName
	if page.File.ContentBaseName == "README" {
		sf.ContentBaseName = "index"
	}
	sf.LogicalName = fmt.Sprintf("%s.%s", sf.ContentBaseName, sf.Ext)
	sf.Dir = page.File.Dir
	sf.Path = path.Join(sf.Dir, sf.LogicalName)
	page.OutFile = sf

	// select our template
	page.TemplateName = "default.html"
	if sf.ContentBaseName == "index" {
		page.TemplateName = "index.html"
	}

	page.RelPermalink = sf.Path
	page.Title = sf.ContentBaseName
	return page
}

// reduceSite processes pages to and adds site aggregate data to each one.
func reduceSite(pages []*Page) []*Page {
	site := NewSite()
	for _, page := range pages {

		if page.Section != "" {
			site.Sections[page.Section] = append(site.Sections[page.Section], page)
		}

		// embed global Site var into each page
		page.Site = site
	}
	return pages
}

// renderWritePage renders final HTML for page and write it to disk.
// TODO: Should be two functions: one render template into a new buffer in Page,
// and the other to write buffer to fs.
func renderWritePage(page *Page, publishDir string, t *template.Template) *Page {
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
	return page
}
