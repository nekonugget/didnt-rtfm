package main

import (
	"github.com/nekonugget/didnt-rtfm/cmd"
	generic "github.com/nekonugget/didnt-rtfm/platforms"
)

func main() {
	generic.CrawlTest()
	cmd.Execute()

}
