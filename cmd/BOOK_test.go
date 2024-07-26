package cmd

import (
	"errors"
	"testing"
)

type mockdata struct {
	inputData [][]string
}

func (md mockdata) ReadData() ([][]string, error) {
	return md.inputData, nil
}

func (md mockdata) WriteData(data [][]string) error {
	return nil
}

func (md mockdata) SeatToIndex(seat string) (row int, col int, err error) {
	row = 0
	col = 0
	err = nil
	return
}

type mockdataReadError struct {
}

func (m mockdataReadError) ReadData() ([][]string, error) {
	return nil, errors.New("could not read file")
}

func (m mockdataReadError) WriteData(data [][]string) error {
	return nil
}

func (m mockdataReadError) SeatToIndex(seat string) (row int, col int, err error) {
	row = 0
	col = 0
	err = nil
	return
}

func TestBook(t *testing.T) {

	data := [][]string{
		{"o", "o", "o", "o", "o", "o", "o", "o"},
		{"o", "o", "o", "o", "o", "o", "o", "o"},
		{"o", "o", "o", "o", "o", "o", "o", "o"},
		{"o", "o", "o", "o", "o", "o", "o", "o"},
		{"o", "o", "o", "o", "o", "o", "o", "o"},
		{"o", "o", "o", "o", "o", "o", "o", "o"},
		{"o", "o", "o", "o", "o", "o", "o", "o"},
		{"o", "o", "o", "o", "o", "o", "o", "o"},
		{"o", "o", "o", "o", "o", "o", "o", "o"},
		{"o", "o", "o", "o", "o", "o", "o", "o"},
		{"o", "o", "o", "o", "o", "o", "o", "o"},
		{"o", "o", "o", "o", "o", "o", "o", "o"},
		{"o", "o", "o", "o", "o", "o", "o", "o"},
		{"o", "o", "o", "o", "o", "o", "o", "o"},
		{"o", "o", "o", "o", "o", "o", "o", "o"},
		{"o", "o", "o", "o", "o", "o", "o", "o"},
		{"o", "o", "o", "o", "o", "o", "o", "o"},
		{"o", "o", "o", "o", "o", "o", "o", "o"},
		{"o", "o", "o", "o", "o", "o", "o", "o"},
		{"o", "o", "o", "o", "o", "o", "o", "o"},
	}

	mockdata := mockdata{inputData: data}

	output := BookSeat([]string{"A0", "4"}, mockdata)

	if output != "SUCCESS" {
		t.Errorf("expected: SUCCESS, actual: %v", output)
	}
}

func TestBookNotEnoughArgs(t *testing.T) {

	mockdata := mockdata{}

	output := BookSeat([]string{"A0"}, mockdata)

	if output != "FAIL" {
		t.Errorf("expected: FAIL, actual: %v", output)
	}
}

func TestBookInvalidArg2(t *testing.T) {

	mockdata := mockdata{}

	output := BookSeat([]string{"A0", "f"}, mockdata)

	if output != "FAIL" {
		t.Errorf("expected: FAIL, actual: %v", output)
	}
}

func TestBookReadError(t *testing.T) {

	mockdata := mockdataReadError{}

	output := BookSeat([]string{"A0", "1"}, mockdata)

	if output != "FAIL" {
		t.Errorf("expected: FAIL, actual: %v", output)
	}
}

func TestBookAlreadyBooked(t *testing.T) {

	mockdata := mockdata{inputData: [][]string{{"x"}}}

	output := BookSeat([]string{"A0", "1"}, mockdata)

	if output != "FAIL" {
		t.Errorf("expected: FAIL, actual: %v", output)
	}
}

func TestBookConsecutiveInvalid(t *testing.T) {

	mockdata := mockdata{inputData: [][]string{{"o", "o", "o", "x"}}}

	output := BookSeat([]string{"A0", "4"}, mockdata)

	if output != "FAIL" {
		t.Errorf("expected: FAIL, actual: %v", output)
	}
}
