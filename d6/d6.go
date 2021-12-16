package d6

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func getData(path string) []int {
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

func sumSlice(slice []int) int {
	sum := 0
	for _, v := range slice {
		sum += v
	}
	return sum
}

func Solve(path string, part int) {

	var nDays int
	switch part {
	case 1:
		nDays = 80
	case 2:
		nDays = 256
	}

	lifes := getData(path)

	count := make([]int, 9)
	for _, l := range lifes {
		count[l]++
	}

	for n := 1; n <= nDays; n++ {
		prev := make([]int, 9)
		copy(prev, count)

		// move left
		for i := 0; i < 8; i++ {
			count[i] = prev[i+1]
		}
		// 0 becomes a 6 and adds a new 8
		count[6] += prev[0]
		count[8] = prev[0]

		//fmt.Printf("After %2d days: %v\n", n, count)
	}
	fmt.Printf("Number of fish at the end: %d\n", sumSlice(count))
}
