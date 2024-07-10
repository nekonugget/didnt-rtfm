package utils

import (
	"os"
)

func checkFile(e err) {

	if err != nil {
		panic(e)
	}
}

//change UA every request

func RandomUA() string {
	var ua_array [10]string
	data, err := os.ReadFile("./agents.txt")
	checkFile(err)

}

/*c := colly.NewCollector()

c.OnRequest(func(r *colly.Request) {
	r.Headers.Set("User-Agent", RandomString())
}) */
