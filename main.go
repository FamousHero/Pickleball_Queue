package main

import (
	"fmt"
	"log"
	"net/http"
	"os"
	"text/template"
)

/*
TODOS:
  - Look into regex matching for path validation
  - Define each templates data struct
  - add sql db for queue storage and retrieval
*/
type Page struct {
	Title string
	Body  []byte
}

func (p *Page) save() error {
	filename := p.Title + ".txt"
	return os.WriteFile(filename, p.Body, 0600)
}

func loadPage(title string) (*Page, error) {
	filename := title + ".txt"
	body, err := os.ReadFile(filename)
	if err != nil {
		return nil, err
	}
	return &Page{Title: title, Body: body}, nil
}

func defaultHandler(w http.ResponseWriter, r *http.Request) {
	renderTemplate(w, "index", &Page{Title: "home page"})
}

func viewHandler(w http.ResponseWriter, r *http.Request) {
	title := r.PathValue("stand_in")
	fmt.Println(title)
	p, err := loadPage(title)
	if err != nil {
		w.WriteHeader(http.StatusNotFound)
		fmt.Fprintf(w, "Error 404: Page not Found")
		return
	}
	renderTemplate(w, "index", p)
}

func editHandler(w http.ResponseWriter, r *http.Request) {
	title := r.URL.Path[len("/edit/"):]
	p, err := loadPage(title)
	if err != nil {
		p = &Page{Title: title}
	}
	renderTemplate(w, title, p)
}

func renderTemplate(w http.ResponseWriter, tmpl string, p *Page) {
	t, err := template.ParseFiles(tmpl + ".html")
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	err = t.Execute(w, p)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}

func main() {
	http.HandleFunc("/view/{stand_in}", viewHandler)
	http.HandleFunc("/{$}", defaultHandler)
	log.Fatal(http.ListenAndServe(":8080", nil))
}
