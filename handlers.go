package main

import (
	"fmt"
	"net/http"
	"time"

	"github.com/FamousHero/Pickleball_Queue/data"
)

func adminHandler(w http.ResponseWriter, r *http.Request) {
	renderTemplate(w, "/admin", &data.AdminPageData{})
}

func cssHandler(w http.ResponseWriter, r *http.Request) {
	http.ServeFile(w, r, "static/css"+r.URL.Path)
}

func playingHandler(w http.ResponseWriter, r *http.Request) {
	renderTemplate(w, "/playing", data.ActiveCourtInfo{
		Group:     *assignGroup,
		StartTime: time.Now(),
	})
}

func defaultHandler(w http.ResponseWriter, r *http.Request) {
	p := r.URL.Path
	if p == "/" {
		p = "/index"
	}
	fmt.Println(p)
	renderTemplate(w, p,
		&data.SignupPageData{
			PlayerInfo: &data.PlayerInfo{
				Name:       "Placeholder",
				Location:   "unknown",
				SkillGroup: "beginner",
			},
			Locations: []string{
				"Orange Grove",
				"Anaheim",
				"Montebello",
				"Commerce",
				"Modesto",
				"Santa Ana",
				"Simi Valley",
				"Pasadena",
			},
		})
}

func defaultPostHandler(w http.ResponseWriter, r *http.Request) {
	// when a form submission is received, parse the data and create
	// player info from it and add time to localPlayers
}

func changeGroupHandler(w http.ResponseWriter, r *http.Request) {
	// receive json with player_id, current_group_id, new_group_id
}

func leaveGroupHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method == http.MethodPost {
		err := r.ParseForm()
		if err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}

		pIdStr := r.FormValue("player-id")
		pId, err := validateArrayIndexFromForm(pIdStr, len(players))

		if err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}
		player = nil

		gIdStr := r.FormValue("group-id")
		gId, err := validateArrayIndexFromForm(gIdStr, len(currentQueue))

		if err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}
		currentQueue[gId].Players[pId] = nil
		delete(players, pId)
		fmt.Println("Player is ", player)

		fmt.Println(len(currentQueue))
		fmt.Println(player)
		http.Redirect(w, r, "/", http.StatusSeeOther)
	}
}

func queueHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Println("calling queueHandler")

	fmt.Println(currentQueue)
	fmt.Println(player)

	location := r.PathValue("location")
	renderTemplate(w, "/queue",
		&data.QueuePageData{
			Location:      location,
			Player:        player,
			AssignedGroup: assignGroup,
			CurrentQueue:  currentQueue, // []GroupInfo{},
			ActiveCourts:  activeCourts,
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
			PlayerInfo:   &data.PlayerInfo{},
			LocalPlayers: localPlayers,
		})
}
