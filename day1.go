package main

import (
	"bufio"
	"log"
	"os"
	"strconv"
)

func day1(part int) (fuel int) {
	lines := loadInput("input1.txt")

	for _, line := range lines {
		lineAsInt, _ := strconv.Atoi(line)
		fuel += fuelForLoad(lineAsInt, part)
	}

	return
}

func fuelForLoad(load, part int) int {
	fuel := load / 3 - 2
	if fuel <= 0 {
		return 0
	}
	if part == 1 {
		return fuel
	}
	return fuel + fuelForLoad(fuel, part)
}

func loadInput(fileName string) []string {
	file, err := os.Open(fileName)
	if err != nil {
		log.Fatal(err)
	}

	defer func() {
		err := file.Close()
		if err != nil {
			log.Fatal(err)
		}
	}()

	var lines []string
	s := bufio.NewScanner(file)
	for s.Scan() {
		lines = append(lines, s.Text())
	}
	return lines
}
