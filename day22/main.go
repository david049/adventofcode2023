package main

import (
	"fmt"
	"os"
	"regexp"
	"slices"
	"sort"
	"strconv"
	"strings"
)

type Block struct {
	startingPoint Point
	endingPoint   Point
	id            int
}

type Point struct {
	x int
	y int
	z int
}

func intersectsXY(block1, block2 Block) bool {
	return max(block1.startingPoint.x, block2.startingPoint.x) <= min(block1.endingPoint.x, block2.endingPoint.x) &&
		max(block1.startingPoint.y, block2.startingPoint.y) <= min(block1.endingPoint.y, block2.endingPoint.y)
}

func wholyContained(input []Block, containedIn []Block) bool {
	entirelyContained := true
	for _, block := range input {
		if !slices.Contains(containedIn, block) {
			return false
		}
	}
	return entirelyContained
}

func main() {
	inputstring, _ := os.ReadFile("input.txt")
	lines := strings.Split(string(inputstring), "\n")

	blocks := []Block{}
	for id, line := range lines {
		pattern := regexp.MustCompile(`(\d+),(\d+),(\d+)~(\d+),(\d+),(\d+)`)
		matches := pattern.FindStringSubmatch(line)
		x1, _ := strconv.Atoi(matches[1])
		y1, _ := strconv.Atoi(matches[2])
		z1, _ := strconv.Atoi(matches[3])
		x2, _ := strconv.Atoi(matches[4])
		y2, _ := strconv.Atoi(matches[5])
		z2, _ := strconv.Atoi(matches[6])
		blocks = append(blocks, Block{Point{x1, y1, z1}, Point{x2, y2, z2}, id})
	}

	sort.SliceStable(blocks, func(i, j int) bool {
		lowestBlock1 := min(blocks[i].endingPoint.z, blocks[i].startingPoint.z)
		lowestBlock2 := min(blocks[j].endingPoint.z, blocks[j].startingPoint.z)
		if lowestBlock1 > lowestBlock2 {
			return false
		}
		if lowestBlock1 < lowestBlock2 {
			return true
		}
		return true
	})

	for ind := range blocks {
		maxZ := 1
		for _, fallenBlock := range blocks[:ind] {
			if intersectsXY(blocks[ind], fallenBlock) {
				maxZ = max(maxZ, fallenBlock.endingPoint.z+1)
			}
		}
		blocks[ind].endingPoint.z -= blocks[ind].startingPoint.z - maxZ
		blocks[ind].startingPoint.z = maxZ
	}

	sort.SliceStable(blocks, func(i, j int) bool {
		lowestBlock1 := min(blocks[i].endingPoint.z, blocks[i].startingPoint.z)
		lowestBlock2 := min(blocks[j].endingPoint.z, blocks[j].startingPoint.z)
		if lowestBlock1 > lowestBlock2 {
			return false
		}
		if lowestBlock1 < lowestBlock2 {
			return true
		}
		return true
	})
	blockSupports := map[int][]Block{}
	blockDependentOn := map[int][]Block{}
	for ind := range blocks {
		for _, fallenBlock := range blocks[:ind] {
			if intersectsXY(blocks[ind], fallenBlock) && blocks[ind].startingPoint.z == fallenBlock.endingPoint.z+1 {
				blockSupports[fallenBlock.id] = append(blockSupports[fallenBlock.id], blocks[ind])
				blockDependentOn[blocks[ind].id] = append(blockDependentOn[blocks[ind].id], fallenBlock)
			}
		}
	}
	numDisintegrated := 0
	for _, block := range blocks {
		onlySupport := false
		for _, supportedBlock := range blockSupports[block.id] {
			if len(blockDependentOn[supportedBlock.id]) < 2 {
				onlySupport = true
				break
			}
		}
		if !onlySupport || len(blockSupports[block.id]) == 0 {
			numDisintegrated += 1
		}
	}

	fmt.Println(numDisintegrated)

	// part 2
	totalFalling := 0
	for _, block := range blocks {
		localFalling := 0
		extraFallingBlocks := []Block{block}
		for i := 0; i < len(extraFallingBlocks); i++ {
			for _, supportedBlock := range blockSupports[extraFallingBlocks[i].id] {
				if slices.Contains(extraFallingBlocks, supportedBlock) {
					continue
				}
				if wholyContained(blockDependentOn[supportedBlock.id], extraFallingBlocks) {
					extraFallingBlocks = append(extraFallingBlocks, supportedBlock)
					localFalling++
				}
			}
		}
		totalFalling += localFalling
	}
	fmt.Println(totalFalling)
}
