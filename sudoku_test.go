package sudoku

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_areIndexesCorrect_ShouldReturnTrueIfIndexesAreCorrect(t *testing.T) {
	assert.Equal(t, true, areIndexesCorrect(0, 0))
	assert.Equal(t, true, areIndexesCorrect(8, 8))
	assert.Equal(t, true, areIndexesCorrect(5, 5))
}

func Test_areIndexesCorrect_ShouldReturnFalseIfIndexesAreIncorrect(t *testing.T) {
	assert.Equal(t, false, areIndexesCorrect(-1, 0))
	assert.Equal(t, false, areIndexesCorrect(9, 0))
	assert.Equal(t, false, areIndexesCorrect(0, -1))
	assert.Equal(t, false, areIndexesCorrect(0, 9))
	assert.Equal(t, false, areIndexesCorrect(-1, -1))
	assert.Equal(t, false, areIndexesCorrect(9, 9))
}

func Test_isNumberCorrect_ShouldReturnTrueIfNumberIsCorrect(t *testing.T) {
	assert.Equal(t, true, isNumberCorrect(1))
	assert.Equal(t, true, isNumberCorrect(9))
	assert.Equal(t, true, isNumberCorrect(5))
}

func Test_isNumberCorrect_ShouldReturnFalseIfNumberIsIncorrect(t *testing.T) {
	assert.Equal(t, false, isNumberCorrect(0))
	assert.Equal(t, false, isNumberCorrect(10))
}

func Test_AddNumber_ShouldAddNumberToSudokuCellIfArgumentsAreCorrect(t *testing.T) {
	s := new(Sudoku)
	s.AddNumber(0, 0, 5)
	s.AddNumber(8, 8, 8)
	s.AddNumber(5, 5, 3)
	expected := [9][9]int{
		{5, 0, 0, 0, 0, 0, 0, 0, 0},
		{0, 0, 0, 0, 0, 0, 0, 0, 0},
		{0, 0, 0, 0, 0, 0, 0, 0, 0},
		{0, 0, 0, 0, 0, 0, 0, 0, 0},
		{0, 0, 0, 0, 0, 0, 0, 0, 0},
		{0, 0, 0, 0, 0, 3, 0, 0, 0},
		{0, 0, 0, 0, 0, 0, 0, 0, 0},
		{0, 0, 0, 0, 0, 0, 0, 0, 0},
		{0, 0, 0, 0, 0, 0, 0, 0, 8},
	}
	assert.Equal(t, expected, s.Numbers)
}

func Test_getColumn_ShouldReturnColumnFromSudokuField(t *testing.T) {
	s := new(Sudoku)
	s.AddNumber(0, 5, 1)
	s.AddNumber(3, 5, 2)
	s.AddNumber(8, 5, 3)
	column := s.getColumn(5)
	assert.Equal(t, line{numbers: [9]int{1, 0, 0, 2, 0, 0, 0, 0, 3}}, column)
}

func Test_getRow_ShouldReturnCorrectRowFromSudokuField(t *testing.T) {
	s := new(Sudoku)
	s.AddNumber(2, 0, 5)
	s.AddNumber(2, 5, 3)
	s.AddNumber(2, 8, 4)
	row := s.getRow(2)
	assert.Equal(t, line{numbers: [9]int{5, 0, 0, 0, 0, 3, 0, 0, 4}}, row)
}

func Test_getSquareIndex_ShouldReturnCorrectSquareIndex(t *testing.T) {
	assert.Equal(t, 0, getSquareIndex(1, 1))
	assert.Equal(t, 1, getSquareIndex(1, 5))
	assert.Equal(t, 4, getSquareIndex(5, 3))
	assert.Equal(t, 8, getSquareIndex(8, 8))
}

func Test_getSquare_ShouldReturnCorrectSquare(t *testing.T) {
	s := new(Sudoku)
	s.AddNumber(3, 6, 1)
	s.AddNumber(3, 8, 2)
	s.AddNumber(5, 7, 3)
	expected := square{
		numbers: [3][3]int{
			{1, 0, 2},
			{0, 0, 0},
			{0, 3, 0},
		},
	}
	assert.Equal(t, expected, s.getSquare(5))
}

