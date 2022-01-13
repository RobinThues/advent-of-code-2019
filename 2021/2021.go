package twentyone

import (
	"github.com/robinthues/advent-of-code/helpers"
	"math"
	"strconv"
	"strings"
)

func d14one() {
	lines, _ := helpers.LoadInput("input-2021-14.txt")
	startSequence := lines[0]

	mappings := map[string]string{}
	for _, line := range lines {
		if !strings.Contains(line, " -> ") {
			continue
		}
		split := strings.Split(line, " -> ")
		mappings[split[0]] = split[1]
	}

	computeNextSequence := func(prevSequence string) string {
		nextSequence := ""
		for i, _ := range prevSequence {
			if i == len(prevSequence)-1 {
				// exit on the last sign, as there is no next one
				nextSequence += string(prevSequence[i])
				continue
			}
			characterAndNext := "" + string(prevSequence[i]) + string(prevSequence[i+1])
			relevantMapping := mappings[characterAndNext]
			nextSequence += string(prevSequence[i]) + relevantMapping
		}
		return nextSequence
	}

	currentSequence := startSequence
	for iterations := 0; iterations < 5; iterations++ {
		currentSequence = computeNextSequence(currentSequence)
		println("Iteration Done: ", iterations, " length: ", len(currentSequence))
	}

	letterOccurances := map[string]int{}
	for _, b := range currentSequence {
		letterOccurances[string(b)] += 1
	}

	max := 0
	min := 10000000000
	for _, v := range letterOccurances {
		if v > max {
			max = v
		}
		if v < min {
			min = v
		}
	}
	maxMinusMin := max - min
	println("Solution: ", maxMinusMin)
}

func d3one() {
	lines, _ := helpers.LoadInput("input-2021-3.txt")

	numberOfLines := len(lines)
	println("numberOfLines", numberOfLines)
	numberOfBits := len(lines[0])
	println("numberOfBits: ", numberOfBits)

	countsOfOne := make([]int, numberOfBits)
	for k, _ := range countsOfOne {
		countsOfOne[k] = 0
	}
	println("countsOfOne: ", len(countsOfOne))

	for _, l := range lines {
		for k, v := range l {
			if v == '1' {
				countsOfOne[k] += 1
			}
		}
	}
	for k, b := range countsOfOne {
		println("Ones in ", k, ": ", b)
	}

	var gammaRate int
	for k, i := range countsOfOne {
		if i > numberOfLines/2 {
			gammaRate += int(math.Pow(2, float64(numberOfBits-k-1)))
		}
	}
	println("Gamma Rate: ", gammaRate)

	var epsilonRate int
	for k, i := range countsOfOne {
		if i <= numberOfLines/2 {
			p := int(math.Pow(2, float64(numberOfBits-k-1)))
			epsilonRate += p
		}
	}
	println("Epsilon Rate: ", epsilonRate)

	println(epsilonRate * gammaRate)
}

func d2two() {
	lines, _ := helpers.LoadInput("input-2021-2.txt")

	depth := 0
	forward := 0
	aim := 0

	for _, l := range lines {
		line := strings.Split(l, " ")

		value, _ := strconv.Atoi(line[1])
		cmd := line[0]

		switch cmd {
		case "forward":
			forward += value
			depth += aim * value
		case "down":
			aim += value
		case "up":
			aim -= value
		}
	}

	println(forward * depth)
}

func d2() {
	lines, _ := helpers.LoadInput("input-2021-2.txt")

	depth := 0
	forward := 0

	for _, l := range lines {
		line := strings.Split(l, " ")

		value, _ := strconv.Atoi(line[1])
		cmd := line[0]

		switch cmd {
		case "forward":
			forward += value
		case "down":
			depth += value
		case "up":
			depth -= value
		}
	}

	println(forward * depth)
}
