package main

import (
	"sort"
)

type pair [2]int
type pairs []pair
type pairSet map[int]int

func (ps pairs) Len() int           { return len(ps) }
func (ps pairs) Swap(i, j int)      { ps[i], ps[j] = ps[j], ps[i] }
func (ps pairs) Less(i, j int) bool { return ps[i][0] < ps[j][0] }

func newPair(a, b int) pair {
	if b < a {
		a, b = b, a
	}

	return pair{a, b}
}

func FindPairIntegers(numbers []int, sum int) []pair {
	var pairsMap pairSet = make(pairSet)
	var output pairs = make(pairs, 0, len(numbers)/2)

	for _, number := range numbers {
		expected := sum - number

		if _, exists := pairsMap[number]; !exists {
			pairsMap[expected] = number
			continue
		}

		output = append(output, newPair(number, expected))
	}

	sort.Sort(output)

	return output
}
