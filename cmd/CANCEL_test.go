package cmd

import (
	"errors"
	"testing"
)

type mockdataCancel struct {
	inputData [][]string
}

func (md mockdataCancel) ReadData() ([][]string, error) {
	return md.inputData, nil
}

func (md mockdataCancel) WriteData(data [][]string) error {
	return nil
}

func (md mockdataCancel) SeatToIndex(seat string) (row int, col int, err error) {
	row = 0
	col = 0
	err = nil
	return
}

type mockdataCancelReadError struct {
}

func (m mockdataCancelReadError) ReadData() ([][]string, error) {
	return nil, errors.New("could not read file")
}

func (m mockdataCancelReadError) WriteData(data [][]string) error {
	return nil
}

func (m mockdataCancelReadError) SeatToIndex(seat string) (row int, col int, err error) {
	row = 0
	col = 0
	err = nil
	return
}

func TestCancel(t *testing.T) {

	data := [][]string{
		{"x", "x", "x", "x", "o", "o", "o", "o"},
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

	mockdataCancel := mockdataCancel{inputData: data}

	output := CancelSeat([]string{"A0", "4"}, mockdataCancel)

	if output != "SUCCESS" {
		t.Errorf("expected: SUCCESS, actual: %v", output)
	}
}

func TestCancelNotEnoughArgs(t *testing.T) {

	mockdataCancel := mockdataCancel{}

	output := CancelSeat([]string{"A0"}, mockdataCancel)

	if output != "FAIL" {
		t.Errorf("expected: FAIL, actual: %v", output)
	}
}

func TestCancelInvalidArg2(t *testing.T) {

	mockdataCancel := mockdataCancel{}

	output := CancelSeat([]string{"A0", "f"}, mockdataCancel)

	if output != "FAIL" {
		t.Errorf("expected: FAIL, actual: %v", output)
	}
}

func TestCancelReadError(t *testing.T) {

	mockdataCancel := mockdataCancelReadError{}

	output := CancelSeat([]string{"A0", "1"}, mockdataCancel)

	if output != "FAIL" {
		t.Errorf("expected: FAIL, actual: %v", output)
	}
}

func TestCancelAlreadyCanceled(t *testing.T) {

	mockdataCancel := mockdataCancel{inputData: [][]string{{"o"}}}

	output := CancelSeat([]string{"A0", "1"}, mockdataCancel)

	if output != "FAIL" {
		t.Errorf("expected: FAIL, actual: %v", output)
	}
}

func TestCancelConsecutiveInvalid(t *testing.T) {

	mockdataCancel := mockdataCancel{inputData: [][]string{{"x", "x", "o", "x"}}}

	output := CancelSeat([]string{"A0", "4"}, mockdataCancel)

	if output != "FAIL" {
		t.Errorf("expected: FAIL, actual: %v", output)
	}
}
