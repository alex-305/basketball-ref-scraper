package main

import (
	"log"

	"github.com/alex-305/basketball-ref-scraper/db"
	"github.com/alex-305/basketball-ref-scraper/scrape"
)

func main() {
	const site = "https://www.basketball-reference.com"
	db := db.Connect()

	defer db.Disconnect()
	scrape.Start(site, db)

	log.Printf("%s", "Finished scraping all players.")

}
