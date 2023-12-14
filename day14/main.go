package main

import (
	"fmt"
	"os"
	"strings"
)

func moveRocksNorth(rocks [][]rune) {
	for row := 0; row < len(rocks); row++ {
		for col := 0; col < len(rocks[0]); col++ {
			if rocks[row][col] == 'O' {
				// scan up
				highestRow := row
				for tmpRow := row - 1; tmpRow >= 0; tmpRow-- {
					if rocks[tmpRow][col] == '.' {
						highestRow = tmpRow
					}
					if rocks[tmpRow][col] == '#' || rocks[tmpRow][col] == 'O' {
						break
					}
				}
				if highestRow != row {
					rocks[highestRow][col] = 'O'
					rocks[row][col] = '.'
				}
			}
		}
	}
}

func moveRocksSouth(rocks [][]rune) {
	for row := len(rocks) - 1; row >= 0; row-- {
		for col := 0; col < len(rocks[0]); col++ {
			if rocks[row][col] == 'O' {
				// scan down
				highestRow := row
				for tmpRow := row + 1; tmpRow < len(rocks); tmpRow++ {
					if rocks[tmpRow][col] == '.' {
						highestRow = tmpRow
					}
					if rocks[tmpRow][col] == '#' || rocks[tmpRow][col] == 'O' {
						break
					}
				}
				if highestRow != row {
					rocks[highestRow][col] = 'O'
					rocks[row][col] = '.'
				}
			}
		}
	}
}

func moveRocksWest(rocks [][]rune) {
	for row := len(rocks) - 1; row >= 0; row-- {
		for col := 0; col < len(rocks[0]); col++ {
			if rocks[row][col] == 'O' {
				// scan down
				highestCol := col
				for tmpCol := col - 1; tmpCol >= 0; tmpCol-- {
					if rocks[row][tmpCol] == '.' {
						highestCol = tmpCol
					}
					if rocks[row][tmpCol] == '#' || rocks[row][tmpCol] == 'O' {
						break
					}
				}
				if highestCol != col {
					rocks[row][highestCol] = 'O'
					rocks[row][col] = '.'
				}
			}
		}
	}
}

func moveRocksEast(rocks [][]rune) {
	for row := len(rocks) - 1; row >= 0; row-- {
		for col := len(rocks[0]) - 1; col >= 0; col-- {
			if rocks[row][col] == 'O' {
				// scan down
				highestCol := col
				for tmpCol := col + 1; tmpCol < len(rocks); tmpCol++ {
					if rocks[row][tmpCol] == '.' {
						highestCol = tmpCol
					}
					if rocks[row][tmpCol] == '#' || rocks[row][tmpCol] == 'O' {
						break
					}
				}
				if highestCol != col {
					rocks[row][highestCol] = 'O'
					rocks[row][col] = '.'
				}
			}
		}
	}
}

func stringifyGrid(arr [][]rune) string {
	var builder strings.Builder
	for _, row := range arr {
		for _, r := range row {
			builder.WriteRune(r)
		}
	}
	result := builder.String()
	return result
}

func main() {
	inputstring, _ := os.ReadFile("input.txt")
	lines := strings.Split(string(inputstring), "\n")
	charArray := make([][]rune, len(lines))
	for i, str := range lines {
		runes := []rune(str)
		charArray[i] = runes
	}

	trackSeen := map[string]int{}
	cycleStart := 0
	cycleEnd := 0
	remainingIterations := -1
	for true {
		moveRocksNorth(charArray)
		moveRocksWest(charArray)
		moveRocksSouth(charArray)
		moveRocksEast(charArray)
		cycleEnd++
		if trackSeen[stringifyGrid(charArray)] != 0 && remainingIterations == -1 {
			cycleStart = trackSeen[stringifyGrid(charArray)]
			cycleLength := cycleEnd - cycleStart
			remainingIterations = (1000000000-cycleStart)%cycleLength + cycleLength + cycleStart
		}
		trackSeen[stringifyGrid(charArray)] = cycleEnd
		if cycleEnd == remainingIterations {
			break
		}
	}

	sum := 0
	for ind, row := range charArray {
		for _, char := range row {
			if char == 'O' {
				sum += len(charArray) - ind
			}
		}
	}
	fmt.Println(sum)
}
