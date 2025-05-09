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
	Players [4]PlayerInfo
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
		Players: [4]PlayerInfo{
			{
				Name:       "Bob",
				Location:   "Modesto",
				SkillGroup: "Beginner+",
			},
			{
				Name:       "Lola",
				Location:   "Modesto",
				SkillGroup: "Beginner+",
			},
			{
				Name:       "Michael",
				Location:   "Modesto",
				SkillGroup: "Beginner+",
			},
			{
				Name:       "Joe",
				Location:   "Modesto",
				SkillGroup: "Beginner+",
			},
		},
	},
	{
		Players: [4]PlayerInfo{
			{
				Name:       "Marco",
				Location:   "Modesto",
				SkillGroup: "Beginner+",
			},
			{
				Name:       "Flint",
				Location:   "Modesto",
				SkillGroup: "Beginner+",
			},
			{
				Name:       "Tori",
				Location:   "Modesto",
				SkillGroup: "Beginner+",
			},
			{
				Name:       "Jerry",
				Location:   "Modesto",
				SkillGroup: "Beginner+",
			},
		},
	},
	{
		Players: [4]PlayerInfo{
			{
				Name:       "Jamie",
				Location:   "Modesto",
				SkillGroup: "Beginner+",
			},
			{
				Name:       "Harry",
				Location:   "Modesto",
				SkillGroup: "Beginner+",
			},
			{
				Name:       "Born",
				Location:   "Modesto",
				SkillGroup: "Beginner+",
			},
			{
				Name:       "William",
				Location:   "Modesto",
				SkillGroup: "Beginner+",
			},
		},
	},
	{
		Players: [4]PlayerInfo{
			{
				Name:       "Cristian",
				Location:   "Modesto",
				SkillGroup: "Beginner+",
			},
			{
				Name:       "Rice",
				Location:   "Modesto",
				SkillGroup: "Beginner+",
			},
			{
				Name:       "Yonie",
				Location:   "Modesto",
				SkillGroup: "Beginner+",
			},
			{
				Name:       "Edgar",
				Location:   "Modesto",
				SkillGroup: "Beginner+",
			},
		},
	},
	{
		Players: [4]PlayerInfo{
			{
				Name:       "Fernando",
				Location:   "Modesto",
				SkillGroup: "Beginner+",
			},
			{
				Name:       "Cumin",
				Location:   "Modesto",
				SkillGroup: "Beginner+",
			},
			{
				Name:       "Jose",
				Location:   "Modesto",
				SkillGroup: "Beginner+",
			},
			{
				Name:       "Donnie",
				Location:   "Modesto",
				SkillGroup: "Beginner+",
			},
		},
	},
	{
		Players: [4]PlayerInfo{
			{
				Name:       "Claudia",
				Location:   "Modesto",
				SkillGroup: "Beginner+",
			},
			{
				Name:       "Pepper",
				Location:   "Modesto",
				SkillGroup: "Beginner+",
			},
			{
				Name:       "Terry",
				Location:   "Modesto",
				SkillGroup: "Beginner+",
			},
			{
				Name:       "Brad",
				Location:   "Modesto",
				SkillGroup: "Beginner+",
			},
		},
	},
	{
		Players: [4]PlayerInfo{
			{
				Name:       "Regina",
				Location:   "Modesto",
				SkillGroup: "Beginner+",
			},
			{
				Name:       "Fix",
				Location:   "Modesto",
				SkillGroup: "Beginner+",
			},
			{
				Name:       "Maddie",
				Location:   "Modesto",
				SkillGroup: "Beginner+",
			},
			{
				Name:       "Mac",
				Location:   "Modesto",
				SkillGroup: "Beginner+",
			},
		},
	},
	{
		Players: [4]PlayerInfo{
			{
				Name:       "Harry",
				Location:   "Modesto",
				SkillGroup: "Beginner+",
			},
			{
				Name:       "Eduard",
				Location:   "Modesto",
				SkillGroup: "Beginner+",
			},
			{
				Name:       "Will",
				Location:   "Modesto",
				SkillGroup: "Beginner+",
			},
			{
				Name:       "Sarah",
				Location:   "Modesto",
				SkillGroup: "Beginner+",
			},
		},
	},
	{
		Players: [4]PlayerInfo{
			{
				Name:       "Olgo",
				Location:   "Modesto",
				SkillGroup: "Beginner+",
			},
			{
				Name:       "Quinn",
				Location:   "Modesto",
				SkillGroup: "Beginner+",
			},
			{
				Name:       "Pole",
				Location:   "Modesto",
				SkillGroup: "Beginner+",
			},
			{
				Name:       "Arnold",
				Location:   "Modesto",
				SkillGroup: "Beginner+",
			},
		},
	},
	{
		Players: [4]PlayerInfo{
			{
				Name:       "Rechi",
				Location:   "Modesto",
				SkillGroup: "Beginner+",
			},
			{
				Name:       "Josh",
				Location:   "Modesto",
				SkillGroup: "Beginner+",
			},
			{
				Name:       "Vick",
				Location:   "Modesto",
				SkillGroup: "Beginner+",
			},
			{
				Name:       "Quote",
				Location:   "Modesto",
				SkillGroup: "Beginner+",
			},
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
				Players: [4]PlayerInfo{
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
			ActiveCourts: ActiveCourtsInfo{},
		})
}
