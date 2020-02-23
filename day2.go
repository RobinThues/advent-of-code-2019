package main

import (
	"errors"
	"fmt"
	"io/ioutil"
	"strconv"
	"strings"
)

func day2() {
	file, err := ioutil.ReadFile("input2.txt")
	if err != nil {
		return
	}
	sf := string(file)
	splittedValues := strings.Split(sf, ",")

	memory := make([]int, len(splittedValues))
	for index, stringAddressValue := range splittedValues {
		intAddressValue, _ := strconv.Atoi(stringAddressValue)
		memory[index] = intAddressValue
	}

	noun := 0
	verb := 0
	for true {
		iterationMemory := make([]int, len(memory))
		copy(iterationMemory, memory)
		fmt.Println(noun, verb)
		iterationMemory[1] = noun
		iterationMemory[2] = verb
		runAllProgramms(iterationMemory)
		if iterationMemory[0] == 19690720 {
			fmt.Println(noun, verb)
			return
		}
		if noun < 99 {
			noun += 1
		} else {
			if verb < 99 {
				verb += 1
			}
		}
		if noun == 99 && verb == 99 {
			return
		}
	}
}

func runAllProgramms(memory []int) {
	instructionPointer := 0
	for instructionPointer < len(memory) - 4 {
		err := executeProgramm(memory[instructionPointer:instructionPointer+4], memory)
		if err != nil {
			instructionPointer = len(memory)
		}
		instructionPointer += 4
	}
}

func executeProgramm(programm []int, memory []int) error {
	firstParam := memory[programm[1]]
	secondParam := memory[programm[2]]
	if programm[0] == 1 {
		memory[programm[3]] = firstParam + secondParam
		return nil
	}
	if programm[0] == 2 {
		memory[programm[3]] = firstParam * secondParam
		return nil
	}
	return errors.New("ende")
}
