package main

import (
	"fmt"
	"log"
	"math/rand"
	"os"
	"strconv"
	"strings"
)

const rounds = 10

func main() {
	if len(os.Args) != 3 {
		fmt.Println("--usage: ./GameOfPig [holdcount1 | rangestart-rangeend] [holdcount2 | rangestart-rangeend]")
		return
	}

	parsedRanges, err := ParseRanges(os.Args[1:])
	if err != nil {
		log.Fatal(err)
	}

	for _, r1 := range parsedRanges[0] {
		for _, r2 := range parsedRanges[1] {
			player1 := Player{strategy: getHoldAtValueStrategy(r1)}
			player2 := Player{strategy: getHoldAtValueStrategy(r2)}
			conductor := Conductor{player1: player1, player2: player2}
			conductor.Conduct(rounds)
		}
	}

}

/* Input Ranges: ["1-10", ["3"]];; Parsed Ranges: [[1,10], [3]] */
func ParseRanges(inputRanges []string) ([][]int, error) {
	parsedRanges := [][]int{}

	for _, ir := range inputRanges {
		parsedRange := []int{}

		inputRange := strings.Split(ir, "-")
		if len(inputRange) > 2 {
			return [][]int{}, fmt.Errorf("invalid range")
		}

		rangeStart, err := strconv.Atoi(inputRange[0])
		if err != nil {
			return [][]int{}, err
		}

		rangeEnd := rangeStart
		if len(inputRange) > 1 {
			rangeEnd, err = strconv.Atoi(inputRange[1])
			if err != nil {
				return [][]int{}, err
			}
		}

		rangeEnd += 1
		if rangeEnd <= rangeStart {
			return [][]int{}, fmt.Errorf("invalid range")
		}

		for i := rangeStart; i < rangeEnd; i++ {
			parsedRange = append(parsedRange, i)
		}

		parsedRanges = append(parsedRanges, parsedRange)
	}

	return parsedRanges, nil
}

/**** Helpers ****/

func RollDie() int {
	return rand.Intn(6) + 1
}

func CompareIntSlice(i1 []int, i2 []int) bool {
	if len(i1) != len(i2) {
		return false
	}

	for i, n1 := range i1 {
		if n1 != i2[i] {
			return false
		}
	}

	return true
}

func CompareIntSlices(i1 [][]int, i2 [][]int) bool {
	if len(i1) != len(i2) {
		return false
	}

	for i, s1 := range i1 {
		if !CompareIntSlice(s1, i2[i]) {
			return false
		}
	}

	return true
}
