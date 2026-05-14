package search

// OxPlow implements the boustrophedon pattern
// Or as the ox plows
func OxPlow(data [][]int, target int) (bool, int) {
	for i, row := range data {
		for j, val := range row {
			if val == target {
				return true, i*len(row) + j
			}
		}
	}
	return false, -1
}
