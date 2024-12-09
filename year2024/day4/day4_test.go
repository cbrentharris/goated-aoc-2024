package day4

import (
	"goated-aoc-2024/year2024"
	"testing"
)

func TestDayFourPartOneExample(t *testing.T) {
	exampleInput := `MMMSXXMASM
MSAMXMSMSA
AMXSXMAAMM
MSAMASMSMX
XMASAMXAMM
XXAMMXXAMA
SMSMSASXSS
SAXAMASAAA
MAMMMXMMMM
MXMXAXMASX`
	result := WordSearch(exampleInput)
	if result != 18 {
		t.Errorf("Expected word search to be 18, found %d", result)
	}
}

func TestDayFourPartOne(t *testing.T) {
	result := WordSearch(year2024.ReadInput("input.txt"))

	if result != 2344 {
		t.Errorf("Expected word search to be 2344, found %d", result)
	}
}

func TestDayFourPartTwoExample(t *testing.T) {
	exampleInput := `MMMSXXMASM
MSAMXMSMSA
AMXSXMAAMM
MSAMASMSMX
XMASAMXAMM
XXAMMXXAMA
SMSMSASXSS
SAXAMASAAA
MAMMMXMMMM
MXMXAXMASX`

	result := WordSearchV2(exampleInput)
	if result != 9 {
		t.Errorf("Expected word search to be 9, found %d", result)
	}
}

func TestDayFourPartTwo(t *testing.T) {
	result := WordSearchV2(year2024.ReadInput("input.txt"))
	if result != 1815 {
		t.Errorf("Expected word search to be 1815, found %d", result)
	}
}
