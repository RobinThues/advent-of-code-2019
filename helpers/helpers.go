package helpers

import (
	"bufio"
	"os"
	"strconv"
)

func StringSliceToIntSlice(s []string) ([]int, error) {
	var intSlice []int

	for _, l := range s {
		i, err := strconv.Atoi(l)
		if err != nil {
			return nil, err
		}
		intSlice = append(intSlice, i)
	}
	return intSlice, nil
}

func LoadInput(fileName string) ([]string, error) {
	file, err := os.Open(fileName)
	if err != nil {
		println("File does not exist: ", fileName)
		return nil, err
	}

	defer func() {
		err := file.Close()
		if err != nil {
			println("File not closeable")
			//log.Fatal(err)
		}
	}()

	var lines []string
	s := bufio.NewScanner(file)
	for s.Scan() {
		lines = append(lines, s.Text())
	}
	return lines, nil
}
