package main

import (
	"fmt"
	"log"
	"net/http"
	"text/template"
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
	PlayerInfo   PlayerInfo
	LocalPlayers []PlayerInfo
}

var (
	players      = make(map[int]PlayerInfo)
	locationInfo = make(map[string][]PlayerInfo)
)

func defaultHandler(w http.ResponseWriter, r *http.Request) {
	renderTemplate(w, "index",
		&SignupPageData{
			PlayerInfo: PlayerInfo{
				Name:       "Placeholder",
				Location:   "unknown",
				SkillGroup: "beginner",
			},
		})
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

func renderTemplate(w http.ResponseWriter, tmpl string, spd *SignupPageData) {
	t, err := template.ParseFiles(tmpl + ".html")
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	err = t.Execute(w, spd)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}

func main() {
	http.HandleFunc("/view/{location}", viewHandler)
	http.HandleFunc("/{$}", defaultHandler)
	log.Fatal(http.ListenAndServe(":8080", nil))
}
