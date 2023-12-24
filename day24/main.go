package main

import (
	"fmt"
	"os"
	"regexp"
	"strconv"
	"strings"
)

type Line struct {
	X         float64
	Y         float64
	Z         float64
	VelocityX float64
	VelocityY float64
	VelocityZ float64
}

func findIntersection(line1 Line, line2 Line) (float64, float64, bool) {
	slope1 := line1.VelocityY / line1.VelocityX
	slope2 := line2.VelocityY / line2.VelocityX

	if slope1 == slope2 {
		return 0, 0, false
	}

	intercept1 := line1.Y - slope1*line1.X
	intercept2 := line2.Y - slope2*line2.X

	intersectX := (intercept2 - intercept1) / (slope1 - slope2)
	intersectY := slope1*intersectX + intercept1

	if intersectX < line1.X && line1.VelocityX >= 0 {
		return 0, 0, false
	}
	if intersectX > line1.X && line1.VelocityX <= 0 {
		return 0, 0, false
	}
	if intersectY < line1.Y && line1.VelocityY >= 0 {
		return 0, 0, false
	}
	if intersectY > line1.Y && line1.VelocityY <= 0 {
		return 0, 0, false
	}

	if intersectX < line2.X && line2.VelocityX >= 0 {
		return 0, 0, false
	}
	if intersectX > line2.X && line2.VelocityX <= 0 {
		return 0, 0, false
	}
	if intersectY < line2.Y && line2.VelocityY >= 0 {
		return 0, 0, false
	}
	if intersectY > line2.Y && line2.VelocityY <= 0 {
		return 0, 0, false
	}

	return intersectX, intersectY, true
}

func main() {
	inputstring, _ := os.ReadFile("input.txt")
	lines := strings.Split(string(inputstring), "\n")
	hailLines := []Line{}
	for _, line := range lines {
		pattern := regexp.MustCompile(`-?\d+`)
		matches := pattern.FindAllString(line, -1)
		x, _ := strconv.ParseFloat(matches[0], 64)
		y, _ := strconv.ParseFloat(matches[1], 64)
		z, _ := strconv.ParseFloat(matches[2], 64)
		vx, _ := strconv.ParseFloat(matches[3], 64)
		vy, _ := strconv.ParseFloat(matches[4], 64)
		vz, _ := strconv.ParseFloat(matches[5], 64)
		hailLines = append(hailLines, Line{x, y, z, vx, vy, vz})
	}

	testMinX := float64(200000000000000)
	testMaxX := float64(400000000000000)
	testMinY := float64(200000000000000)
	testMaxY := float64(400000000000000)

	// part 1
	numIntersect := 0
	for i := range hailLines {
		for j := i + 1; j < len(hailLines); j++ {
			x, y, intersects := findIntersection(hailLines[i], hailLines[j])
			if !intersects {
				continue
			}
			if x >= testMinX && x <= testMaxX && y >= testMinY && y <= testMaxY {
				numIntersect++
			}
		}
	}
	fmt.Println(numIntersect)
}
