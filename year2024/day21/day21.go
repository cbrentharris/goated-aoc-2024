package day21

import (
	"fmt"
	"goated-aoc-2024/year2024"
	"math"
	"regexp"
	"strconv"
	"strings"
)

var numberRegex = regexp.MustCompile("\\d+")
var numberGrid = make([][]rune, 4)
var directionalGrid = make([][]rune, 2)

var (
	zero     = '0'
	one      = '1'
	two      = '2'
	three    = '3'
	four     = '4'
	five     = '5'
	six      = '6'
	seven    = '7'
	eight    = '8'
	nine     = '9'
	activate = 'A'
	up       = '^'
	down     = 'v'
	left     = '<'
	right    = '>'
	invalid  = '!'
)

func FindComplexitiesOfShortestButtonPresses(input string, numRobots int) int {
	codes := strings.Split(input, "\n")
	totalComplexity := 0
	populateNumberGrid()
	populateDirectionalGrid()
	for _, code := range codes {
		totalComplexity += findComplexity(code, numRobots)

	}
	return totalComplexity
}

func populateNumberGrid() {
	for y, _ := range numberGrid {
		row := make([]rune, 3)
		numberGrid[y] = row
		switch y {
		case 0:
			row[0] = seven
			row[1] = eight
			row[2] = nine
		case 1:
			row[0] = four
			row[1] = five
			row[2] = six
		case 2:
			row[0] = one
			row[1] = two
			row[2] = three
		case 3:
			row[0] = invalid
			row[1] = zero
			row[2] = activate
		}
	}
}
func populateDirectionalGrid() {
	for y, _ := range directionalGrid {
		row := make([]rune, 3)
		directionalGrid[y] = row
		switch y {
		case 0:
			row[0] = invalid
			row[1] = up
			row[2] = activate
		case 1:
			row[0] = left
			row[1] = down
			row[2] = right
		}
	}
}

func findComplexity(code string, robots int) int {
	numericPortion, _ := strconv.Atoi(numberRegex.FindAllString(code, 1)[0])
	shortestPath := findShortestPathOfButtonPresses([]rune(code), robots)
	return numericPortion * shortestPath
}

type RobotChain struct {
	C             *year2024.Coordinate
	Previous      *RobotChain
	ShortestPaths *map[year2024.Coordinate]DijkstraResult
}

func findShortestPathOfButtonPresses(code []rune, robots int) int {
	numberRuneToCoordinate := make(map[rune]year2024.Coordinate)
	directionRuneToCoordinate := make(map[rune]year2024.Coordinate)
	coordinateToDirectionRune := make(map[year2024.Coordinate]rune)

	for y, row := range numberGrid {
		for x, col := range row {
			if col != invalid {
				numberRuneToCoordinate[col] = year2024.Coordinate{X: x, Y: y}
			}
		}
	}

	allShortestPaths := make(map[year2024.Coordinate]DijkstraResult)
	for y, row := range directionalGrid {
		for x, col := range row {
			if col != invalid {
				coordinate := year2024.Coordinate{X: x, Y: y}
				directionRuneToCoordinate[col] = coordinate
				coordinateToDirectionRune[coordinate] = col
			}
		}
	}

	for y, row := range directionalGrid {
		for x, col := range row {
			if col != invalid {
				//fmt.Printf("%c: %v\n", col, year2024.Coordinate{X: x, Y: y})
				coordinate := year2024.Coordinate{X: x, Y: y}
				allShortestPaths[coordinate] = dijkstras(coordinate, directionalGrid)
			}
		}
	}

	numButtonPresses := 0
	start := numberRuneToCoordinate['A']
	newMemo := make(map[MemoKey]int)
	for _, r := range code {
		current := numberRuneToCoordinate[r]
		numButtonPresses += buttonPressesFromNumericGrid(start, current, numberGrid, &allShortestPaths, &directionRuneToCoordinate, &coordinateToDirectionRune, robots-1, &newMemo)
		start = current
	}
	return numButtonPresses
}

type DijkstraResult struct {
	Dist map[year2024.Coordinate]int
	Prev map[year2024.Coordinate]*[][]year2024.Coordinate
}

func buttonPressesFromNumericGrid(source year2024.Coordinate, destination year2024.Coordinate, grid [][]rune, shortestPaths *map[year2024.Coordinate]DijkstraResult, directionToCoordinate *map[rune]year2024.Coordinate, coordinateToDirection *map[year2024.Coordinate]rune, depth int, memo *map[MemoKey]int) int {
	paths := dijkstras(source, grid)
	minCost := math.MaxInt
	for _, path := range *paths.Prev[destination] {
		var sequence []rune
		var totalCost int
		sequence = buildDirectionalSequenceFromPath(source, path)
		for i, direction := range sequence {
			var previousCoordinate year2024.Coordinate
			if i == 0 {
				previousCoordinate = (*directionToCoordinate)[activate]
			} else {
				previousDirection := sequence[i-1]
				previousCoordinate = (*directionToCoordinate)[previousDirection]
			}
			totalCost += buttonPressesFromDirectionalGrid(previousCoordinate, (*directionToCoordinate)[direction], shortestPaths, directionToCoordinate, coordinateToDirection, depth-1, memo)
		}
		if totalCost < minCost {
			minCost = totalCost
		}
	}
	return minCost
}

