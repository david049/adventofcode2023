package main

import (
	"fmt"
	"os"
	"strings"
)

type Point struct {
	row int
	col int
}

func neighbours(input [][]rune, row int, col int) []Point {
	points := []Point{}
	if row-1 >= 0 && input[row-1][col] != '#' {
		points = append(points, Point{row - 1, col})
	}
	if row+1 < len(input) && input[row+1][col] != '#' {
		points = append(points, Point{row + 1, col})
	}
	if col+1 < len(input[0]) && input[row][col+1] != '#' {
		points = append(points, Point{row, col + 1})
	}
	if col-1 >= 0 && input[row][col-1] != '#' {
		points = append(points, Point{row, col - 1})
	}
	return points
}

func dfs(input [][]rune, visited [][]int, row int, col int, curr int) int {
	if row == len(input)-1 && col == len(input)-2 {
		return curr
	}
	if visited[row][col] == 1 {
		return -1
	}
	visited[row][col] = 1
	neighbours := neighbours(input, row, col)
	longestPath := 0
	for _, neighbour := range neighbours {
		longestPath = max(longestPath, dfs(input, visited, neighbour.row, neighbour.col, curr+1))
	}
	visited[row][col] = 0
	return longestPath
}

func main() {
	inputstring, _ := os.ReadFile("input.txt")
	lines := strings.Split(string(inputstring), "\n")
	charArray := make([][]rune, len(lines))
	intArray := make([][]int, len(lines))
	for i, str := range lines {
		runes := []rune(str)
		charArray[i] = runes
		zeroInts := make([]int, len(str))
		intArray[i] = zeroInts
	}

	fmt.Println(dfs(charArray, intArray, 0, 1, 0))
}
