package main

import (
	"testing"
	"reflect"
)

func Test(t *testing.T) {
	for _, tc := range findPairsTestCases {
		t.Run(tc.description, func(t *testing.T) {
			actual := FindPairIntegers(tc.numbers, tc.sum)

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
