package data

import (
	"errors"
	"testing"
)

func TestSeatToIndex(t *testing.T) {
	tests := []struct {
		name        string
		seat        string
		expectedRow int
		expectedCol int
		gotError    error
	}{
		{
			name:        "valid seat maps to correct row/col",
			seat:        "A0",
			expectedRow: 0,
			expectedCol: 0,
			gotError:    nil,
		},
		{
			name:        "valid seat at upper limit maps to correct row/col",
			seat:        "T7",
			expectedRow: 19,
			expectedCol: 7,
			gotError:    nil,
		},
		{
			name:        "error because seat argument is too long",
			seat:        "ASD1234",
			expectedRow: -1,
			expectedCol: -1,
			gotError:    errors.New("seat argument is too long"),
		},
		{
			name:        "error because row is not in range",
			seat:        "X0",
			expectedRow: 23,
			expectedCol: 0,
			gotError:    errors.New("seat row must be between A and T"),
		},
		{
			name:        "error because col is not in range",
			seat:        "A9",
			expectedRow: 0,
			expectedCol: 8,
			gotError:    errors.New("seat column must be between 0 and 7"),
		},
	}

	var data Data = AraData{}

	for _, v := range tests {
		t.Run(v.name, func(t *testing.T) {
			row, col, err := data.SeatToIndex(v.seat)

			if v.gotError != nil {
				if err.Error() != v.gotError.Error() {
					t.Errorf(
						"expected an error: %v, but got: %v",
						v.gotError,
						err,
					)
				}

				return
			}

			if row != v.expectedRow || col != v.expectedCol {
				t.Fatalf("expected row=%v, col=%v; got row=%v, col=%v", v.expectedRow, v.expectedCol, row, col)
			}
		})
	}
}
