package main

import (
	"testing"
)

func TestManhattanDistance(t *testing.T) {
	x := 0
	y := 0
	d := manhattanDistance(x, y)
	if d != 0 {
		t.Errorf("d = %d; want 0", d)
	}

	x = 1
	d = manhattanDistance(x, y)
	if d != 1 {
		t.Errorf("d = %d; want 1", d)
	}

	y = 1
	d = manhattanDistance(x, y)
	if d != 2 {
		t.Errorf("d = %d; want 2", d)
	}

	x = -2
	d = manhattanDistance(x, y)
	if d != 3 {
		t.Errorf("d = %d; want 3", d)
	}
}

func TestFollowRoute(t *testing.T) {
	w := newWire()

	testRoute := Route{
		Path{direction: DIRECTION_DOWN, length: 1},
		Path{direction: DIRECTION_LEFT, length: 2},
		Path{direction: DIRECTION_UP, length: 3},
		Path{direction: DIRECTION_RIGHT, length: 4},
	}
	expectedPosition := Position{2, 2}

	w.route = testRoute
	w.followRoute()

	if w.trail[len(w.trail)-1] != expectedPosition {
		t.Errorf("Position is %v; want %v", w.trail[len(w.trail)-1], expectedPosition)
	}
}
func TestCrossingAt(t *testing.T) {
	w1 := newWire()
	w2 := newWire()
	w1.trail = append(w1.trail, Trail{Position{0, 1}, Position{1, 1}}...)
	w2.trail = append(w2.trail, Trail{Position{1, 0}, Position{1, 1}}...)
	i := w1.crossingAt(w2)
	if i != 2 {
		t.Errorf("CrossingAt is %d; want 2", i)
	}
}
