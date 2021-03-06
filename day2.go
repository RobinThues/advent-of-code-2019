package main

import (
	"fmt"
	"io/ioutil"
	"strconv"
	"strings"
)

const(
	OPCODE_ADD = 1
	OPCODE_MULTIPLY = 2
	OPCODE_HALT = 99
	PROGRAM_LENGTH = 4
)

func loadProgramFromInput(fileName string) []int {
	file, err := ioutil.ReadFile(fileName)
	if err != nil {
		return nil
	}
	stringProgram := string(file)
	intCodes := strings.Split(stringProgram, ",")

	memory := make([]int, len(intCodes))
	for index, stringAddressValue := range intCodes {
		intAddressValue, _ := strconv.Atoi(stringAddressValue)
		memory[index] = intAddressValue
	}
	return memory
}

type IntcodeComputer struct {
	memory []int
	instructionPointer int
}

func (c IntcodeComputer) runProgram() int {
	firstParameter := c.memory[c.instructionPointer + 1]
	secondParameter := c.memory[c.instructionPointer + 2]
	thirdParameter := c.memory[c.instructionPointer + 3]

	switch c.memory[c.instructionPointer] {
	case OPCODE_ADD:
		c.memory[thirdParameter] = c.memory[firstParameter] + c.memory[secondParameter]
		return 1
	case OPCODE_MULTIPLY:
		c.memory[thirdParameter] = c.memory[firstParameter] * c.memory[secondParameter]
		return 1
	case OPCODE_HALT:
		return -1
	default:
		return -1
	}
}

func (c IntcodeComputer) runAllPrograms() {
	for c.instructionPointer < len(c.memory) - PROGRAM_LENGTH {
		programStatus := c.runProgram()
		if programStatus == -1 {
			return
		}
		c.instructionPointer += PROGRAM_LENGTH
	}
}

func NewIntcodeComputer(memory []int) *IntcodeComputer {
	c := IntcodeComputer{
		memory:             memory,
		instructionPointer: 0,
	}
	return &c
}

func day2(part int) {
	memory := loadProgramFromInput("input2.txt")

	if part == 1 {
		day2part1(memory)
	}
	if part == 2 {
		day2part2(memory)
	}
}

func day2part1(memory []int) {
	memory[1] = 12
	memory[2] = 2
	c := NewIntcodeComputer(memory)
	c.runAllPrograms()
	fmt.Println("Result for Part1:",memory[0])
}

func day2part2(memory []int) {
	for noun := 0; noun < 99; noun++ {
		for verb := 0; verb < 99; verb++ {
			iterationMemory := make([]int, len(memory))
			copy(iterationMemory, memory)
			iterationMemory[1] = noun
			iterationMemory[2] = verb

			c := NewIntcodeComputer(iterationMemory)
			c.runAllPrograms()
			if c.memory[0] == 19690720 {
				fmt.Println("Result for Part2:", 100 * noun + verb)
				return
			}
		}
	}
}