package day20

import (
	"goated-aoc-2024/year2024"
	"math"
	"strings"
	"sync"
)

var (
	start = 'S'
	end   = 'E'
	track = '.'
	wall  = '#'
)

type Cheat struct {
	Start         year2024.Coordinate
	End           year2024.Coordinate
	SecondsPassed int
	StartCost     int
}

type Task struct {
	Coordinate     year2024.Coordinate
	Grid           *[][]rune
	Dist           *sync.Map
	MaxCheatLength int
}

func CheatsV2ThatSaveAtLeast(input string, picoseconds int, maxCheatLength int) int {
	lines := strings.Split(input, "\n")

	grid := make([][]rune, len(lines))

	for i, line := range lines {
		grid[i] = []rune(line)
	}

	var startCoordinate year2024.Coordinate

	for y, row := range grid {
		for x, col := range row {
			coordinate := year2024.Coordinate{X: x, Y: y}
			switch col {
			case start:
				startCoordinate = coordinate
			}
		}
	}

	dist := dijkstras(startCoordinate, grid)

	var sm sync.Map

	for k, v := range dist {
		sm.Store(k, v)
	}
	cheats := make(chan Cheat)
	numWorkers := 8
	tasks := make(chan Task, numWorkers*2)
	var wg sync.WaitGroup
	for i := 0; i < numWorkers; i++ {
		wg.Add(1)
		go func() {
			defer wg.Done()
			for task := range tasks {
				findPossibleCheats(task.Coordinate, task.Grid, task.Dist, task.MaxCheatLength, &cheats)
			}
		}()
	}

	go func() {
		for y, row := range grid {
			for x, col := range row {
				coordinate := year2024.Coordinate{X: x, Y: y}
				if col == start || col == track {
					tasks <- Task{
						Coordinate:     coordinate,
						Grid:           &grid,
						Dist:           &sm,
						MaxCheatLength: maxCheatLength,
					}
				}
			}
		}
		close(tasks)
	}()

	total := 0
	go func() {
		wg.Wait()
		close(cheats)
	}()
	processedCheats := year2024.NewHashSet[Cheat]()
	for cheat := range cheats {
		if processedCheats.Contains(cheat) {
			continue
		}
		processedCheats.Add(cheat)
		delta := calcChangeWithCheat(cheat, dist)
		if delta >= picoseconds {
			total++
		}
	}

	return total
}

type CoordinateAndCount struct {
	Coordinate year2024.Coordinate
	Count      int
}

func dijkstras(start year2024.Coordinate, grid [][]rune) map[year2024.Coordinate]int {
	dist := make(map[year2024.Coordinate]int)

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

	q.Offer(CoordinateAndCount{Coordinate: start, Count: 0})

	for !q.IsEmpty() {
		current, _ := q.Remove()
		cost := dist[current.Coordinate]
		for _, next := range year2024.AdjacentCoordinates(current.Coordinate, grid) {
			nextIsTrack := grid[next.Y][next.X] == track
			nextIsEnd := grid[next.Y][next.X] == end
			canMove := nextIsTrack || nextIsEnd

			if canMove {
				nextCost := dist[next]
				if cost+1 < nextCost {
					dist[next] = cost + 1
					q.Offer(CoordinateAndCount{Coordinate: next, Count: cost + 1})
				}
			}
		}
	}

	return dist
}

func findPossibleCheats(coordinate year2024.Coordinate, grid *[][]rune, dist *sync.Map, maxCheatLength int, cheats *chan Cheat) {
	distanceToCurrentCoordinate, _ := dist.Load(coordinate)
	currentCheatLength := 2

	for currentCheatLength <= maxCheatLength {
		xDelta := 0
		yDelta := currentCheatLength

		for xDelta <= currentCheatLength {
			xPlus := coordinate.X + xDelta
			xMinus := coordinate.X - xDelta
			yPlus := coordinate.Y + yDelta
			yMinus := coordinate.Y - yDelta
			for _, c := range []year2024.Coordinate{
				{X: xPlus, Y: yPlus},
				{X: xPlus, Y: yMinus},
				{X: xMinus, Y: yPlus},
				{X: xMinus, Y: yMinus},
			} {
				if year2024.OffTheMap2(c, grid) {
					continue
				}
				value := (*grid)[c.Y][c.X]
				if value == start || value == wall {
					continue
				}
				manHattanDistance := year2024.ManhattanDistance(coordinate, c)
				distanceWithoutCheating, _ := dist.Load(c)
				if manHattanDistance+distanceToCurrentCoordinate.(int) < distanceWithoutCheating.(int) && manHattanDistance <= maxCheatLength {
					*cheats <- Cheat{Start: coordinate, End: c, StartCost: distanceToCurrentCoordinate.(int), SecondsPassed: manHattanDistance}
				}
			}
			xDelta++
			yDelta--
		}
		currentCheatLength++
	}
}

func calcChangeWithCheat(cheat Cheat, dist map[year2024.Coordinate]int) int {
	withoutCheat := dist[cheat.End]
	startCost := dist[cheat.Start]
	return withoutCheat - (startCost + cheat.SecondsPassed)
}
