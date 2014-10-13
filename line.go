package gosudoku

type line struct {
	numbers [9]int
}

// Returns true if all numbers in the slice (except zero) are unique
func (l *line) isUnique() bool {
	for i := 0; i < 8; i++ {
		for j := i + 1; j < 9; j++ {
			if l.numbers[i] != 0 && l.numbers[i] == l.numbers[j] {
				return false
			}
		}
	}
	return true
}