func Test_canBePut_ShouldReturnTrueIfNumberCanBePutIntoTheCell(t *testing.T) {
	s := new(Sudoku)
	s.AddNumber(0, 0, 1)
	s.AddNumber(0, 1, 2)
	s.AddNumber(0, 2, 3)
	s.AddNumber(1, 0, 4)
	s.AddNumber(1, 2, 6)
	s.AddNumber(2, 0, 7)
	s.AddNumber(2, 1, 8)
	s.AddNumber(2, 2, 9)
	assert.Equal(t, true, s.canBePut(1, 1, 5))
}

func Test_canBePut_ShouldReturnFalseIfPotentialNumberIsNotUniqueInsideTheSquare(t *testing.T) {
	s := new(Sudoku)
	s.AddNumber(0, 0, 1)
	assert.Equal(t, false, s.canBePut(2, 2, 1))
}

func Test_canBePut_ShouldReturnFalseIfPotentialNumberIsNotUniqueInsideTheRow(t *testing.T) {
	s := new(Sudoku)
	s.AddNumber(3, 3, 3)
	assert.Equal(t, false, s.canBePut(3, 8, 3))
}

func Test_canBePut_ShouldReturnFalseIfPotentialNumberIsNotUniqueInsideTheColumn(t *testing.T) {
	s := new(Sudoku)
	s.AddNumber(4, 2, 1)
	assert.Equal(t, false, s.canBePut(7, 2, 1))
}

func Test_canBePut_ShouldReturnFalseIfThereIsAnotherNumberThanZeroInTheCell(t *testing.T) {
	s := new(Sudoku)
	s.AddNumber(0, 0, 1)
	assert.Equal(t, false, s.canBePut(0, 0, 2))
}

func Test_getPossibleNumberAmount_ShouldReturnOneForCellWhereOnlyOneNumberCanBePut(t *testing.T) {
	s := new(Sudoku)
	s.AddNumber(0, 0, 1)
	s.AddNumber(0, 1, 2)
	s.AddNumber(0, 2, 3)
	s.AddNumber(0, 3, 4)
	s.AddNumber(0, 4, 5)
	s.AddNumber(0, 5, 6)
	s.AddNumber(0, 6, 7)
	s.AddNumber(0, 7, 8)
	assert.Equal(t, 1, s.getPossibleNumberAmount(0, 8))
}

func Test_getPossibleNumberAmount_ShouldReturnThreeForCellWhereOnlyThreeNumbersCanBePut(t *testing.T) {
	s := new(Sudoku)
	s.AddNumber(0, 0, 1)
	s.AddNumber(1, 0, 2)
	s.AddNumber(2, 0, 3)
	s.AddNumber(3, 0, 4)
	s.AddNumber(4, 5, 5)
	s.AddNumber(4, 6, 6)
	assert.Equal(t, 3, s.getPossibleNumberAmount(4, 0))
}

func Test_getMostNormalizedCell_ShouldReturnMostNormalizedCell(t *testing.T) {
	s := new(Sudoku)
	s.Numbers = [9][9]int{
		{2, 4, 8, 3, 9, 5, 7, 1, 6},
		{5, 7, 1, 6, 2, 8, 3, 4, 9},
		{9, 3, 0, 7, 4, 1, 5, 8, 2},
		{6, 8, 2, 5, 3, 9, 1, 7, 4},
		{3, 5, 9, 1, 7, 4, 6, 2, 8},
		{7, 1, 4, 8, 6, 2, 9, 5, 3},
		{8, 6, 3, 4, 1, 7, 2, 9, 5},
		{1, 9, 5, 2, 8, 6, 4, 3, 7},
		{4, 2, 7, 9, 5, 3, 8, 6, 1},
	}
	actualCount, actualCell := s.getMostNormalizedCell()
	assert.Equal(t, 1, actualCount)
	assert.Equal(t, cell{2, 2}, actualCell)
}

