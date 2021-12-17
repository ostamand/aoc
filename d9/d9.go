package d9

import (
	"bufio"
	"fmt"
	"os"
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

func Solve(path string, part int) {
	switch part {
	case 1:
		grid := getData(path)
		riskLevel := 0
		//helpers.PrintGrid(grid)
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
					riskLevel += value + 1
				}
			}
		}
		fmt.Printf("Sum of the risk level: %d\n", riskLevel)
	}
}
