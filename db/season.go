package db

import (
	"log"
	"scrape/models"
)

func (db *DB) CreateSeasonTable() {
	stmt, err := db.Prepare(`
	CREATE TABLE IF NOT EXISTS seasons (
		"teamid" TEXT NOT NULL,
		"playerid" TEXT NOT NULL,
		"year" TEXT NOT NULL,
		"gp" INTEGER,
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
