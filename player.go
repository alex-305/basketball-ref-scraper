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
