package main

import (
	"bytes"
	"fmt"
	"io"
	"io/ioutil"
	"log"
	"net/url"
	"os"
	"path"
	"path/filepath"
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
	Markdown *bytes.Buffer
	// MarkdownHTML is rendered from GitFile content.
	MarkdownHTML *bytes.Buffer
	// HTML is content rendered from template.
	HTML *bytes.Buffer
	// Metadata is read from YAML Front Matter.
	Metadata map[string]interface{}
	// Site contains information about the entire site and all pages.
	Site *Site
	// File is simple metadata regarding original input file of page.
	File *File

	TemplateName string

	Permalink    string
	RelPermalink string
	Section      string
	Title        string
}

// NewPage creates a new Page from path and bytes.
// Bytes represent the raw Markdown content.
func NewPage(path string, b []byte) *Page {
	// Should this be NewPage()?
	page := &Page{
		File:     NewFile(path),
		Markdown: bytes.NewBuffer(b),
	}
	return page
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
		// replace README with index
		newDest = strings.Replace(newDest, "README", "index", 1)

		link.Destination = []byte(newDest)

		return ast.WalkContinue, nil
	})
}

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

	tmplFilePaths, err := filepath.Glob(".goliki/layouts/*.html")
	if err != nil {
		log.Fatalln(err)
	}

	// load and parse templates
	t := template.Must(template.ParseFiles(tmplFilePaths...))
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

	pages := []*Page{}
	paths, err := listFilePaths(".", filePattern)
	if err != nil {
		log.Fatal(err)
	}
	for _, path := range paths {
		page, err := readPage(path)
		if err != nil {
			log.Println(err)
			continue
		}
		addGit(page, r)
		renderMarkdown(page, markdown)
		pages = append(pages, page)
	}
	reduceSite(pages)
	for _, page := range pages {
		renderPage(page, t)
		writePage(page, publishDir)
	}
}

// listFilepaths traverses the file tree at rootDir and returns an array of paths.
func listFilePaths(rootDir string, filePattern string) ([]string, error) {
	// fpr is the filepath regex
	fpr, err := regexp.Compile(filePattern)
	if err != nil {
		log.Println(err)
		return nil, err
	}

	paths := []string{}
	// https://pkg.go.dev/path/filepath@go1.15.2#WalkFunc
	err = filepath.Walk(rootDir, func(path string, info os.FileInfo, err error) error {
		if err != nil {
			return err
		}
		if !info.IsDir() && fpr.MatchString(path) {
			log.Println(path)
			paths = append(paths, path)
		}
		return nil
	})
	if err != nil {
		log.Println(err)
		return paths, err
	}

	return paths, err
}

func readPage(path string) (*Page, error) {
	b, err := ioutil.ReadFile(path)
	if err != nil {
		return nil, err
	}
	log.Printf("Read %s\n", path)
	return NewPage(path, b), nil
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
		log.Println(err)
		log.Println("OK")
		return page
	}

	lastCommit, err := commitIter.Next()
	if err != nil {
		if err == io.EOF {
			log.Printf("No commit history: %s\n", page.File.Path)
		}
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
	if err := markdown.Convert(page.Markdown.Bytes(), page.MarkdownHTML, parser.WithContext(context)); err != nil {
		log.Println(err)
		return page
	}

	// TODO: Metadata is too raw
	page.Metadata = meta.Get(context)

	// will this error?
	page.Section = strings.Split(page.File.Dir, "/")[0]

	filename := fmt.Sprintf("%s.%s", page.File.ContentBaseName, "html")

	// FIXME: This relative URL is only correct at root
	page.RelPermalink = path.Join(page.File.Dir, filename)
	page.Permalink = path.Join("/", page.File.Dir, filename)
	page.Title = page.File.ContentBaseName
	return page
}

// reduceSite processes pages to and adds site aggregate data to each one.
func reduceSite(pages []*Page) []*Page {
	site := NewSite()
	for _, page := range pages {

		// do not add README/index to Section list
		if page.Section != "" && page.File.ContentBaseName != "README" {
			site.Sections[page.Section] = append(site.Sections[page.Section], page)
		}

		// embed global Site var into each page
		page.Site = site
	}
	return pages
}

// renderPage renders HTML for page.
func renderPage(page *Page, t *template.Template) *Page {

	// set template if it has not been defined
	if page.TemplateName == "" {
		page.TemplateName = "default.html"
		if page.File.LogicalName == "README.md" {
			page.TemplateName = "README.html"
		}
		// This is readme @ root
		if page.File.Path == "README.md" {
			page.TemplateName = "index.html"
		}
	}

	// write file with rendered template
	page.HTML = new(bytes.Buffer)
	err := t.ExecuteTemplate(page.HTML, page.TemplateName, page)
	if err != nil {
		// This is not good
		log.Fatal(err)
	}
	return page
}

// writePage writes page HTML to path at publishDir.
// it ensures the path exists
func writePage(page *Page, publishDir string) *Page {
	// FIXME: check that publishDir exists

	// ensure destination exists
	pageDir := path.Join(publishDir, page.File.Dir)
	err := os.MkdirAll(pageDir, os.ModePerm)
	if err != nil {
		log.Println(err)
		return nil
	}

	// README.md -> index.html
	contentBaseName := page.File.ContentBaseName
	if contentBaseName == "README" {
		contentBaseName = "index"
	}

	logicalName := fmt.Sprintf("%s.%s", contentBaseName, "html")
	pagePath := path.Join(pageDir, logicalName)

	f, err := os.Create(pagePath)
	if err != nil {
		log.Println(err)
		return nil
	}
	defer f.Close()
	f.Write(page.HTML.Bytes())

	return page
}
