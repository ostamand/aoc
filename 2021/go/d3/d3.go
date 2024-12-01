package d3

import (
	"bufio"
	"fmt"
	"math"
	"os"
	"strings"
)

type rating int
type calculationType int

const (
	mostCommon  calculationType = 0
	leastCommon                 = 1
)

const (
	oxygen rating = 0
	co2           = 1
)

func getData(path string) [][]int {
	data := make([][]int, 0)

	f, _ := os.Open(path)
	defer f.Close()

	scanner := bufio.NewScanner(f)
	for scanner.Scan() {
		text := strings.TrimSpace(scanner.Text())
		bits := make([]int, len(text))
		for i, c := range text {
			switch c {
			case '1':
				bits[i] = 1
			case '0':
				bits[i] = 0
			}
		}
		data = append(data, bits)
	}
	return data
}

func applyDiagnostic(measurements [][]int, c calculationType) []int {
	sum := make([]int, len(measurements[0]))
	for _, number := range measurements {
		for i, x := range number {
			sum[i] += x
		}
	}
	n := len(measurements)
	ints := make([]int, len(sum))
	for i, x := range sum {
		switch c {
		case mostCommon:
			// if 0 and 1 are equally common, keep values with a 1
			if x >= n-x {
				ints[i] = 1
			}
		case leastCommon:
			// if 0 and 1 are equally common, keep values with a 0
			if x >= n-x {
				ints[i] = 0
			} else {
				ints[i] = 1
			}
		}
	}
	return ints
}

func invertBits(ints []int) []int {
	intsInverted := make([]int, len(ints))
	for i, x := range ints {
		switch x {
		case 0:
			intsInverted[i] = 1
		case 1:
			intsInverted[i] = 0
		}
	}
	return intsInverted
}

func binaryToDecimal(bits []int) int {
	decimal := 0
	n := len(bits) - 1
	for i := 0; i <= n; i++ {
		decimal += bits[i] * int(math.Pow(2, float64(n-i)))
	}
	return decimal
}

func getRating(bits [][]int, r rating) int {
	value := 0
	for i := 0; i < len(bits[0]); i++ {
		var ints []int

		switch r {
		case oxygen:
			ints = applyDiagnostic(bits, mostCommon)
		case co2:
			ints = applyDiagnostic(bits, leastCommon)
		}

		toKeep := make([][]int, 0)
		for _, b := range bits {
			if b[i] == ints[i] {
				toKeep = append(toKeep, b)
			}
		}

		if len(toKeep) == 1 {
			value = binaryToDecimal(toKeep[0])
			break
		} else {
			bits = toKeep
		}
	}
	return value
}

func Solve(path string, part int) {
	switch part {
	case 1:
		data := getData(path)

		intsMost := applyDiagnostic(data, mostCommon)
		intsLeast := applyDiagnostic(data, leastCommon)

		gammaRate := binaryToDecimal(intsMost)
		epsilonRate := binaryToDecimal(intsLeast)

		fmt.Printf("gamma rate * epsilon rate = %d\n", gammaRate*epsilonRate)

	case 2:
		data := getData(path)
		oxygenRating := getRating(data, oxygen)
		co2Rating := getRating(data, co2)

		fmt.Printf("Oxygen rating: %d\n", oxygenRating)
		fmt.Printf("C02 rating: %d\n", co2Rating)
		fmt.Printf("Life support rating: %d\n", oxygenRating*co2Rating)
	}
}
