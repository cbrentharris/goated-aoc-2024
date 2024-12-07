package day7

import (
	"goated-aoc-2024/year2024"
	"testing"
)

var exampleInput = `190: 10 19
3267: 81 40 27
83: 17 5
156: 15 6
7290: 6 8 6 15
161011: 16 10 13
192: 17 8 14
21037: 9 7 18 13
292: 11 6 16 20`

func TestDaySevenPartOneExample(t *testing.T) {
	result := GuessCalibrationEquations(exampleInput, false)
	if result != 3749 {
		t.Errorf("Expected 3749 calibration result, found %d", result)
	}
}
func TestDaySevenPartOne(t *testing.T) {
	result := GuessCalibrationEquations(year2024.ReadInput("input.txt"), false)
	if result != 4998764814652 {
		t.Errorf("Expected 4998764814652 calibration result, found %d", result)
	}
}

func TestDaySevenPartTwoExample(t *testing.T) {
	result := GuessCalibrationEquations(exampleInput, true)
	if result != 11387 {
		t.Errorf("Expected 11387 calibration result, found %d", result)
	}
}

func TestDaySevenPartTwo(t *testing.T) {
	result := GuessCalibrationEquations(year2024.ReadInput("input.txt"), true)
	if result != 4998764814652 {
		t.Errorf("Expected 4998764814652 calibration result, found %d", result)
	}
}
