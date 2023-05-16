package main

var findPairsTestCases = []struct {
	description string
	sum         int
	numbers     []int
	expected    pairs
}{
	{
		description: "using example",
		sum:         12,
		numbers:     []int{1, 9, 5, 0, 20, -4, 12, 16, 7},
		expected: pairs{
			{12, 0},
			{5, 7},
			{16, -4},
		},
	},
}
