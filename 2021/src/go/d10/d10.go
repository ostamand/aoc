package d10

import (
	"bufio"
	"fmt"
	"os"
	"sort"
)

func getData(path string) [][]int {
	var data [][]int
	f, _ := os.Open(path)
	defer f.Close()
	scanner := bufio.NewScanner(f)
	for scanner.Scan() {
		text := scanner.Text()
		var line []int
		for _, c := range text {
			line = append(line, int(c))
		}
		data = append(data, line)
	}
	return data
}

func isIn(list []int, code int) (bool, int) {
	for idx, l := range list {
		if l == code {
			return true, idx
		}
	}
	return false, 0
}

func lint(line []int) (bool, []int) {
	var opened []int

	var openCodes []int
	for _, c := range "([{<" {
		openCodes = append(openCodes, int(c))
	}

	var closeCodes []int
	for _, c := range ")]}>" {
		closeCodes = append(closeCodes, int(c))
	}

	for _, c := range line {
		// check for open codes
		if flag, _ := isIn(openCodes, c); flag {
			opened = append(opened, c)
		}
		// check for close codes
		if flag, idx := isIn(closeCodes, c); flag {
			// check if fits with latest open code
			if openCodes[idx] == opened[len(opened)-1] {
				// ok remove from opened
				opened = opened[:len(opened)-1]
			} else {
				return true, []int{c}
			}
		}
	}

	var fill []int
	for i := len(opened) - 1; i >= 0; i-- {
		_, idx := isIn(openCodes, opened[i])
		fill = append(fill, closeCodes[idx])
	}

	return false, fill
}

func Solve(path string, part int) {
	lines := getData(path)
	switch part {
	case 1:
		points := make(map[int]int)

		p := [...]int{3, 57, 1197, 25137}
		for i, c := range ")]}>" {
			points[int(c)] = p[i]
		}

		totalPoints := 0
		for _, line := range lines {
			corrupted, codes := lint(line)
			if corrupted {
				totalPoints += points[codes[0]]
			}
		}

		fmt.Printf("Total syntax error: %d\n", totalPoints)

	case 2:
		points := make(map[int]int)

		for i, c := range ")]}>" {
			points[int(c)] = i + 1
		}

		var scores []int
		for _, line := range lines {
			corrupted, codes := lint(line)
			if !corrupted {
				score := 0
				for _, code := range codes {
					score *= 5
					score += points[code]
				}
				scores = append(scores, score)
			}
		}
		sort.Ints(scores)
		fmt.Printf("Middle score: %d\n", scores[int(len(scores)/2)])
	}
}
