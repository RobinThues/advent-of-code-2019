package main

import (
	"testing"
)

func TestNewIntcodeComputer(t *testing.T) {
	memory := []int{1, 2, 3, 4}
	c := NewIntcodeComputer(memory)

	if c.instructionPointer != 0 {
		t.Errorf("c.instructionPointer = %d; want 0", c.instructionPointer)
	}
	if compareMemories(memory, c.memory) != true {
		t.Errorf("c.memory = %v; want %v", c.memory, memory)
	}
}

func TestRunProgram(t *testing.T) {
	memory := []int{1, 0, 0, 1}
	c := NewIntcodeComputer(memory)
	c.runProgram()
	if c.memory[1] != 2 {
		t.Errorf("memory[1] = %d; want 2", c.memory[1])
	}
}

func compareMemories(a, b []int) bool {
	if len(a) != len(b) {
		return false
	}

	for i, v := range a {
		if v != b[i] {
			return false
		}
	}
	return true
}

func TestCompareMemories(t *testing.T) {
	a := []int{1, 2}
	b := []int{0}
	if compareMemories(a, b) != false {
		t.Errorf("compareMemories(a, b) = %v; want %v", compareMemories(a, b), false)
	}

	b = []int{0, 2}
	if compareMemories(a, b) != false {
		t.Errorf("compareMemories(a, b) = %v; want %v", compareMemories(a, b), false)
	}

	b = []int{1, 2}
	if compareMemories(a, b) != true {
		t.Errorf("compareMemories(a, b) = %v; want %v", compareMemories(a, b), true)
	}
}