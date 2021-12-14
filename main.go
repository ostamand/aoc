package main

import (
	"flag"

	"github.com/ostamand/aoc21/d1"
	"github.com/ostamand/aoc21/d2"
	"github.com/ostamand/aoc21/d3"
	"github.com/ostamand/aoc21/d4"
)

/*

Day 1:
	Part 1: go build . && ./aoc21 -d 1 -p 1 -data inputs/day1.txt
	Part 2: go build . && ./aoc21 -d 1 -p 2 -data inputs/day1.txt
Day 2:
	Part 1: go build . && ./aoc21 -d 2 -p 1 -data inputs/day2.txt
	Part 2: go build . && ./aoc21 -d 2 -p 2 -data inputs/day2.txt
Day 3:
	Part 1: go build . && ./aoc21 -d 3 -p 1 -data inputs/day3.txt
	Part 2: go build . && ./aoc21 -d 3 -p 2 -data inputs/day3.txt
Day 4:
	Part 1: go build . && ./aoc21 -d 4 -p 1 -data inputs/day4.txt

*/
func main() {
	var d = flag.Int("d", 1, "day number")
	var inputPath = flag.String("data", "inputs/day1.txt", "path to input file")
	var part = flag.Int("p", 1, "part 1 or 2 of the problem")
	flag.Parse()

	switch *d {
	case 1:
		d1.Solve(*inputPath, *part)
	case 2:
		d2.Solve(*inputPath, *part)
	case 3:
		d3.Solve(*inputPath, *part)
	case 4:
		d4.Solve(*inputPath, *part)
	}
}
