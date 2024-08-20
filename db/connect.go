package db

import (
	"database/sql"
	"log"

	_ "github.com/mattn/go-sqlite3"
)

type DB struct {
	*sql.DB
}

func Connect() DB {
	sqlDb, err := sql.Open("sqlite3", "./nba.db")
	if err != nil {
		log.Fatal(err)
	}
	db := DB{sqlDb}
	db.Exec("PRAGMA foreign_keys = ON;")
	db.CreateTables()

	return db
}

func (db *DB) CreateTables() {
	db.CreateTeamTable()
	db.CreateTeamSeasonTable()
	db.CreatePlayerTable()
	db.CreatePlayerSeasonTable()
}

func (db *DB) Disconnect() {
	db.Close()
}
