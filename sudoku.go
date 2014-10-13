package gosudoku

import (
	"fmt"
	"math"
)

type Sudoku struct {
	Numbers [9][9]int
}

// Prints sudoku to console
func (s *Sudoku) Print() {
	for i := 0; i < 9; i++ {
		for j := 0; j < 9; j++ {
			// Printing group of 3 numbers
			fmt.Print(s.Numbers[i][j], " ")
			// Inserting space after each 3 columns
			if (j+1)%3 == 0 {
				fmt.Print(" ")
			}
		}
		// Inserting line break after each 3 rows
		fmt.Println()
		if (i+1)%3 == 0 {
			fmt.Println()
		}
	}
}

// Adds number to the Sudoku's cell
func (s *Sudoku) AddNumber(rowIndex int, colIndex int, number int) {
	if areIndexesCorrect(rowIndex, colIndex) && isNumberCorrect(number) {
		s.Numbers[rowIndex][colIndex] = number
	}
}

// Resolves Sudoku
func (s *Sudoku) Resolve() bool {
	// Do next step if game is not over
	if s.getUnresolvedCellAmount() > 0 {
		count, cell := s.getMostNormalizedCell()

		index := 0
		result := false
		values := s.getPossibleNumbers(cell)

		// Looping until putting correct number
		for !result {
			// Something went wrong. Cleaning cell and returning false
			if index == count {
				s.Numbers[cell.rowIndex][cell.colIndex] = 0
				return false
			}
			number := values[index]
			index++
			s.Numbers[cell.rowIndex][cell.colIndex] = number

			// Trying to do next step
			result = s.Resolve()
		}
	}
	return true
}

// Returns amount of cells equal to zero
func (s *Sudoku) getUnresolvedCellAmount() int {
	result := 0
	for i := 0; i < 9; i++ {
		for j := 0; j < 9; j++ {
			if s.Numbers[i][j] == 0 {
				result++
			}
		}
	}
	return result
}

// Gets column by it's index
func (s *Sudoku) getColumn(index int) line {
	column := line{}
	current := 0
	for i := 0; i < 9; i++ {
		column.numbers[current] = s.Numbers[i][index]
		current++
	}
	return column
}

// Gets row by it's index
func (s *Sudoku) getRow(index int) line {
	row := line{}
	row.numbers = s.Numbers[index]
	return row
}

// Gets square by it's index
func (s *Sudoku) getSquare(index int) square {
	sqr := square{}

	// Row and Column indexes on sudoku field to start copying from
	rowIndex := int(math.Floor(float64(index)/3.0)) * 3
	colIndex := (index - rowIndex) * 3

	x := 0
	y := 0

	// Copying values from sudoku field to square
	for i := rowIndex; i < rowIndex+3; i++ {
		for j := colIndex; j < colIndex+3; j++ {
			sqr.numbers[y][x] = s.Numbers[i][j]
			x++
		}
		y++
		x = 0
	}
	return sqr
}

// Returns true if number can be put inside the cell and be still unique
// inside the row / column / square
func (s *Sudoku) canBePut(rowIndex int, colIndex int, number int) bool {
	if s.Numbers[rowIndex][colIndex] != 0 {
		return false
	}
	// Creating a copy of sudoku
	tempSudoku := Sudoku{Numbers: s.Numbers}
	tempSudoku.AddNumber(rowIndex, colIndex, number)

	// Getting row, column and square that will be checked
	row := tempSudoku.getRow(rowIndex)
	col := tempSudoku.getColumn(colIndex)
	sqrIndex := getSquareIndex(rowIndex, colIndex)
	sqr := tempSudoku.getSquare(sqrIndex)

	// Checking row, column and square of the sudoku's copy
	if row.isUnique() && col.isUnique() && sqr.isUnique() {
		return true
	}
	return false
}

// Returns cell with minimal possible number amount to put
func (s *Sudoku) getMostNormalizedCell() (int, cell) {
	minPossibilities := 10
	rowIndex := -1
	colIndex := -1

	// Finding the cell with minimal possible number amount to put
	for i := 0; i < 9; i++ {
		for j := 0; j < 9; j++ {

			// Only if cell is currently empty
			if s.Numbers[i][j] == 0 {
				amount := s.getPossibleNumberAmount(i, j)

				// If new minimal amount found, replacing result and indexes
				if amount < minPossibilities {
					minPossibilities = amount
					rowIndex = i
					colIndex = j
				}
			}
		}
	}
	return minPossibilities, cell{rowIndex, colIndex}
}

// Returns possible amount of numbers for the cell with indexes
func (s *Sudoku) getPossibleNumberAmount(rowIndex int, colIndex int) int {
	count := 0

	// How much numbers can be put inside the cell
	for i := 1; i <= 9; i++ {
		if s.canBePut(rowIndex, colIndex, i) {
			count++
		}
	}
	return count
}

// Returns list of possible numbers for the cell
func (s *Sudoku) getPossibleNumbers(c cell) []int {
	possibleNumbers := []int{}

	// The list of numbers can be put inside the cell currently
	for i := 1; i <= 9; i++ {
		if s.canBePut(c.rowIndex, c.colIndex, i) {
			possibleNumbers = append(possibleNumbers, i)
		}
	}
	return possibleNumbers
}

// Column and row indexes must be in diaposone 0..8
func areIndexesCorrect(colIndex int, rowIndex int) bool {
	if colIndex < 0 || colIndex > 8 || rowIndex < 0 || rowIndex > 8 {
		return false
	}
	return true
}

// The number must be in diaposone 1..9
func isNumberCorrect(number int) bool {
	if number < 1 || number > 9 {
		return false
	}
	return true
}

// Gets square index by cell indexes
func getSquareIndex(rowIndex int, colIndex int) (sqrIndex int) {
	return 3*int((math.Floor(float64(rowIndex)/3.0))) + int((math.Floor(float64(colIndex) / 3.0)))
}
