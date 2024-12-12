package day12

import (
	"goated-aoc-2024/year2024"
	"strings"
)

type Region struct {
	Coordinates *year2024.HashSet[year2024.Coordinate]
	Plant       int32
}

func (r *Region) Area() int {
	return r.Coordinates.Size()
}

func (r *Region) Perimeter() int {
	perimeter := 0
	for coordinate := range r.Coordinates.Iterator() {
		adjacencies := year2024.AdjacentCoordinatesIncludingOffTheMap(coordinate)
		for _, adjacency := range adjacencies {
			if !r.Coordinates.Contains(adjacency) {
				perimeter++
			}
		}
	}
	return perimeter
}

func (r *Region) Cost() int {
	return r.Area() * r.Perimeter()
}

func (r *Region) CostWithSides() int {
	return r.Area() * r.Corners()
}

func (r *Region) Corners() int {
	edges := r.Edges()
	total := 0
	for edge := range edges.Iterator() {
		up := year2024.Coordinate{X: edge.X, Y: edge.Y - 1}
		down := year2024.Coordinate{X: edge.X, Y: edge.Y + 1}
		left := year2024.Coordinate{X: edge.X - 1, Y: edge.Y}
		right := year2024.Coordinate{X: edge.X + 1, Y: edge.Y}
		upLeft := year2024.Coordinate{X: edge.X - 1, Y: edge.Y - 1}
		upRight := year2024.Coordinate{X: edge.X + 1, Y: edge.Y - 1}
		downLeft := year2024.Coordinate{X: edge.X - 1, Y: edge.Y + 1}
		downRight := year2024.Coordinate{X: edge.X + 1, Y: edge.Y + 1}
		corners := 0
		for _, coordinates := range [][]year2024.Coordinate{
			{up, left, upLeft},
			{up, right, upRight},
			{down, left, downLeft},
			{down, right, downRight}} {
			// A corner is defined as:
			// Any time two adjacencies are not in the segment, i.e.
			//   B A B
			//   B C B
			//   B B B
			// left of C is B, top of B is A, thats a corner.
			// Additionally, a corner can be if it has two adjacencies but the diagnoal is different, so
			//  B A B
			//  A A B
			//  C C C
			// So the middle A has top A and left A, but top left is B, so thats a corner
			doesNotHaveTheFirstTwo := !(r.Coordinates.Contains(coordinates[0]) || r.Coordinates.Contains(coordinates[1]))
			if doesNotHaveTheFirstTwo {
				corners++
			}
			hasTheFirstTwo := r.Coordinates.Contains(coordinates[0]) && r.Coordinates.Contains(coordinates[1])
			if !r.Coordinates.Contains(coordinates[2]) {
				if hasTheFirstTwo {
					corners++
				}
			}
		}
		total += corners
	}
	return total
}

func (r *Region) Edges() *year2024.HashSet[year2024.Coordinate] {
	edges := year2024.NewHashSet[year2024.Coordinate]()
	for coordinate := range r.Coordinates.Iterator() {
		adjacencies := year2024.FullAdjacenciesIncludingOffTheMap(coordinate)
		for _, adjacency := range adjacencies {
			if !r.Coordinates.Contains(adjacency) {
				edges.Add(coordinate)
			}
		}
	}
	return edges
}

func CalculateFenceCost(input string, useSides bool) int {
	visited := year2024.NewHashSet[year2024.Coordinate]()
	grid := year2024.ToRuneSlice(strings.Split(input, "\n"))

	regions := year2024.NewHashSet[Region]()

	for y, row := range grid {
		for x, plant := range row {
			coordinate := year2024.Coordinate{X: x, Y: y}
			region := Region{Coordinates: year2024.NewHashSet[year2024.Coordinate](), Plant: plant}
			crawlRegion(plant, coordinate, grid, visited, &region)
			if region.Coordinates.Size() > 0 {
				regions.Add(region)
			}
		}
	}

	totalCost := 0
	for region := range regions.Iterator() {
		if useSides {
			totalCost += region.CostWithSides()
		} else {
			totalCost += region.Cost()
		}
	}

	return totalCost
}

func crawlRegion(plant rune, coordinate year2024.Coordinate, grid [][]rune, visited *year2024.HashSet[year2024.Coordinate], region *Region) {
	if visited.Contains(coordinate) {
		return
	}
	if plant == region.Plant {
		region.Coordinates.Add(coordinate)
		visited.Add(coordinate)
		adjacencies := year2024.AdjacentCoordinates(coordinate, grid)
		for _, adjacency := range adjacencies {
			adjacencyPlant := grid[adjacency.Y][adjacency.X]
			crawlRegion(adjacencyPlant, adjacency, grid, visited, region)
		}
	}
}
