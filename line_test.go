package sudoku

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_isUnique_ShouldReturnTrueIfAllNumbersExceptZeroAreUnique(t *testing.T) {
	l1 := line{numbers: [9]int{1, 2, 3, 4, 5, 6, 7, 8, 9}}
	l2 := line{numbers: [9]int{1, 2, 0, 0, 0, 0, 0, 0, 9}}
	assert.Equal(t, true, l1.isUnique())
	assert.Equal(t, true, l2.isUnique())
}

func Test_isUnique_ShouldReturnFalseIfNumbersAreNotUnique(t *testing.T) {
	l1 := line{numbers: [9]int{1, 2, 3, 3, 4, 5, 6, 7, 8}}
	l2 := line{numbers: [9]int{1, 0, 0, 0, 0, 0, 0, 0, 1}}
	l3 := line{numbers: [9]int{0, 0, 0, 0, 0, 0, 0, 1, 1}}
	assert.Equal(t, false, l1.isUnique())
	assert.Equal(t, false, l2.isUnique())
	assert.Equal(t, false, l3.isUnique())
}
