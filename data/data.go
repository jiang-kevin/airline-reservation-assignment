/*
Copyright © 2024 Kevin Jiang <kevin@kevinjiang.dev>
*/
package data

import (
	"encoding/csv"
	"errors"
	"os"

	"github.com/spf13/viper"
)

type Data interface {
	ReadData() (data [][]string, err error)
	WriteData(records [][]string) error
	SeatToIndex(seat string) (row int, col int, err error)
}

type AraData struct{}

func (d AraData) ReadData() (data [][]string, err error) {

	filepath := viper.GetString("data_filename")

	file, err := os.Open(filepath)

	if err != nil {
		return
	}
	defer file.Close()

	csvReader := csv.NewReader(file)
	data, err = csvReader.ReadAll()
	if err != nil {
		return
	}

	return
}

func (d AraData) WriteData(records [][]string) error {

	filepath := viper.GetString("data_filename")
	file, err := os.Create(filepath)
	if err != nil {
		return err
	}

	defer file.Close()

	csvWriter := csv.NewWriter(file)
	defer csvWriter.Flush()
	err = csvWriter.WriteAll(records)
	if err != nil {
		return err
	}

	return nil
}

// Converting characters to row/column numbers using ASCII values - 'A' = 65, '0' = 48
func (d AraData) SeatToIndex(seat string) (row int, col int, err error) {
	if len(seat) > 2 {
		row = -1
		col = -1
		err = errors.New("seat argument is too long")
		return
	}

	row = int(seat[0] - 65)
	col = int(seat[1] - 48)

	if row < 0 || row > 19 {
		err = errors.New("seat row must be between A and T")
		return
	}

	if col < 0 || col > 7 {
		err = errors.New("seat column must be between 0 and 7")
		return
	}

	return
}
