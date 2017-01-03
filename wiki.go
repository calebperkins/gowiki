package main

import (
	"errors"
	"html/template"
	"log"
	"net/http"
	"regexp"

	"github.com/gorilla/mux"
)

var templates = template.Must(template.ParseFiles("templates/edit.html", "templates/view.html"))
var validPath = regexp.MustCompile("^/(edit|save|view)/([a-zA-Z0-9]+)$")

func getTitle(w http.ResponseWriter, r *http.Request) (string, error) {
	m := validPath.FindStringSubmatch(r.URL.Path)
	if m == nil {
		http.NotFound(w, r)
		return "", errors.New("Invalid page title")
	}
	return m[2], nil
}

func pageTitle(r *http.Request) string {
	return mux.Vars(r)["title"]
}

func main() {
	r := NewRouter()
	log.Fatal(http.ListenAndServe(":8080", r))
}
