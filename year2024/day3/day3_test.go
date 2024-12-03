package day3

import (
    "goated-aoc-2024/year2024"
    "testing"
)

func TestDayOnePartOneExample(t *testing.T) {
    exampleInput := "xmul(2,4)%&mul[3,7]!@^do_not_mul(5,5)+mul(32,64]then(mul(11,8)mul(8,5))"
    result := CorruptedProgramExecutor(exampleInput)

    if result != 161 {
        t.Errorf("Expected 161 as an answer for the example input, got %d", result)
    }
}

func TestDayOnePartOne(t *testing.T) {
    result := CorruptedProgramExecutor(year2024.ReadInput("input.txt"))

    if result != 160672468 {
        t.Errorf("Expected 160672468 as an answer for the input, got %d", result)
    }
}

func TestDayOnePartTwoExample(t *testing.T) {
    exampleInput := "xmul(2,4)&mul[3,7]!^don't()_mul(5,5)+mul(32,64](mul(11,8)undo()?mul(8,5))"
    result := CorruptedProgramExecutorV2(exampleInput)

    if result != 48 {
        t.Errorf("Expected 48 as an answer for the example input, got %d", result)
    }
}

func TestDayOnePartTwo(t *testing.T) {
    result := CorruptedProgramExecutorV2(year2024.ReadInput("input.txt"))

    if result != 84893551 {
        t.Errorf("Expected 84893551 as an answer for the input, got %d", result)
    }
}
