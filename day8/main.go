package main

import (
	"fmt"
	"os"
	"strings"
)

type Destinations struct {
	left  string
	right string
}

func findGCD(num1 int, num2 int) int {
	if num2 == 0 {
		return num1
	}
	return findGCD(num2, num1%num2)
}

func findLCM(num1 int, num2 int) int {
	return (num1 / findGCD(num1, num2)) * num2
}

func findLCMArray(nums []int) int {
	lcm := nums[0]
	for i := 1; i < len(nums); i++ {
		lcm = findLCM(lcm, nums[i])
	}
	return lcm
}

func main() {
	inputstring, _ := os.ReadFile("input.txt")
	lines := strings.Split(string(inputstring), "\n")
	instructions := lines[0]
	rest := lines[2:]

	locationMap := make(map[string]Destinations)
	var startingNodes []string

	for _, location := range rest {
		splitLocation := strings.Fields(location)
		left := splitLocation[2][1:4]
		right := splitLocation[3][0:3]
		locationMap[splitLocation[0]] = Destinations{left: left, right: right}
		if splitLocation[0][2] == 'A' {
			startingNodes = append(startingNodes, splitLocation[0])
		}
	}

	numSteps := []int{}

	for _, node := range startingNodes {
		currentLocation := node
		steps := 0
		currentInstructionIndex := 0
		for currentLocation[2] != 'Z' {
			if instructions[currentInstructionIndex%len(instructions)] == 'L' {
				currentLocation = locationMap[currentLocation].left
			}
			if instructions[currentInstructionIndex%len(instructions)] == 'R' {
				currentLocation = locationMap[currentLocation].right
			}
			currentInstructionIndex++
			steps++
		}
		numSteps = append(numSteps, steps)
	}
	fmt.Println(findLCMArray(numSteps))
}
