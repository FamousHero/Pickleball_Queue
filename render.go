package main

import (
	"html/template"
	"net/http"
	"strings"
)

func renderTemplate(w http.ResponseWriter, tmpl string, data any) {
	funcMap := template.FuncMap{
		"inc": func(i int) int {
			return i + 1
		},
		"concatLocation": func(loc string) string {
			return strings.Replace(loc, " ", "-", -1)
		},
	}
	t, err := template.New(tmpl[1:] + ".html").Funcs(funcMap).ParseFiles("templates" + tmpl + ".html")
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	err = t.Execute(w, data)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}
