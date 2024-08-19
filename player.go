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

func getSeason(e *colly.HTMLElement, playerID string) models.Season {
	season := models.Season{}
	season.TeamID = e.ChildText("td[data-stat=team_id] a")
	season.PlayerID = playerID
	season.Year = getYear(e.Attr("id"))

	season.GamesPlayed = getIntStat(e.ChildText("td[data-stat=g]"))
	season.Age = getIntStat(e.ChildText("td[data-stat=age]"))
	mpg, err := getFloatStat("td[data-stat=mp_per_g]")

	if err != nil {
		season.MinutesPlayed = &mpg
	}

	rpg, err := getFloatStat(e.ChildText("td[data-stat=trb_per_g]"))
	if err == nil {
		season.ReboundsPerGame = &rpg
	}
	ppg, err := getFloatStat(e.ChildText("td[data-stat=pts_per_g]"))
	if err == nil {
		season.PointsPerGame = &ppg
	}
	apg, err := getFloatStat(e.ChildText("td[data-stat=ast_per_g]"))
	if err == nil {
		season.AssistsPerGame = &apg
	}
	bpg, err := getFloatStat(e.ChildText("td[data-stat=blk_per_g]"))
	if err == nil {
		season.BlocksPerGame = &bpg
	}
	spg, err := getFloatStat(e.ChildText("td[data-stat=stl_per_g]"))
	if err == nil {
		season.StealsPerGame = &spg
	}
	return season
}
