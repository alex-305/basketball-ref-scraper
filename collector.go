package main

import (
	"time"

	"github.com/gocolly/colly/v2"
	"github.com/gocolly/colly/v2/queue"
)

func NewCollyCollector(site string) (*colly.Collector, *queue.Queue) {
	c := colly.NewCollector()

	c.SetRequestTimeout(10 * time.Second)
	c.Limit(&colly.LimitRule{
		DomainGlob:  "*",
		Delay:       6 * time.Second,
		RandomDelay: 1 * time.Second,
		Parallelism: 1,
	})

	q, _ := queue.New(
		1,
		&queue.InMemoryQueueStorage{MaxSize: 10000},
	)

	return c, q
}
