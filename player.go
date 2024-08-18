package main

import (
	"strconv"
	"strings"

	"scrape/models"

	"github.com/gocolly/colly/v2"
)

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

func getStat(str string) (float32, error) {
	stat, err := strconv.ParseFloat(str, 32)

	if err != nil {
		return 0.0, err
	}
	return float32(stat), nil
}

func getSeason(e *colly.HTMLElement, playerID string) models.Season {
	season := models.Season{}
	season.TeamID = e.ChildText("td[data-stat=team_id] a")
	season.PlayerID = playerID
	season.Year = getYear(e.Attr("id"))

	rpg, err := getStat(e.ChildText("td[data-stat=trb_per_g]"))
	if err == nil {
		season.ReboundsPerGame = &rpg
	}
	ppg, err := getStat(e.ChildText("td[data-stat=pts_per_g]"))
	if err == nil {
		season.PointsPerGame = &ppg
	}
	apg, err := getStat(e.ChildText("td[data-stat=ast_per_g]"))
	if err == nil {
		season.AssistsPerGame = &apg
	}
	bpg, err := getStat(e.ChildText("td[data-stat=blk_per_g]"))
	if err == nil {
		season.BlocksPerGame = &bpg
	}
	spg, err := getStat(e.ChildText("td[data-stat=stl_per_g]"))
	if err == nil {
		season.StealsPerGame = &spg
	}
	return season
}

func GetPlayer(site, href, name string) models.Player {

	c := NewCollyCollector(site)

	link := site + href

	player := models.Player{
		Id:   getID(link),
		Name: name,
	}

	c.OnHTML("#per_game > tbody", func(e *colly.HTMLElement) {
		e.ForEach("tr", func(i int, h *colly.HTMLElement) {
			player.Seasons = append(player.Seasons, getSeason(h, player.Id))
		})
	})
	c.Visit(link)

	return player
}
