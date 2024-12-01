package d14

import (
	"bufio"
	"fmt"
	"os"
	"sort"
	"strings"
)

func getData(path string) (template []string, pairs map[string]string) {
	pairs = make(map[string]string)

	f, _ := os.Open(path)
	defer f.Close()

	scanner := bufio.NewScanner(f)

	i := 0
	for scanner.Scan() {
		text := scanner.Text()
		switch {
		case i == 0:
			for _, c := range text {
				template = append(template, string(c))
			}
		case i > 1:
			splits := strings.Split(text, "->")
			pairs[strings.TrimSpace(splits[0])] = strings.TrimSpace(splits[1])
		}
		i++
	}
	return
}

func countsPerElement(template []string) (counts map[string]int) {
	counts = make(map[string]int)
	for _, element := range template {
		counts[element]++
	}
	return
}

func simulate(template []string, pairs map[string]string, nSteps int) []string {
	var polymer []string
	for i := 0; i < nSteps; i++ {
		fmt.Printf("Step %d\n", i+1)
		polymer = make([]string, 0)
		for n := 0; n < len(template)-1; n += 1 {
			key := template[n] + template[n+1]
			polymer = append(polymer, template[n], pairs[key])
		}
		polymer = append(polymer, template[len(template)-1])
		template = make([]string, len(polymer))
		copy(template, polymer)
	}
	return polymer
}

func Solve(path string, part int) {
	nStep := 10
	if part > 1 {
		nStep = 40
	}
	template, pairs := getData(path)
	polymer := simulate(template, pairs, nStep)
	elements := countsPerElement(polymer)
	counts := make([]int, len(elements))
	i := 0
	for _, count := range elements {
		counts[i] = count
		i++
	}
	sort.Ints(counts)
	fmt.Printf("After step %d\n", nStep)
	fmt.Printf("Length of polymer: %d\n", len(polymer))
	fmt.Printf("Quantity most - least: %d\n", counts[len(counts)-1]-counts[0])
}
