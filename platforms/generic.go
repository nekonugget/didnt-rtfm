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
func CheckUrl(string urlInput) (err, url *URL) {
	parsed := url.Parse(userInput)
	return parsed
}
