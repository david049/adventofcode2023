package main

import (
	"fmt"
	"os"
	"strconv"
	"strings"
)

func createKey(input string, index int, currentNum int, hashIndex int, final []int) string {
	return input + fmt.Sprintf(",%d,%d,%d,%v", index, currentNum, hashIndex, final)
}

func findGroupsOfHashtags(input string) []int {
	var result []int
	count := 0
	for _, char := range input {
		if char == '#' {
			count++
		} else if count > 0 {
			result = append(result, count)
			count = 0
		}
	}
	if count > 0 {
		result = append(result, count)
	}
	return result
}

var cache = map[string]int{}

func checkEquality(nums1 []int, nums2 []int) bool {
	if len(nums1) != len(nums2) {
		return false
	}
	for i := 0; i < len(nums1); i++ {
		if nums1[i] != nums2[i] {
			return false
		}
	}
	return true
}

func checkCombinations(input string, index int, currentNum int, hashIndex int, final []int) int {
	key := createKey(input, index, currentNum, hashIndex, final)
	if val, ok := cache[key]; ok {
		return val
	}
	if index == len(input) {
		if len(final) == hashIndex {
			return 1
		}
		return 0
	}
	if input[index] == '#' {
		cache[key] = checkCombinations(input, index+1, currentNum+1, hashIndex, final)
		return cache[key]
	} else if input[index] == '.' || hashIndex == len(final) {
		if hashIndex < len(final) && currentNum == final[hashIndex] {
			cache[key] = checkCombinations(input, index+1, 0, hashIndex+1, final)
			return cache[key]
		} else if currentNum == 0 {
			cache[key] = checkCombinations(input, index+1, 0, hashIndex, final)
			return cache[key]
		} else {
			cache[key] = 0
			return cache[key]
		}
	} else {
		if currentNum == final[hashIndex] {
			cache[key] = checkCombinations(input, index+1, 0, hashIndex+1, final) + checkCombinations(input, index+1, currentNum+1, hashIndex, final)
			return cache[key]
		} else if currentNum == 0 {
			cache[key] = checkCombinations(input, index+1, 0, hashIndex, final) + checkCombinations(input, index+1, currentNum+1, hashIndex, final)
			return cache[key]
		} else {
			cache[key] = checkCombinations(input, index+1, currentNum+1, hashIndex, final)
			return cache[key]
		}
	}
}

func determineNumberOfCombinations(input string, nums []int) int {
	return checkCombinations(input, 0, 0, 0, nums)
}

func quintupleString(input string) string {
	returnString := ""
	for i := 0; i < 5; i++ {
		newString := input
		returnString += newString
		if i != 4 {
			returnString += "?"
		}
	}
	return returnString + "."
}

func quintupleInts(input []int) []int {
	newInput := []int{}
	for i := 0; i < 5; i++ {
		for _, num := range input {
			newInput = append(newInput, num)
		}
	}
	return newInput
}

func main() {
	inputstring, _ := os.ReadFile("input.txt")
	lines := strings.Split(string(inputstring), "\n")
	sum := 0
	for _, line := range lines {
		nums := []int{}
		splitSpring := strings.Fields(line)
		splitStringInts := strings.Split(splitSpring[1], ",")
		for _, num := range splitStringInts {
			number, _ := strconv.Atoi(num)
			nums = append(nums, number)
		}
		sum += determineNumberOfCombinations(quintupleString(splitSpring[0]), quintupleInts(nums))
	}
	fmt.Println(sum)
}
