package main

import (
	"fmt" 
	"unicode"
	"strings"
	"io/ioutil"
)

func trebuchet(input string) (int){
	input = strings.Replace(input, "three", "t3e", -1)
	input = strings.Replace(input, "four", "f4r", -1)
	input = strings.Replace(input, "five", "f5e", -1)
	input = strings.Replace(input, "seven", "s7n", -1)
	input = strings.Replace(input, "eight", "e8t", -1)
	input = strings.Replace(input, "nine", "n9e", -1)
	input = strings.Replace(input, "six", "s6x", -1)
	input = strings.Replace(input, "one", "o1e", -1)
	input = strings.Replace(input, "two", "t2o", -1)
	digit1 := -1
	digit2 := -1
	fmt.Println(input)
	for _, char := range input {
		if unicode.IsDigit(char) {
			if digit1 == -1 {
				digit1 = int(char - '0')
			}
			digit2 = int(char - '0')
		}
	}
	return digit1*10 + digit2
}

func main() {
	inputstring, _ := ioutil.ReadFile("input.txt")
	sum := 0
	lines := strings.Split(string(inputstring), "\n")
	for _, line := range lines {
		sum += trebuchet(line)
	}
	fmt.Println(sum)
}