package day20

import (
	"goated-aoc-2024/year2024"
	"testing"
)

var exampleInput = `###############
#...#...#.....#
#.#.#.#.#.###.#
#S#...#.#.#...#
#######.#.#.###
#######.#.#...#
#######.#.###.#
###..E#...#...#
###.#######.###
#...###...#...#
#.#####.#.###.#
#.#...#.#.#...#
#.#.#.#.#.#.###
#...#...#...###
###############`

func TestDayTwentyPartOneExample(t *testing.T) {

	result := CheatsV2ThatSaveAtLeast(exampleInput, 20, 2)
	if result != 5 {
		t.Errorf("Expected at least 5 cheats to save 20 picoseconds, instead found %d", result)
	}
}

func TestDayTwentyPartOne(t *testing.T) {
	result := CheatsV2ThatSaveAtLeast(year2024.ReadInput("input.txt"), 100, 2)
	if result != 1289 {
		t.Errorf("Expected at least 1289 cheats to save 100 picoseconds, instead found %d", result)
	}
}

func BenchmarkCheatsThatSaveAtLeast(b *testing.B) {
	for i := 0; i < b.N; i++ {
		CheatsV2ThatSaveAtLeast(year2024.ReadInput("input.txt"), 100, 20)
	}
}

func TestDayTwentyPartTwoExample(t *testing.T) {

	result := CheatsV2ThatSaveAtLeast(exampleInput, 72, 20)

	if result != 29 {
		t.Errorf("Expected at least 29 cheats to save at least 72 picoseconds, instead found %d", result)
	}
}

func TestDayTwentyPartTwo(t *testing.T) {

	result := CheatsV2ThatSaveAtLeast(year2024.ReadInput("input.txt"), 100, 20)
	if result != 982425 {
		t.Errorf("Expected more than 982425 cheats to save 100 picoseconds, instead found %d", result)
	}
}
