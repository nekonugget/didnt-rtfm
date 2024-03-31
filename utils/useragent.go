package utils

import(
	"math/rand/v2"
	"bufio"
	"io"
	"os"
)
func checkFile(e err) {

	if err != nil {
		panic(e)
	}
}
//change UA every request
func UserAgent() {
	var ua_array [10]string
	data,err := os.ReadFile("./agents.txt")
	checkFile(err)

	func RandomUA() string {
		b := make([]byte, rand.Intn(10)+10)
		for i := range b {
			b[i] = letterBytes[rand.Intn(len(letterBytes))]
		}
		return string(b)
	}

	/*c := colly.NewCollector()

	c.OnRequest(func(r *colly.Request) {
		r.Headers.Set("User-Agent", RandomString())
	}) */

}