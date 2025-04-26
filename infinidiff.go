package infinidiff

import "fmt"

// Diff should be a function that returns
// equal should be a function returning true if a[ai] == b[bi]
// equal must be transitive, that is if a == b a == c, b == c.
func Diff(lengths []int, equal func(a, ai, b, bi int) bool) [][]bool {
	memo := map[string][][]bool{}
	return diffMemo(lengths, equal, memo)
}

func diffMemo(lengths []int, equal func(a, ai, b, bi int) bool, memo map[string][][]bool) [][]bool {
	memoKey := fmt.Sprint(lengths)
	if result, ok := memo[memoKey]; ok {
		return result
	}

	checked := make([]bool, len(lengths))

	var best [][]bool
	var bestChoice []bool
	for i := range checked {
		if checked[i] { // already in a checked set
			continue
		}

		if lengths[i] == 0 {
			continue
		}

		// set of matching lines to try using
		choice := make([]bool, len(lengths))
		choice[i] = true
		for j := i + 1; j < len(lengths); j++ {
			if lengths[j] != 0 && equal(i, lengths[i]-1, j, lengths[j]-1) {
				choice[j] = true
				checked[j] = true
			}
		}

		subLengths := make([]int, len(lengths))
		for j := range lengths {
			subLengths[j] = lengths[j]
			if choice[j] {
				subLengths[j]--
			}
		}

		subResult := diffMemo(subLengths, equal, memo)
		if best == nil || len(subResult) < len(best) {
			best = subResult
			bestChoice = choice
		}
	}

	if best == nil { // there were no choices considered, meaning length is all zeros.
		memo[memoKey] = [][]bool{}
		return [][]bool{}
	}

	result := make([][]bool, len(best)+1)
	copy(result[:len(best)], best)
	result[len(best)] = bestChoice

	memo[memoKey] = result
	return result
}
