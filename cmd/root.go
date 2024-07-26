/*
Copyright © 2024 Kevin Jiang <kevin@kevinjiang.dev>
*/
package cmd

import (
	"github.com/spf13/cobra"
)

// rootCmd represents the base command when called without any subcommands
func RootCmd() *cobra.Command {
	return &cobra.Command{
		Use:   "ara",
		Short: "Book/cancel reservations for seats on flights",
		Long: `This CLI can be used to book or cancel seats on flights. There are two actions: BOOK and CANCEL. To see the relevant documentation for each command, use the -h flag like so:
	
	ara BOOK -h`,
	}
}

func Execute(cmd *cobra.Command) error {
	cmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
	return cmd.Execute()
}
