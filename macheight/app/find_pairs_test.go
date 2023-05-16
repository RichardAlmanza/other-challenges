package main

import (
	"testing"
	"reflect"
)

func Test(t *testing.T) {
	for _, tC := range findPairsTestCases {
		t.Run(tC.description, func(t *testing.T) {
			actual := FindPairIntegers(tC.numbers, tC.sum)

			if !reflect.DeepEqual(actual, tC.expected) {
				t.Fatalf("FindPairIntegers(%v, %v) = %v, want: %v", tC.numbers, tC.sum, actual, tC.expected)
			}
		})
	}
}
