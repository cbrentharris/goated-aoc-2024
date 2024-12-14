package day14

import (
	"goated-aoc-2024/year2024"
	"testing"
)

var example = `p=0,4 v=3,-3
p=6,3 v=-1,-3
p=10,3 v=-1,2
p=2,0 v=2,-1
p=0,0 v=1,3
p=3,0 v=-2,-2
p=7,6 v=-1,-3
p=3,0 v=-1,-2
p=9,3 v=2,3
p=7,3 v=-1,2
p=2,4 v=2,-3
p=9,5 v=-3,-3`

func TestDayOnePartOneExample(t *testing.T) {
	result := CalculateSafetyScore(example, 11, 7, 100)

	if result != 12 {
		t.Errorf("Expected safety score of 12, found %d instead.", result)
	}
}

func TestDayOnePartOne(t *testing.T) {
	result := CalculateSafetyScore(year2024.ReadInput("input.txt"), 101, 103, 100)

	if result <= 218699784 {
		t.Errorf("Expected safety score to be higher than 218699784, found %d instead.", result)
	}
	if result != 231019008 {
		t.Errorf("Expected safety score to be higher than 231019008, found %d instead.", result)
	}
}

func TestDayOnePartTwo(t *testing.T) {
	result := FindChristmasTree(year2024.ReadInput("input.txt"), 101, 103)
	if result != 8280 {
		t.Errorf("Expected seconds of finding christmas tree to be 8280, instead got %d", result)
	}
}
