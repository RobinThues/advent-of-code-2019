package main

import (
	"fmt"
	"strconv"
	"strings"
)

const(
	DIRECTION_UP = "U"
	DIRECTION_DOWN = "D"
	DIRECTION_LEFT = "L"
	DIRECTION_RIGHT = "R"
)

type Wires = []Wire

type Wire struct {
	steps []Step
}

type Step struct {
	direction string
	length int
}

func day3(part int) {
	wires := readDay3Input("input3.txt")
	if part == 1 {
		buildMap(wires)
	}
}

func readDay3Input(fileName string) Wires {
	lines := loadInput(fileName)

	var wires Wires

	for _, l := range lines {
		splittedLine := strings.Split(l, ",")
		wire := Wire{}

		for _, split := range splittedLine {
			s, _ := strconv.Atoi(split[1:])
			wire.steps = append(wire.steps, Step{
				string(split[0]),
				s,
			})
		}
		wires = append(wires, wire)
	}
	return wires
}

type Map = [][]int

func buildMap(wires Wires) Map {
	maxUp := 0
	maxDown := 0
	maxLeft := 0
	maxRight := 0

	for _, wire := range wires {
		currVertical := 0
		currHorizontal := 0

		for _, step := range wire.steps {
			switch step.direction {
			case DIRECTION_UP:
				currVertical += step.length
				if currVertical > maxUp {
					maxUp = currVertical
				}
			case DIRECTION_DOWN:
				currVertical -= step.length
				if -currVertical > maxDown {
					maxDown = -currVertical
				}
			case DIRECTION_LEFT:
				currHorizontal -= step.length
				if -currHorizontal > maxLeft {
					maxLeft = -currHorizontal
				}
			case DIRECTION_RIGHT:
				currHorizontal += step.length
				if currHorizontal > maxRight {
					maxRight = currHorizontal
				}
			}
		}
	}

	fmt.Println("MaxUp:", maxUp)
	fmt.Println("maxDown:", maxDown)
	fmt.Println("maxLeft:", maxLeft)
	fmt.Println("maxRight:", maxRight)

	mapHorizontal := make([][]int, maxLeft+maxRight)
	for i := range mapHorizontal {
		mapHorizontal[i] = make([]int, maxUp+maxDown)
	}
	return mapHorizontal
}

func drawWiresOnMap() {
	// for the first wire
	// -- walk path and mark with X
	// for the second wire
	// -- walk path, mark with Y if empty, mark with Z if cross
}

func findClosestCrossing() {
	// for each crossing
	// -- calculate horizontal and vertical difference to middle, add up
	// choose crossing with minimum
}
