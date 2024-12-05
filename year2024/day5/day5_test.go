package day5

import (
    "goated-aoc-2024/year2024"
    "testing"
)

var exampleInput = `47|53
97|13
97|61
97|47
75|29
61|13
75|53
29|13
97|29
53|29
61|53
97|53
61|29
47|13
75|47
97|75
47|61
75|61
47|29
75|13
53|13

75,47,61,53,29
97,61,53,29,13
75,29,13
75,97,47,61,53
61,13,29
97,13,75,29,47`

func TestDayFivePartOneExample(t *testing.T) {
    result := MiddlePageCount(exampleInput, false)
    if result != 143 {
        t.Errorf("Expected 143, got %d", result)
    }
}

func TestDayFivePartOne(t *testing.T) {
    result := MiddlePageCount(year2024.ReadInput("input.txt"), false)
    if result != 5955 {
        t.Errorf("Expected 5955, got %d", result)
    }
}
func TestDayFivePartTwoExample(t *testing.T) {
    result := MiddlePageCount(exampleInput, true)
    if result != 123 {
        t.Errorf("Expected 123, got %d", result)
    }
}
func TestDayFivePartTwo(t *testing.T) {
    result := MiddlePageCount(year2024.ReadInput("input.txt"), true)
    if result != 4030 {
        t.Errorf("Expected 4030, got %d", result)
    }
}
