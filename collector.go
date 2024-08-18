package main

import (
	"time"

	"github.com/gocolly/colly/v2"
)

func NewCollyCollector(site string) *colly.Collector {
	c := colly.NewCollector()

	c.SetRequestTimeout(10 * time.Second)
	c.Limit(&colly.LimitRule{
		DomainGlob:  site,
		Delay:       6 * time.Second,
		RandomDelay: 1 * time.Second,
		Parallelism: 1,
	})
	return c
}
