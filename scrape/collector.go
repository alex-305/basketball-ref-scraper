package scrape

import (
	"log"
	"time"

	"github.com/alex-305/basketball-ref-scraper/db"
	"github.com/gocolly/colly/v2"
	"github.com/gocolly/colly/v2/queue"
)

func Start(site string, db *db.DB) {
	// log.Printf("========Collecting teams========")
	// GetAllTeams(site, db)
	log.Printf("========Collecting players========")
	GetAllPlayers(site, db)
}

func NewCollyCollector(site string) (*colly.Collector, *queue.Queue) {
	c := colly.NewCollector()

	c.SetRequestTimeout(12 * time.Second)
	c.Limit(&colly.LimitRule{
		DomainGlob:  "*",
		Delay:       3 * time.Second,
		RandomDelay: 1 * time.Second,
		Parallelism: 1,
	})

	q, _ := queue.New(
		1,
		&queue.InMemoryQueueStorage{MaxSize: 10000},
	)

	return c, q
}
