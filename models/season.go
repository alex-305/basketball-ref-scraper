package models

type Season struct {
	TeamID          string
	PlayerID        string
	PointsPerGame   *float32
	AssistsPerGame  *float32
	ReboundsPerGame *float32
	BlocksPerGame   *float32
	StealsPerGame   *float32
	Year            string
}
