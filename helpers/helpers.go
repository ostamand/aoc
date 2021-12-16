package helpers

import (
	"bufio"
	"fmt"
	"math"
	"os"
	"strconv"
	"strings"
)

func PrintGrid(grid [][]int) {
	for _, line := range grid {
		for _, x := range line {
			fmt.Printf("%4d", x)
		}
		fmt.Println()
	}
}

func ReadInts(path string) []int {
	var data []int

	f, _ := os.Open(path)
	defer f.Close()

	scanner := bufio.NewScanner(f)
	for scanner.Scan() {
		// only one line in he file
		text := scanner.Text()
		splits := strings.Split(text, ",")
		for _, s := range splits {
			if i, err := strconv.Atoi(s); err == nil {
				data = append(data, i)
			}
		}
	}
	return data
}

func MaxInts(ints []int) int {
	max := math.MinInt
	for _, v := range ints {
		if v > max {
			max = v
		}
	}
	return max
}

func MinInts(ints []int) int {
	min := math.MaxInt
	for _, v := range ints {
		if v < min {
			min = v
		}
	}
	return min
}

func AbsInt(x int) int {
	if x < 0 {
		return -x
	}
	return x
}
