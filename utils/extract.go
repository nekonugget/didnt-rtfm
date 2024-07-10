package utils

type Searcher struct {
	urlSearch  string //[]URLsearch
	credSearch string //[]CredSearch
}

// github.com/smacker/go-tree-sitter/tree/master/html
// will move this to a go test ... laaater. zz
func testExtract() *Searcher {
	return &Searcher{
		urlSearch:  "idk",
		credSearch: "idk",
	}
}

func ExtractURLs() {
	// Util using Colly to get URLs.

}
func ExtractJSON() {
	// Util parsing JSON data from site or file.
}
func ExtractPDF() {

}
