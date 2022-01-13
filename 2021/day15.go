package twentyone

type edge struct {
	from   *vertex
	to     *vertex
	weight int
}

type vertex struct {
	isStart           bool
	isGoal            bool
	visitedAlready    bool
	distanceFromStart int
	edges             []*edge
}

type graph struct {
	vertices []*vertex
	edges    []*edge
}

func newGraphFromIntSlices(ints [][]int) graph {
	g := graph{}
	return g
}

func dijkstra() {

}

func inputToIntSlices(input []string) [][]int {
	var rows [][]int

	for _, line := range input {
		row := make([]int, len(line))
		for k, char := range line {
			row[k] = int(char)
		}
		rows = append(rows, row)
	}
	return rows
}

func D15one(input []string) {
	i := inputToIntSlices(input)
	g := newGraphFromIntSlices(i)
	g = g
}
