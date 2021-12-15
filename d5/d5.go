package d5

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

type option int

const (
	straightOnly option = 0
	all                 = 1
)

type line struct {
	p1 [2]int
	p2 [2]int
}

func getData(path string) []line {
	var lines []line

	f, _ := os.Open(path)
	defer f.Close()

	// function to go from x,y to line coordinates
	toCoord := func(text string) [2]int {
		var coord [2]int
		splits := strings.Split(strings.TrimSpace(text), ",")
		for i := range [...]int{1, 2} {
			if v, err := strconv.Atoi(splits[i]); err == nil {
				coord[i] = v
			}
		}
		return coord
	}

	scanner := bufio.NewScanner(f)
	for scanner.Scan() {
		text := scanner.Text()
		splits := strings.Split(text, "->")
		l := line{toCoord(splits[0]), toCoord(splits[1])}
		lines = append(lines, l)
	}
	return lines
}

func maxOf(x1 int, x2 int) int {
	if x1 > x2 {
		return x1
	}
	return x2
}

func minOf(x1 int, x2 int) int {
	if x1 < x2 {
		return x1
	}
	return x2
}

func getSizeOfGrid(lines []line) ([2]int, [2]int) {
	var (
		x [2]int
		y [2]int
	)

	// a bit hacky, could be done cleaner
	for _, line := range lines {
		x[0] = minOf(line.p1[0], x[0])
		x[0] = minOf(line.p2[0], x[0])

		x[1] = maxOf(line.p1[0], x[1])
		x[1] = maxOf(line.p2[0], x[1])

		y[0] = minOf(line.p1[1], y[0])
		y[0] = minOf(line.p2[1], y[0])

		y[1] = maxOf(line.p1[1], y[1])
		y[1] = maxOf(line.p2[1], y[1])
	}
	return x, y
}

func isStraightLine(l line) bool {
	if l.p1[0] == l.p2[0] || l.p1[1] == l.p2[1] {
		return true
	}
	return false
}

func pointsFromLine(l line, o option) [][]int {
	var points [][]int

	if x := l.p1[0]; x == l.p2[0] {
		// horizontal line
		i0 := minOf(l.p1[1], l.p2[1])
		i1 := maxOf(l.p1[1], l.p2[1])
		for i := i0; i <= i1; i++ {
			points = append(points, []int{x, i})
		}
	} else if y := l.p1[1]; y == l.p2[1] {
		// vertical line
		i0 := minOf(l.p1[0], l.p2[0])
		i1 := maxOf(l.p1[0], l.p2[0])
		for i := i0; i <= i1; i++ {
			points = append(points, []int{i, y})
		}
	} else if o == all {
		// diagonal line
		slopeX := -1
		slopeY := -1
		if l.p2[0] >= l.p1[0] {
			slopeX = 1
		}
		if l.p2[1] >= l.p1[1] {
			slopeY = 1
		}

		i := 0
		for {
			p := []int{l.p1[0] + (i * slopeX), l.p1[1] + (i * slopeY)}
			points = append(points, p)
			if p[0] == l.p2[0] {
				if p[0] == 991 || p[1] == 991 {
					fmt.Println("test")
				}
				break
			}
			i++
		}
	}

	return points
}

func Solve(path string, part int) {
	var o option

	switch part {
	case 1:
		o = straightOnly
	case 2:
		o = all
	}

	lines := getData(path)

	gridX, gridY := getSizeOfGrid(lines)

	// build grid
	maxY := gridY[1] + 1
	maxX := gridX[1] + 1

	grid := make([][]int, maxY)
	for i := 0; i < maxY; i++ {
		grid[i] = make([]int, maxX)
	}

	for _, line := range lines {
		points := pointsFromLine(line, o)
		// 590
		for _, p := range points {
			grid[p[1]][p[0]] += 1
		}
	}

	// the number of points where at least two lines overlap
	overlaps := 0
	for _, line := range grid {
		for _, x := range line {
			if x > 1 {
				overlaps++
			}
		}
	}

	fmt.Printf("Grid X: [%d %d]\n", gridX[0], gridX[1])
	fmt.Printf("Grid Y: [%d %d]\n", gridY[0], gridY[1])

	//fmt.Printf("Grid\n")
	//helpers.PrintGrid(grid)

	fmt.Printf("# Overlaps: %d\n", overlaps)

}
