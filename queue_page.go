package main

import (
	"net/http"
	"time"
)

type QueuePageData struct {
	Player        PlayerInfo
	AssignedGroup GroupInfo
	CurrentQueue  []GroupInfo
	ActiveCourts  ActiveCourtsInfo
}

type GroupInfo struct {
	Player1, Player2, Player3, Player4 PlayerInfo
}

type ActiveCourtsInfo struct {
	Players   []GroupInfo
	StartTime time.Time
}

func queueHandler(w http.ResponseWriter, r *http.Request) {

	renderTemplate(w, "/queue",
		&QueuePageData{
			Player: PlayerInfo{
				Name: "Test",
			},
			AssignedGroup: GroupInfo{
				Player1: PlayerInfo{
					Name:     "Player1",
					Location: "Test Local",
				},
				Player2: PlayerInfo{
					Name:     "Player2",
					Location: "Test Local",
				},
				Player3: PlayerInfo{
					Name:     "Player3",
					Location: "Test Local",
				},
				Player4: PlayerInfo{
					Name:     "Player4",
					Location: "Test Local",
				},
			},
			CurrentQueue: []GroupInfo{},
			ActiveCourts: ActiveCourtsInfo{},
		})
}
