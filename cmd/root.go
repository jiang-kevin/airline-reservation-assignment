/*
Copyright © 2024 Kevin Jiang <kevin@kevinjiang.dev>
*/
package cmd

import (
	"os"

	"github.com/spf13/cobra"
)

// rootCmd represents the base command when called without any subcommands
var rootCmd = &cobra.Command{
	Use:   "ara",
	Short: "Book/cancel reservations for seats on flights",
	Long: `This CLI can be used to book or cancel seats on flights. There are two actions: BOOK and CANCEL. To see the relevant documentation for each command, use the -h flag like so:

ara BOOK -h`,
}

func Execute() {
	err := rootCmd.Execute()
	if err != nil {
		os.Exit(1)
	}
}

func init() {
	rootCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
