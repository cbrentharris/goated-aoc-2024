package day16

import (
	"fmt"
	"goated-aoc-2024/year2024"
	"math"
	"strings"
)

var (
	startMarker = 'S'
	endMarker   = 'E'
	obstacle    = '#'
	empty       = '.'
	up          = '^'
	down        = 'v'
	left        = '<'
	right       = '>'
)

type MinValue struct {
	Value int
}

func BestPathScore(input string) (int, int) {
	grid := year2024.ToRuneSlice(strings.Split(input, "\n"))
	start := year2024.Coordinate{X: 1, Y: len(grid) - 2}
	end := year2024.Coordinate{X: len(grid[0]) - 2, Y: 1}
	validateStartAndEnd(start, end, grid)
	numObstacles := 0
	for _, row := range grid {
		for _, col := range row {
			if col == obstacle {
				numObstacles++
			}
		}
	}
	memo := make(map[MemoKeyForLowestScore]MemoValue)
	lowestScore, _ := findLowestScore(start, end, grid, 0, right, &memo, numObstacles)

	coordinatesMemo := make(map[MemoKeyForFindingPoints]bool)
	coordinates := year2024.NewHashSet[year2024.Coordinate]()
	findAllCoordinatesToLowestScore(start, end, grid, 0, right, &coordinatesMemo, numObstacles, lowestScore, 0, coordinates)
	return lowestScore, coordinates.Size()
}

func validateStartAndEnd(start year2024.Coordinate, end year2024.Coordinate, grid [][]rune) {
	if grid[start.Y][start.X] != startMarker {
		panic("bad grid input")
	}
	if grid[end.Y][end.X] != endMarker {
		panic("bad grid input")
	}
}

type MemoKeyForLowestScore struct {
	Coordinate   year2024.Coordinate
	Direction    int32
	VisitedDepth int
}

type MemoValue struct {
	Score    int
	CanReach bool
}

func findLowestScore(current year2024.Coordinate, end year2024.Coordinate, grid [][]rune, depth int, currentDirection int32, memo *map[MemoKeyForLowestScore]MemoValue, numObstacles int) (int, bool) {
	if current == end {
		return 0, true
	}

	currentValue := grid[current.Y][current.X]

	if currentValue == obstacle {
		return math.MaxInt, false
	}

	if year2024.OffTheMap(current, grid) {
		return math.MaxInt, false
	}

	tooDeep := depth >= (len(grid)*len(grid[0])-numObstacles)/2
	if tooDeep {
		return math.MaxInt, false
	}

	depth++
	memoKey := MemoKeyForLowestScore{Coordinate: current, Direction: currentDirection, VisitedDepth: depth}

	value, contains := (*memo)[memoKey]
	if contains {
		return value.Score, value.CanReach
	}

	minPath := math.MaxInt
	clockwiseCoordinate, clockwiseDirection := rotateClockwise(current, currentDirection)
	counterClockwiseCoordinate, counterClockwiseDirection := rotateCounterClockwise(current, currentDirection)
	straight := nextCoordinate(currentDirection, current)

	clockwiseMin, clockwiseCanReach := findLowestScore(clockwiseCoordinate, end, grid, depth, clockwiseDirection, memo, numObstacles)
	counterClockwiseMin, counterClockwiseCanReach := findLowestScore(counterClockwiseCoordinate, end, grid, depth, counterClockwiseDirection, memo, numObstacles)
	straightMin, straightCanReach := findLowestScore(straight, end, grid, depth, currentDirection, memo, numObstacles)

	if !clockwiseCanReach && !counterClockwiseCanReach && !straightCanReach {
		memoValue := MemoValue{Score: math.MaxInt, CanReach: false}
		(*memo)[memoKey] = memoValue
		return memoValue.Score, false
	}

	if clockwiseCanReach {
		clockwiseMin += 1 + 1000
	}

	if counterClockwiseCanReach {
		counterClockwiseMin += 1 + 1000
	}

	if straightCanReach {
		straightMin += 1
	}

	minPath = min(clockwiseMin, counterClockwiseMin, straightMin)
	memoValue := MemoValue{Score: minPath, CanReach: true}
	(*memo)[memoKey] = memoValue

	return minPath, true
}

