/*
 */
package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

var platformCmd = &cobra.Command{
	Use:   "platform",
	Short: "Have user specify the platform rather than guessing",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("platform called")
	},
}

func init() {
	rootCmd.AddCommand(platformCmd)
}
