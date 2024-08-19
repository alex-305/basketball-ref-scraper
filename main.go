package main

import "log"

func main() {
	const site = "https://www.basketball-reference.com"

	getAllPlayers(site)

	log.Printf("%s", "Finished scraping all players.")

}
