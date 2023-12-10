package main

import (
	"fmt"
	"os"
	"strings"
)

func getSym(row int, col int, grid []string) rune {
	if row < 0 || row >= len(grid) {
		return '.'
	}
	if col < 0 || col >= len(grid[row]) {
		return '.'
	}
	return rune(grid[row][col])
}

func findConnection(grid []string, startRow int, startCol int) ([]int, []int) {
	rows := []int{}
	cols := []int{}
	char := getSym(startRow, startCol, grid)
	if char == 'S' {
		if getSym(startRow-1, startCol, grid) == '|' { // north
			rows = append(rows, startRow-1)
			cols = append(cols, startCol) // north
		}
		if getSym(startRow+1, startCol, grid) == '|' { // south
			rows = append(rows, startRow+1)
			cols = append(cols, startCol) // south
		}
		if getSym(startRow, startCol+1, grid) == 'J' || getSym(startRow, startCol+1, grid) == '7' || getSym(startRow, startCol+1, grid) == '-' { // west
			rows = append(rows, startRow)
			cols = append(cols, startCol+1) // east
		}
		if getSym(startRow, startCol-1, grid) == 'F' || getSym(startRow, startCol-1, grid) == '7' || getSym(startRow, startCol-1, grid) == '-' {
			rows = append(rows, startRow)
			cols = append(cols, startCol-1) // west
		}
	}
	if char == '|' {
		rows = append(rows, startRow-1)
		cols = append(cols, startCol) // north
		rows = append(rows, startRow+1)
		cols = append(cols, startCol) // south
	}
	if char == 'L' {
		rows = append(rows, startRow-1)
		cols = append(cols, startCol) // north
		rows = append(rows, startRow)
		cols = append(cols, startCol+1) // east
	}
	if char == 'J' {
		rows = append(rows, startRow-1)
		cols = append(cols, startCol) // north
		rows = append(rows, startRow)
		cols = append(cols, startCol-1) // west
	}
	if char == '7' {
		rows = append(rows, startRow)
		cols = append(cols, startCol-1) // west
		rows = append(rows, startRow+1)
		cols = append(cols, startCol) // south
	}
	if char == 'F' {
		rows = append(rows, startRow)
		cols = append(cols, startCol+1) // east
		rows = append(rows, startRow+1)
		cols = append(cols, startCol) // south
	}
	if char == '-' {
		rows = append(rows, startRow)
		cols = append(cols, startCol+1) // east
		rows = append(rows, startRow)
		cols = append(cols, startCol-1) // west
	}
	return rows, cols
}

func main() {
	inputstring, _ := os.ReadFile("input.txt")
	lines := strings.Split(string(inputstring), "\n")
	grid := make([]string, len(lines))
	startCol := -1
	startRow := -1
	for row, line := range lines {
		grid[row] = line
		for col, chr := range line {
			if chr == 'S' { // traverse to find S
				startCol = col
				startRow = row
			}
		}
	}

	numGrid := make([][]int, len(lines))
	for i := 0; i < len(lines); i++ {
		numGrid[i] = make([]int, len(grid[0]))
	}

	numGrid[startRow][startCol] = 1

	connectionRow, connectionCol := findConnection(grid, startRow, startCol)

	currentRow := connectionRow[0]
	currentCol := connectionCol[0]
	connectionRow, connectionCol = findConnection(grid, currentRow, currentCol)
	row1 := connectionRow[0]
	row2 := connectionRow[1]
	col1 := connectionCol[0]
	col2 := connectionCol[1]

	for numGrid[row1][col1] == 0 || numGrid[row2][col2] == 0 {
		numGrid[currentRow][currentCol] = 1
		goRows, goCols := findConnection(grid, currentRow, currentCol)
		row1 = goRows[0]
		row2 = goRows[1]
		col1 = goCols[0]
		col2 = goCols[1]
		if numGrid[row1][col1] > 0 && numGrid[row2][col2] == 0 {
			currentRow = row2
			currentCol = col2
		}
		if numGrid[row1][col1] == 0 && numGrid[row2][col2] > 0 {
			currentRow = row1
			currentCol = col1
		}
	}

	count := 0
	for row, rowNums := range numGrid {
		for col, num := range rowNums {
			if num == 0 {
				tmpCol := col
				numUp := 0
				for tmpCol >= 0 {
					if (grid[row][tmpCol] == '|' || grid[row][tmpCol] == 'J' || grid[row][tmpCol] == 'L') && numGrid[row][tmpCol] == 1 {
						numUp++
					}
					tmpCol--
				}
				if numUp%2 == 1 {
					numGrid[row][col] = 2
					count++
				}

			}
		}
	}
	fmt.Println(count)
}
