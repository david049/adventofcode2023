package main

import (
	"fmt"
	"os"
	"slices"
	"strings"
)

type Point struct {
	row int
	col int
}

func getSym(row int, col int, grid [][]rune) rune {
	if row < 0 || row >= len(grid) {
		return '#'
	}
	if col < 0 || col >= len(grid[row]) {
		return '#'
	}
	return rune(grid[row][col])
}

func expandOutward(input [][]rune, intGrid [][]int, row int, col int) {
	intGrid[row][col] -= 1
	if intGrid[row][col] < 0 {
		intGrid[row][col] = 0
	}
	if getSym(row-1, col, input) == '.' {
		intGrid[row-1][col] += 1
	}
	if getSym(row+1, col, input) == '.' {
		intGrid[row+1][col] += 1
	}
	if getSym(row, col+1, input) == '.' {
		intGrid[row][col+1] += 1
	}
	if getSym(row, col-1, input) == '.' {
		intGrid[row][col-1] += 1
	}
}

func checkZeros(input []int) bool {
	for _, num := range input {
		if num != 0 {
			return false
		}
	}
	return true
}

func calculateDifferencesForward(input []int) int {
	var differences []int
	for i := 1; i < len(input); i++ {
		differences = append(differences, input[i]-input[i-1])
	}
	if checkZeros(differences) {
		return input[len(input)-1]
	}
	return calculateDifferencesForward(differences) + input[len(input)-1]
}

func main() {
	inputstring, _ := os.ReadFile("input.txt")
	lines := strings.Split(string(inputstring), "\n")

	multiplier := 5 // run on a larger grid so we can extrapolate
	charArray := make([][]rune, len(lines)*multiplier)
	intGrid := make([][]int, len(lines)*multiplier)
	for j := 0; j < 5; j++ {
		for i, str := range lines {
			newStr := strings.Repeat(str, multiplier)
			runes := []rune(newStr)
			charArray[i+j*len(lines)] = runes
			toAdd := []int{}
			for j := 0; j < len(newStr); j++ {
				toAdd = append(toAdd, 0)
			}
			intGrid[i+j*len(lines)] = toAdd
		}
	}

	startingRows := []int{}
	startingCols := []int{}

	for row, line := range charArray {
		for col := range line {
			if charArray[row][col] == 'S' {
				charArray[row][col] = '.'
				if !slices.Contains(startingRows, row) {
					startingRows = append(startingRows, row)
				}
				if !slices.Contains(startingCols, col) {
					startingCols = append(startingCols, col)
				}
			}
		}
	}

	pointsToExpand := []Point{{startingRows[len(startingRows)/2], startingCols[len(startingCols)/2]}}
	dataPoints := []int{}
	stepSize := len(charArray) / multiplier
	for i := 0; i < 327; i++ {
		for row, line := range charArray {
			for col := range line {
				if intGrid[row][col] > 0 {
					intGrid[row][col] = 1
					pointsToExpand = append(pointsToExpand, Point{row, col})
				}
			}
		}
		for _, point := range pointsToExpand {
			expandOutward(charArray, intGrid, point.row, point.col)
		}
		pointsToExpand = []Point{}
		if i == 64 || i == 64+stepSize || i == 64+stepSize*2 {
			numPoints := 0
			for row, line := range charArray {
				for col := range line {
					if intGrid[row][col] > 0 {
						numPoints++
					}
				}
			}
			dataPoints = append(dataPoints, numPoints)
		}
	}

	for i := 2; i < 202300; i++ { // 26501365 = 65 + 202300 * 131
		dataPoints = append(dataPoints, calculateDifferencesForward(dataPoints)) // day 9 solution!
		if len(dataPoints) > 5 {                                                 // only need to keep the last 5 points to extrapolate, maybe less
			dataPoints = dataPoints[len(dataPoints)-5:]
		}
	}
	fmt.Println(dataPoints[len(dataPoints)-1])
}
