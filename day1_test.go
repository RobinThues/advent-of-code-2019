package main

import (
	"testing"
)

func TestFuelForLoad(t *testing.T) {
	gotFor0 := fuelForLoad(0, 1)
	if gotFor0 != 0 {
		t.Errorf("fuelForLoad(0,1) = %d; want 0", gotFor0)
	}
	gotFor1 := fuelForLoad(1, 1)
	if gotFor1 != 0 {
		t.Errorf("fuelForLoad(1,1) = %d; want 0", gotFor1)
	}
	gotFor9 := fuelForLoad(9, 1)
	if gotFor9 != 1 {
		t.Errorf("fuelForLoad(9,1) = %d; want 1", gotFor9)
	}
	gotFor0part2 := fuelForLoad(0, 2)
	if gotFor0part2 != 0 {
		t.Errorf("fuelForLoad(0,2) = %d; want 1", gotFor0part2)
	}
	gotFor1part2 := fuelForLoad(1, 2)
	if gotFor1part2 != 0 {
		t.Errorf("fuelForLoad(1,2) = %d; want 0", gotFor1part2)
	}
	gotFor9part2 := fuelForLoad(9, 2)
	if gotFor9part2 != 1 {
		t.Errorf("fuelForLoad(9,2) = %d; want 1", gotFor9part2)
	}
}

func TestDay1Part1(t *testing.T) {
	v := day1(1)
	if v != 3393938 {
		t.Errorf("failed day1 part1")
	}
}

func TestLoadInvalidInput(t *testing.T) {
	v := loadInput("")
	if v != nil {
		t.Errorf("loaded invalid input somehow")
	}
}

func BenchmarkFuelForLoadPartOne(b *testing.B) {
	for i := 0; i < b.N; i++ {
		fuelForLoad(i, 1)
	}
}

func BenchmarkFuelForLoadPartTwo(b *testing.B) {
	for i := 0; i < b.N; i++ {
		fuelForLoad(i, 2)
	}
}