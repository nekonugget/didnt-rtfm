package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

var authCmd = &cobra.Command{
	Use:   "auth",
	Short: "Login deets",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("platform called")
	},
}

func init() {
	rootCmd.AddCommand(authCmd)
}
