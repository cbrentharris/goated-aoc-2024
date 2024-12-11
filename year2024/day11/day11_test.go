package day11

import (
	"goated-aoc-2024/year2024"
	"testing"
)

var exampleInput = `125 17`

func TestDayElevenPartOneExample(t *testing.T) {
	result := CountStones(exampleInput, 25)
	if result != 55312 {
		t.Errorf("Expected 55312 stones, found %d", result)
	}
}

func TestDayElevenPartOneTest(t *testing.T) {
	result := CountStones(exampleInput, 6)
	if result != 22 {
		t.Errorf("Expected 55312 stones, found %d", result)
	}
}

func TestDayElevenPartOne(t *testing.T) {
	result := CountStones(year2024.ReadInput("input.txt"), 25)
	if result != 216996 {
		t.Errorf("Expected 216996 stones, found %d", result)
	}
}

func TestDayElevenPartTwo(t *testing.T) {
	result := CountStones(year2024.ReadInput("input.txt"), 75)
	if result != 257335372288947 {
		t.Errorf("Expected 257335372288947 stones, found %d", result)
	}
}

func BenchmarkCountStones(b *testing.B) {
	for i := 0; i < b.N; i++ {
		CountStones(year2024.ReadInput("input.txt"), 75)
	}
}
