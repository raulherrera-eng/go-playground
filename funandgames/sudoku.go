package funandgames

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

			if cell == Empty {
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

// Calculates the index of the 3x3 sub-grid
func subGridIndex(row, col int) int {
	return (row/SubGrid)*SubGrid + (col / SubGrid)
}

func isDigit(b byte) bool {
	return b >= '1' && b <= '9'
}

func NewSudoku(grid [][]byte) *Sudoku {
	return &Sudoku{
		grid: grid,
	}
}
