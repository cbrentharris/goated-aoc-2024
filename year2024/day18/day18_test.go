package day18

import (
	"goated-aoc-2024/year2024"
	"testing"
)

var exampleInput = `5,4
4,2
4,5
3,0
2,1
6,3
2,4
1,5
0,6
3,3
2,6
5,1
1,2
5,5
2,5
6,5
1,4
0,4
6,4
1,1
6,1
1,0
0,5
1,6
2,0`

func TestDayEighteenPartOneExampmle(t *testing.T) {
	result := MinStepsToExitMemoryLocation(exampleInput, 7, 12)

	if result != 22 {
		t.Errorf("Expected 22 steps to exit memory location, instead found %d", result)
	}
}

func TestDayEighteenPartOne(t *testing.T) {
	result := MinStepsToExitMemoryLocation(year2024.ReadInput("input.txt"), 71, 1024)

	if result != 308 {
		t.Errorf("Expected 308 steps to exit memory location, instead found %d", result)
	}
}

func BenchmarkMinStepsToExitMemoryLocation(b *testing.B) {
	for i := 0; i < b.N; i++ {
		MinStepsToExitMemoryLocation(year2024.ReadInput("input.txt"), 71, 1024)
	}
}

func TestDayEighteenPartTwoExample(t *testing.T) {
	result := CoordinateThatCutsOff(exampleInput, 7, 12)

	if result != "6,1" {
		t.Errorf("Expected coordinate 6,1 to cut off, instead found %s", result)
	}
}

func TestDayEighteenPartTwo(t *testing.T) {
	result := CoordinateThatCutsOff(year2024.ReadInput("input.txt"), 71, 1024)

	if result != "46,28" {
		t.Errorf("Expected coordinate 46,28 to cut off, instead found %s", result)
	}
}
