package main

import (
	"fmt"
	"os"
	"strings"
)

func isReflectionRow(input []string, row int) bool {
	misMatches := 0
	for col := 0; col < len(input[0]); col++ {
		for iterator := 0; iterator < len(input); iterator++ {
			if row-iterator < 0 {
				break
			}
			if row+iterator+1 >= len(input) {
				break
			}
			if input[row-iterator][col] != input[row+iterator+1][col] {
				misMatches++
			}
		}
	}
	if misMatches == 1 {
		return true
	}
	return false
}

func isReflectionCol(input []string, col int) bool {
	misMatches := 0
	for row := 0; row < len(input); row++ {
		for iterator := 0; iterator < len(input); iterator++ {
			if col-iterator < 0 {
				break
			}
			if col+iterator+1 >= len(input[0]) {
				break
			}
			if input[row][col-iterator] != input[row][col+iterator+1] {
				misMatches++
			}
		}
	}
	if misMatches == 1 {
		return true
	}
	return false
}

func scanForColRow(input string) int {
	lines := strings.Split(input, "\n")
	for i := 0; i < len(lines)-1; i++ {
		if isReflectionRow(lines, i) {
			return (i + 1) * 100
		}
	}
	for i := 0; i < len(lines[0])-1; i++ {
		if isReflectionCol(lines, i) {
			return i + 1
		}
	}
	return 0
}

func main() {
	inputstring, _ := os.ReadFile("input.txt")
	lines := strings.Split(string(inputstring), "\n\n")
	sum := 0
	for _, line := range lines {
		sum += scanForColRow(line)
	}
	fmt.Println(sum)
}
