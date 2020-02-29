package main

import (
	"fmt"
	"strconv"
	"strings"
)

const input string = "136818 - 685979"

func day4() {
	min, max := readDay4Input()
	count := 0
	for i := min; i <= max; i++ {
		if matchesCriteria(i) {
			count += 1
		}
	}
	// Solution Part1: 1919
	fmt.Println(count)
}

func matchesCriteria(input int) bool {
	stringNumber := strconv.Itoa(input)
	digitsCount := len(stringNumber) - 1
	isOnlyIncreasing := true

	for i := 0; i < digitsCount; i++ {
		// hacking around nasty runes
		d1, _ := strconv.Atoi(stringNumber[i : i+1])
		d2, _ := strconv.Atoi(stringNumber[i+1 : i+2])

		if d1 > d2 {
			isOnlyIncreasing = false
		}
	}

	hasPair := false
	for i := 0; i < digitsCount; i++ {
		if stringNumber[i] == stringNumber[i+1] {
			hasPair = true
		}
	}

	return isOnlyIncreasing && hasPair
}

func readDay4Input() (int, int) {
	splitted := strings.Split(input, " - ")
	min, _ := strconv.Atoi(splitted[0])
	max, _ := strconv.Atoi(splitted[1])
	return min, max
}
