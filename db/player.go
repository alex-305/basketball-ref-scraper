package db

import (
	"log"

	"github.com/alex-305/basketball-ref-scraper/models"
)

func (db *DB) CreatePlayerTable() {
	stmt, err := db.Prepare(`
	CREATE TABLE IF NOT EXISTS players (
		"id" TEXT NOT NULL PRIMARY KEY,
		"name" TEXT NOT NULL
	);
	`)

	if err != nil {
		log.Fatal(err)
	}

	stmt.Exec()
}

func (db *DB) CreatePlayerSeasonTable() {
	stmt, err := db.Prepare(`
	CREATE TABLE IF NOT EXISTS player_seasons (
		"teamid" TEXT NOT NULL,
		"playerid" TEXT NOT NULL,
		"year" TEXT NOT NULL,
		"position" TEXT NOT NULL,
		"gp" INTEGER,
		"age" INTEGER,
		"mp" REAL,
		"ppg" REAL,
		"apg" REAL,
		"rpg" REAL,
		"bpg" REAL,
		"spg" REAL,
		FOREIGN KEY(playerid) REFERENCES players(id),
		PRIMARY KEY(year, teamid, playerid)
	);
	`)
	if err != nil {
		log.Fatal(err)
	}

	stmt.Exec()
}

func (db *DB) InsertPlayer(player models.Player) error {
	query := `
	INSERT OR REPLACE INTO players(id, name) VALUES($1, $2);
	`
	_, err := db.Exec(query, player.Id, player.Name)

	if err != nil {
		return err
	}

	for _, season := range player.Seasons {
		err := db.InsertPlayerSeason(season)
		if err != nil {
			return err
		}
	}

	log.Printf("Inserted %s.", player.Name)

	return nil
}

func (db *DB) IDAvailable(id string) bool {
	query := `
	SELECT COUNT(*) FROM players WHERE id=$1;`
	row := db.QueryRow(query, id)
	var num int
	row.Scan(&num)
	return num == 0
}

func (db *DB) InsertPlayerSeason(season models.PlayerSeason) error {
	query := `
	INSERT OR REPLACE INTO player_seasons(teamid, playerid, year, ppg, apg, rpg, bpg, spg, gp, age, mp, position) VALUES($1, $2, $3, $4, $5, $6, $7, $8, $9, $10, $11, $12)
	`
	_, err := db.Exec(query, season.TeamID, season.PlayerID, season.Year, season.PointsPerGame, season.AssistsPerGame, season.ReboundsPerGame, season.BlocksPerGame, season.StealsPerGame, season.GamesPlayed, season.Age, season.MinutesPlayed, season.Position)

	if err != nil {
		return err
	}
	return nil
}
