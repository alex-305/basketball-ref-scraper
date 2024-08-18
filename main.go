package main

import (
	"fmt"
	"log"
	"scrape/db"

	"github.com/gocolly/colly/v2"
)

func main() {
	const site = "https://www.basketball-reference.com"

	c := NewCollyCollector(site)
	db := db.Connect()

	c.OnHTML("table#players tbody", func(e *colly.HTMLElement) {
		e.ForEach("tr th a", func(i int, a *colly.HTMLElement) {
			href := a.Attr("href")
			player := GetPlayer(site, href, a.Text)
			err := db.InsertPlayer(player)

			if err != nil {
				log.Fatal(err)
			}
		})
	})

	c.OnError(func(r *colly.Response, err error) {
		fmt.Printf("Error:%s", err)
	})

	letter := 'a'
	for letter <= 'z' {

		c.Visit(site + "/players/" + string(letter) + "/")

		if letter+1 == 'x' {
			letter += 1
		}
		letter += 1
	}
}
