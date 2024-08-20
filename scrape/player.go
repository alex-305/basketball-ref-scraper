package scrape

import (
	"fmt"
	"log"

	"github.com/alex-305/basketball-ref-scraper/db"
	"github.com/alex-305/basketball-ref-scraper/models"

	"github.com/gocolly/colly/v2"
)

func GetAllPlayers(site string) {
	c, q := NewCollyCollector(site)
	db := db.Connect()

	defer db.Disconnect()

	c.OnHTML("table#players tbody", func(e *colly.HTMLElement) {
		e.ForEach("tr th a", func(i int, a *colly.HTMLElement) {
			href := a.Attr("href")

			link := site + href

			player := models.Player{
				Id:   getPlayerIDFromHref(link),
				Name: a.Text,
			}

			playerCollector := c.Clone()

			playerCollector.OnHTML("#per_game > tbody", func(e *colly.HTMLElement) {
				e.ForEach("tr", func(i int, h *colly.HTMLElement) {
					season, ok := getPlayerSeason(h, player.Id)

					if ok {
						player.Seasons = append(player.Seasons, season)
					}
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

func getPlayerSeason(e *colly.HTMLElement, playerID string) (models.PlayerSeason, bool) {
	season := models.PlayerSeason{}
	league := e.ChildText(statToAttr("lg_id") + " a")
	if league != "NBA" {
		return models.PlayerSeason{}, false
	}
	season.TeamID = e.ChildText(statToAttr("team_id") + " a")

	if season.TeamID == "" {
		return models.PlayerSeason{}, false
	}

	season.PlayerID = playerID
	season.Year = getYearFromID(e.Attr("id"))

	season.GamesPlayed = getIntStat(e.ChildText(statToAttr("g")))
	season.Age = getIntStat(e.ChildText(statToAttr("age")))
	mpg, err := getFloatStat(statToAttr("mp_per_g"))

	season.Position = e.ChildText(statToAttr("pos"))

	if err != nil {
		season.MinutesPlayed = &mpg
	}

	ppg, err := getFloatStat(e.ChildText(statToAttr("pts_per_g")))

	if err != nil {
		return models.PlayerSeason{}, false
	}
	season.PointsPerGame = &ppg

	rpg, err := getFloatStat(e.ChildText(statToAttr("trb_per_g")))
	if err == nil {
		season.ReboundsPerGame = &rpg
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
	return season, true
}
