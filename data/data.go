package data

type SignupPageData struct {
	Locations    []string
	PlayerInfo   *PlayerInfo   `json: "playerInfo"`
	LocalPlayers []*PlayerInfo `json: "localPlayers` // Current Players at that same location
}

type QueuePageData struct {
	Location      string
	Player        *PlayerInfo // can check queue even if not signed up
	AssignedGroup *GroupInfo
	CurrentQueue  []*GroupInfo       // Change to Queue[GroupInfo]
	ActiveCourts  []*ActiveCourtInfo // If nobody on a court, set that court to nil
}

type AdminPageData struct {
}