func findAllCoordinatesToLowestScore(current year2024.Coordinate, end year2024.Coordinate, grid [][]rune, depth int, currentDirection int32, memo *map[MemoKeyForFindingPoints]bool, numObstacles int, target int, runningTotal int, visited *year2024.HashSet[year2024.Coordinate]) bool {
	if current == end && runningTotal == target {
		visited.Add(current)
		return true
	}

	if current == end {
		return false
	}

	if runningTotal > target {
		return false
	}

	currentValue := grid[current.Y][current.X]

	if currentValue == obstacle {
		// Hit an obstacle
		return false
	}

	if year2024.OffTheMap(current, grid) {
		// Off the map
		return false
	}

	visitedDepth := depth
	depth++
	memoKey := MemoKeyForFindingPoints{Coordinate: current, Direction: currentDirection, Depth: visitedDepth, Score: runningTotal}

	value, contains := (*memo)[memoKey]
	if contains {
		return value
	}

	clockwiseCoordinate, clockwiseDirection := rotateClockwise(current, currentDirection)
	counterClockwiseCoordinate, counterClockwiseDirection := rotateCounterClockwise(current, currentDirection)
	straight := nextCoordinate(currentDirection, current)

	clockwiseCanReach := findAllCoordinatesToLowestScore(clockwiseCoordinate, end, grid, depth, clockwiseDirection, memo, numObstacles, target, runningTotal+1+1000, visited)
	counterClockwiseCanReach := findAllCoordinatesToLowestScore(counterClockwiseCoordinate, end, grid, depth, counterClockwiseDirection, memo, numObstacles, target, runningTotal+1+1000, visited)
	straightCanReach := findAllCoordinatesToLowestScore(straight, end, grid, depth, currentDirection, memo, numObstacles, target, runningTotal+1, visited)

	canReach := clockwiseCanReach || counterClockwiseCanReach || straightCanReach
	if canReach {
		(*memo)[memoKey] = true
		visited.Add(current)
		return true
	} else {
		(*memo)[memoKey] = false
		return false
	}
}

type MemoKeyForFindingPoints struct {
	Coordinate year2024.Coordinate
	Direction  int32
	Depth      int
	Score      int
}

func rotateClockwise(coordinate year2024.Coordinate, direction int32) (year2024.Coordinate, int32) {
	switch direction {
	case up:
		return year2024.Coordinate{X: coordinate.X + 1, Y: coordinate.Y}, right
	case down:
		return year2024.Coordinate{X: coordinate.X - 1, Y: coordinate.Y}, left
	case left:
		return year2024.Coordinate{X: coordinate.X, Y: coordinate.Y - 1}, up
	case right:
		return year2024.Coordinate{X: coordinate.X, Y: coordinate.Y + 1}, down
	default:
		panic(fmt.Sprintf("unknown direction %c", direction))
	}
}

func rotateCounterClockwise(coordinate year2024.Coordinate, direction int32) (year2024.Coordinate, int32) {
	switch direction {
	case up:
		return year2024.Coordinate{X: coordinate.X - 1, Y: coordinate.Y}, left
	case down:
		return year2024.Coordinate{X: coordinate.X + 1, Y: coordinate.Y}, right
	case left:
		return year2024.Coordinate{X: coordinate.X, Y: coordinate.Y + 1}, down
	case right:
		return year2024.Coordinate{X: coordinate.X, Y: coordinate.Y - 1}, up
	default:
		panic(fmt.Sprintf("unknown direction %c", direction))
	}
}

func nextCoordinate(direction int32, coordinate year2024.Coordinate) year2024.Coordinate {
	switch direction {
	case up:
		return year2024.Coordinate{X: coordinate.X, Y: coordinate.Y - 1}
	case right:
		return year2024.Coordinate{X: coordinate.X + 1, Y: coordinate.Y}
	case down:
		return year2024.Coordinate{X: coordinate.X, Y: coordinate.Y + 1}
	case left:
		return year2024.Coordinate{X: coordinate.X - 1, Y: coordinate.Y}
	}
	panic(fmt.Sprintf("Unknown direction: %c", direction))
}
