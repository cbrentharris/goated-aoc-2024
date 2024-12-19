package day19

import (
	"goated-aoc-2024/year2024"
	"testing"
)

var exampleInput = `r, wr, b, g, bwu, rb, gb, br

brwrr
bggr
gbbr
rrbgbr
ubwu
bwurrg
brgr
bbrgwb`

func TestDayNineteenPartOneExample(t *testing.T) {
	result := CountPossibleDesigns(exampleInput)

	if result != 6 {
		t.Errorf("Expected 6 possible designs, found %d", result)
	}
}

func TestDayNineteenPartOne(t *testing.T) {
	result := CountPossibleDesigns(year2024.ReadInput("input.txt"))

	if result != 317 {
		t.Errorf("Expected 317 possible designs, found %d", result)
	}
}

func TestDayNineteenPartTwoExample(t *testing.T) {
	result := CountPossibleDistinctDesigns(exampleInput)

	if result != 16 {
		t.Errorf("Expected 16 possible designs, found %d", result)
	}
}

func TestDayNineteenPartTwo(t *testing.T) {
	result := CountPossibleDistinctDesigns(year2024.ReadInput("input.txt"))

	if result != 883443544805484 {
		t.Errorf("Expected 883443544805484 possible designs, found %d", result)
	}
}

func BenchmarkCountPossibleDistinctDesigns(b *testing.B) {
	for i := 0; i < b.N; i++ {
		CountPossibleDistinctDesigns(year2024.ReadInput("input.txt"))
	}
}
