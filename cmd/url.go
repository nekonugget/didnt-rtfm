package cmd

import (
	"errors"

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
	Args: func(cmd *cobra.Command, args []string) error {
		if len(args) < 1 {
			return errors.New("url required")
		}
		return nil
	},
	Run: func(cmd *cobra.Command, args []string) {
		url := args[0]
		generic.CrawlTest(url)

	},
}
