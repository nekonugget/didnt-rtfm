// generic crawler
package generic

import (
	"log"

	"github.com/gocolly/colly"
	"github.com/nekonugget/didnt-rtfm/utils"
)

func CrawlTest(url string) {
	log.Println("Hello world")
	purl, err := utils.CheckUrl(url)
	if err != nil {
		return
	}
	log.Println(purl)
	c := colly.NewCollector()
	c.Visit(url)
}
