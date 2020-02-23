package main

func readDay3Input(fileName string) {
	// read line
	// separate by comma
	// first letter is direction
	// rest to int is distance
}

func buildMap() {
	// for each wire
	// add/sub ups and downs, save max ups and max downs
	// add/sub lefts and rights, save max right and max left
	// create map with maxUps+Maxdowns and maxLefts+maxRights
	// mark the middle
}

func drawWiresOnMap() {
	// for the first wire
	// -- walk path and mark with X
	// for the second wire
	// -- walk path, mark with Y if empty, mark with Z if cross
}

func findClosestCrossing() {
	// for each crossing
	// -- calculate horizontal and vertical difference to middle, add up
	// choose crossing with minimum
}
