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

func newWire() *Wire {
	w := Wire{
		trail: Trail{Position{0, 0}},
	}
	return &w
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

func (w *Wire) crossingAt(other *Wire) int {
	min := 100000
	for _, p := range w.trail {
		d := manhattanDistance(p.x, p.y)
		if d > 0 && d < min {
			for _, otherP := range other.trail {
				if p == otherP {
					min = d
				}
			}
		}
	}
	// Solution: 209
	return min
}

// manhattanDistance computes the manhattan distance to the center (0,0)
func manhattanDistance(x, y int) int {
	return int(math.Abs(float64(x)) + math.Abs(float64(y)))
}

func (w *Wire) closestStepCrossingAt(other *Wire) {
	min := 100000
	for iw, p := range w.trail {
		if iw < min {
			for io, otherP := range other.trail {
				if p == otherP {
					distance := iw + io
					if distance > 0 && distance < min {
						min = distance
					}
				}
			}
		}
	}
	// Solution: 43258
	fmt.Println("min", min)
}

func day3(part int) {
	wires := readDay3Input("input3.txt")
	for _, w := range wires {
		w.followRoute()
	}
	if part == 1 {
		fmt.Println("min", wires[0].crossingAt(wires[1]))
	} else {
		wires[0].closestStepCrossingAt(wires[1])
	}
}

func readDay3Input(fileName string) Wires {
	lines := loadInput(fileName)

	var wires Wires

	for _, l := range lines {
		splittedLine := strings.Split(l, ",")
		wire := newWire()

		for _, split := range splittedLine {
			s, _ := strconv.Atoi(split[1:])
			wire.route = append(wire.route, Path{
				string(split[0]),
				s,
			})
		}
		wires = append(wires, wire)
	}
	return wires
}
