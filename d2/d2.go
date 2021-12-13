package d2

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

// forward, down, up

type command struct {
	name  string
	units int
}

func getData(path string) []command {
	commands := make([]command, 0)

	f, _ := os.Open(path)
	defer f.Close()

	scanner := bufio.NewScanner(f)
	for scanner.Scan() {
		fields := strings.Fields(scanner.Text())
		value, _ := strconv.Atoi(fields[1])
		c := command{name: fields[0], units: value}
		commands = append(commands, c)
	}

	return commands
}

func moveFrom(origin []int, vector []int, units int) []int {
	for i := 0; i <= 1; i++ {
		origin[i] += vector[i] * units
	}
	return origin
}

func Solve(path string, part int) {
	switch part {

	case 1:
		commands := getData(path)

		// horizontal position, depth
		position := []int{0, 0}
		moves := map[string][]int{"forward": {1, 0}, "down": {0, 1}, "up": {0, -1}}

		for _, c := range commands {
			position = moveFrom(position, moves[c.name], c.units)
		}

		fmt.Printf("Position after planned course: %v\n", position)
		fmt.Printf("Final horizontal * final depth: %d\n", position[0]*position[1])
	}
}
