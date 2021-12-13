package d1

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

func sum(l []int) int {
	sum := 0
	for _, x := range l {
		sum += x
	}
	return sum
}

func getData(path string) []int {
	// read input file
	data := make([]int, 0)

	f, _ := os.Open(path)
	defer f.Close()

	scanner := bufio.NewScanner(f)
	for scanner.Scan() {
		if i, err := strconv.Atoi(scanner.Text()); err == nil {
			data = append(data, i)
		}
	}

	return data
}

func Solve(inputPath string, part int) {
	data := getData(inputPath)

	if part == 1 {
		// check measurements greater than previous
		count := 0
		for i := 1; i < len(data); i++ {
			if data[i] > data[i-1] {
				count++
			}
		}
		fmt.Printf("Measurements are larger than the previous measurement: %d\n", count)
	} else if part == 2 {

		count := 0
		for i := 0; i < len(data)-3; i++ {

			m1 := data[i : i+3]
			m2 := data[i+1 : i+4]

			if sum(m2) > sum(m1) {
				count++
			}
		}
		fmt.Printf("Sums larger than the previous sum: %d\n", count)
	}
}
