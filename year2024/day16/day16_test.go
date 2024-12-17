package day16

import (
	"goated-aoc-2024/year2024"
	"testing"
)

var (
	exampleOne = `###############
#.......#....E#
#.#.###.#.###.#
#.....#.#...#.#
#.###.#####.#.#
#.#.#.......#.#
#.#.#####.###.#
#...........#.#
###.#.#####.#.#
#...#.....#.#.#
#.#.#.###.#.#.#
#.....#...#.#.#
#.###.#.#.#.#.#
#S..#.....#...#
###############`
	exampleTwo = `#################
#...#...#...#..E#
#.#.#.#.#.#.#.#.#
#.#.#.#...#...#.#
#.#.#.#.###.#.#.#
#...#.#.#.....#.#
#.#.#.#.#.#####.#
#.#...#.#.#.....#
#.#.#####.#.###.#
#.#.#.......#...#
#.#.###.#####.###
#.#.#...#.....#.#
#.#.#.#####.###.#
#.#.#.........#.#
#.#.#.#########.#
#S#.............#
#################`
)

func TestDaySixteenPartOneExample(t *testing.T) {
	result, _ := BestPathScore(exampleOne)
	if result != 7036 {
		t.Errorf("Expected best path for example one to have score of 7036, instead found %d", result)
	}
}

func TestDaySixteenPartOneExampleTwo(t *testing.T) {
	result, _ := BestPathScore(exampleTwo)
	if result != 11048 {
		t.Errorf("Expected best path for example one to have score of 11048, instead found %d", result)
	}
}

func TestDaySixteenPartTwoExample(t *testing.T) {
	_, result := BestPathScore(exampleOne)
	if result != 45 {
		t.Errorf("Expected best path for example one to have score of 45, instead found %d", result)
	}
}

func TestDaySixteenPartTwoExampleTwo(t *testing.T) {
	_, result := BestPathScore(exampleTwo)
	if result != 64 {
		t.Errorf("Expected best path for example one to have score of 11048, instead found %d", result)
	}
}

func TestDaySixteenPartOne(t *testing.T) {
	result, _ := BestPathScore(year2024.ReadInput("input.txt"))
	if result != 72428 {
		t.Errorf("Expected best path for example one to have score of 72428, instead found %d", result)
	}
}

func TestDaySixteenPartTwo(t *testing.T) {
	_, result := BestPathScore(year2024.ReadInput("input.txt"))
	// 428 wrong
	if result != 456 {
		t.Errorf("Expected best path for example one to have score of 72428, instead found %d", result)
	}
}

func BenchmarkTestDaySixteen(b *testing.B) {
	for i := 0; i < b.N; i++ {
		BestPathScore(year2024.ReadInput("input.txt"))
	}
}
