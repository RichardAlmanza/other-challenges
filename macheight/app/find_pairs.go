package main

type pair [2]int
type pairSet map[int]int

func FindPairIntegers(numbers []int, sum int) []pair {
	var pairs pairSet = make(pairSet)
	var output []pair = make([]pair, 0, len(numbers)/2)

	for _, number := range numbers {
		expected := sum - number

		if _, exists := pairs[number]; !exists {
			pairs[expected] = number
			continue
		}

		output = append(output, pair{number, expected})
	}

	return output
}
