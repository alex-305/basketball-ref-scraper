package db

import (
	"log"

	"github.com/alex-305/basketball-ref-scraper/models"
)

func (db *DB) CreateTeamTable() {
	stmt, err := db.Prepare(`
	CREATE TABLE IF NOT EXISTS teams (
		"id" TEXT NOT NULL PRIMARY KEY,
		"name" TEXT NOT NULL
	);
	`)

	if err != nil {
		log.Fatal(err)
	}

	stmt.Exec()
}

func (db *DB) CreateTeamSeasonTable() {
	stmt, err := db.Prepare(`
	CREATE TABLE IF NOT EXISTS team_seasons (
		"teamid" TEXT NOT NULL,
		"year" TEXT NOT NULL,
		"wins" INTEGER NOT NULL,    
		"losses" INTEGER NOT NULL,
		FOREIGN KEY(teamid) REFERENCES teams(id),
		PRIMARY KEY(teamid, year)
	);
	`)
	if err != nil {
		log.Fatal(err)
	}
	stmt.Exec()
}

func (db *DB) InsertTeam(team models.Team) error {
	query := `
	INSERT OR REPLACE INTO teams(id, name) VALUES($1, $2);
	`
	_, err := db.Exec(query, team.Id, team.Name)

	if err != nil {
		return err
	}

	for _, season := range team.Seasons {
		err := db.InsertTeamSeason(season)
		if err != nil {
			return err
		}
	}

	log.Printf("Inserted %s.", team.Name)

	return nil
}

func (db *DB) InsertTeamSeason(season models.TeamSeason) error {
	return nil
}
