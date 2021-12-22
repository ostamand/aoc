package helpers

import (
	"bufio"
	"fmt"
	"image"
	"image/color"
	"image/png"
	"math"
	"os"
	"strconv"
	"strings"
)

func PrintGrid(grid [][]int) {
	for _, line := range grid {
		for _, x := range line {
			fmt.Printf("%4d", x)
		}
		fmt.Println()
	}
}

func ReadInts(path string) []int {
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

func MaxInts(ints []int) int {
	max := math.MinInt
	for _, v := range ints {
		if v > max {
			max = v
		}
	}
	return max
}

func MinInts(ints []int) int {
	min := math.MaxInt
	for _, v := range ints {
		if v < min {
			min = v
		}
	}
	return min
}

func SumOfGrid(ints [][]int) (sum int) {
	for _, line := range ints {
		for _, x := range line {
			sum += x
		}
	}
	return
}

func AbsInt(x int) int {
	if x < 0 {
		return -x
	}
	return x
}

func GridToImage(grid [][]int, cellSize int, outputFile string) {
	height := len(grid)*cellSize + 1
	width := len(grid[0])*cellSize + 1

	// define image with rectangle
	upLeft := image.Point{0, 0}
	lowRight := image.Point{width, height}
	img := image.NewRGBA(image.Rectangle{upLeft, lowRight})

	// set color of each pixel
	for x := 0; x < width; x++ {
		for y := 0; y < height; y++ {
			i := y / cellSize
			j := x / cellSize

			c := color.RGBA{0, 0, 0, 0xff}
			if x%cellSize == 0 || y%cellSize == 0 {
				c = color.RGBA{100, 200, 200, 0xff}
			} else if grid[i][j] > 0 {
				c = color.RGBA{255, 255, 255, 0xff}
			}
			img.Set(x, y, c)
		}
	}
	f, _ := os.Create(outputFile)
	png.Encode(f, img)
}
