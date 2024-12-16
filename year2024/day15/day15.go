package day15

import (
	"fmt"
	"goated-aoc-2024/year2024"
	"strings"
)

var (
	up           = '^'
	down         = 'v'
	left         = '<'
	right        = '>'
	robot        = '@'
	box          = 'O'
	edge         = '#'
	empty        = '.'
	boxLeftHalf  = '['
	boxRightHalf = ']'
)

func SumGPSCoordinates(input string) int {
	grid, directions := createBasicGridAndDirections(input)
	var initialRobotCoordinate year2024.Coordinate
	for y, row := range grid {
		for x, col := range row {
			if col == robot {
				initialRobotCoordinate = year2024.Coordinate{X: x, Y: y}
			}
		}
	}

	for _, direction := range directions {
		initialRobotCoordinate = move(initialRobotCoordinate, direction, grid)
	}

	gpsCoordinateSum := 0
	for y, row := range grid {
		for x, col := range row {
			if col == box {
				gpsCoordinateSum += y*100 + x
			}
		}
	}
	return gpsCoordinateSum
}

func SumGPSCoordinatesExpanded(input string, printWarehouse bool) int {
	grid, directions := createBasicGridAndDirections(input)

	expandedGrid := make([][]rune, len(grid))
	var initialRobotCoordinate year2024.Coordinate
	for y, row := range grid {
		expandedGrid[y] = make([]rune, len(row)*2)
		for x, col := range row {
			firstIndex := x * 2
			secondIndex := x*2 + 1
			var firstValue, secondValue rune
			switch col {
			case robot:
				initialRobotCoordinate = year2024.Coordinate{X: firstIndex, Y: y}
				firstValue = robot
				secondValue = empty
			case empty:
				firstValue = empty
				secondValue = empty
			case box:
				firstValue = boxLeftHalf
				secondValue = boxRightHalf
			case edge:
				firstValue = edge
				secondValue = edge
			default:
				panic("unknown character when expanding grid.")
			}
			expandedGrid[y][firstIndex] = firstValue
			expandedGrid[y][secondIndex] = secondValue
		}
	}

	for _, direction := range directions {
		if printWarehouse {
			for _, row := range expandedGrid {
				for _, col := range row {
					fmt.Printf("%c", col)
				}
				fmt.Println("")
			}
			fmt.Printf("Next direction: %c\n", direction)
		}
		initialRobotCoordinate = moveExpanded(initialRobotCoordinate, direction, expandedGrid)
	}

	gpsCoordinateSum := 0
	for y, row := range expandedGrid {
		for x, col := range row {
			if col == boxLeftHalf {
				gpsCoordinateSum += x
			}
			if col == boxRightHalf {
				gpsCoordinateSum += y * 100
			}
		}
	}
	return gpsCoordinateSum
}

func moveExpanded(coordinate year2024.Coordinate, direction rune, grid [][]rune) year2024.Coordinate {
	next := nextCoordinate(direction, coordinate)
	value := grid[next.Y][next.X]
	if value == edge {
		return coordinate
	} else if value == empty {
		grid[coordinate.Y][coordinate.X] = empty
		grid[next.Y][next.X] = robot
		return next
	} else {
		if direction == left || direction == right {
			nextEmptySpace := findEmptySpace(direction, next, grid)
			if nextEmptySpace == nil {
				return coordinate
			}
			grid[coordinate.Y][coordinate.X] = empty
			// When moving left or right, our position is like
			// @[]
			// So our next position, if our position is X, is X + 1 (or X -1)
			// This means, we must move [ or ] to X + 2 or X - 2, hence tmp = next of next
			// We create tmp, because we need to retain next to eventually return
			// as the next robot value
			tmp := nextCoordinate(direction, next)
			for tmp != nextCoordinate(direction, *nextEmptySpace) {
				switch direction {
				case left:
					grid[tmp.Y][tmp.X] = boxRightHalf
					tmp = nextCoordinate(direction, tmp)
					grid[tmp.Y][tmp.X] = boxLeftHalf
					tmp = nextCoordinate(direction, tmp)
				case right:
					grid[tmp.Y][tmp.X] = boxLeftHalf
					tmp = nextCoordinate(direction, tmp)
					grid[tmp.Y][tmp.X] = boxRightHalf
					tmp = nextCoordinate(direction, tmp)
				}
			}
			grid[next.Y][next.X] = robot
			return next
		} else {
			pushed := moveStackedBoxes(direction, coordinate, grid)
			if pushed {
				return next
			} else {
				return coordinate
			}
		}
	}
}

