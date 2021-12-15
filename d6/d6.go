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

func Solve(path string, part int) {

	var nDays int
	switch part {
	case 1:
		nDays = 80
	case 2:
		nDays = 256
	}

	lifes := getData(path)

	for n := 1; n <= nDays; n++ {
		nToAdd := 0
		for idx, value := range lifes {
			if value == 0 {
				nToAdd++
				lifes[idx] = 6
			} else {
				lifes[idx] = value - 1
			}
		}
		for i := 0; i < nToAdd; i++ {
			lifes = append(lifes, 8)
		}
		//fmt.Printf("After %2d days: %v\n", n, lifes)
	}
	fmt.Printf("Number of fish at the end: %d\n", len(lifes))
}
