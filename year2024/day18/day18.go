package day18

import (
	"fmt"
	"goated-aoc-2024/year2024"
	"math"
	"strconv"
	"strings"
)

type CoordinateAndSteps struct {
	Coordinate year2024.Coordinate
	Steps      int
	Visited    *year2024.HashSet[year2024.Coordinate]
}

func MinStepsToExitMemoryLocation(input string, gridSize int, bytesToSimulate int) int {
	lines := strings.Split(input, "\n")
	coordinates := make([]year2024.Coordinate, len(lines))
	for i, line := range lines {
		split := strings.Split(line, ",")
		x, _ := strconv.Atoi(split[0])
		y, _ := strconv.Atoi(split[1])
		coordinates[i] = year2024.Coordinate{X: x, Y: y}
	}
	coordinatesPerStep := make(map[int]*year2024.HashSet[year2024.Coordinate])
	for i := 1; i <= len(coordinates); i++ {
		coordinatesPerStep[i] = year2024.NewHashSet[year2024.Coordinate]()
		for j := 0; j < i; j++ {
			coordinatesPerStep[i].Add(coordinates[j])
		}
	}
	start := year2024.Coordinate{X: 0, Y: 0}
	fakeGrid := make([][]rune, gridSize)
	fakeGrid[0] = make([]rune, gridSize)
	exit := year2024.Coordinate{X: gridSize - 1, Y: gridSize - 1}
	current := CoordinateAndSteps{Coordinate: start, Steps: 0, Visited: year2024.NewHashSet[year2024.Coordinate]()}
	current.Visited.Add(start)
	fallen := coordinatesPerStep[bytesToSimulate]
	return dijkstras(start, exit, fakeGrid, fallen)
}

func CoordinateThatCutsOff(input string, gridSize int, bytesToSimulate int) string {
	lines := strings.Split(input, "\n")
	coordinates := make([]year2024.Coordinate, len(lines))
	for i, line := range lines {
		split := strings.Split(line, ",")
		x, _ := strconv.Atoi(split[0])
		y, _ := strconv.Atoi(split[1])
		coordinates[i] = year2024.Coordinate{X: x, Y: y}
	}
	coordinatesPerStep := make(map[int]*year2024.HashSet[year2024.Coordinate])
	for i := 1; i <= len(coordinates); i++ {
		coordinatesPerStep[i] = year2024.NewHashSet[year2024.Coordinate]()
		for j := 0; j < i; j++ {
			coordinatesPerStep[i].Add(coordinates[j])
		}
	}
	start := year2024.Coordinate{X: 0, Y: 0}
	fakeGrid := make([][]rune, gridSize)
	fakeGrid[0] = make([]rune, gridSize)
	exit := year2024.Coordinate{X: gridSize - 1, Y: gridSize - 1}
	current := CoordinateAndSteps{Coordinate: start, Steps: 0, Visited: year2024.NewHashSet[year2024.Coordinate]()}
	current.Visited.Add(start)
	neighborsMatch := true
	startIndex := bytesToSimulate + 1
	end := len(coordinates)
	delta := end - startIndex
	mid := startIndex + delta/2
	for neighborsMatch {
		delta = end - startIndex
		mid = startIndex + delta/2
		left := mid - 1
		right := mid
		leftBlocks := false
		rightBlocks := false
		if dijkstras(start, exit, fakeGrid, coordinatesPerStep[left]) == math.MaxInt {
			leftBlocks = true
		}
		if dijkstras(start, exit, fakeGrid, coordinatesPerStep[right]) == math.MaxInt {
			rightBlocks = true
		}
		neighborsMatch = leftBlocks == rightBlocks
		if neighborsMatch {
			if leftBlocks {
				end = mid
			} else {
				startIndex = mid
			}
		}
	}
	c := coordinates[mid-1]
	return fmt.Sprintf("%d,%d", c.X, c.Y)
}

type CoordinateSteps struct {
	Coordinate year2024.Coordinate
	Steps      int
}

func dijkstras(start, end year2024.Coordinate, grid [][]rune, fallen *year2024.HashSet[year2024.Coordinate]) int {
	dist := make([][]int, len(grid))
	heap := year2024.NewMinHeap[CoordinateSteps](func(a, b CoordinateSteps) bool {
		return a.Steps < b.Steps
	})
	heap.Offer(CoordinateSteps{Coordinate: start, Steps: 0})
	for y := 0; y < len(grid); y++ {
		dist[y] = make([]int, len(grid[0]))
		for x := 0; x < len(grid[0]); x++ {
			c := year2024.Coordinate{X: x, Y: y}
			if c == start {
				continue
			}
			dist[y][x] = math.MaxInt
		}
	}
	for !heap.IsEmpty() {
		current, _ := heap.Remove()
		for _, adj := range year2024.AdjacentCoordinates(current.Coordinate, grid) {
			if fallen.Contains(adj) {
				continue
			}
			currentCost := dist[current.Coordinate.Y][current.Coordinate.X]
			nextCost := dist[adj.Y][adj.X]
			if currentCost+1 < nextCost {
				dist[adj.Y][adj.X] = currentCost + 1
				heap.Offer(CoordinateSteps{Coordinate: adj, Steps: currentCost + 1})
			}
		}
	}
	return dist[end.Y][end.X]
}

func bfs(start, end year2024.Coordinate, fallen *year2024.HashSet[year2024.Coordinate], grid [][]rune) int {
	level := year2024.NewHashSet[year2024.Coordinate]()
	steps := 0
	visited := year2024.NewHashSet[year2024.Coordinate]()
	level.Add(start)
	visited.Add(start)
	for !level.Contains(end) {
		nextLevel := year2024.NewHashSet[year2024.Coordinate]()
		for c := range level.Iterator() {
			for _, adj := range year2024.AdjacentCoordinates(c, grid) {
				if visited.Contains(adj) {
					continue
				}
				if fallen.Contains(adj) {
					continue
				}
				visited.Add(adj)
				nextLevel.Add(adj)
			}
		}
		level = nextLevel
		steps++
	}
	return steps
}
