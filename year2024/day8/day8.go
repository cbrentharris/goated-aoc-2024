package day8

import (
	"goated-aoc-2024/year2024"
	"strings"
)

var empty = '.'

func CountAntiNodeLocations(input string, incorporateResonantHarmonics bool) int {
	lines := strings.Split(input, "\n")
	grid := year2024.ToRuneSlice(lines)
	antennaToCoordinates := make(map[rune][]year2024.Coordinate)
	antiNodeCoordinates := make(map[year2024.Coordinate]struct{})
	for y, row := range grid {
		for x, col := range row {
			if col == empty {
				continue
			}
			antennaToCoordinates[col] = append(antennaToCoordinates[col], year2024.Coordinate{X: x, Y: y})
		}
	}

	for _, coordinates := range antennaToCoordinates {
		for i, a := range coordinates {
			for j, b := range coordinates {
				if i == j {
					continue
				}
				for _, antiNode := range computeAntiNodeCoordinates(a, b, grid, incorporateResonantHarmonics) {
					antiNodeCoordinates[antiNode] = struct{}{}
				}
			}
		}
	}

	return len(antiNodeCoordinates)
}

func computeAntiNodeCoordinates(a, b year2024.Coordinate, grid [][]rune, incorporateResonantHarmonics bool) []year2024.Coordinate {
	var antiNodeCoordinates []year2024.Coordinate
	if incorporateResonantHarmonics {
		antiNodeCoordinates = []year2024.Coordinate{a, b}
	}
	xDelta := year2024.Abs(a.X, b.X)
	yDelta := year2024.Abs(a.Y, b.Y)

	moveBottomRight := func(coordinate year2024.Coordinate) year2024.Coordinate {
		return year2024.Coordinate{X: coordinate.X + xDelta, Y: coordinate.Y + yDelta}
	}
	moveTopLeft := func(coordinate year2024.Coordinate) year2024.Coordinate {
		return year2024.Coordinate{X: coordinate.X - xDelta, Y: coordinate.Y - yDelta}
	}
	moveBottomLeft := func(coordinate year2024.Coordinate) year2024.Coordinate {
		return year2024.Coordinate{X: coordinate.X - xDelta, Y: coordinate.Y + yDelta}
	}
	moveTopRight := func(coordinate year2024.Coordinate) year2024.Coordinate {
		return year2024.Coordinate{X: coordinate.X + xDelta, Y: coordinate.Y - yDelta}
	}
	generateAntiNodes := func(coordinate year2024.Coordinate, nextAntiNodeGenerator func(c year2024.Coordinate) year2024.Coordinate) {
		nextAntiNode := nextAntiNodeGenerator(coordinate)
		for !year2024.OffTheMap(nextAntiNode, grid) {
			antiNodeCoordinates = append(antiNodeCoordinates, nextAntiNode)
			nextAntiNode = nextAntiNodeGenerator(nextAntiNode)
			if !incorporateResonantHarmonics {
				break
			}
		}
	}
	switch {
	case a.X > b.X && a.Y > b.Y:
		// diagonal slope up left, a is in the bottom right
		// for the bottom right, we have to add the x delta and the y delta
		// for the top left, we have to subtract the x delta and the y delta
		generateAntiNodes(a, moveBottomRight)
		generateAntiNodes(b, moveTopLeft)
	case a.X > b.X && a.Y < b.Y:
		// diagonal slope up right, a is in the top right
		// for the top right, we have to subtract the y delta and add the x delta
		// for the bottom left, we have to add the y delta and subtract the x delta
		generateAntiNodes(a, moveTopRight)
		generateAntiNodes(b, moveBottomLeft)
	case a.X < b.X && a.Y > b.Y:
		// diagonal slope up right, as in the bottom left
		// for the top right, we have to subtract the y delta and add the x delta
		// for the bottom left, we have to add the y delta and subtract the x delta
		generateAntiNodes(b, moveTopRight)
		generateAntiNodes(a, moveBottomLeft)
	case a.X < b.X && a.Y < b.Y:
		// diagonal slope up left, a is in the top left
		// for the bottom right, we have to add the x delta and the y delta
		// for the top left, we have to subtract the x delta and the y delta
		generateAntiNodes(b, moveBottomRight)
		generateAntiNodes(a, moveTopLeft)
	}
	return antiNodeCoordinates
}
