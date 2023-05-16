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
	{
		description: "50 numbers",
		sum:         -50,
		numbers: []int{
			-34, -22, 34, 37, 14,
			-11, -29, 53, 7, 1,
			65, 67, 74, 23, 29,
			46, -32, 17, 30, -16,
			78, -26, 71, 25, 0,
			31, 42, 61, -3, 3,
			-15, -36, -8, 5, 72,
			55, -17, 50, 12, 9,
			59, -31, 48, 76, 15,
			-13, 33, 10, 13, -28,
		},
		expected: pairs{
			{-34, -16},
			{-22, -28},
		},
	},
	{
		description: "one number",
		sum:         2,
		numbers:     []int{2},
		expected:    pairs{},
	},
	{
		description: "empty output",
		sum:         2,
		numbers:     []int{1, 9, 5, 0, 20, -4, 12, 16, 7},
		expected:    pairs{},
	},
	{
		description: "empty input",
		sum:         2,
		numbers:     []int{},
		expected:    pairs{},
	},
	{
		description: "sum 0 empty input",
		sum:         0,
		numbers:     []int{},
		expected:    pairs{},
	},
	{
		description: "sum 0 empty output",
		sum:         0,
		numbers:     []int{9, 3, 20, 28, 65},
		expected:    pairs{},
	},
	{
		description: "sum 0",
		sum:         0,
		numbers: []int{
			9, 3, 20, 28, 65,
			29, -40, -8, 55, 80,
			70, 38, 64, -13, 8,
			51, 30, 67, 60, 59,
		},
		expected: pairs{
			{8, -8},
		},
	},
}
