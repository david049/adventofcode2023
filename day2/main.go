package main

import (
	"fmt"
	"io/ioutil"
	"strings"
	"strconv"
)

func checkValidGame(input string) (int, int, int) {
	marbles := strings.Split(input, ",")
	red := 0
	blue := 0
	green := 0

	for _, colour := range marbles {
		combo := strings.Split(colour, " ")
		num, _ := strconv.Atoi(combo[1])
		if strings.Contains(colour, "red") {
			red = num
		}
		if strings.Contains(colour, "green") {
			green = num
		}
		if strings.Contains(colour, "blue") {
			blue = num
		}
	}
	return red, green, blue
}

func main() {
	inputstring, _ := ioutil.ReadFile("input.txt")
	lines := strings.Split(string(inputstring), "\n")
	sum := 0
	for _, line := range lines {
		line1 := strings.Split(line, ":")
		//gamestr := line1[0][5:]
		//gamenum, _ := strconv.Atoi(gamestr)
		lines2 := strings.Split(line1[1], ";")
	    maxRed := 0
		maxBlue := 0
		maxGreen := 0
		for _, game := range(lines2) {
			red, green, blue := checkValidGame(game)
			if red > maxRed {
				maxRed = red
			}
			if green > maxGreen {
				maxGreen = green
			}
			if blue > maxBlue {
				maxBlue = blue
			}
		}
		sum += maxRed*maxBlue*maxGreen
	}
	fmt.Println(sum)
}