package funandgames

import (
	"errors"
	"reflect"
)

const Rows = 9
const Cols = 9
const SubGrid = 3
const Empty = '.'

type Sudoku struct {
	grid [][]byte
}

func (s *Sudoku) IsValid() bool {
	if len(s.grid) != Rows || len(s.grid[0]) != Cols {
		return false
	}

	// Row, column, and sub-grid validation arrays
	rowSeen := make([][]bool, Rows)
	colSeen := make([][]bool, Cols)
	subGridSeen := make([][]bool, SubGrid*SubGrid)

	// initialise arrays
	for i := 0; i < Rows; i++ {
		rowSeen[i] = make([]bool, 9)
		colSeen[i] = make([]bool, 9)
		subGridSeen[i] = make([]bool, 9)
	}

	for rowIndex := 0; rowIndex < Rows; rowIndex++ {
		for colIndex := 0; colIndex < Cols; colIndex++ {
			cell := s.grid[rowIndex][colIndex]

			if isEmpty(cell) {
				continue
			} else if !isDigit(cell) {
				return false
			}

			digitIndex := cell - '1' // the position in the array for the cell's value
			subgridIndex := subGridIndex(rowIndex, colIndex)

			// Check row, column, and sub-grid
			if rowSeen[rowIndex][digitIndex] || colSeen[colIndex][digitIndex] || subGridSeen[subgridIndex][digitIndex] {
				return false
			}

			rowSeen[rowIndex][digitIndex] = true
			colSeen[colIndex][digitIndex] = true
			subGridSeen[subgridIndex][digitIndex] = true
		}
	}

	return true
}

func (s *Sudoku) Solve() (*Sudoku, error) {

	if !s.IsValid() {
		return nil, errors.New("cannot solve a non-valid Sudoku")
	}

	row, col := findNextEmptyCell(s.grid)
	if row < 0 || col < 0 {
		return s, nil // no more empty cells + valid
	}

	for i := '1'; i <= '9'; i++ {

		s.grid[row][col] = byte(i)
		if s, _ := s.Solve(); s != nil {
			return s, nil
		}

		s.grid[row][col] = byte(Empty)

	}

	return nil, nil

}

func findNextEmptyCell(grid [][]byte) (int, int) {
	for rowIndex := 0; rowIndex < Rows; rowIndex++ {
		for colIndex := 0; colIndex < Cols; colIndex++ {
			cell := grid[rowIndex][colIndex]
			if isEmpty(cell) {
				return rowIndex, colIndex
			}
		}
	}
	return -1, -1
}

// Calculates the index of the 3x3 sub-grid
func subGridIndex(row, col int) int {
	return (row/SubGrid)*SubGrid + (col / SubGrid)
}

func isDigit(b byte) bool {
	return b >= '1' && b <= '9'
}

func isEmpty(b byte) bool {
	return b == Empty
}

func (s *Sudoku) Equals(sudoku *Sudoku) bool {
	return reflect.DeepEqual(s.grid, sudoku.grid)
}

func NewSudoku(grid [][]byte) *Sudoku {
	return &Sudoku{
		grid: grid,
	}
}
