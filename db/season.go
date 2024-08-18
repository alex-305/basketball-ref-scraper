package db

import (
	"log"
	"scrape/models"
)

type Season struct {
	TeamID          string
	PlayerID        string
	PointsPerGame   float32
	AssistsPerGame  float32
	ReboundsPerGame float32
	BlocksPerGame   float32
	StealsPerGame   float32
	Year            string
}

func (db *DB) CreateSeasonTable() {
	stmt, err := db.Prepare(`
	CREATE TABLE IF NOT EXISTS seasons (
		"teamid" TEXT NOT NULL,
		"playerid" TEXT NOT NULL,
		"year" TEXT NOT NULL,
		"ppg" REAL,
		"apg" REAL,
		"rpg" REAL,
		"bpg" REAL,
		"spg" REAL,
		PRIMARY KEY(year, teamid, playerid)
	);
	`)
	if err != nil {
		log.Fatal(err)
	}

	stmt.Exec()
}

func (db *DB) InsertSeason(season models.Season) error {
	query := `
	INSERT OR REPLACE INTO seasons(teamid, playerid, year, ppg, apg, rpg, bpg, spg) VALUES($1, $2, $3, $4, $5, $6, $7, $8)
	`
	_, err := db.Exec(query, season.TeamID, season.PlayerID, season.Year, season.PointsPerGame, season.AssistsPerGame, season.ReboundsPerGame, season.BlocksPerGame, season.StealsPerGame)

	if err != nil {
		return err
	}
	return nil
}
