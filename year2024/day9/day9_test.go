package day9

import (
	"goated-aoc-2024/year2024"
	"testing"
)

var exampleInput = "2333133121414131402"

func TestDayOnePartOneExample(t *testing.T) {
	checksum := OptimizeContiguousFreeSpace(exampleInput)
	if checksum != 1928 {
		t.Errorf("Expected checksum 1928, got %d", checksum)
	}
}

func TestDayOnePartOne(t *testing.T) {
	checksum := OptimizeContiguousFreeSpace(year2024.ReadInput("input.txt"))
	if checksum <= 86205133941 {
		t.Errorf("Expected checksum 86205133941, got %d", checksum)
	}
	if checksum != 6398608069280 {
		t.Errorf("Expected checksum 6398608069280, got %d", checksum)
	}
}

func TestDayOnePartOneExampleTwo(t *testing.T) {
	secondExample := "12345"
	result := OptimizeContiguousFreeSpace(secondExample)
	if result != 60 {
		t.Errorf("Expected checksum 30, got %d", result)
	}
}

func TestDayOnePartTwoExample(t *testing.T) {
	checksum := OptimizeContiguousFreeSpaceWithoutFragmentation(exampleInput)
	if checksum != 2858 {
		t.Errorf("Expected checksum 2858, got %d", checksum)
	}
}
func TestDayOnePartTwo(t *testing.T) {
	checksum := OptimizeContiguousFreeSpaceWithoutFragmentation(year2024.ReadInput("input.txt"))
	if checksum != 6427437134372 {
		t.Errorf("Expected checksum 6427437134372, got %d", checksum)
	}
}