func Test_getMostNormalizedCell_ShouldReturnFirstCellIfSudokuFieldIsEmpty(t *testing.T) {
	s := new(Sudoku)
	s.Numbers = [9][9]int{
		{0, 0, 0, 0, 0, 0, 0, 0, 0},
		{0, 0, 0, 0, 0, 0, 0, 0, 0},
		{0, 0, 0, 0, 0, 0, 0, 0, 0},
		{0, 0, 0, 0, 0, 0, 0, 0, 0},
		{0, 0, 0, 0, 0, 0, 0, 0, 0},
		{0, 0, 0, 0, 0, 0, 0, 0, 0},
		{0, 0, 0, 0, 0, 0, 0, 0, 0},
		{0, 0, 0, 0, 0, 0, 0, 0, 0},
		{0, 0, 0, 0, 0, 0, 0, 0, 0},
	}
	actualCount, actualCell := s.getMostNormalizedCell()
	assert.Equal(t, 9, actualCount)
	assert.Equal(t, cell{0, 0}, actualCell)
}

func Test_getMostNormalizedCell_ShouldReturnZeroAmountForTheOnlyOneEmptyCellOfBrokenSudoku(t *testing.T) {
	s := new(Sudoku)
	s.Numbers = [9][9]int{
		{9, 9, 9, 9, 9, 9, 9, 9, 9},
		{1, 2, 3, 4, 5, 0, 6, 7, 8},
		{9, 9, 9, 9, 9, 9, 9, 9, 9},
		{9, 9, 9, 9, 9, 9, 9, 9, 9},
		{9, 9, 9, 9, 9, 9, 9, 9, 9},
		{9, 9, 9, 9, 9, 9, 9, 9, 9},
		{9, 9, 9, 9, 9, 9, 9, 9, 9},
		{9, 9, 9, 9, 9, 9, 9, 9, 9},
		{9, 9, 9, 9, 9, 9, 9, 9, 9},
	}
	actualCount, actualCell := s.getMostNormalizedCell()
	assert.Equal(t, 0, actualCount)
	assert.Equal(t, cell{1, 5}, actualCell)
}

func Test_getUnresolvedCellAmount_ShouldReturnZeroIfAllCellsAreFilledWithnumbers(t *testing.T) {
	s := new(Sudoku)
	s.Numbers = [9][9]int{
		{2, 4, 8, 3, 9, 5, 7, 1, 6},
		{5, 7, 1, 6, 2, 8, 3, 4, 9},
		{9, 3, 6, 7, 4, 1, 5, 8, 2},
		{6, 8, 2, 5, 3, 9, 1, 7, 4},
		{3, 5, 9, 1, 7, 4, 6, 2, 8},
		{7, 1, 4, 8, 6, 2, 9, 5, 3},
		{8, 6, 3, 4, 1, 7, 2, 9, 5},
		{1, 9, 5, 2, 8, 6, 4, 3, 7},
		{4, 2, 7, 9, 5, 3, 8, 6, 1},
	}
	assert.Equal(t, 0, s.getUnresolvedCellAmount())
}

func Test_getUnresolvedCellAmount_ShouldReturn81IfSudokuFieldIsEmpty(t *testing.T) {
	s := new(Sudoku)
	s.Numbers = [9][9]int{
		{0, 0, 0, 0, 0, 0, 0, 0, 0},
		{0, 0, 0, 0, 0, 0, 0, 0, 0},
		{0, 0, 0, 0, 0, 0, 0, 0, 0},
		{0, 0, 0, 0, 0, 0, 0, 0, 0},
		{0, 0, 0, 0, 0, 0, 0, 0, 0},
		{0, 0, 0, 0, 0, 0, 0, 0, 0},
		{0, 0, 0, 0, 0, 0, 0, 0, 0},
		{0, 0, 0, 0, 0, 0, 0, 0, 0},
		{0, 0, 0, 0, 0, 0, 0, 0, 0},
	}
	assert.Equal(t, 81, s.getUnresolvedCellAmount())
}

