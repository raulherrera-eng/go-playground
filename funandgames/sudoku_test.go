package funandgames

import (
	"testing"
)

func TestSudoku_IsValid(t *testing.T) {

	tests := []struct {
		grid  [][]byte
		valid bool
	}{
		{
			valid: true,
			grid: [][]byte{
				{'5', '3', '.', '.', '7', '.', '.', '.', '.'},
				{'6', '.', '.', '1', '9', '5', '.', '.', '.'},
				{'.', '9', '8', '.', '.', '.', '.', '6', '.'},
				{'8', '.', '.', '.', '6', '.', '.', '.', '3'},
				{'4', '.', '.', '8', '.', '3', '.', '.', '1'},
				{'7', '.', '.', '.', '2', '.', '.', '.', '6'},
				{'.', '6', '.', '.', '.', '.', '2', '8', '.'},
				{'.', '.', '.', '4', '1', '9', '.', '.', '5'},
				{'.', '.', '.', '.', '8', '.', '.', '7', '9'},
			},
		},
		{
			valid: false,
			grid: [][]byte{
				{'5', '3', '.', '.', '7', '.', '.', '.', '.'},
				{'6', '.', '3', '1', '9', '5', '.', '.', '.'}, // error here
				{'.', '9', '8', '.', '.', '.', '.', '6', '.'},
				{'8', '.', '.', '.', '6', '.', '.', '.', '3'},
				{'4', '.', '.', '8', '.', '3', '.', '.', '1'},
				{'7', '.', '.', '.', '2', '.', '.', '.', '6'},
				{'.', '6', '.', '.', '.', '.', '2', '8', '.'},
				{'.', '.', '.', '4', '1', '9', '.', '.', '5'},
				{'.', '.', '.', '.', '8', '.', '.', '7', '9'},
			},
		},
	}

	for _, test := range tests {

		// Test logic here
		sudoku := NewSudoku(test.grid)
		actual := sudoku.IsValid()
		if actual != test.valid {
			t.Error("Expected", test, "got", actual)
		}

	}

}

func TestSudoku_Solve(t *testing.T) {

	tests := []struct {
		grid  [][]byte
		valid bool
	}{
		{
			valid: true,
			grid: [][]byte{
				{'5', '3', '.', '.', '7', '.', '.', '.', '.'},
				{'6', '.', '.', '1', '9', '5', '.', '.', '.'},
				{'.', '9', '8', '.', '.', '.', '.', '6', '.'},
				{'8', '.', '.', '.', '6', '.', '.', '.', '3'},
				{'4', '.', '.', '8', '.', '3', '.', '.', '1'},
				{'7', '.', '.', '.', '2', '.', '.', '.', '6'},
				{'.', '6', '.', '.', '.', '.', '2', '8', '.'},
				{'.', '.', '.', '4', '1', '9', '.', '.', '5'},
				{'.', '.', '.', '.', '8', '.', '.', '7', '9'},
			},
		},
	}

	for _, test := range tests {

		// Test logic here
		sudoku := NewSudoku(test.grid)
		solved, err := sudoku.Solve()
		if test.valid && err != nil {
			t.Error("failed to solve valid Sudoku with error")
		} else if test.valid && solved == nil {
			t.Error("failed to solve valid Sudoku without error")
		} else if !test.valid && (err == nil || solved != nil) {
			t.Error("solved a non-valid Sudoku")
		} else if test.valid && solved != nil {
			// all good :)
		}

	}

}
