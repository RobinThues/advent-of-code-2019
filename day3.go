package main

import (
	"fmt"
	"math"
	"strconv"
	"strings"
)

const (
	DIRECTION_UP    = "U"
	DIRECTION_DOWN  = "D"
	DIRECTION_LEFT  = "L"
	DIRECTION_RIGHT = "R"
)

type Wires = []*Wire
type Route = []Path

type Position struct {
	x int
	y int
}

type Trail = []Position

type Wire struct {
	route Route
	trail Trail
}

type Path struct {
	direction string
	length    int
}

func (w *Wire) followRoute() {
	for _, p := range w.route {
		stepsGone := 0
		for stepsGone < p.length {
			switch p.direction {
			case DIRECTION_LEFT:
				w.trail = append(w.trail, Position{
					x: w.trail[len(w.trail)-1].x - 1,
					y: w.trail[len(w.trail)-1].y,
				})
			case DIRECTION_RIGHT:
				w.trail = append(w.trail, Position{
					x: w.trail[len(w.trail)-1].x + 1,
					y: w.trail[len(w.trail)-1].y,
				})
			case DIRECTION_UP:
				w.trail = append(w.trail, Position{
					x: w.trail[len(w.trail)-1].x,
					y: w.trail[len(w.trail)-1].y + 1,
				})
			case DIRECTION_DOWN:
				w.trail = append(w.trail, Position{
					x: w.trail[len(w.trail)-1].x,
					y: w.trail[len(w.trail)-1].y - 1,
				})
			}
			stepsGone += 1
		}
	}
}

func (w *Wire) crossingAt(other *Wire) {
	min := 100000
	for _, p := range w.trail {
		for _, otherP := range other.trail {
			if p == otherP {
				distance := int(math.Abs(float64(p.x)) + math.Abs(float64(p.y)))
				if distance > 0 && distance < min {
					min = distance
				}
			}
		}
	}
	fmt.Println("min", min)
}

func (w *Wire) closestStepCrossingAt(other *Wire) {
	min := 100000
	for iw, p := range w.trail {
		for io, otherP := range other.trail {
			if p == otherP {
				distance := iw + io
				if distance > 0 && distance < min {
					min = distance
				}
			}
		}
	}
	// Solution: 209
	fmt.Println("min", min)
}

func day3(part int) {
	wires := readDay3Input("input3.txt")
	for _, w := range wires {
		w.followRoute()
		fmt.Println("Final Position", w.trail[len(w.trail)-1])
	}
	// Solution: 43258
	wires[0].closestStepCrossingAt(wires[1])
}

func readDay3Input(fileName string) Wires {
	lines := loadInput(fileName)

	var wires Wires

	for _, l := range lines {
		splittedLine := strings.Split(l, ",")
		wire := Wire{}

		wire.trail = append(wire.trail, Position{0, 0})

		for _, split := range splittedLine {
			s, _ := strconv.Atoi(split[1:])
			wire.route = append(wire.route, Path{
				string(split[0]),
				s,
			})
		}
		wires = append(wires, &wire)
	}
	return wires
}
