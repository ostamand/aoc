package main

import (
	"flag"

	"github.com/ostamand/aoc/2021/src/go/d1"
	"github.com/ostamand/aoc/2021/src/go/d10"
	"github.com/ostamand/aoc/2021/src/go/d12"
	"github.com/ostamand/aoc/2021/src/go/d13"
	"github.com/ostamand/aoc/2021/src/go/d14"
	"github.com/ostamand/aoc/2021/src/go/d2"
	"github.com/ostamand/aoc/2021/src/go/d3"
	"github.com/ostamand/aoc/2021/src/go/d4"
	"github.com/ostamand/aoc/2021/src/go/d5"
	"github.com/ostamand/aoc/2021/src/go/d6"
	"github.com/ostamand/aoc/2021/src/go/d7"
	"github.com/ostamand/aoc/2021/src/go/d9"
)

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
	case 5:
		d5.Solve(*inputPath, *part)
	case 6:
		d6.Solve(*inputPath, *part)
	case 7:
		d7.Solve(*inputPath, *part)
	case 9:
		d9.Solve(*inputPath, *part)
	case 10:
		d10.Solve(*inputPath, *part)
	case 12:
		d12.Solve(*inputPath, *part)
	case 13:
		d13.Solve(*inputPath, *part)
	case 14:
		d14.Solve(*inputPath, *part)
	}
}
