package data

import "time"

type PlayerInfo struct {
	Name       string
	Location   string
	SkillGroup string
}

type GroupInfo struct {
	Players   [4]PlayerInfo
	CreatedAt time.Time
}

type ActiveCourtsInfo struct {
	Players   []GroupInfo
	StartTime time.Time
}
