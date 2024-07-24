/*
Copyright © 2024 Kevin Jiang <kevin@kevinjiang.dev>
*/
package cmd

import (
	"encoding/csv"
	"errors"
	"os"
)

func readData() (data [][]string, err error) {

	filepath := "data.txt"
	file, err := os.Open(filepath)
	if err != nil {
		return
	}
	defer file.Close()

	reader := csv.NewReader(file)
	data, err = reader.ReadAll()
	if err != nil {
		return
	}

	return
}

func writeData(records [][]string) error {

	filepath := "data.txt"
	file, err := os.Create(filepath)
	if err != nil {
		return err
	}
	defer file.Close()

	writer := csv.NewWriter(file)
	defer writer.Flush()
	err = writer.WriteAll(records)
	if err != nil {
		return err
	}

	return nil
}

// Converting characters to row/column numbers using ASCII values - 'A' = 65, '0' = 48
func seatToIndex(seat string) (row int, col int, err error) {
	if len(seat) > 2 {
		err = errors.New("seat argument is too long")
		return
	}

	row = int(seat[0] - 65)
	col = int(seat[1] - 48)

	if row < 0 || row > 20 {
		err = errors.New("seat row must be between A and T")
		return
	}

	if col < 0 || col > 8 {
		err = errors.New("seat column must be between 0 and 8")
		return
	}

	return
}
