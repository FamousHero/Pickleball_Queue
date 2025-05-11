package main

import (
	"fmt"
	"net/http"

	"github.com/FamousHero/Pickleball_Queue/data"
)

func adminHandler(w http.ResponseWriter, r *http.Request) {
	renderTemplate(w, "/admin", &data.AdminPageData{})
}

func cssHandler(w http.ResponseWriter, r *http.Request) {
	http.ServeFile(w, r, "static/css"+r.URL.Path)
}

func defaultHandler(w http.ResponseWriter, r *http.Request) {
	p := r.URL.Path
	if p == "/" {
		p = "/index"
	}
	fmt.Println(p)
	renderTemplate(w, p,
		&data.SignupPageData{
			PlayerInfo: data.PlayerInfo{
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
func queueHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Println("calling queueHandler")
	renderTemplate(w, "/queue",
		&data.QueuePageData{
			Player: data.PlayerInfo{
				Name: "Test",
			},
			AssignedGroup: data.GroupInfo{
				Players: [4]data.PlayerInfo{
					{
						Name:     "Player1",
						Location: "Test Local",
					},
					{
						Name:     "Player2",
						Location: "Test Local",
					},
					{
						Name:     "Player3",
						Location: "Test Local",
					},
					{
						Name:     "Player4",
						Location: "Test Local",
					},
				},
			},
			CurrentQueue: currentQueue, // []GroupInfo{},
			ActiveCourts: activeCourts,
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
		&data.SignupPageData{
			PlayerInfo:   data.PlayerInfo{},
			LocalPlayers: localPlayers,
		})
}
