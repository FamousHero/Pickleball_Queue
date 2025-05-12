package main

import (
	"log"
	"net/http"
	"slices"

	"github.com/FamousHero/Pickleball_Queue/data"
)

/*
TODOS:
  - Look into regex matching for path validation
  - Define each templates data struct
  - add sql db for queue storage and retrieval
*/

var (
	players      = make(map[int]*data.PlayerInfo)
	locationInfo = make(map[string][]*data.PlayerInfo)
)

func main() {
	playerCount := 0
	players[playerCount] = player
	player.PlayerId = playerCount
	playerCount++

	assignGroup.GroupId = len(currentQueue) - 3
	currentQueue = slices.Insert(currentQueue, assignGroup.GroupId, assignGroup)

	http.Handle("/static/", http.StripPrefix("/static", http.FileServer(http.Dir("static"))))

	http.HandleFunc("POST /change-group", changeGroupHandler)
	http.HandleFunc("POST /leave-group", leaveGroupHandler)

	http.HandleFunc("/view/{location}", viewHandler)
	http.HandleFunc("/queue/{location}", queueHandler)
	http.HandleFunc("/playing", playingHandler)
	http.HandleFunc("/admin", adminHandler)
	http.HandleFunc("/{$}", defaultHandler)
	log.Fatal(http.ListenAndServe(":8080", nil))
}
