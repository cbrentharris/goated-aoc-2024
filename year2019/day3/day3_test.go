package day3

import (
	"strings"
	"testing"
)

func TestPartOne(t *testing.T) {
	input, err := DayOne()
	if err != nil {
		t.Error(err)
		return
	}
	expected := 248
	if input != expected {
		t.Errorf("Expected %d, got %d", expected, input)
	}
}

func TestPartTwo(t *testing.T) {
	input, err := DayTwo()
	if err != nil {
		t.Error(err)
		return
	}
	expected := 28580
	if input != expected {
		t.Errorf("Expected %d, got %d", expected, input)
	}
}

func TestPartTwoExample(t *testing.T) {
	exampleWireOne := strings.Split("R75,D30,R83,U83,L12,D49,R71,U7,L72", ",")
	exampleWireTwo := strings.Split("U62,R66,U55,R34,D71,R55,D58,R83", ",")
	intersectionPoints := findIntersectionPoints(exampleWireOne, exampleWireTwo)
	minimumDistance := Min(findStepDistances(exampleWireOne, exampleWireTwo, intersectionPoints))
	if minimumDistance != 610 {
		t.Errorf("Expected %d to be 610.", minimumDistance)
	}
}
func TestPartTwoExampleOne(t *testing.T) {
	exampleWireOne := strings.Split("R8,U5,L5,D3", ",")
	exampleWireTwo := strings.Split("U7,R6,D4,L4", ",")
	intersectionPoints := findIntersectionPoints(exampleWireOne, exampleWireTwo)
	minimumDistance := Min(findStepDistances(exampleWireOne, exampleWireTwo, intersectionPoints))
	if minimumDistance != 30 {
		t.Errorf("Expected %d to be 30.", minimumDistance)
	}
}
func TestPartTwoExampleThree(t *testing.T) {
	exampleWireOne := strings.Split("R98,U47,R26,D63,R33,U87,L62,D20,R33,U53,R51", ",")
	exampleWireTwo := strings.Split("U98,R91,D20,R16,D67,R40,U7,R15,U6,R7", ",")
	intersectionPoints := findIntersectionPoints(exampleWireOne, exampleWireTwo)
	minimumDistance := Min(findStepDistances(exampleWireOne, exampleWireTwo, intersectionPoints))
	if minimumDistance != 410 {
		t.Errorf("Expected %d to be 410.", minimumDistance)
	}
}
