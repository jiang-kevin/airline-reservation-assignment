/*
Copyright © 2024 Kevin Jiang <kevin@kevinjiang.dev>
*/
package cmd

import (
	"fmt"
	"strconv"

	"github.com/jiang-kevin/airline-reservation-assignment/util"

	"github.com/spf13/cobra"
)

// CANCELCmd represents the CANCEL command
var CANCELCmd = &cobra.Command{
	Use:   "CANCEL <starting seat> <number of consecutive seats>",
	Short: "Cancel seats for the flight",
	Long: `This command can be used to cancel one or more seats on a flight for customers.
	
Example: ara BOOK A0 2

Note: The command will attempt to book seats consecutively counting up from the starting seat, up to the specified number. It will fail if it exceeds the maximum range of seats.
	`,
	Run: func(cmd *cobra.Command, args []string) {
		if len(args) < 2 {
			fmt.Println("FAIL")
			return
		}

		seat := args[0]
		num, err := strconv.Atoi(args[1])
		if err != nil {
			fmt.Println("FAIL")
			return
		}

		records, err := util.ReadData()
		if err != nil {
			fmt.Println("FAIL")
			fmt.Println(err)
			return
		}

		row, col, err := util.SeatToIndex(seat)
		if err != nil {
			fmt.Println("FAIL")
			return
		}

		if records[row][col] == "o" {
			fmt.Println("FAIL")
			return
		} else {
			records[row][col] = "o"
		}

		if num > 1 {
			for i := col + 1; i < col+num; i++ {
				if i < len(records[row]) && records[row][i] != "o" {
					records[row][i] = "o"
				} else {
					fmt.Println("FAIL")
					return
				}
			}
		}

		err = util.WriteData(records)
		if err != nil {
			fmt.Println("FAIL")
			fmt.Println(err)
			return
		}

		fmt.Println("SUCCESS")
	},
}

func init() {
	rootCmd.AddCommand(CANCELCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// CANCELCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// CANCELCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
