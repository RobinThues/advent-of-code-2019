package main

import (
	"testing"
	h "github.com/robinthues/go-helpers"
)

func TestNewIntcodeComputer(t *testing.T) {
	memory := []int{1, 2, 3, 4}
	c := NewIntcodeComputer(memory)

	if c.instructionPointer != 0 {
		t.Errorf("c.instructionPointer = %d; want 0", c.instructionPointer)
	}
	if h.CompareIntSlices(memory, c.memory) != true {
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
