package models

type Team struct {
	Id      string
	Name    string
	Seasons []TeamSeason
}

type TeamSeason struct {
	TeamID string
	Year   string
	Wins   int
	Losses int
}
