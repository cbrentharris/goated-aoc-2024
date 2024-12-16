package day15

import (
	"goated-aoc-2024/year2024"
	"testing"
)

var firstExample = `########
#..O.O.#
##@.O..#
#...O..#
#.#.O..#
#...O..#
#......#
########

<^^>>>vv<v>>v<<`

var secondExample = `##########
#..O..O.O#
#......O.#
#.OO..O.O#
#..O@..O.#
#O#..O...#
#O..O..O.#
#.OO.O.OO#
#....O...#
##########

<vv>^<v^>v>^vv^v>v<>v^v<v<^vv<<<^><<><>>v<vvv<>^v^>^<<<><<v<<<v^vv^v>^
vvv<<^>^v^^><<>>><>^<<><^vv^^<>vvv<>><^^v>^>vv<>v<<<<v<^v>^<^^>>>^<v<v
><>vv>v^v^<>><>>>><^^>vv>v<^^^>>v^v^<^^>v^^>v^<^v>v<>>v^v^<v>v^^<^^vv<
<<v<^>>^^^^>>>v^<>vvv^><v<<<>^^^vv^<vvv>^>v<^^^^v<>^>vvvv><>>v^<<^^^^^
^><^><>>><>^^<<^^v>>><^<v>^<vv>>v>>>^v><>^v><<<<v>>v<v<v>vvv>^<><<>^><
^>><>^v<><^vvv<^^<><v<<<<<><^v<<<><<<^^<v<^^^><^>>^<v^><<<^>>^v<v^v<v^
>^>>^v>vv>^<<^v<>><<><<v<<v><>v<^vv<<<>^^v^>^^>>><<^v>>v^v><^^>>^<>vv^
<><^^>^^^<><vvvvv^v<v<<>^v<v>v<<^><<><<><<<^^<<<^<<>><<><^^^>^^<>^>v<>
^^>vv<^v^v<vv>^<><v<^v>^^^>>>^^vvv^>vvv<>>>^<^>>>>>^<<^v>^vvv<>^<><<v>
v^^>>><<^^<>>^v^<v^vv<>v^<<>^<^v^v><^<<<><<^<v><v<>vv>>v><v^<vv<>v^<<^`

func TestDayFifteenPartOneExampleOne(t *testing.T) {
	result := SumGPSCoordinates(firstExample)
	if result != 2028 {
		t.Errorf("Expected GPS coordinate sum in first example to equal 2028, instead found %d", result)
	}
}

func TestDayFifteenPartOneExampleTwo(t *testing.T) {
	result := SumGPSCoordinates(secondExample)
	if result != 10092 {
		t.Errorf("Expected GPS coordinate sum in first example to equal 10092, instead found %d", result)
	}
}

func TestDayFifteenPartOne(t *testing.T) {
	result := SumGPSCoordinates(year2024.ReadInput("input.txt"))
	if result != 1383666 {
		t.Errorf("Expected GPS coordinate sum in first example to equal 1383666, instead found %d", result)
	}
}

func TestDayFifteenPartTwoExampleOne(t *testing.T) {
	result := SumGPSCoordinatesExpanded(secondExample, false)
	if result != 9021 {
		t.Errorf("Expected GPS coordinate sum in first example to equal 9021, instead found %d", result)
	}
}
func TestDayFifteenPartTwo(t *testing.T) {
	result := SumGPSCoordinatesExpanded(year2024.ReadInput("input.txt"), false)
	if result != 1412866 {
		t.Errorf("Expected GPS coordinate sum in first example to equal 1412866, instead found %d", result)
	}
}
