package main

import (
	"errors"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

const (
	UseExample = `

==> Use example <==
app <numbers separated by comma> <target sum>
app 1,9,5,0,20,-4,12,16,7 12
===================

`
)

var (
	ErrExpectedArgument error = errors.New("the program receives only two arguments")
	ErrParsingNumbers   error = errors.New("an error occurred while parsing the given numbers")
)

func main() {
	var args = os.Args[1:]

	if len(args) != 2 {
		log.Fatal(ErrExpectedArgument, UseExample)
	}

	numbers, errN := parseNumbers(args[0])
	targetsSum, errT := parseNumbers(args[1])

	err := errors.Join(errN, errT)

	if err != nil {
		log.Fatal(err, UseExample)
	}

	for _, p := range FindPairIntegers(numbers, targetsSum[0]) {
		fmt.Printf("%d,%d\n", p[0], p[1])
	}

}

func parseNumbers(input string) ([]int, error) {
	var numbers []int

	for _, number := range strings.Split(input, ",") {
		n, err := strconv.Atoi(number)
		
		if err != nil {
			return nil, errors.Join(ErrParsingNumbers, err)
		}

		numbers = append(numbers, n)
	}

	return numbers, nil
}
