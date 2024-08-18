package db

import (
	"log"
	"scrape/models"
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

func (db *DB) InsertPlayer(player models.Player) error {
	query := `
	INSERT OR REPLACE INTO players(id, name) VALUES($1, $2);
	`
	_, err := db.Exec(query, player.Id, player.Name)

	if err != nil {
		return err
	}

	for _, season := range player.Seasons {
		err := db.InsertSeason(season)
		if err != nil {
			return err
		}
	}

	log.Printf("Inserted %s into db", player.Name)

	return nil
}
