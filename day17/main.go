package main

import (
	"container/heap"
	"fmt"
	"os"
	"strconv"
	"strings"
)

type Direction int

const (
	UP Direction = iota
	DOWN
	LEFT
	RIGHT
)

type State struct {
	row   int
	col   int
	dir   Direction
	moves int
}

type Item struct {
	state    State // The value of the item.
	heatLoss int   // The heat loss of the item, can use as priority in queue
	index    int   // The index of the item in the heap.
}

func getNextDirection(direction Direction, row int, col int) (int, int) {
	switch direction {
	case UP:
		return row - 1, col
	case DOWN:
		return row + 1, col
	case LEFT:
		return row, col - 1
	case RIGHT:
		return row, col + 1
	}
	panic("Impossible")
}

type PriorityQueue []*Item

func (pq PriorityQueue) Len() int { return len(pq) }

func (pq PriorityQueue) Less(i, j int) bool {
	return pq[i].heatLoss < pq[j].heatLoss
}

func (pq PriorityQueue) Swap(i, j int) {
	pq[i], pq[j] = pq[j], pq[i]
	pq[i].index = i
	pq[j].index = j
}

func (pq *PriorityQueue) Push(x interface{}) {
	item := x.(*Item)
	item.index = len(*pq)
	*pq = append(*pq, item)
}

func (pq *PriorityQueue) Pop() interface{} {
	old := *pq
	n := len(old)
	item := old[n-1]
	item.index = -1
	*pq = old[0 : n-1]
	return item
}

func neighbours(input [][]int, direction Direction, movesSoFar int, row int, col int) []State {
	if movesSoFar < 4 {
		nextRow, nextCol := getNextDirection(direction, row, col)
		return []State{{row: nextRow, col: nextCol, dir: direction, moves: movesSoFar + 1}}
	}
	if movesSoFar >= 10 {
		if direction == UP || direction == DOWN {
			leftRow, leftCol := getNextDirection(LEFT, row, col)
			rightRow, rightCol := getNextDirection(RIGHT, row, col)
			return []State{{row: leftRow, col: leftCol, dir: LEFT, moves: 1}, {row: rightRow, col: rightCol, dir: RIGHT, moves: 1}}
		} else {
			upRow, upCol := getNextDirection(UP, row, col)
			downRow, downCol := getNextDirection(DOWN, row, col)
			return []State{{row: upRow, col: upCol, dir: UP, moves: 1}, {row: downRow, col: downCol, dir: DOWN, moves: 1}}
		}
	} else {
		upRow, upCol := getNextDirection(UP, row, col)
		downRow, downCol := getNextDirection(DOWN, row, col)
		leftRow, leftCol := getNextDirection(LEFT, row, col)
		rightRow, rightCol := getNextDirection(RIGHT, row, col)
		turnUpState := State{row: upRow, col: upCol, dir: UP, moves: 1}
		turnDownState := State{row: downRow, col: downCol, dir: DOWN, moves: 1}
		turnLeftState := State{row: leftRow, col: leftCol, dir: LEFT, moves: 1}
		turnRightState := State{row: rightRow, col: rightCol, dir: RIGHT, moves: 1}
		if direction == UP {
			return []State{{row: upRow, col: upCol, dir: UP, moves: movesSoFar + 1}, turnLeftState, turnRightState}
		} else if direction == DOWN {
			return []State{{row: downRow, col: downCol, dir: DOWN, moves: movesSoFar + 1}, turnLeftState, turnRightState}
		} else if direction == LEFT {
			return []State{{row: leftRow, col: leftCol, dir: LEFT, moves: movesSoFar + 1}, turnUpState, turnDownState}
		} else {
			return []State{{row: rightRow, col: rightCol, dir: RIGHT, moves: movesSoFar + 1}, turnUpState, turnDownState}
		}
	}
}

func dijkstra(input [][]int, maxConseutive int) int {
	startRight := State{row: 0, col: 0, dir: RIGHT, moves: 1}
	startDown := State{row: 0, col: 0, dir: DOWN, moves: 1}
	firstStateRight := Item{heatLoss: 0, state: startRight, index: 0}
	firstStateDown := Item{heatLoss: 0, state: startDown, index: 1}
	pq := PriorityQueue{
		&firstStateDown,
		&firstStateRight,
	}
	minCost := map[State]int{startRight: 0, startDown: 0}
	for len(pq) > 0 {
		currentItem := heap.Pop(&pq).(*Item)
		if minCost[currentItem.state] < currentItem.heatLoss {
			continue
		}
		if currentItem.state.row == len(input)-1 && currentItem.state.col == len(input[0])-1 && currentItem.state.moves >= 4 {
			return currentItem.heatLoss
		}
		neighbours := neighbours(input, currentItem.state.dir, currentItem.state.moves, currentItem.state.row, currentItem.state.col)
		for _, neighbour := range neighbours {
			if neighbour.row < 0 || neighbour.row >= len(input) || neighbour.col < 0 || neighbour.col >= len(input[0]) {
				continue
			}
			neighbourHeatLoss := input[neighbour.row][neighbour.col]
			if _, ok := minCost[neighbour]; ok && minCost[neighbour] <= currentItem.heatLoss+neighbourHeatLoss {
				continue
			}
			minCost[neighbour] = currentItem.heatLoss + neighbourHeatLoss
			heap.Push(&pq, &Item{heatLoss: currentItem.heatLoss + neighbourHeatLoss, state: neighbour})
		}
	}
	panic("I didn't finish?!")
}

func main() {
	inputstring, _ := os.ReadFile("input.txt")
	lines := strings.Split(string(inputstring), "\n")
	intGrid := [][]int{}
	pq := make(PriorityQueue, 0)
	heap.Init(&pq)
	for _, line := range lines {
		newLine := []int{}
		for _, char := range line {
			newInt, _ := strconv.Atoi(string(char))
			newLine = append(newLine, newInt)
		}
		intGrid = append(intGrid, newLine)
	}
	fmt.Println(dijkstra(intGrid, 3))
}
