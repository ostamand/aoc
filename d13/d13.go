package d13

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"

	"github.com/ostamand/aoc21/helpers"
)

func getData(path string) (grid [][]int, folds [][]int) {
	f, _ := os.Open(path)
	defer f.Close()

	scanner := bufio.NewScanner(f)

	var dots [][]int
	maxs := [...]int{0, 0}
	flag := false
	for scanner.Scan() {
		text := scanner.Text()
		if strings.Compare(text, "") == 0 {
			flag = true
		} else if !flag {
			splits := strings.Split(text, ",")
			coord := make([]int, 2)
			for i, s := range splits {
				// x, y
				if x, err := strconv.Atoi(s); err == nil {
					coord[i] = x
					if x > maxs[i] {
						maxs[i] = x
					}
				}
			}
			dots = append(dots, coord)
		} else {
			// read folds
			text = strings.Replace(text, "fold along ", "", -1)
			splits := strings.Split(text, "=")
			data := make([]int, 2)
			switch splits[0] {
			case "x":
				data[0] = 0
			case "y":
				data[0] = 1
			}
			i, _ := strconv.Atoi(splits[1])
			data[1] = i
			folds = append(folds, data)
		}
	}

	// build grid
	grid = make([][]int, maxs[1]+1)
	for i := range grid {
		grid[i] = make([]int, maxs[0]+1)
	}

	// write dots
	for _, coord := range dots {
		grid[coord[1]][coord[0]] = 1
	}

	return
}

/*
To apply a fold
copy to line/column x to 0 moving away from fold line
*/

func addTo(array *[]int, x []int) {
	// assuming the sizes are OK
	for i := range x {
		(*array)[i] += x[i]
	}
}

func fold(grid *[][]int, foldAxis int, foldLine int) {
	switch foldAxis {
	// fold along y
	case 1:
		for i := foldLine + 1; i < len(*grid); i++ {
			idx := foldLine - (i - foldLine)
			addTo(&(*grid)[idx], (*grid)[i])
		}
		*grid = (*grid)[:foldLine]
	// fold along x
	case 0:
		for _, line := range *grid {
			for i := foldLine + 1; i < len(line); i++ {
				idx := foldLine - (i - foldLine)
				line[idx] += line[i]
			}
		}
		// remove columns
		for i := range *grid {
			(*grid)[i] = (*grid)[i][:foldLine]
		}
	}

	// cleanup
	for _, line := range *grid {
		for i, x := range line {
			if x > 1 {
				line[i] = 1
			}
		}
	}
}

func Solve(path string, part int) {
	grid, folds := getData(path)
	fmt.Println("Original grid:")
	//helpers.PrintGrid(grid)
	//GridToImage(grid, 5, "images/d13_original.png")
	for i := range folds {
		fold(&grid, folds[i][0], folds[i][1])

		sum := helpers.SumOfGrid(grid)

		fmt.Printf("After fold %d:\n", i+1)
		fmt.Printf("# of dots: %d\n", sum)
	}
	helpers.GridToImage(grid, 10, "images/d13_final.png")
}
