package main

import (
	"fmt"
	"os"
	"strings"
)

type Direction int

const (
	NORTH Direction = iota
	SOUTH
	EAST
	WEST
)

type Visit struct {
	direction Direction
	row       int
	col       int
}

func getNextDirection(direction Direction, row int, col int) (int, int) {
	switch direction {
	case NORTH:
		return row - 1, col
	case SOUTH:
		return row + 1, col
	case EAST:
		return row, col + 1
	case WEST:
		return row, col - 1
	}
	panic("Impossible")
}

func energizeAndSpread(input [][]rune, energizedTiles [][]int, visitedLocations map[Visit]int, direction Direction, row int, col int) {
	if row < 0 || row >= len(input) {
		return
	}
	if col < 0 || col >= len(input[0]) {
		return
	}
	if visitedLocations[Visit{direction: direction, row: row, col: col}] == 1 {
		return
	} else {
		energizedTiles[row][col] = 1
		visitedLocations[Visit{direction: direction, row: row, col: col}] = 1
		if input[row][col] == '.' {
			nextRow, nextCol := getNextDirection(direction, row, col)
			energizeAndSpread(input, energizedTiles, visitedLocations, direction, nextRow, nextCol)
		}
		if input[row][col] == '|' {
			if direction == EAST || direction == WEST {
				northRow, northCol := getNextDirection(NORTH, row, col)
				southRow, southCol := getNextDirection(SOUTH, row, col)
				energizeAndSpread(input, energizedTiles, visitedLocations, NORTH, northRow, northCol)
				energizeAndSpread(input, energizedTiles, visitedLocations, SOUTH, southRow, southCol)
			} else {
				nextRow, nextCol := getNextDirection(direction, row, col)
				energizeAndSpread(input, energizedTiles, visitedLocations, direction, nextRow, nextCol)
			}
		}
		if input[row][col] == '-' {
			if direction == NORTH || direction == SOUTH {
				eastRow, eastCol := getNextDirection(EAST, row, col)
				westRow, westCol := getNextDirection(WEST, row, col)
				energizeAndSpread(input, energizedTiles, visitedLocations, EAST, eastRow, eastCol)
				energizeAndSpread(input, energizedTiles, visitedLocations, WEST, westRow, westCol)
			} else {
				nextRow, nextCol := getNextDirection(direction, row, col)
				energizeAndSpread(input, energizedTiles, visitedLocations, direction, nextRow, nextCol)
			}
		}
		if input[row][col] == '/' {
			newDirection := NORTH
			switch direction {
			case NORTH:
				newDirection = EAST
			case SOUTH:
				newDirection = WEST
			case EAST:
				newDirection = NORTH
			case WEST:
				newDirection = SOUTH
			}
			nextRow, nextCol := getNextDirection(newDirection, row, col)
			energizeAndSpread(input, energizedTiles, visitedLocations, newDirection, nextRow, nextCol)
		}
		if input[row][col] == '\\' {
			newDirection := NORTH
			switch direction {
			case NORTH:
				newDirection = WEST
			case SOUTH:
				newDirection = EAST
			case EAST:
				newDirection = SOUTH
			case WEST:
				newDirection = NORTH
			}
			nextRow, nextCol := getNextDirection(newDirection, row, col)
			energizeAndSpread(input, energizedTiles, visitedLocations, newDirection, nextRow, nextCol)
		}
	}
}

func main() {
	inputstring, _ := os.ReadFile("input.txt")
	lines := strings.Split(string(inputstring), "\n")
	charArray := make([][]rune, len(lines))
	energizedTiles := [][]int{}
	for i, str := range lines {
		zeroInts := make([]int, len(str))
		energizedTiles = append(energizedTiles, zeroInts)
		runes := []rune(str)
		charArray[i] = runes
	}

	maxSum := 0
	// go north from south
	for col := 0; col < len(charArray[0]); col++ {
		energizeAndSpread(charArray, energizedTiles, make(map[Visit]int), NORTH, len(charArray)-1, col)
		sum := 0
		for row, line := range energizedTiles {
			for col, num := range line {
				if num == 1 {
					sum++
					energizedTiles[row][col] = 0
				}
			}
		}
		if sum > maxSum {
			maxSum = sum
		}
	}
	// go south from north
	for col := 0; col < len(charArray[0]); col++ {
		energizeAndSpread(charArray, energizedTiles, make(map[Visit]int), SOUTH, 0, col)
		sum := 0
		for row, line := range energizedTiles {
			for col, num := range line {
				if num == 1 {
					sum++
					energizedTiles[row][col] = 0
				}
			}
		}
		if sum > maxSum {
			maxSum = sum
		}
	}
	// go west from east
	for row := 0; row < len(charArray); row++ {
		energizeAndSpread(charArray, energizedTiles, make(map[Visit]int), WEST, row, len(charArray[0])-1)
		sum := 0
		for row, line := range energizedTiles {
			for col, num := range line {
				if num == 1 {
					sum++
					energizedTiles[row][col] = 0
				}
			}
		}
		if sum > maxSum {
			maxSum = sum
		}
	}
	// go east from west
	for row := 0; row < len(charArray); row++ {
		energizeAndSpread(charArray, energizedTiles, make(map[Visit]int), EAST, row, 0)
		sum := 0
		for row, line := range energizedTiles {
			for col, num := range line {
				if num == 1 {
					sum++
					energizedTiles[row][col] = 0
				}
			}
		}
		if sum > maxSum {
			maxSum = sum
		}
	}

	energizeAndSpread(charArray, energizedTiles, make(map[Visit]int), EAST, 0, 0)
	sum := 0
	for _, line := range energizedTiles {
		for _, num := range line {
			if num == 1 {
				sum++
			}
		}
	}

	fmt.Println(sum)
	fmt.Println(maxSum)
}
