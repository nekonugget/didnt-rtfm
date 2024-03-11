// generic crawler
package generic

import (
	"log"

	"github.com/gocolly/colly"
)

func CrawlTest() {
	log.Println("Hello world")
	c := colly.NewCollector()
	c.Visit("example.com")
}
