package main

import (
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	inputstring, _ := os.ReadFile("input.txt")
	lines := strings.Split(string(inputstring), "\n")
	times := strings.Fields(lines[0])[1:]
	distances := strings.Fields(lines[1])[1:]
	numWays := 0
	time := ""
	dist := ""
	for i := 0; i < len(times); i++ {
		time += times[i]
		dist += distances[i]
	}
	totalTime, _ := strconv.Atoi(time)
	totalDist, _ := strconv.Atoi(dist)
	for j := 1; j < totalTime; j++ {
		tempTime := totalTime - j
		tempDist := tempTime * j
		if tempDist > totalDist {
			numWays++
		}
	}
	fmt.Println(numWays)
}
