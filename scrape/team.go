package scrape

import (
	"fmt"
	"log"

	"github.com/alex-305/basketball-ref-scraper/db"
	"github.com/alex-305/basketball-ref-scraper/models"

	"github.com/gocolly/colly/v2"
)

func GetAllTeams(site string) {
	c, _ := NewCollyCollector(site)
	db := db.Connect()

	defer db.Disconnect()

	c.OnHTML("table#teams_active tbody", func(e *colly.HTMLElement) {
		e.ForEach("tr th a", func(i int, a *colly.HTMLElement) {
			href := a.Attr("href")

			link := site + href

			team := models.Team{
				Id:   getTeamIDFromHref(href),
				Name: a.Text,
			}

			teamCollector := c.Clone()

			teamCollector.OnHTML("#"+team.Id+" > tbody", func(e *colly.HTMLElement) {
				e.ForEach("tr", func(i int, h *colly.HTMLElement) {
					season, ok := getTeamSeason(h, team.Id)

					if ok {
						team.Seasons = append(team.Seasons, season)
					}
				})
				err := db.InsertTeam(team)
				if err != nil {
					log.Fatal(err)
				}
			})
			teamCollector.OnError(func(r *colly.Response, err error) {
				fmt.Printf("Error:%s", err)
			})
			teamCollector.Visit(link)
		})
	})

	c.OnError(func(r *colly.Response, err error) {
		fmt.Printf("Error:%s", err)
	})
	c.Visit(site + "/teams")
}

func getTeamSeason(e *colly.HTMLElement, teamID string) (models.TeamSeason, bool) {
	season := models.TeamSeason{
		TeamID: teamID,
	}
	yearHref := e.Attr(statToAttr("season") + " a href")
	season.Year = getYearFromHref(yearHref)

	season.Wins = getIntStat(e.Attr(statToAttr("wins")))
	season.Losses = getIntStat(e.Attr(statToAttr("losses")))

	return season, true
}
