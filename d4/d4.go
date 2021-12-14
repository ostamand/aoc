package d4

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func lineToArray(line string, sep string) []int {
	var values []int
	splits := strings.Split(line, sep)
	for _, s := range splits {
		if v, err := strconv.Atoi(s); err == nil {
			values = append(values, v)
		}
	}
	return values
}

func printGrid(grid [][]int) {
	for _, line := range grid {
		for _, x := range line {
			fmt.Printf("%4d", x)
		}
		fmt.Println()
	}
}

func getData(path string) ([]int, [][][]int) {
	var randomNumbers []int
	var grids [][][]int

	f, _ := os.Open(path)
	defer f.Close()

	// some flags to manage reading
	i := 0
	readingGrid := false
	var grid [][]int

	scanner := bufio.NewScanner(f)
	for scanner.Scan() {

		switch i {
		case 0:
			// first line contains the random numbers
			randomNumbers = lineToArray(scanner.Text(), ",")
			i++
		default:
			text := scanner.Text()

			// new grid starts
			if !readingGrid && text != "" {
				grid = make([][]int, 0)
				readingGrid = true
			}

			// grid is done
			if readingGrid && text == "" {
				grids = append(grids, grid)
				readingGrid = false
			}

			if readingGrid && text != "" {
				grid = append(grid, lineToArray(text, " "))
			}
		}
	}

	// add the last grid
	grids = append(grids, grid)

	return randomNumbers, grids
}

func isIn(x int, list []int) bool {
	for _, i := range list {
		if i == x {
			return true
		}
	}
	return false
}

func allIn(x []int, in []int) bool {
	flag := true
	for _, x := range x {
		flag = flag && isIn(x, in)
	}
	return flag
}

func checkWin(grid [][]int, marked []int) bool {

	// check horizontals
	for _, line := range grid {
		if win := allIn(line, marked); win {
			return true
		}
	}

	// check columns
	nColumns := len(grid[0])
	for j := 0; j < nColumns; j++ {

		var column []int
		for _, line := range grid {
			column = append(column, line[j])
		}

		if win := allIn(column, marked); win {
			return true
		}
	}

	return false
}

func sumAllNotIn(grid [][]int, in []int) int {
	// get list of numbers not in
	var notIn []int
	for _, line := range grid {
		for _, x := range line {
			if !isIn(x, in) {
				notIn = append(notIn, x)
			}
		}
	}

	// sum the list
	sum := 0
	for _, x := range notIn {
		sum += x
	}

	return sum
}

func Solve(path string, part int) {
	randomNumbers, grids := getData(path)

	switch part {
	case 1:
		var currentlyMarked []int
		var winningIdx int
		var winningNumber int
		var winningSum int
		for _, r := range randomNumbers {
			currentlyMarked = append(currentlyMarked, r)

			// check if one of the grid wins
			isWinning := false
			for i, grid := range grids {
				if isWinning = checkWin(grid, currentlyMarked); isWinning {
					winningIdx = i
					winningNumber = r
					break
				}
			}
			if isWinning {
				winningSum = sumAllNotIn(grids[winningIdx], currentlyMarked)
				break
			}
		}
		fmt.Printf("Winning board: %d\n", winningIdx+1)
		fmt.Printf("Winning number: %d\n", winningNumber)
		fmt.Printf("Winning sum: %d\n", winningSum)
		fmt.Printf("Final score: %d\n", winningSum*winningNumber)

	case 2:
		var currentlyMarked []int
		var loosingNumber int
		var loosingSum int

		for _, r := range randomNumbers {
			currentlyMarked = append(currentlyMarked, r)

			// check if one of the grid wins
			var loosingGrids [][][]int
			for _, grid := range grids {

				isWinning := checkWin(grid, currentlyMarked)

				if !isWinning {
					loosingGrids = append(loosingGrids, grid)
				}

				if isWinning && len(grids) == 1 {
					// last grid finally wins
					loosingNumber = r
					loosingSum = sumAllNotIn(grids[0], currentlyMarked)
					break
				}
			}
			grids = loosingGrids
			if len(grids) == 0 {
				break
			}
		}
		fmt.Printf("Loosing number: %d\n", loosingNumber)
		fmt.Printf("Loosing sum: %d\n", loosingSum)
		fmt.Printf("Final score: %d\n", loosingSum*loosingNumber)
	}
}
