package d3

import (
	"bufio"
	"fmt"
	"math"
	"os"
	"strings"
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

func mostCommon(measurements [][]int) []int {
	sum := make([]int, len(measurements[0]))
	for _, number := range measurements {
		for i, x := range number {
			sum[i] += x
		}
	}
	// most common bit
	n := len(measurements)
	ints := make([]int, len(sum))
	for i, x := range sum {
		if x > n-x {
			ints[i] = 1
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

func Solve(path string, part int) {
	switch part {
	case 1:
		data := getData(path)

		intsMost := mostCommon(data)
		intsLeast := invertBits(intsMost)

		gammaRate := binaryToDecimal(intsMost)
		epsilonRate := binaryToDecimal(intsLeast)

		fmt.Printf("gamma rate * epsilon rate = %d\n", gammaRate*epsilonRate)
	}
}
