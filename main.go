package main

import (
	"fmt"
	"html/template"
	"log"
	"net/http"
)

/*
TODOS:
  - Look into regex matching for path validation
  - Define each templates data struct
  - add sql db for queue storage and retrieval
*/

type PlayerInfo struct {
	Name       string
	Location   string
	SkillGroup string
}
type SignupPageData struct {
	PlayerInfo   PlayerInfo   `json: "playerInfo"`
	LocalPlayers []PlayerInfo `json: "localPlayers` // Current Players at that same location
}

var (
	players      = make(map[int]PlayerInfo)
	locationInfo = make(map[string][]PlayerInfo)
)

func defaultHandler(w http.ResponseWriter, r *http.Request) {
	p := r.URL.Path
	if p == "/" {
		p = "/index"
	}
	fmt.Println(p)
	renderTemplate(w, p,
		&SignupPageData{
			PlayerInfo: PlayerInfo{
				Name:       "Placeholder",
				Location:   "unknown",
				SkillGroup: "beginner",
			},
		})
}

func defaultPostHandler(w http.ResponseWriter, r *http.Request) {
	// when a form submission is received, parse the data and create
	// player info from it and add time to localPlayers
}
func viewHandler(w http.ResponseWriter, r *http.Request) {
	location := r.PathValue("location")
	fmt.Println(location)
	localPlayers, ok := locationInfo[location]
	if !ok {
		w.WriteHeader(http.StatusNotFound)
		fmt.Fprintf(w, "Error 404: Page not Found")
		return
	}
	renderTemplate(w, "index",
		&SignupPageData{
			PlayerInfo:   PlayerInfo{},
			LocalPlayers: localPlayers,
		})
}

func renderTemplate(w http.ResponseWriter, tmpl string, data any) {
	funcMap := template.FuncMap{
		"inc": func(i int) int {
			return i + 1
		},
	}
	t, err := template.New("queue.html").Funcs(funcMap).ParseFiles("static/templates" + tmpl + ".html")
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	err = t.Execute(w, data)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}

func main() {

	http.HandleFunc("/view/{location}", viewHandler)
	http.HandleFunc("/queue", queueHandler)
	http.HandleFunc("/admin", adminHandler)
	http.HandleFunc("/{$}", defaultHandler)
	log.Fatal(http.ListenAndServe(":8080", nil))
}
