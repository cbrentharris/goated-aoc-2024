package day6

import (
	"goated-aoc-2024/year2024"
	"testing"
)

var exampleInput = `....#.....
.........#
..........
..#.......
.......#..
..........
.#..^.....
........#.
#.........
......#...`

func TestDaySixPartOneExample(t *testing.T) {
	result := CountGuardMovements(exampleInput)
	if result != 41 {
		t.Errorf("Expected 41 guard movements, found %d", result)
	}
}

func TestDaySixPartOne(t *testing.T) {
	result := CountGuardMovements(year2024.ReadInput("input.txt"))
	if result != 4374 {
		t.Errorf("Expected 4374 guard movements, found %d", result)
	}
}

func TestDaySixPartTwoExample(t *testing.T) {
	result := CountPositionsThatCreateLoops(exampleInput)
	if result != 6 {
		t.Errorf("Expected 6 loops, found %d", result)
	}
}

func TestDaySixPartTwo(t *testing.T) {
	result := CountPositionsThatCreateLoops(year2024.ReadInput("input.txt"))
	if result != 1705 {
		t.Errorf("Expected 1705 loops, found %d", result)
	}
}
