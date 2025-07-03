package model

type User struct {
	ID         int    `json:"id"`
	FName      string `json:"fname"`
	LName      string `json:"lname"`
	Email      string `json:"email"`
	Handle     string `json:"ig"`
	Passport   string `json:"citizen"`
	GameType   string `json:"mode"`
	StrongHand string `json:"hand"`
	RecentTeam string `json:"team"`
}
