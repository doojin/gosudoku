package sudoku

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_isUnique_ShouldReturnTrueIfAllNumbersAreUniqueInTheSquare(t *testing.T) {
	s1 := square{
		// All numbers but zero
		numbers: [3][3]int{
			{1, 2, 3},
			{4, 5, 6},
			{7, 8, 9},
		},
	}
	// All zero
	s2 := square{
		numbers: [3][3]int{
			{0, 0, 0},
			{0, 0, 0},
			{0, 0, 0},
		},
	}
	// Zero + Non-zero numbers
	s3 := square{
		numbers: [3][3]int{
			{0, 0, 0},
			{0, 1, 0},
			{0, 3, 2},
		},
	}
	assert.Equal(t, true, s1.isUnique())
	assert.Equal(t, true, s2.isUnique())
	assert.Equal(t, true, s3.isUnique())
}

func Test_isUnique_ShouldReturnFalseIfNumbersAreNotUniqueInTheSquare(t *testing.T) {
	// Not unique numbers on one row
	s1 := square{
		numbers: [3][3]int{
			{0, 0, 0},
			{1, 0, 1},
			{0, 0, 0},
		},
	}
	// Not unique numbers on one col
	s2 := square{
		numbers: [3][3]int{
			{1, 0, 0},
			{0, 0, 0},
			{1, 0, 0},
		},
	}
	// Not unique numbers in custom place
	s3 := square{
		numbers: [3][3]int{
			{1, 0, 0},
			{0, 2, 0},
			{0, 0, 1},
		},
	}
	assert.Equal(t, false, s1.isUnique())
	assert.Equal(t, false, s2.isUnique())
	assert.Equal(t, false, s3.isUnique())
}
