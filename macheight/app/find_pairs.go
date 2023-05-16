package main

type pair [2]int
type pairs []pair
type pairSet map[int]int

func FindPairIntegers(numbers []int, sum int) pairs {
	var pairsMap pairSet = make(pairSet)
	var output pairs = make(pairs, 0, len(numbers)/2)

	for _, number := range numbers {
		expected := sum - number

		if _, exists := pairsMap[number]; !exists {
			pairsMap[expected] = number
			continue
		}

		output = append(output, pair{number, expected})
	}

	return output
}
