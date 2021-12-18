package d9

import (
	"bufio"
	"fmt"
	"os"
	"sort"
	"strconv"
)

func getData(path string) [][]int {
	var grid [][]int

	f, _ := os.Open(path)
	defer f.Close()

	scanner := bufio.NewScanner(f)
	for scanner.Scan() {
		text := scanner.Text()
		var line []int
		for _, c := range text {
			if i, err := strconv.Atoi(string(c)); err == nil {
				line = append(line, i)
			}
		}
		grid = append(grid, line)
	}
	return grid
}

func getLowPoints(grid [][]int) [][]int {
	var lowPoints [][]int

	for i := 0; i < len(grid); i++ {
		line := grid[i]
		for j := 0; j < len(line); j++ {
			value := line[j]
			isLower := true
			// left
			if j-1 >= 0 && line[j-1] <= value {
				isLower = false
			}
			// right
			if j+1 < len(line) && line[j+1] <= value {
				isLower = false
			}
			// up
			if i-1 >= 0 && grid[i-1][j] <= value {
				isLower = false
			}
			// down
			if i+1 < len(grid) && grid[i+1][j] <= value {
				isLower = false
			}
			if isLower {
				lowPoints = append(lowPoints, []int{i, j})
			}
		}
	}
	return lowPoints
}

func isNotIn(list *[][]int, i int, j int) bool {
	for _, l := range *list {
		if l[0] == i && l[1] == j {
			return false
		}
	}
	return true
}

func topThree(list []int) []int {
	sort.Ints(list)

	result := []int{0, 0, 0}
	for _, l := range list {
		if l >= result[0] {
			c := []int{0, 0, 0}
			copy(c, result)
			result[1] = c[0]
			result[2] = c[1]
			result[0] = l
		}
	}
	return result
}

func floodFill(grid [][]int, i int, j int, filled *[][]int) {
	*filled = append(*filled, []int{i, j})
	// left
	if idx := j - 1; idx >= 0 && grid[i][idx] != 9 && isNotIn(filled, i, idx) {
		floodFill(grid, i, idx, filled)
	}
	// right
	if idx := j + 1; idx < len(grid[i]) && grid[i][idx] != 9 && isNotIn(filled, i, idx) {
		floodFill(grid, i, idx, filled)
	}
	// up
	if idx := i - 1; idx >= 0 && grid[idx][j] != 9 && isNotIn(filled, idx, j) {
		floodFill(grid, idx, j, filled)
	}
	// down
	if idx := i + 1; idx < len(grid) && grid[idx][j] != 9 && isNotIn(filled, idx, j) {
		floodFill(grid, idx, j, filled)
	}
}

func Solve(path string, part int) {
	grid := getData(path)
	lowPoints := getLowPoints(grid)

	switch part {
	case 1:
		riskLevel := 0
		for _, lowPoint := range lowPoints {
			riskLevel += grid[lowPoint[0]][lowPoint[1]] + 1
		}
		fmt.Printf("Sum of the risk level: %d\n", riskLevel)
	case 2:
		var basins []int
		for _, lowPoint := range lowPoints {
			i := lowPoint[0]
			j := lowPoint[1]
			// recursive flood fill
			var basinIdx [][]int
			floodFill(grid, i, j, &basinIdx)
			if len(basinIdx) > 0 {
				basins = append(basins, len(basinIdx))
			}
		}
		tops := topThree(basins)
		fmt.Printf("All basins sizes: %d\n", basins)
		fmt.Printf("Top three basin: %d\n", tops)
		fmt.Printf("Top three basin multiplied sizes: %d\n", tops[0]*tops[1]*tops[2])
	}
}
