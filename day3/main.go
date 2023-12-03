package main

import (
	"fmt"
	"io/ioutil"
	"strings"
	"unicode"
)

func convToNum(nums []int, len int) (int) {
	num := 0
	mag := 1
	for i := 1; i < len; i++ {
		mag *= 10
	}
	for i := 0; i < len; i++ {
		num += mag * nums[i]
		mag /= 10
	}
	return num
}

func getSym(row int, col int, grid []string) (rune) {
	if row < 0 || row >= len(grid) {
		return '.'
	}
	if col < 0 || col >= len(grid[row]) {
		return '.'
	}
	return rune(grid[row][col])
}

func checkSyms(len int, startCol int, row int, grid []string, gearRow int, gearCol int) (bool) {
	for i := startCol - 1; i <= startCol + len; i++ {
		above := getSym(row - 1, i, grid)
		below := getSym(row + 1, i, grid)
		if !unicode.IsDigit(above) && above != '.' && above == '*' && row - 1 == gearRow && i == gearCol{
			return true
		}
		if !unicode.IsDigit(below) && below != '.' && below == '*' && row + 1 == gearRow && i == gearCol{
			return true
		}
	}
	left := getSym(row, startCol - 1, grid)
	right := getSym(row, startCol + len, grid)
	if !unicode.IsDigit(left) && left != '.' && left == '*' && row == gearRow && startCol - 1 == gearCol {
		return true
	}
	if !unicode.IsDigit(right) && right != '.' && right == '*' && row == gearRow && startCol + len == gearCol{
		return true
	}
	return false
}

func checkGear(row int, col int, grid []string) (int) {
	sum, nums := checkLine(row - 1, grid, row, col)
	sum2, nums2 := checkLine(row + 1, grid, row, col)
	sum3, nums3 := checkLine(row, grid, row, col)
	if nums + nums2 + nums3 != 2 {
		return 0
	}
	return sum * sum2 * sum3
}
func checkLine(ind int, grid []string, gearRow int, gearCol int) (int, int) {
	if ind < 0 || ind >= len(grid) {
		return 1, 0
	}
	numLen := 0
	ints := []int{0, 0, 0}
	gearRatio := 1
	startNum := 0
	numNums := 0
	for col, char := range grid[ind] {
		if unicode.IsDigit(char) {
			ints[numLen] = int(char - '0')
			if numLen == 0 {
				startNum = col
			}
			numLen += 1
		} else {
			if numLen > 0 {
				if (checkSyms(numLen, startNum, ind, grid, gearRow, gearCol)) {
					gearRatio *= convToNum(ints, numLen)
					numNums += 1
				}
			}
			numLen = 0
			startNum = 0
		}
		if col == len(grid[ind]) - 1 {
			if numLen > 0 {
				if (checkSyms(numLen, startNum, ind, grid, gearRow, gearCol)) {
					gearRatio *= convToNum(ints, numLen)
					numNums += 1
				}
			}
			numLen = 0
			startNum = 0
		}
	}
	return gearRatio, numNums
}

func main() {
	inputstring, _ := ioutil.ReadFile("input.txt")
	lines := strings.Split(string(inputstring), "\n")
	grid := make([]string, len(lines))
	sum := 0
	for ind, line := range lines {
		grid[ind] = line
	}
	for ind, line := range grid {
		for col, char := range line {
			if char == '*' {
				sum += checkGear(ind, col, grid)
			}
		}
		
	}
	fmt.Println(sum)
}