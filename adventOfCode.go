package main

import (
	twentyone "github.com/robinthues/advent-of-code/2021"
	"github.com/robinthues/advent-of-code/helpers"
)

func main() {
	lines, err := helpers.LoadInput("2021/input-2021-5.txt")
	if err != nil {
		println(err)
	}

	twentyone.D5one(lines)
}
