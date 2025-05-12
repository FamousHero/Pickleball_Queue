package data

import "time"

type PlayerInfo struct {
	PlayerId   int
	Name       string
	Location   string
	SkillGroup string
}

type GroupInfo struct {
	GroupId    int
	Players    [4]*PlayerInfo
	SkillLevel string
	CreatedAt  time.Time
}

type ActiveCourtInfo struct {
	Group     GroupInfo
	StartTime time.Time
}
