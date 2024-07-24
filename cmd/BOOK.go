/*
Copyright © 2024 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

// BOOKCmd represents the BOOK command
var BOOKCmd = &cobra.Command{
	Use:   "BOOK",
	Short: "Book a seat for the flight",
	Long:  `This command can be used to book a seat on a flight for customers.`,
	Run: func(cmd *cobra.Command, args []string) {
		var seat = args[0]
		// var num = args[1]

		records, err := readData()
		if err != nil {
			fmt.Println("FAIL")
			fmt.Println(err)
			return
		}

		row, col := seatToIndex(seat)

		if records[row][col] == "x" {
			fmt.Println("FAIL")
			return
		}

		records[row][col] = "x"

		err = writeData(records)
		if err != nil {
			fmt.Println("FAIL")
			fmt.Println(err)
			return
		}

		fmt.Println("SUCCESS")
	},
}

func init() {
	rootCmd.AddCommand(BOOKCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// BOOKCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// BOOKCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
