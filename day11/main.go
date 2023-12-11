package main

import (
	"fmt"
	"math"
	"os"
	"slices"
	"strings"
)

type Galaxy struct {
	row int
	col int
}

func getEmptyRows(lines []string) []int {
	indices := []int{}
	for index, line := range lines {
		noGalaxies := true
		for _, char := range line {
			if char == '#' {
				noGalaxies = false
				break
			}
		}
		if noGalaxies {
			indices = append(indices, index)
		}
	}
	return indices
}

func getEmptyCols(lines []string) []int {
	indices := []int{}
	for i := 0; i < len(lines[0]); i++ {
		noGalaxies := true
		for row := range lines {
			if lines[row][i] == '#' {
				noGalaxies = false
				break
			}
		}
		if noGalaxies {
			indices = append(indices, i)
		}
	}
	return indices
}

func main() {
	inputstring, _ := os.ReadFile("input.txt")
	lines := strings.Split(string(inputstring), "\n")
	expandingSpace := int64(1000000)
	galaxies := []Galaxy{}
	for row, line := range lines {
		for col, char := range line {
			if char == '#' {
				galaxies = append(galaxies, Galaxy{row: row, col: col})
			}
		}
	}
	sum := int64(0)
	emptyRows := getEmptyRows(lines)
	emptyCols := getEmptyCols(lines)
	for i := 0; i < len(galaxies); i++ {
		for galNum := i; galNum < len(galaxies); galNum++ {
			for rowWalk := galaxies[i].row; rowWalk < galaxies[galNum].row; rowWalk++ {
				if slices.Contains(emptyRows, rowWalk) {
					sum += expandingSpace
				} else {
					sum += 1
				}
			}
			col1 := int(math.Max(float64(galaxies[i].col), float64(galaxies[galNum].col)))
			col2 := int(math.Min(float64(galaxies[i].col), float64(galaxies[galNum].col)))
			for colWalk := col2; colWalk < col1; colWalk++ {
				if slices.Contains(emptyCols, colWalk) {
					sum += expandingSpace
				} else {
					sum += 1
				}
			}
		}
	}

	fmt.Println(sum)
}
