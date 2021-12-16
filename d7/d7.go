package d7

import (
	"fmt"
	"math"

	"github.com/ostamand/aoc21/helpers"
)

func Solve(path string, part int) {
	positions := helpers.ReadInts(path)

	n := helpers.MinInts(positions)
	m := helpers.MaxInts(positions)

	leastFuel := math.MaxInt
	for i := n; i <= m; i++ {
		fuel := 0
		for _, p := range positions {

			distance := helpers.AbsInt(p - i)
			if part == 1 {
				fuel += distance
			} else if part == 2 {
				for i := 1; i <= distance; i++ {
					fuel += i
				}
			}
		}
		if leastFuel > fuel {
			leastFuel = fuel
		}
	}
	fmt.Printf("Minimum fuel to align: %d\n", leastFuel)
}
