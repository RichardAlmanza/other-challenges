package main

var findPairsTestCases = []struct {
	description string
	sum         int
	numbers     []int
	expected    []pair
}{
	{
		description: "using example",
		sum:         12,
		numbers:     []int{1, 9, 5, 0, 20, -4, 12, 16, 7},
		expected: []pair{
			{-4, 16},
			{0, 12},
			{5, 7},
		},
	},
}
