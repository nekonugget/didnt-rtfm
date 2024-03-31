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
	c.OnHTML("a[href]", func(e *colly.HTMLElement) {
		e.Request.Visit(e.Attr("href"))
	})
	c.OnHTML("span", func(e *colly.HTMLElement) {
		log.Println(e.Text)
	})
	c.OnRequest(func(r *colly.Request) {
		log.Println("Visiting: ", r.URL)
	})
	c.Visit(url)

	c.OnError(func(_ *colly.Response, err error) {
		log.Println("Something went wrong:", err)
	})
}

func ApiUrls() {

}
