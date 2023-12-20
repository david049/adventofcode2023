package main

import (
	"fmt"
	"os"
	"slices"
	"strings"
)

type ModuleType int

const (
	BROADCASTER ModuleType = iota
	FLIPFLOP
	CONJUNCTION
)

type Module struct {
	moduleType           ModuleType
	status               int
	mostRecent           int
	destination          []string
	input                []string
	mostRecentlyReceived map[string]int
}

type Signal struct {
	signal      int
	destination string
	origin      string
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
	modules := map[string]Module{}
	moduleNames := []string{}
	moduleToWatch := ""
	for _, line := range lines {
		splitLine := strings.Split(line, "->")
		name := strings.Fields(splitLine[0])[0]
		moduleType := BROADCASTER
		if name[0] == '%' {
			name = name[1:]
			moduleType = FLIPFLOP
		} else if name[0] == '&' {
			name = name[1:]
			moduleType = CONJUNCTION
		}
		destinationArray := strings.Split(splitLine[1], ",")
		resultDestinations := []string{}
		for _, dest := range destinationArray {
			resultDestinations = append(resultDestinations, dest[1:])
			if dest[1:] == "rx" {
				moduleToWatch = name
			}
		}
		moduleNames = append(moduleNames, name)
		modules[name] = Module{moduleType: moduleType, status: 0, mostRecent: 0, destination: resultDestinations, input: []string{}, mostRecentlyReceived: map[string]int{}}
	}

	for _, name := range moduleNames {
		for _, dest := range modules[name].destination {
			if modules[dest].moduleType == CONJUNCTION {
				modules[dest] = Module{moduleType: modules[dest].moduleType, mostRecent: 0, destination: modules[dest].destination, input: append(modules[dest].input, name), mostRecentlyReceived: modules[dest].mostRecentlyReceived}
			}
		}

	}

	lowSignals := 0
	highSignals := 0

	for j := 0; j < 1000; j++ {
		processQueue := []Signal{{0, "broadcaster", "button"}}
		for i := 0; i < len(processQueue); i++ {
			signal := processQueue[i].signal
			if signal == 1 {
				highSignals++
			} else {
				lowSignals++
			}
			name := processQueue[i].destination
			receivingModule := modules[name]
			switch receivingModule.moduleType {
			case BROADCASTER:
				for _, dest := range receivingModule.destination {
					processQueue = append(processQueue, Signal{0, dest, "broadcaster"})
					if signal == 1 {
					}
				}
			case FLIPFLOP:
				if signal == 1 {
					continue
				} else {
					currentStatus := receivingModule.status
					newStatus := 0
					if currentStatus == 0 {
						newStatus = 1
					}
					for _, dest := range receivingModule.destination {
						processQueue = append(processQueue, Signal{newStatus, dest, name})
					}
					modules[name] = Module{moduleType: receivingModule.moduleType, status: newStatus, destination: receivingModule.destination}
				}
			case CONJUNCTION:
				if slices.Contains(modules[name].input, processQueue[i].origin) {
					modules[name].mostRecentlyReceived[processQueue[i].origin] = signal
				}
				allHigh := true
				for _, inputName := range modules[name].input {
					if modules[name].mostRecentlyReceived[inputName] == 0 {
						allHigh = false
						break
					}
				}
				if allHigh {
					for _, dest := range receivingModule.destination {
						processQueue = append(processQueue, Signal{0, dest, name})
					}
				} else {
					for _, dest := range receivingModule.destination {
						processQueue = append(processQueue, Signal{1, dest, name})
					}
				}
			}
		}
	}

	lcmMap := map[string]int{}
	buttonPresses := 1000 // instead of resetting modules, just start from 1000
	for true {
		buttonPresses++
		allDefined := true
		for _, name := range modules[moduleToWatch].input {
			if lcmMap[name] == 0 {
				allDefined = false
			}
		}
		if allDefined {
			break
		}
		processQueue := []Signal{{0, "broadcaster", "button"}}
		for i := 0; i < len(processQueue); i++ {
			signal := processQueue[i].signal
			name := processQueue[i].destination

			if slices.Contains(modules[moduleToWatch].input, processQueue[i].origin) && lcmMap[processQueue[i].origin] == 0 && signal == 1 {
				lcmMap[processQueue[i].origin] = buttonPresses
				break
			}

			if signal == 1 {
				highSignals++
			} else {
				lowSignals++
			}

			receivingModule := modules[name]
			switch receivingModule.moduleType {
			case BROADCASTER:
				for _, dest := range receivingModule.destination {
					processQueue = append(processQueue, Signal{0, dest, "broadcaster"})
					if signal == 1 {
					}
				}
			case FLIPFLOP:
				if signal == 1 {
					continue
				} else {
					currentStatus := receivingModule.status
					newStatus := 0
					if currentStatus == 0 {
						newStatus = 1
					}
					for _, dest := range receivingModule.destination {
						processQueue = append(processQueue, Signal{newStatus, dest, name})
					}
					modules[name] = Module{moduleType: receivingModule.moduleType, status: newStatus, destination: receivingModule.destination}
				}
			case CONJUNCTION:
				if slices.Contains(modules[name].input, processQueue[i].origin) {
					modules[name].mostRecentlyReceived[processQueue[i].origin] = signal
				}
				allHigh := true
				for _, inputName := range modules[name].input {
					if modules[name].mostRecentlyReceived[inputName] == 0 {
						allHigh = false
						break
					}
				}
				if allHigh {
					for _, dest := range receivingModule.destination {
						processQueue = append(processQueue, Signal{0, dest, name})
					}
				} else {
					for _, dest := range receivingModule.destination {
						processQueue = append(processQueue, Signal{1, dest, name})
					}
				}
			}

		}
	}

	fmt.Println("part1:", lowSignals*highSignals)
	lcmArray := []int{}
	for _, name := range modules[moduleToWatch].input {
		lcmArray = append(lcmArray, lcmMap[name])
	}
	fmt.Println("part2:", findLCMArray(lcmArray))
}
