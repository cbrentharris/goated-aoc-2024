package day13

import (
	"goated-aoc-2024/year2024"
	"testing"
)

var example = `Button A: X+94, Y+34
Button B: X+22, Y+67
Prize: X=8400, Y=5400

Button A: X+26, Y+66
Button B: X+67, Y+21
Prize: X=12748, Y=12176

Button A: X+17, Y+86
Button B: X+84, Y+37
Prize: X=7870, Y=6450

Button A: X+69, Y+23
Button B: X+27, Y+71
Prize: X=18641, Y=10279`

func TestDayOnePartOneExample(t *testing.T) {
	result := TokensRequiredToWinAllPrizes(example, false)

	if result != 480 {
		t.Errorf("Expected 480 total tokens used to achieve prizes, found %d", result)
	}
}

func TestDayOnePartOne(t *testing.T) {
	result := TokensRequiredToWinAllPrizes(year2024.ReadInput("input.txt"), false)

	if result != 28262 {
		t.Errorf("Expected 28262 total tokens used to achieve prizes, found %d", result)
	}
}

func TestDayOnePartTwo(t *testing.T) {
	result := TokensRequiredToWinAllPrizes(year2024.ReadInput("input.txt"), true)

	if result != 101406661266314 {
		t.Errorf("Expected 101406661266314 total tokens used to achieve prizes, found %d", result)
	}
}
