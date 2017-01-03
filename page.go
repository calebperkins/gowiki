package main

import (
	"io/ioutil"
	"path/filepath"
)

// Page represents a wiki article.
type Page struct {
	Title string
	Body  string
}

func pagePath(title string) string {
	return filepath.Join("data", title+".txt")
}

func (p *Page) save() error {
	filename := pagePath(p.Title)
	return ioutil.WriteFile(filename, []byte(p.Body), 0600)
}

func loadPage(title string) (*Page, error) {
	filename := pagePath(title)
	body, err := ioutil.ReadFile(filename)
	if err != nil {
		return nil, err
	}
	return &Page{Title: title, Body: string(body)}, nil
}
