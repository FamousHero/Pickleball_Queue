package main

import (
	"net/http"
)

type AdminPageData struct {
}

func adminHandler(w http.ResponseWriter, r *http.Request) {
	renderTemplate(w, "/admin", &AdminPageData{})
}
