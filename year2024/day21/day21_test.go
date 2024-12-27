package day21

import (
	"goated-aoc-2024/year2024"
	"testing"
)

var exampleInput = `029A
980A
179A
456A
379A`

func TestDayTwentyOnePartOneExample(t *testing.T) {
	result := FindComplexitiesOfShortestButtonPresses(exampleInput, 4)
	if result != 126384 {
		t.Errorf("Expected 126834 complexity for example input, instead found %d", result)
	}
}

func TestDayTwentyOnePartOneExampleOne(t *testing.T) {
	result := FindComplexitiesOfShortestButtonPresses("029A", 4)
	if result != 68*29 {
		t.Errorf("Expected 68 length for 029A, instead found %d", result/29)
	}
}

func TestDayTwentyOnePartOneExampleTwo(t *testing.T) {
	result := FindComplexitiesOfShortestButtonPresses("379A", 4)
	if result != 64*379 {
		t.Errorf("Expected 64 length for 379A, instead found %d", result/379)
	}
}

func TestDayTwentyOnePartOneExampleThree(t *testing.T) {
	result := FindComplexitiesOfShortestButtonPresses("980A", 4)
	if result != 60*980 {
		t.Errorf("Expected 60 length for 980A, instead found %d", result/980)
	}
}

func TestDayTwentyOnePartOneExampleFour(t *testing.T) {
	result := FindComplexitiesOfShortestButtonPresses("179A", 4)
	if result != 68*179 {
		t.Errorf("Expected 68 length for 980A, instead found %d", result/179)
	}
}

func TestDayTwentyOnePartOneExampleFive(t *testing.T) {
	result := FindComplexitiesOfShortestButtonPresses("456A", 4)
	if result != 64*456 {
		t.Errorf("Expected 64 length for 456A, instead found %d", result/456)
	}
}

func TestDayTwentyOnePartOne(t *testing.T) {
	result := FindComplexitiesOfShortestButtonPresses(year2024.ReadInput("input.txt"), 4)
	if result != 128962 {
		t.Errorf("Expected 128962 complexity for input, instead found %d", result)
	}
}

func TestDayTwentyOnePartTwo(t *testing.T) {
	result := FindComplexitiesOfShortestButtonPresses(year2024.ReadInput("input.txt"), 27)
	if result != 159684145150108 {
		t.Errorf("Expected 159684145150108 complexity for input, instead found %d", result)
	}
}

func BenchmarkFindComplexitiesOfShortestButtonPresses(b *testing.B) {
	for i := 0; i < b.N; i++ {
		FindComplexitiesOfShortestButtonPresses(year2024.ReadInput("input.txt"), 4)
	}
}
