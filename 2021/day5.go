package twentyone

import (
	"math"
	"strconv"
	"strings"
)

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}

func dist(a, b int) int {
	return int(math.Abs(float64(b - a)))
}

/*
func markLine(x1, y1, x2, y2 int, oceanFloor [][]int) {
	xDist := math.Abs(float64(x2 - x1))
	yDist := math.Abs(float64(y2 - y1))
	xMin := min(x1, x2)
	yMin := min(y1, y2)
}
*/

const size = 1000

func D5one(lines []string) {
	var oceanFloor [][]int

	for i := 0; i < size; i++ {
		oceanFloorRow := make([]int, size)
		oceanFloor = append(oceanFloor, oceanFloorRow)
	}
	println("OceanFloor initiated")

	for _, l := range lines {
		lineSplitted := strings.Split(l, " -> ")
		startSplitted := strings.Split(lineSplitted[0], ",")
		x1, _ := strconv.Atoi(startSplitted[0])
		y1, _ := strconv.Atoi(startSplitted[1])

		endSplitted := strings.Split(lineSplitted[1], ",")
		x2, _ := strconv.Atoi(endSplitted[0])
		y2, _ := strconv.Atoi(endSplitted[1])

		xDist := dist(x2, x1)
		yDist := dist(y2, y1)

		if xDist > 0 && yDist == xDist {
			/*for d := 0; d <= xDist; d++ {
				xChange := 0
				if x1 > x2 {
					xChange = 1
				} else {
					xChange = -1
				}
			}*/
		}

		if xDist == 0 {
			yMin := min(y1, y2)
			for d := 0; d <= yDist; d++ {
				oceanFloor[yMin+d][x1] += 1

			}
		}
		if yDist == 0 {
			xMin := min(x1, x2)
			for d := 0; d <= xDist; d++ {
				oceanFloor[y1][xMin+d] += 1
			}
		}
	}

	cnt := 0
	for i := 0; i < size; i++ {
		for j := 0; j < size; j++ {
			if oceanFloor[i][j] >= 2 {
				cnt += 1

			}
			//fmt.Print(oceanFloor[i][j])
		}
		//println()
	}
	println("Count: ", cnt)
}
