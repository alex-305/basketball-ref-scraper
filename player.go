package main

import (
	"fmt"
	"log"
	"strconv"
	"strings"

	"scrape/db"
	"scrape/models"

	"github.com/gocolly/colly/v2"
)

func getAllPlayers(site string) {
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

func getID(link string) string {
	parts := strings.Split(link, "/")
	last := parts[len(parts)-1]

	return strings.TrimSuffix(last, ".html")
}

func getYear(id string) string {
	parts := strings.Split(id, ".")
	dateString := parts[len(parts)-1]

	return dateString
}

func getFloatStat(str string) (float32, error) {
	stat, err := strconv.ParseFloat(str, 32)

	if err != nil {
		return 0.0, err
	}
	return float32(stat), nil
}

func getIntStat(str string) int {
	stat, _ := strconv.ParseInt(str, 10, 32)
	return int(stat)
}

func statToAttr(stat string) string {
	return "td[data-stat=" + stat + "]"
}

func getSeason(e *colly.HTMLElement, playerID string) models.Season {
	season := models.Season{}
	season.TeamID = e.ChildText(statToAttr("team_id") + " a")
	season.PlayerID = playerID
	season.Year = getYear(e.Attr("id"))

	season.GamesPlayed = getIntStat(e.ChildText(statToAttr("g")))
	season.Age = getIntStat(e.ChildText(statToAttr("age")))
	mpg, err := getFloatStat(statToAttr("mp_per_g"))

	if err != nil {
		season.MinutesPlayed = &mpg
	}

	rpg, err := getFloatStat(e.ChildText(statToAttr("trb_per_g")))
	if err == nil {
		season.ReboundsPerGame = &rpg
	}
	ppg, err := getFloatStat(e.ChildText(statToAttr("pts_per_g")))
	if err == nil {
		season.PointsPerGame = &ppg
	}
	apg, err := getFloatStat(e.ChildText(statToAttr("ast_per_g")))
	if err == nil {
		season.AssistsPerGame = &apg
	}
	bpg, err := getFloatStat(e.ChildText(statToAttr("blk_per_g")))
	if err == nil {
		season.BlocksPerGame = &bpg
	}
	spg, err := getFloatStat(e.ChildText(statToAttr("stl_per_g")))
	if err == nil {
		season.StealsPerGame = &spg
	}
	return season
}
