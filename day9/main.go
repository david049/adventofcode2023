package main

import (
	"fmt"
	"os"
	"strconv"
	"strings"
)

func checkZeros(input []int) bool {
	for _, num := range input {
		if num != 0 {
			return false
		}
	}
	return true
}

func calculateDifferencesForward(input []int) int {
	var differences []int
	for i := 1; i < len(input); i++ {
		differences = append(differences, input[i]-input[i-1])
	}
	if checkZeros(differences) {
		return input[len(input)-1]
	}
	return calculateDifferencesForward(differences) + input[len(input)-1]
}

func calculateDifferencesBackward(input []int) int {
	var differences []int
	for i := 1; i < len(input); i++ {
		differences = append(differences, input[i]-input[i-1])
	}
	if checkZeros(differences) {
		return input[0]
	}
	return input[0] - calculateDifferencesBackward(differences)
}

func main() {
	inputstring, _ := os.ReadFile("input.txt")
	lines := strings.Split(string(inputstring), "\n")
	sum := 0
	for _, line := range lines {
		var ints []int
		for _, num := range strings.Fields(line) {
			toAdd, _ := strconv.Atoi(num)
			ints = append(ints, toAdd)
		}
		sum += calculateDifferencesBackward(ints)
	}
	fmt.Println(sum)
}
