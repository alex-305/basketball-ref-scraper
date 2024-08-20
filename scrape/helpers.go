package scrape

import (
	"strconv"
	"strings"
)

func getTeamIDFromHref(href string) string {
	parts := strings.Split(href, "/")
	return parts[len(parts)-2]
}

func getYearFromID(id string) string {
	parts := strings.Split(id, ".")
	dateString := parts[len(parts)-1]

	return dateString
}

func getYearFromHref(href string) string {
	parts := strings.Split(href, "/")
	last := parts[len(parts)-1]

	return strings.TrimSuffix(last, ".html")
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
