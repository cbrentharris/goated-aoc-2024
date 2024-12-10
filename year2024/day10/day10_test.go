package day10

import (
	"goated-aoc-2024/year2024"
	"testing"
)

var exampleInput = `89010123
78121874
87430965
96549874
45678903
32019012
01329801
10456732`

func TestDayTenPartOneExample(t *testing.T) {
	totalScores := ComputeAllTrailHeadScoresV1(exampleInput)
	if totalScores != 36 {
		t.Errorf("Expected a total score of 36, found %d", totalScores)
	}
}

func TestDayTenPartOne(t *testing.T) {
	totalScores := ComputeAllTrailHeadScoresV1(year2024.ReadInput("input.txt"))
	if totalScores != 644 {
		t.Errorf("Expected a total score of 644, found %d", totalScores)
	}
}

func TestDayTenPartTwoExample(t *testing.T) {
	totalScores := ComputeAllTrailHeadScores(exampleInput)
	if totalScores != 81 {
		t.Errorf("Expected a total score of 81, found %d", totalScores)
	}
}

func TestDayTenPartTwo(t *testing.T) {
	totalScores := ComputeAllTrailHeadScores(year2024.ReadInput("input.txt"))
	if totalScores != 1366 {
		t.Errorf("Expected a total score of 1366, found %d", totalScores)
	}
}
