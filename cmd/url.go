package cmd

import (
	generic "github.com/nekonugget/didnt-rtfm/platforms"
	"github.com/spf13/cobra"
)

func init() {
	rootCmd.AddCommand(urlCmd)
}

var urlCmd = &cobra.Command{
	Use:   "url",
	Short: "site to crawl",
	Long:  `target page to crawl`,
	Run: func(cmd *cobra.Command, args []string) (parsedUrl *URL) {
		urlInput, e := generic.CheckUrl(args[0])
		if e != nil {
			panic(e)
		}
		return urlInput
	},
}
