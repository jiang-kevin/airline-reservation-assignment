/*
Copyright © 2024 Kevin Jiang <kevin@kevinjiang.dev>
*/
package cmd

import (
	"fmt"
	"strconv"

	"github.com/jiang-kevin/airline-reservation-assignment/data"

	"github.com/spf13/cobra"
)

// CANCELCmd represents the CANCEL command
func CancelCmd(data data.Data) *cobra.Command {
	return &cobra.Command{
		Use:   "CANCEL <starting seat> <number of consecutive seats>",
		Short: "Cancel seats for the flight",
		Long: `This command can be used to cancel one or more seats on a flight for customers.
		
	Example: ara BOOK A0 2
	
	Note: The command will attempt to book seats consecutively counting up from the starting seat, up to the specified number. It will fail if it exceeds the maximum range of seats.
		`,
		Run: func(cmd *cobra.Command, args []string) {
			fmt.Println(CancelSeat(args, data))
		},
	}
}

func CancelSeat(args []string, data data.Data) string {
	if len(args) < 2 {
		return "FAIL"
	}

	seat := args[0]
	num, err := strconv.Atoi(args[1])
	if err != nil {
		return "FAIL"
	}

	records, err := data.ReadData()
	if err != nil {
		return "FAIL"
	}

	row, col, err := data.SeatToIndex(seat)
	if err != nil {
		return "FAIL"
	}

	if records[row][col] == "o" {
		return "FAIL"
	} else {
		records[row][col] = "o"
	}

	if num > 1 {
		for i := col + 1; i < col+num; i++ {
			if i < len(records[row]) && records[row][i] != "o" {
				records[row][i] = "o"
			} else {
				return "FAIL"
			}
		}
	}

	err = data.WriteData(records)
	if err != nil {
		return "FAIL"
	}

	return "SUCCESS"
}
