package main

import (
	"fmt"
	"os"
	"slices"

	"sort"
	"strconv"
	"strings"
)

type Range struct {
	start int
	end   int
}

func main() {
	inputstring, _ := os.ReadFile("input.txt")
	lines := strings.Split(string(inputstring), "\n")
	seeds := strings.Fields(lines[0])[1:]
	var seedNums []int
	var ranges []Range
	for _, num := range seeds {
		intNum, _ := strconv.Atoi(num)
		seedNums = append(seedNums, intNum)
	}
	for i := 0; i < len(seedNums); i += 2 {
		ranges = append(ranges, Range{start: seedNums[i], end: seedNums[i] + seedNums[i+1]})
	}

	for _, line := range strings.Split(string(inputstring), "\n\n")[1:] {
		vals := strings.Split(line, ":\n")[1]
		var newRanges []Range
		var doneRanges []Range
		for _, val := range strings.Split(vals, "\n") {
			mapRow := strings.Fields(val)
			end, _ := strconv.Atoi(mapRow[0])
			start, _ := strconv.Atoi(mapRow[1])
			valRange, _ := strconv.Atoi(mapRow[2])
			for _, seedRange := range ranges {
				if slices.Contains(doneRanges, seedRange) {
					continue
				}
				if seedRange.start >= start+valRange {
					continue
				}
				if seedRange.end <= start {
					continue
				}
				if seedRange.start >= start {
					if seedRange.end <= start+valRange {
						newRanges = append(newRanges, Range{start: seedRange.start - start + end, end: seedRange.end - start + end})
						doneRanges = append(doneRanges, seedRange)
					}
					if seedRange.end > start+valRange {
						newRanges = append(newRanges, Range{start: seedRange.start - start + end, end: end + valRange})
						doneRanges = append(doneRanges, seedRange)
						ranges = append(ranges, Range{start: start + valRange, end: seedRange.end})
					}
				}
				if seedRange.start < start && seedRange.end <= start+valRange {
					newRanges = append(newRanges, Range{start: end, end: seedRange.end - start + end})
					ranges = append(ranges, Range{start: seedRange.start, end: start})
					doneRanges = append(doneRanges, seedRange)
					continue
					// starts before, ends in, results in 2 intervals
				}
				if seedRange.start < start && seedRange.end >= start+valRange {
					ranges = append(ranges, Range{start: start + valRange, end: seedRange.end})
					ranges = append(ranges, Range{start: seedRange.start, end: start})
					newRanges = append(newRanges, Range{start: end, end: end + valRange})
					doneRanges = append(doneRanges, seedRange)
					continue
					// starts 	before, ends outside, results in 3 intervals
				}
			}
		}
		for _, seedRange := range ranges {
			if !slices.Contains(doneRanges, seedRange) {
				newRanges = append(newRanges, seedRange)
			}
		}
		ranges = newRanges
	}
	var startIntervals []int
	for _, seedRange := range ranges {
		startIntervals = append(startIntervals, seedRange.start)
	}

	sort.Ints(startIntervals)
	fmt.Println(startIntervals[0])
}
