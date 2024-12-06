package day6

import (
	"fmt"
	"goated-aoc-2024/year2024"
	"strings"
)

var (
	up          = '^'
	down        = 'v'
	left        = '<'
	right       = '>'
	obstruction = '#'
	empty       = '.'
)

func CountGuardMovements(input string) int {
	rows := strings.Split(input, "\n")
	grid := make([][]rune, len(rows))
	for rowIndex, row := range rows {
		grid[rowIndex] = []rune(row)
	}
	direction, coordinate := findInitialCoordinateAndDirection(grid)
	coordinatesTraversed := make(map[year2024.Coordinate]struct{})
	coordinatesTraversed[coordinate] = struct{}{}
	for !offTheMap(coordinate, grid) {
		direction, coordinate = nextDirectionAndCoordinate(direction, coordinate, grid)
		coordinatesTraversed[coordinate] = struct{}{}
	}
	return len(coordinatesTraversed) - 1 // last coordinate was off the map
}

func nextDirection(direction int32) int32 {
	switch direction {
	case up:
		return right
	case right:
		return down
	case down:
		return left
	case left:
		return up
	}
	panic(fmt.Sprintf("Unknown direction: %c", direction))
}

func nextDirectionAndCoordinate(direction int32, coordinate year2024.Coordinate, grid [][]rune) (int32, year2024.Coordinate) {
	nextCoord := nextCoordinate(direction, coordinate)
	if offTheMap(nextCoord, grid) {
		return direction, nextCoord
	}
	nextRune := grid[nextCoord.Y][nextCoord.X]
	if nextRune == obstruction {
		return nextDirection(direction), coordinate
	} else {
		return direction, nextCoord
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

func offTheMap(coordinate year2024.Coordinate, grid [][]rune) bool {
	return coordinate.Y >= len(grid) || coordinate.Y < 0 || coordinate.X >= len(grid[0]) || coordinate.X < 0
}

func findInitialCoordinateAndDirection(grid [][]rune) (int32, year2024.Coordinate) {
	for y, row := range grid {
		for x, col := range row {
			if col == up || col == down || col == left || col == right {
				return col, year2024.Coordinate{X: x, Y: y}
			}
		}
	}
	panic("Unable to find starting position in grid.")
}

func CountPositionsThatCreateLoops(input string) int {
	rows := strings.Split(input, "\n")
	grid := make([][]rune, len(rows))
	for rowIndex, row := range rows {
		grid[rowIndex] = []rune(row)
	}

	totalLoops := 0
	for y, row := range grid {
		for x, col := range row {
			if col != empty {
				continue
			}

			if canCreateLoop(year2024.Coordinate{X: x, Y: y}, grid) {
				totalLoops++
			}
		}
	}

	return totalLoops
}

func canCreateLoop(loopCandidate year2024.Coordinate, grid [][]rune) bool {
	grid[loopCandidate.Y][loopCandidate.X] = obstruction
	loopDetected := false
	direction, coordinate := findInitialCoordinateAndDirection(grid)
	totalSteps := 0
	for !offTheMap(coordinate, grid) {
		direction, coordinate = nextDirectionAndCoordinate(direction, coordinate, grid)
		totalSteps++
		if totalSteps > 2*len(grid)*len(grid[0]) {
			loopDetected = true
			break
		}
	}
	grid[loopCandidate.Y][loopCandidate.X] = empty
	return loopDetected
}
