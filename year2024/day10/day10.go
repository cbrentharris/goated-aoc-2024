package day10

import (
	"goated-aoc-2024/year2024"
	"strings"
)

type PotentialTrail struct {
	TrailHead year2024.Coordinate
	Current   year2024.Coordinate
}

func ComputeAllTrailHeadScores(input string) int {
	grid := toGrid(input)
	queue := parsePotentialTrails(grid)
	scores := make(map[year2024.Coordinate]int)
	endStates := findEndStates(queue, grid)
	for _, endState := range endStates {
		scores[endState.TrailHead]++
	}
	totalScores := 0
	for _, score := range scores {
		totalScores += score
	}
	return totalScores
}

func ComputeAllTrailHeadScoresV1(input string) int {
	grid := toGrid(input)
	queue := parsePotentialTrails(grid)
	endStates := findEndStates(queue, grid)
	uniqueEndStates := year2024.NewHashSet[PotentialTrail]()
	for _, endState := range endStates {
		uniqueEndStates.Add(endState)
	}
	return uniqueEndStates.Size()
}

func parsePotentialTrails(grid [][]int) year2024.Deque[PotentialTrail] {
	queue := year2024.Deque[PotentialTrail]{}
	for y, row := range grid {
		for x, col := range row {
			if col == 0 {
				current := year2024.Coordinate{X: x, Y: y}
				queue.Enqueue(PotentialTrail{TrailHead: current, Current: current})
			}
		}
	}
	return queue
}

func toGrid(input string) [][]int {
	split := strings.Split(input, "\n")
	grid := make([][]int, len(split))
	for i, row := range split {
		grid[i] = year2024.ToIntSlice(strings.Split(row, ""))
	}
	return grid
}

func findEndStates(queue year2024.Deque[PotentialTrail], grid [][]int) []PotentialTrail {
	var endStates []PotentialTrail
	for !queue.IsEmpty() {
		next, _ := queue.RemoveFirst()
		currentValue := grid[next.Current.Y][next.Current.X]
		if currentValue == 9 {
			endStates = append(endStates, next)
		} else {
			for _, adjacency := range year2024.AdjacentCoordinates(next.Current, grid) {
				adjacencyValue := grid[adjacency.Y][adjacency.X]
				if adjacencyValue == currentValue+1 {
					queue.Enqueue(PotentialTrail{next.TrailHead, adjacency})
				}
			}
		}
	}
	return endStates
}
