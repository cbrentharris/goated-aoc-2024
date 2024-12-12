package day12

import (
	"goated-aoc-2024/year2024"
	"testing"
)

var (
	example1 = `AAAA
BBCD
BBCC
EEEC`
	example2 = `OOOOO
OXOXO
OOOOO
OXOXO
OOOOO`
	example3 = `RRRRIICCFF
RRRRIICCCF
VVRRRCCFFF
VVRCCCJFFF
VVVVCJJCFE
VVIVCCJJEE
VVIIICJJEE
MIIIIIJJEE
MIIISIJEEE
MMMISSJEEE`
)

func TestDayTwelvePartOneExample(t *testing.T) {
	result := CalculateFenceCost(example1, false)
	if result != 140 {
		t.Errorf("Expected cost of 140 for example 1, found %d", result)
	}
}

func TestDayTwelvePartOne(t *testing.T) {
	result := CalculateFenceCost(year2024.ReadInput("input.txt"), false)
	if result != 1446042 {
		t.Errorf("Expected cost of 1446042 for example 1, found %d", result)
	}
}

func TestDayTwelvePartTwoExample(t *testing.T) {
	result := CalculateFenceCost(example1, true)
	if result != 80 {
		t.Errorf("Expected cost of 80 for example 1, found %d", result)
	}
}

func TestDayTwelvePartTwoExampleTwo(t *testing.T) {
	partTwoSecondExample := `EEEEE
EXXXX
EEEEE
EXXXX
EEEEE`
	result := CalculateFenceCost(partTwoSecondExample, true)
	if result != 236 {
		t.Errorf("Expected cost of 236 for example 2, found %d", result)
	}
}

func TestDayTwelvePartTwoExampleThree(t *testing.T) {
	exampleThree := `AAAAAA
AAABBA
AAABBA
ABBAAA
ABBAAA
AAAAAA`
	result := CalculateFenceCost(exampleThree, true)
	if result != 368 {
		t.Errorf("Expected cost of 368 for example 3, found %d", result)
	}
}

func TestDayTwelvePartTwoExampleFour(t *testing.T) {
	result := CalculateFenceCost(example3, true)
	if result != 1206 {
		t.Errorf("Expected cost of 1206 for example 4, found %d", result)
	}
}

func TestDayTwelvePartTwoExampleFive(t *testing.T) {
	result := CalculateFenceCost(example2, true)
	if result != 436 {
		t.Errorf("Expected cost of 436 for example five, found %d", result)
	}
}

func TestDayTwelvePartTwo(t *testing.T) {
	result := CalculateFenceCost(year2024.ReadInput("input.txt"), true)
	if result != 902742 {
		t.Errorf("Expected cost to be 902742 for part two, found %d", result)
	}
}
