package main

import (
	"fmt"
	"os"
	"strconv"
	"strings"
)

type Lens struct {
	label       string
	focalLength int
}

func hashString(input string) int {
	sum := 0
	for i := 0; i < len(input); i++ {
		sum = ((int(input[i]) + sum) * 17) % 256
	}
	return sum
}

func main() {
	inputstring, _ := os.ReadFile("input.txt")
	lines := strings.Split(string(inputstring), "\n")
	charVals := []string{}

	// part 1
	sum := 0
	for _, line := range lines {
		commaSeparatedLines := strings.Split(line, ",")
		for _, val := range commaSeparatedLines {
			if val == "" {
				continue
			}
			sum += hashString(val)
			charVals = append(charVals, val)
		}
	}
	fmt.Println(sum)

	// part 2
	sum = 0
	finalBoxes := map[int][]Lens{}

	for _, str := range charVals {
		if strings.Contains(str, "-") {
			label := str[:len(str)-1]
			for i := 0; i < len(finalBoxes[hashString(label)]); i++ {
				if finalBoxes[hashString(label)][i].label == label {
					finalBoxes[hashString(label)] = append(finalBoxes[hashString(label)][:i], finalBoxes[hashString(label)][i+1:]...)
					break
				}
			}
		}
		if strings.Contains(str, "=") {
			splitStr := strings.Split(str, "=")
			label := splitStr[0]
			focalLength, _ := strconv.Atoi(splitStr[1])
			found := false
			for i := 0; i < len(finalBoxes[hashString(label)]); i++ {
				if finalBoxes[hashString(label)][i].label == label {
					found = true
					finalBoxes[hashString(label)][i].focalLength = focalLength
				}
			}
			if !found {
				finalBoxes[hashString(label)] = append(finalBoxes[hashString(label)], Lens{label: label, focalLength: focalLength})
			}
		}
	}
	for i := 0; i < 256; i++ {
		if len(finalBoxes[i]) != 0 {
			for slot := 0; slot < len(finalBoxes[i]); slot++ {
				sum += (i + 1) * (slot + 1) * finalBoxes[i][slot].focalLength
			}
		}
	}
	fmt.Println(sum)
}
