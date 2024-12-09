package day8

import (
	"goated-aoc-2024/year2024"
	"testing"
)

var exampleInput = `............
........0...
.....0......
.......0....
....0.......
......A.....
............
............
........A...
.........A..
............
............`

func TestDayEightPartOneExample(t *testing.T) {
	result := CountAntiNodeLocations(exampleInput, false)
	if result != 14 {
		t.Errorf("Expected 14 unique antinode locations for the example input, found %d", result)
	}
}

func TestDayEightPartOne(t *testing.T) {
	result := CountAntiNodeLocations(year2024.ReadInput("input.txt"), false)
	if result != 336 {
		t.Errorf("Expected 336 unique antinode locations for the example input, found %d", result)
	}
}

func TestDayEightPartTwoExample(t *testing.T) {
	result := CountAntiNodeLocations(exampleInput, true)
	if result != 34 {
		t.Errorf("Expected 34 unique antinode locations for the example input, found %d", result)
	}
}

func TestDayEightPartTwo(t *testing.T) {
	result := CountAntiNodeLocations(year2024.ReadInput("input.txt"), true)
	if result != 1131 {
		t.Errorf("Expected 1131 unique antinode locations for the example input, found %d", result)
	}
}
