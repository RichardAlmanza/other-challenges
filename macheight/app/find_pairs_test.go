package main

import (
	"reflect"
	"sort"
	"testing"
)

func (ps pairs) Len() int           { return len(ps) }
func (ps pairs) Swap(i, j int)      { ps[i], ps[j] = ps[j], ps[i] }
func (ps pairs) Less(i, j int) bool { return ps[i][0] < ps[j][0] }

func (ps pairs) sortPairs() {
	for i := range ps {
		if ps[i][1] < ps[i][0] {
			ps[i][0], ps[i][1] = ps[i][1], ps[i][0]
		}
	}
}


func Test(t *testing.T) {
	for _, tc := range findPairsTestCases {
		t.Run(tc.description, func(t *testing.T) {
			actual := FindPairIntegers(tc.numbers, tc.sum)

			actual.sortPairs()
			sort.Sort(actual)

			tc.expected.sortPairs()
			sort.Sort(tc.expected)

			if !reflect.DeepEqual(actual, tc.expected) {
				t.Fatalf("FindPairIntegers(%v, %v) = %v, want: %v", tc.numbers, tc.sum, actual, tc.expected)
			}
		})
	}
}

func Benchmark(b *testing.B) {
	if testing.Short() {
		b.Skip("skipping benchmark in short mode.")
	}

	for i := 0; i < b.N; i++ {
		for _, tc := range findPairsTestCases {
			FindPairIntegers(tc.numbers, tc.sum)
		}
	}
}
