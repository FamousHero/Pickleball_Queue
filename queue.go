package main

import (
	"time"

	"github.com/FamousHero/Pickleball_Queue/data"
)

// Players should be made private, also set a createdAt field to know
// which group is ahead of which

// //////////////////////////////////
// //////// Queue Imp ///////////////
// //////////////////////////////////
type Queue[T any] struct {
	container []T
}

func (q *Queue[T]) Push(x T) {
	q.container = append(q.container, x)
}

func (q *Queue[T]) Pop() T {
	v := q.container[0]
	q.container = q.container[1:]
	return v
}

var activeCourts = []*data.ActiveCourtInfo{
	{
		StartTime: time.Now(),
		Group: data.GroupInfo{
			SkillLevel: "Intermediate",
			Players: [4]*data.PlayerInfo{
				{
					Name:       "Millie",
					Location:   "Ohio",
					SkillGroup: "Intermediate",
				},
				{
					Name:       "Bobby",
					Location:   "Ohio",
					SkillGroup: "Intermediate",
				},
				{
					Name:       "Brown",
					Location:   "Ohio",
					SkillGroup: "Intermediate",
				},
				{
					Name:       "Junior",
					Location:   "Ohio",
					SkillGroup: "Intermediate",
				},
			},
		},
	},
}

// ///////////////////////////////////
// ///////// test data ///////////////
// ///////////////////////////////////
var currentQueue = []data.GroupInfo{
	{
		SkillLevel: "Beginner+",
		Players: [4]*data.PlayerInfo{
			{
				Name:       "Bob",
				Location:   "Modesto",
				SkillGroup: "Beginner+",
			},
			nil,
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
		SkillLevel: "Beginner+",
		Players: [4]*data.PlayerInfo{
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
		SkillLevel: "Beginner+",
		Players: [4]*data.PlayerInfo{
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
			nil,
			{
				Name:       "William",
				Location:   "Modesto",
				SkillGroup: "Beginner+",
			},
		},
	},
	{
		SkillLevel: "Beginner+",
		Players: [4]*data.PlayerInfo{
			{
				Name:       "Cristian",
				Location:   "Modesto",
				SkillGroup: "Beginner+",
			},
			nil,
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
		SkillLevel: "Beginner+",
		Players: [4]*data.PlayerInfo{
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
		SkillLevel: "Beginner+",
		Players: [4]*data.PlayerInfo{
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
		SkillLevel: "Beginner+",
		Players: [4]*data.PlayerInfo{
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
		SkillLevel: "Beginner+",
		Players: [4]*data.PlayerInfo{
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
		SkillLevel: "Beginner+",
		Players: [4]*data.PlayerInfo{
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
		SkillLevel: "Beginner+",
		Players: [4]*data.PlayerInfo{
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
