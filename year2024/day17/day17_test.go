package day17

import (
	"goated-aoc-2024/year2024"
	"testing"
)

var exampleInput = `Register A: 729
Register B: 0
Register C: 0

Program: 0,1,5,4,3,0`

var example2 = `Register A: 2024
Register B: 0
Register C: 0

Program: 0,3,5,4,3,0`

func TestDaySeventeenPartOneExample(t *testing.T) {
	result := ExecuteProgram(exampleInput)
	if result != "4,6,3,5,6,3,5,2,1,0" {
		t.Errorf("Expected output 4,6,3,5,6,3,5,2,1,0, but instead got %s", result)
	}
}

func TestDaySeventeenPartOne(t *testing.T) {
	result := ExecuteProgram(year2024.ReadInput("input.txt"))
	if result != "1,5,0,1,7,4,1,0,3" {
		t.Errorf("Expected output 1,5,0,1,7,4,1,0,3 but instead got %s", result)
	}
}

func TestDaySeventeenPartTwoExample(t *testing.T) {
	result := FindCorrectRegisterValue(example2)
	if result != 117440 {
		t.Errorf("Expected output 117440 but instead got %d", result)
	}
}

func TestDaySeventeenPartTwo(t *testing.T) {
	result := FindCorrectRegisterValue(year2024.ReadInput("input.txt"))
	if result != 47910079998866 {
		t.Errorf("Expected output 47910079998866 but instead got %d", result)
	}
}
