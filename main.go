package main

import (
	"fmt"
	"log"
	"os"
	"path"
	"strings"

	"github.com/go-git/go-git/v5"
	"github.com/go-git/go-git/v5/plumbing/object"
	"github.com/yuin/goldmark"
	"github.com/yuin/goldmark-meta"
	"github.com/yuin/goldmark/extension"
	"github.com/yuin/goldmark/parser"
)

// Document represents a goliki file.
type Document struct {
	File       *object.File
	LastCommit *object.Commit
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
		goldmark.WithExtensions(
			extension.GFM,
			meta.Meta,
		),
	)

	// for each doc we need to calculate the resulting path
	// e.g. f/a.md > build/f/a.html
	// and write rendered contents here
	for _, d := range docs {

		//
		relpath := strings.Replace(d.File.Name, ".md", ".html", 1)
		buildpath := path.Join("build", relpath)

		fmt.Println(buildpath)

		// mkdir -p
		os.MkdirAll(path.Dir(buildpath), os.ModePerm)

		// get file content
		contents, err := d.File.Contents()
		if err != nil {
			log.Fatalf("could not get contents: %s", d.File.Name)
			continue
		}

		f, err := os.Create(buildpath)
		if err != nil {
			panic(err)
		}
		defer f.Close()

		// TODO: Custom Link renderer
		// TODO: Add https://github.com/yuin/goldmark-highlighting
		context := parser.NewContext()
		if err := markdown.Convert([]byte(contents), f, parser.WithContext(context)); err != nil {
			panic(err)
		}

	}

}
