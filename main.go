package main

import (
	"fmt"
	"log"
	"scrape/db"
	"scrape/models"

	"github.com/gocolly/colly/v2"
)

func main() {
	const site = "https://www.basketball-reference.com"

	c, q := NewCollyCollector(site)
	db := db.Connect()

	defer db.Disconnect()

	c.OnHTML("table#players tbody", func(e *colly.HTMLElement) {
		e.ForEach("tr th a", func(i int, a *colly.HTMLElement) {
			href := a.Attr("href")

			link := site + href

			player := models.Player{
				Id:   getID(link),
				Name: a.Text,
			}

			playerCollector := c.Clone()

			playerCollector.OnHTML("#per_game > tbody", func(e *colly.HTMLElement) {
				e.ForEach("tr", func(i int, h *colly.HTMLElement) {
					player.Seasons = append(player.Seasons, getSeason(h, player.Id))
				})
				err := db.InsertPlayer(player)
				if err != nil {
					log.Fatal(err)
				}
			})
			playerCollector.OnError(func(r *colly.Response, err error) {
				fmt.Printf("Error:%s", err)
			})
			playerCollector.Visit(link)
		})
	})

	c.OnError(func(r *colly.Response, err error) {
		fmt.Printf("Error:%s", err)
	})

	letter := 'a'
	for letter <= 'z' {

		q.AddURL(site + "/players/" + string(letter) + "/")
		letter += 1
	}
	q.Run(c)

}
