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
	up          = '^'
	down        = 'v'
	left        = '<'
	right       = '>'
)

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
	lowestScore, size := dijkstras(start, end, grid)

	return lowestScore, size
}

func validateStartAndEnd(start year2024.Coordinate, end year2024.Coordinate, grid [][]rune) {
	if grid[start.Y][start.X] != startMarker {
		panic("bad grid input")
	}
	if grid[end.Y][end.X] != endMarker {
		panic("bad grid input")
	}
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

type CoordinateAndDirection struct {
	Coordinate year2024.Coordinate
	Distance   int
	Direction  int32
}

func dijkstras(start, end year2024.Coordinate, grid [][]rune) (int, int) {
	heap := year2024.NewMinHeap[CoordinateAndDirection](func(a, b CoordinateAndDirection) bool {
		return a.Distance < b.Distance
	})
	heap.Offer(CoordinateAndDirection{Coordinate: start, Distance: 0, Direction: right})
	dist := make(map[CoordinateAndDirection]int)
	dist[CoordinateAndDirection{Coordinate: start, Direction: right}] = 0
	prev := make(map[CoordinateAndDirection]*year2024.HashSet[CoordinateAndDirection])
	for y, row := range grid {
		for x, col := range row {
			if col == obstacle {
				continue
			}
			c := year2024.Coordinate{X: x, Y: y}
			if c == start {
				continue
			}
			for _, d := range []int32{up, left, right, down} {
				dist[CoordinateAndDirection{Coordinate: c, Direction: d}] = math.MaxInt
			}
		}
	}
	for !heap.IsEmpty() {
		current, _ := heap.Remove()
		clockwiseCoordinate, clockwiseDirection := rotateClockwise(current.Coordinate, current.Direction)
		counterClockwiseCoordinate, counterClockwiseDirection := rotateCounterClockwise(current.Coordinate, current.Direction)
		straight := nextCoordinate(current.Direction, current.Coordinate)
		currentCost, _ := dist[CoordinateAndDirection{Coordinate: current.Coordinate, Direction: current.Direction}]
		newClockwiseCost := currentCost + 1000 + 1
		newCounterClockwiseCost := currentCost + 1000 + 1
		newStraightCost := currentCost + 1
		clockwiseKey := CoordinateAndDirection{Coordinate: clockwiseCoordinate, Direction: clockwiseDirection}
		counterClockwiseKey := CoordinateAndDirection{Coordinate: counterClockwiseCoordinate, Direction: counterClockwiseDirection}
		straightKey := CoordinateAndDirection{Coordinate: straight, Direction: current.Direction}
		clockwiseCost, clockwiseExists := dist[clockwiseKey]
		counterClockwiseCost, counterClockwiseExists := dist[counterClockwiseKey]
		straightCost, straightExists := dist[straightKey]
		if clockwiseExists && clockwiseCost > newClockwiseCost {
			dist[clockwiseKey] = newClockwiseCost
			heap.Offer(CoordinateAndDirection{Coordinate: clockwiseCoordinate, Direction: clockwiseDirection, Distance: newClockwiseCost})
			prev[clockwiseKey] = year2024.NewHashSet[CoordinateAndDirection]()
			prev[clockwiseKey].Add(current)
		}
		if clockwiseCost == newClockwiseCost {
			prev[clockwiseKey].Add(current)
		}
		if counterClockwiseExists && counterClockwiseCost > newCounterClockwiseCost {
			dist[counterClockwiseKey] = newCounterClockwiseCost
			heap.Offer(CoordinateAndDirection{Coordinate: counterClockwiseCoordinate, Direction: counterClockwiseDirection, Distance: newCounterClockwiseCost})
			prev[counterClockwiseKey] = year2024.NewHashSet[CoordinateAndDirection]()
			prev[counterClockwiseKey].Add(current)
		}
		if counterClockwiseCost == newCounterClockwiseCost {
			prev[counterClockwiseKey].Add(current)
		}
		if straightExists && straightCost > newStraightCost {
			dist[straightKey] = newStraightCost
			heap.Offer(CoordinateAndDirection{Coordinate: straight, Direction: current.Direction, Distance: newStraightCost})
			prev[straightKey] = year2024.NewHashSet[CoordinateAndDirection]()
			prev[straightKey].Add(current)
		}
		if straightCost == newStraightCost {
			prev[straightKey].Add(current)
		}
	}
	minEnd := math.MaxInt
	for k, v := range dist {
		if k.Coordinate == end && v < minEnd {
			minEnd = v
		}
	}
	coordinates := year2024.NewHashSet[year2024.Coordinate]()
	for k, v := range dist {
		if v == minEnd && k.Coordinate == end {
			queue := year2024.Deque[CoordinateAndDirection]{}
			queue.Enqueue(k)
			for !queue.IsEmpty() {
				current, _ := queue.RemoveFirst()
				coordinates.Add(current.Coordinate)
				prevKey := CoordinateAndDirection{Coordinate: current.Coordinate, Direction: current.Direction}
				p, exists := prev[prevKey]
				if exists {
					for c := range p.Iterator() {
						queue.Enqueue(c)
					}
				}
			}
		}
	}
	return minEnd, coordinates.Size()
}