type RuneAndCoordinate struct {
	Rune       int32
	Coordinate year2024.Coordinate
}

func moveStackedBoxes(direction rune, coordinate year2024.Coordinate, grid [][]rune) bool {
	// Going along the y-axis here, either up or down
	initialElementsPartOfTheStack := year2024.Deque[year2024.Coordinate]{}
	initialElementsPartOfTheStack.Enqueue(coordinate)
	allElementsPartOfTheStack := year2024.Deque[RuneAndCoordinate]{}

	coordinatesToClear := year2024.NewHashSet[year2024.Coordinate]()
	coordinatesToClear.Add(coordinate)

	for !initialElementsPartOfTheStack.IsEmpty() {
		c, _ := initialElementsPartOfTheStack.RemoveFirst()
		currentValue := grid[c.Y][c.X]
		next := nextCoordinate(direction, c)
		nextValue := grid[next.Y][next.X]
		switch nextValue {
		case boxLeftHalf:
			initialElementsPartOfTheStack.Enqueue(next)
			if currentValue == boxRightHalf || currentValue == robot {
				initialElementsPartOfTheStack.Enqueue(year2024.Coordinate{X: next.X + 1, Y: next.Y})
				coordinatesToClear.Add(year2024.Coordinate{X: next.X + 1, Y: next.Y})
			}
			allElementsPartOfTheStack.Enqueue(RuneAndCoordinate{Rune: currentValue, Coordinate: c})
		case boxRightHalf:
			initialElementsPartOfTheStack.Enqueue(next)
			if currentValue == boxLeftHalf || currentValue == robot {
				initialElementsPartOfTheStack.Enqueue(year2024.Coordinate{X: next.X - 1, Y: next.Y})
				coordinatesToClear.Add(year2024.Coordinate{X: next.X - 1, Y: next.Y})
			}
			allElementsPartOfTheStack.Enqueue(RuneAndCoordinate{Rune: currentValue, Coordinate: c})
		case empty:
			allElementsPartOfTheStack.Enqueue(RuneAndCoordinate{Rune: currentValue, Coordinate: c})
		case robot:
			initialElementsPartOfTheStack.Enqueue(next)
			allElementsPartOfTheStack.Enqueue(RuneAndCoordinate{Rune: currentValue, Coordinate: c})
		case edge:
			// If we ever encounter an edge on the stack, cannot move
			return false
		default:
			panic("Unknown nextValue")
		}
	}

	nextStack := year2024.NewHashSet[year2024.Coordinate]()

	for !allElementsPartOfTheStack.IsEmpty() {
		c, _ := allElementsPartOfTheStack.RemoveFirst()
		next := nextCoordinate(direction, c.Coordinate)
		grid[next.Y][next.X] = c.Rune
		nextStack.Add(next)
	}

	for c := range coordinatesToClear.Iterator() {
		if !nextStack.Contains(c) {
			grid[c.Y][c.X] = empty
		}
	}
	return true
}

func move(coordinate year2024.Coordinate, direction rune, grid [][]rune) year2024.Coordinate {
	next := nextCoordinate(direction, coordinate)
	value := grid[next.Y][next.X]
	if value == edge {
		return coordinate
	} else if value == empty {
		grid[coordinate.Y][coordinate.X] = empty
		grid[next.Y][next.X] = robot
		return next
	} else {
		nextEmptySpace := findEmptySpace(direction, next, grid)
		if nextEmptySpace == nil {
			return coordinate
		}
		grid[coordinate.Y][coordinate.X] = empty
		grid[nextEmptySpace.Y][nextEmptySpace.X] = box
		grid[next.Y][next.X] = robot
		return next
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

func findEmptySpace(direction int32, coordinate year2024.Coordinate, grid [][]rune) *year2024.Coordinate {
	next := nextCoordinate(direction, coordinate)
	value := grid[next.Y][next.X]
	for value != edge {
		if value == empty {
			return &next
		}
		next = nextCoordinate(direction, next)
		value = grid[next.Y][next.X]
	}
	return nil
}

func createBasicGridAndDirections(input string) ([][]rune, []rune) {
	lines := strings.Split(input, "\n")
	wareHouseBorder := lines[0]
	lastIndexOfWarehouse := -1

	for index, line := range lines {
		if index == 0 {
			continue
		}
		if line == wareHouseBorder {
			lastIndexOfWarehouse = index
			break
		}
	}
	grid := year2024.ToRuneSlice(lines[:lastIndexOfWarehouse+1])
	directions := []rune(strings.Join(lines[lastIndexOfWarehouse+1:], ""))
	return grid, directions
}
