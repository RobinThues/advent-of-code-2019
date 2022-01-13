package main

import (
	"bufio"
	"fmt"
	"os"
	"path/filepath"
	"strconv"
)

func main() {
	lines := readFileByLines(filepath.Join("2018", "input1.txt"))
	var results []int
	result := 0
	found := false
	for _, line := range lines {
		if found {
			return
		}
		results = append(results, result)
		result += inputLineToInteger(line)
		for _, r := range results {
			if r == result {
				fmt.Println(r)
				found = true
			}
		}
	}
	fmt.Println(result)
}

func readFileByLines(fileName string) []string {
	file, err := os.Open(fileName)
	if err != nil {
		//log.Fatal(err)
	}

	defer func() {
		err := file.Close()
		if err != nil {
			//log.Fatal(err)
		}
	}()

	var lines []string
	s := bufio.NewScanner(file)
	for s.Scan() {
		lines = append(lines, s.Text())
	}
	return lines
}

func inputLineToInteger(line string) int {
	i, _ := strconv.Atoi(line)
	return i
}
