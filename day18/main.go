package main

import (
	"fmt"
	"os"
	"strconv"
	"strings"
)

type Point struct {
	row int64
	col int64
}

func determinant(x1 int64, y1 int64, x2 int64, y2 int64) int64 {
	return x1*y2 - x2*y1
}

func determineArea(points []Point, border int64) int64 {
	sum := int64(0)
	for i := 0; i < len(points); i++ {
		firstIndex := i
		secondIndex := (i + 1) % len(points)
		sum += determinant(points[firstIndex].col, points[firstIndex].row, points[secondIndex].col, points[secondIndex].row)
	}
	return sum/2 - border/2 + 1 + border
}

func main() {
	inputstring, _ := os.ReadFile("input.txt")
	lines := strings.Split(string(inputstring), "\n")
	currentRow := int64(0)
	currentCol := int64(0)
	border := int64(0)
	points := []Point{}
	for _, line := range lines {
		instructionLine := strings.Fields(line)
		num := ""
		for i := 2; i < 7; i++ {
			num += string(instructionLine[2][i])
		}
		intNum, _ := strconv.ParseInt(num, 16, 64)
		switch instructionLine[2][7] {
		case '0':
			currentCol += intNum
		case '2':
			currentCol -= intNum
		case '1':
			currentRow += intNum
		case '3':
			currentRow -= intNum
		}
		border += intNum
		points = append(points, Point{row: currentRow, col: currentCol})
	}

	fmt.Println(determineArea(points, border))
}
