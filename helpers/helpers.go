package helpers

import "fmt"

func PrintGrid(grid [][]int) {
	for _, line := range grid {
		for _, x := range line {
			fmt.Printf("%4d", x)
		}
		fmt.Println()
	}
}
