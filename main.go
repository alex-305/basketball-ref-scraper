package main

import (
	"log"

	"github.com/alex-305/basketball-ref-scraper/scrape"
)

func main() {
	const site = "https://www.basketball-reference.com"
	scrape.Start(site)

	log.Printf("%s", "Finished scraping all players.")

}