type CoordinateAndCount struct {
	Coordinate year2024.Coordinate
	Count      int
	Direction  int32
	Path       *[]year2024.Coordinate
}

func dijkstras(start year2024.Coordinate, grid [][]rune) DijkstraResult {
	dist := make(map[year2024.Coordinate]int)
	prev := make(map[year2024.Coordinate]*[][]year2024.Coordinate)

	dist[start] = 0

	q := year2024.NewMinHeap[CoordinateAndCount](func(a, b CoordinateAndCount) bool {
		return a.Count < b.Count
	})

	for y, row := range grid {
		for x, _ := range row {
			coordinate := year2024.Coordinate{X: x, Y: y}
			if coordinate != start {
				dist[coordinate] = math.MaxInt
			}
		}
	}

	q.Offer(CoordinateAndCount{Coordinate: start, Count: 0, Path: &[]year2024.Coordinate{start}})

	for !q.IsEmpty() {
		current, _ := q.Remove()
		cost := dist[current.Coordinate]
		for _, next := range year2024.AdjacentCoordinates(current.Coordinate, grid) {
			nextIsInvalid := grid[next.Y][next.X] == invalid
			canMove := !nextIsInvalid
			newNextCost := cost + 1
			nextDir := nextDirection(current.Coordinate, next)
			if canMove {
				nextCost := dist[next]
				if newNextCost < nextCost {
					dist[next] = newNextCost
					newPath := make([]year2024.Coordinate, len(*current.Path))
					copy(newPath, *current.Path)
					newPath = append(newPath, next)
					prev[next] = &[][]year2024.Coordinate{newPath}
					q.Offer(CoordinateAndCount{Coordinate: next, Count: newNextCost, Direction: nextDir, Path: &newPath})
				} else if newNextCost == nextCost {
					newPath := make([]year2024.Coordinate, len(*current.Path))
					copy(newPath, *current.Path)
					newOuter := make([][]year2024.Coordinate, len(*prev[next]))
					copy(newOuter, *prev[next])
					newPath = append(newPath, next)
					newOuter = append(newOuter, newPath)
					prev[next] = &newOuter
					q.Offer(CoordinateAndCount{Coordinate: next, Count: newNextCost, Direction: nextDir, Path: &newPath})
				}
			}
		}
	}

	return DijkstraResult{Prev: prev, Dist: dist}
}
func nextDirection(coordinate year2024.Coordinate, next year2024.Coordinate) int32 {
	xDelta := next.X - coordinate.X
	yDelta := next.Y - coordinate.Y
	switch {
	case xDelta > 0:
		return right
	case xDelta < 0:
		return left
	case yDelta > 0:
		return down
	case yDelta < 0:
		return up
	}
	panic("bad state")
}

type MemoKey struct {
	Source      year2024.Coordinate
	Destination year2024.Coordinate
	Depth       int
}

func buttonPressesFromDirectionalGrid(source year2024.Coordinate, destination year2024.Coordinate, shortestPaths *map[year2024.Coordinate]DijkstraResult, directionToCoordinate *map[rune]year2024.Coordinate, coordinateToDirection *map[year2024.Coordinate]rune, depth int, memo *map[MemoKey]int) int {
	if source == destination {
		return 1
	}
	memoKey := MemoKey{Source: source, Destination: destination, Depth: depth}
	result, exists := (*memo)[memoKey]
	if exists {
		return result
	}
	paths := (*shortestPaths)[source].Prev[destination]
	if paths == nil {
		panic(fmt.Errorf("bad path for source: %v, desination: %v -- %v", source, destination, (*shortestPaths)[source]))
	}
	if len(*paths) == 0 {
		panic(fmt.Errorf("no path found from source to destination: source: %v, destination: %v", source, destination))
	}
	if depth == 1 {
		return len((*paths)[0])
	}
	minCost := math.MaxInt
	for _, path := range *paths {
		var sequence []rune
		var totalCost int
		sequence = buildDirectionalSequenceFromPath(source, path)
		for i, direction := range sequence {
			var previousCoordinate year2024.Coordinate
			if i == 0 {
				previousCoordinate = (*directionToCoordinate)[activate]
			} else {
				previousDirection := sequence[i-1]
				previousCoordinate = (*directionToCoordinate)[previousDirection]
			}
			totalCost += buttonPressesFromDirectionalGrid(previousCoordinate, (*directionToCoordinate)[direction], shortestPaths, directionToCoordinate, coordinateToDirection, depth-1, memo)
		}
		if totalCost < minCost {
			minCost = totalCost
		}
	}
	(*memo)[memoKey] = minCost
	return minCost
}

func buildDirectionalSequenceFromPath(source year2024.Coordinate, path []year2024.Coordinate) []rune {
	var currentCoordinate year2024.Coordinate
	var sequence []rune
	for i, coordinate := range path[1:] {
		if i == 0 {
			currentCoordinate = coordinate
			currentDirection := nextDirection(source, coordinate)
			sequence = append(sequence, currentDirection)
		} else {
			currentDirection := nextDirection(currentCoordinate, coordinate)
			sequence = append(sequence, currentDirection)
			currentCoordinate = coordinate
		}
	}
	sequence = append(sequence, activate)
	return sequence
}
