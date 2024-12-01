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

func moveWithAim(origin []int, vector []int, units int) []int {
	// update depth
	origin[0] += vector[0] * units
	// check if moving forward
	if vector[1] == 1 {
		origin[1] += vector[1] * units // horizontal position
		origin[2] += origin[0] * units // depth
	}
	return origin
}

func Solve(path string, part int) {
	commands := getData(path)

	switch part {

	case 1:
		// horizontal position, depth
		position := []int{0, 0}
		moves := map[string][]int{"forward": {1, 0}, "down": {0, 1}, "up": {0, -1}}

		for _, c := range commands {
			position = moveFrom(position, moves[c.name], c.units)
		}

		fmt.Printf("Position after planned course: %v\n", position)
		fmt.Printf("Final horizontal * final depth: %d\n", position[0]*position[1])

	case 2:
		// aim, horizontal position, depth
		position := []int{0, 0, 0}
		moves := map[string][]int{"forward": {0, 1, 0}, "down": {1, 0, 0}, "up": {-1, 0, 0}}

		for _, c := range commands {
			position = moveWithAim(position, moves[c.name], c.units)
		}

		fmt.Printf("Position after planned course: %v\n", position[1:])
		fmt.Printf("Final horizontal * final depth: %d\n", position[1]*position[2])
	}
}
