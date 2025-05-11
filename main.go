package main

import (
	"log"
	"net/http"

	"github.com/FamousHero/Pickleball_Queue/data"
)

/*
TODOS:
  - Look into regex matching for path validation
  - Define each templates data struct
  - add sql db for queue storage and retrieval
*/

var (
	players      = make(map[int]data.PlayerInfo)
	locationInfo = make(map[string][]data.PlayerInfo)
)

func main() {

	http.Handle("/static/", http.StripPrefix("/static", http.FileServer(http.Dir("static"))))
	http.HandleFunc("/view/{location}", viewHandler)
	http.HandleFunc("/queue", queueHandler)
	http.HandleFunc("/admin", adminHandler)
	http.HandleFunc("/{$}", defaultHandler)
	log.Fatal(http.ListenAndServe(":8080", nil))
}
