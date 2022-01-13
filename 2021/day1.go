package twentyone

import "github.com/robinthues/advent-of-code/helpers"

func D1(lines []string) {
	intLines, err := helpers.StringSliceToIntSlice(lines)
	if err != nil {
		println(err)
	}

	//current := slidingWindowSum(intLines[0], intLines[1], intLines[2])
	current := intLines[0]
	var last int
	numberOfIncreases := 0

	for k, _ := range intLines {
		if k > (len(intLines) - 1) {
			continue
		}

		last = current
		//current = slidingWindowSum(intLines[k], intLines[k+1], intLines[k+2])
		current = intLines[k]

		if current > last {
			numberOfIncreases += 1
		}

		println("k: ", k, "Last: ", last, "Current:", current, "Increase:", numberOfIncreases)
	}
	println(numberOfIncreases)
}

func slidingWindowSum(a, b, c int) int {
	return a + b + c
}
