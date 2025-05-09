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

var (
	activeCourts []ActiveCourtsInfo
)

// test data
var currentQueue = []GroupInfo{
	{
		Player1: PlayerInfo{
			Name:       "Bob",
			Location:   "Modesto",
			SkillGroup: "Beginner+",
		},
		Player2: PlayerInfo{
			Name:       "Lola",
			Location:   "Modesto",
			SkillGroup: "Beginner+",
		},
		Player3: PlayerInfo{
			Name:       "Michael",
			Location:   "Modesto",
			SkillGroup: "Beginner+",
		},
		Player4: PlayerInfo{
			Name:       "Joe",
			Location:   "Modesto",
			SkillGroup: "Beginner+",
		},
	},
	{
		Player1: PlayerInfo{
			Name:       "Marco",
			Location:   "Modesto",
			SkillGroup: "Beginner+",
		},
		Player2: PlayerInfo{
			Name:       "Flint",
			Location:   "Modesto",
			SkillGroup: "Beginner+",
		},
		Player3: PlayerInfo{
			Name:       "Tori",
			Location:   "Modesto",
			SkillGroup: "Beginner+",
		},
		Player4: PlayerInfo{
			Name:       "Jerry",
			Location:   "Modesto",
			SkillGroup: "Beginner+",
		},
	},
	{
		Player1: PlayerInfo{
			Name:       "Jamie",
			Location:   "Modesto",
			SkillGroup: "Beginner+",
		},
		Player2: PlayerInfo{
			Name:       "Harry",
			Location:   "Modesto",
			SkillGroup: "Beginner+",
		},
		Player3: PlayerInfo{
			Name:       "Born",
			Location:   "Modesto",
			SkillGroup: "Beginner+",
		},
		Player4: PlayerInfo{
			Name:       "William",
			Location:   "Modesto",
			SkillGroup: "Beginner+",
		},
	},
	{
		Player1: PlayerInfo{
			Name:       "Cristian",
			Location:   "Modesto",
			SkillGroup: "Beginner+",
		},
		Player2: PlayerInfo{
			Name:       "Rice",
			Location:   "Modesto",
			SkillGroup: "Beginner+",
		},
		Player3: PlayerInfo{
			Name:       "Yonie",
			Location:   "Modesto",
			SkillGroup: "Beginner+",
		},
		Player4: PlayerInfo{
			Name:       "Edgar",
			Location:   "Modesto",
			SkillGroup: "Beginner+",
		},
	},
	{
		Player1: PlayerInfo{
			Name:       "Fernando",
			Location:   "Modesto",
			SkillGroup: "Beginner+",
		},
		Player2: PlayerInfo{
			Name:       "Cumin",
			Location:   "Modesto",
			SkillGroup: "Beginner+",
		},
		Player3: PlayerInfo{
			Name:       "Jose",
			Location:   "Modesto",
			SkillGroup: "Beginner+",
		},
		Player4: PlayerInfo{
			Name:       "Donnie",
			Location:   "Modesto",
			SkillGroup: "Beginner+",
		},
	},
	{
		Player1: PlayerInfo{
			Name:       "Claudia",
			Location:   "Modesto",
			SkillGroup: "Beginner+",
		},
		Player2: PlayerInfo{
			Name:       "Pepper",
			Location:   "Modesto",
			SkillGroup: "Beginner+",
		},
		Player3: PlayerInfo{
			Name:       "Terry",
			Location:   "Modesto",
			SkillGroup: "Beginner+",
		},
		Player4: PlayerInfo{
			Name:       "Brad",
			Location:   "Modesto",
			SkillGroup: "Beginner+",
		},
	},
	{
		Player1: PlayerInfo{
			Name:       "Regina",
			Location:   "Modesto",
			SkillGroup: "Beginner+",
		},
		Player2: PlayerInfo{
			Name:       "Fix",
			Location:   "Modesto",
			SkillGroup: "Beginner+",
		},
		Player3: PlayerInfo{
			Name:       "Maddie",
			Location:   "Modesto",
			SkillGroup: "Beginner+",
		},
		Player4: PlayerInfo{
			Name:       "Mac",
			Location:   "Modesto",
			SkillGroup: "Beginner+",
		},
	},
	{
		Player1: PlayerInfo{
			Name:       "Harry",
			Location:   "Modesto",
			SkillGroup: "Beginner+",
		},
		Player2: PlayerInfo{
			Name:       "Eduard",
			Location:   "Modesto",
			SkillGroup: "Beginner+",
		},
		Player3: PlayerInfo{
			Name:       "Will",
			Location:   "Modesto",
			SkillGroup: "Beginner+",
		},
		Player4: PlayerInfo{
			Name:       "Sarah",
			Location:   "Modesto",
			SkillGroup: "Beginner+",
		},
	},
	{
		Player1: PlayerInfo{
			Name:       "Olgo",
			Location:   "Modesto",
			SkillGroup: "Beginner+",
		},
		Player2: PlayerInfo{
			Name:       "Quinn",
			Location:   "Modesto",
			SkillGroup: "Beginner+",
		},
		Player3: PlayerInfo{
			Name:       "Pole",
			Location:   "Modesto",
			SkillGroup: "Beginner+",
		},
		Player4: PlayerInfo{
			Name:       "Arnold",
			Location:   "Modesto",
			SkillGroup: "Beginner+",
		},
	},
	{
		Player1: PlayerInfo{
			Name:       "Rechi",
			Location:   "Modesto",
			SkillGroup: "Beginner+",
		},
		Player2: PlayerInfo{
			Name:       "Josh",
			Location:   "Modesto",
			SkillGroup: "Beginner+",
		},
		Player3: PlayerInfo{
			Name:       "Vick",
			Location:   "Modesto",
			SkillGroup: "Beginner+",
		},
		Player4: PlayerInfo{
			Name:       "Quote",
			Location:   "Modesto",
			SkillGroup: "Beginner+",
		},
	},
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
			CurrentQueue: currentQueue, // []GroupInfo{},
			ActiveCourts: ActiveCourtsInfo{},
		})
}
