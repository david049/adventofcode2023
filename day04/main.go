package main

import (
	"fmt"
	"os"
	"slices"
	"strconv"
	"strings"
)

func main() {
	inputstring, _ := os.ReadFile("input.txt")
	lines := strings.Split(string(inputstring), "\n")
	scratchMap := map[int]int{}
	numCards := map[int]int{}
	for _, line := range lines {
		matches := 0
		game := strings.Split(line, ":")
		gameID := strings.Split(game[0], " ")
		gameNum, _ := strconv.Atoi(gameID[len(gameID) - 1])
		gameNums := strings.Split(strings.Trim(game[1], " "), "|")
		winningNums := strings.Split(strings.Trim(gameNums[0], " "), " ")
		myNums := strings.Split(strings.Trim(gameNums[1], " "), " ")
		for _, num := range myNums {
			if slices.Contains(winningNums, num) {
				if strings.Trim(num, " ") == "" {
					continue
				}
				matches += 1
				numCards[gameNum + matches] += 1
			}
		}
	    extraIterations := numCards[gameNum]
		for i := 0; i < extraIterations; i++ {
			matches = 0
			for _, num := range myNums {
				if slices.Contains(winningNums, num) {
					if strings.Trim(num, " ") == "" {
						continue
					}
					matches += 1
					numCards[gameNum + matches] += 1
				}
			}
		}
		scratchMap[gameNum] = matches
	}
	numCards[1] = 0
	sum := 0
	for i := 1; i <= len(lines); i++ {
		sum += numCards[i] + 1
	}
	fmt.Println(sum)
}
