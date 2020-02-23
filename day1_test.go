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