package gosudoku

type square struct {
	numbers [3][3]int
}

// Returns true if all numbers in the square (except zero) are unique
func (s *square) isUnique() bool {
	l := new(line)
	index := 0
	for i := 0; i < 3; i++ {
		for j := 0; j < 3; j++ {
			l.numbers[index] = s.numbers[i][j]
			index++
		}
	}
	return l.isUnique()
}