func Test_getUnresolvedCellAmount_ShouldReturnCorrectAmountOfUnresolvedCells(t *testing.T) {
	s := new(Sudoku)
	s.Numbers = [9][9]int{
		{2, 4, 8, 3, 9, 5, 7, 1, 6},
		{5, 7, 1, 6, 2, 8, 3, 4, 9},
		{9, 3, 6, 0, 4, 1, 5, 8, 2},
		{6, 8, 2, 5, 3, 9, 1, 7, 4},
		{3, 5, 9, 1, 7, 0, 6, 2, 8},
		{7, 1, 4, 8, 6, 2, 9, 5, 3},
		{8, 6, 0, 4, 1, 7, 2, 9, 5},
		{1, 9, 5, 2, 8, 6, 4, 3, 7},
		{4, 0, 7, 9, 5, 3, 8, 0, 1},
	}
	assert.Equal(t, 5, s.getUnresolvedCellAmount())
}

func Test_getPossibleNumbers_ShouldReturnSliceOfAllPossibleNumbers(t *testing.T) {
	s := new(Sudoku)
	s.Numbers = [9][9]int{
		{2, 4, 8, 3, 9, 5, 7, 1, 6},
		{5, 7, 1, 6, 2, 8, 3, 4, 9},
		{9, 3, 0, 7, 4, 1, 5, 8, 2},
		{6, 8, 2, 5, 3, 9, 1, 7, 4},
		{3, 5, 9, 1, 7, 4, 6, 2, 8},
		{7, 1, 4, 8, 6, 2, 9, 5, 3},
		{8, 6, 3, 4, 1, 7, 2, 9, 5},
		{1, 9, 5, 2, 8, 6, 4, 3, 7},
		{4, 2, 7, 9, 5, 3, 8, 6, 1},
	}
	assert.Equal(t, []int{6}, s.getPossibleNumbers(cell{2, 2}))
}

func Test_Resolve_ShouldResolveSudoku(t *testing.T) {
	s := new(Sudoku)
	s.Numbers = [9][9]int{
		{2, 4, 8, 3, 9, 5, 7, 1, 6},
		{5, 7, 1, 6, 2, 8, 3, 4, 9},
		{9, 3, 0, 7, 4, 1, 5, 8, 2},
		{6, 8, 2, 5, 3, 9, 1, 7, 4},
		{3, 5, 9, 1, 7, 4, 6, 2, 8},
		{7, 1, 4, 8, 6, 2, 9, 5, 3},
		{8, 6, 3, 4, 1, 7, 2, 9, 5},
		{1, 9, 5, 2, 8, 6, 4, 3, 7},
		{4, 2, 7, 9, 5, 3, 8, 6, 1},
	}
	expected := new(Sudoku)
	expected.Numbers = [9][9]int{
		{2, 4, 8, 3, 9, 5, 7, 1, 6},
		{5, 7, 1, 6, 2, 8, 3, 4, 9},
		{9, 3, 6, 7, 4, 1, 5, 8, 2},
		{6, 8, 2, 5, 3, 9, 1, 7, 4},
		{3, 5, 9, 1, 7, 4, 6, 2, 8},
		{7, 1, 4, 8, 6, 2, 9, 5, 3},
		{8, 6, 3, 4, 1, 7, 2, 9, 5},
		{1, 9, 5, 2, 8, 6, 4, 3, 7},
		{4, 2, 7, 9, 5, 3, 8, 6, 1},
	}
	s.Resolve()
	assert.Equal(t, expected, s)
}

func Test_Resolve_ShouldReturnFalseIfSudokuCantBeResolved(t *testing.T) {
	s := new(Sudoku)
	s.Numbers = [9][9]int{
		{8, 5, 0, 3, 0, 0, 0, 1, 2},
		{0, 0, 0, 0, 0, 0, 5, 0, 0},
		{0, 0, 0, 8, 2, 0, 6, 0, 3},

		{0, 8, 0, 0, 6, 0, 9, 0, 0},
		{0, 4, 0, 7, 0, 2, 0, 8, 0},
		{0, 0, 2, 0, 3, 0, 0, 7, 0},

		{5, 0, 1, 0, 9, 6, 0, 0, 0},
		{0, 0, 7, 0, 0, 0, 0, 0, 0},
		{8, 9, 0, 0, 0, 3, 0, 0, 0},
	}
	assert.Equal(t, false, s.Resolve())
}
