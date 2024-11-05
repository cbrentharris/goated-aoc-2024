package day1

import (
	"testing"
)

func TestDayOne(t *testing.T) {
	input, err := PartOne()
	if err != nil {
		t.Error(err)
		return
	}
	// Assert input is expected value
	expected := 3297866
	if input != expected {
		t.Errorf("Expected %d, got %d", 0, input)
	}
}

func TestDayTwo(t *testing.T) {
	input, err := PartTwo()
	if err != nil {
		t.Error(err)
		return
	}
	unexpected := 4943804
	if input == unexpected {
		t.Errorf("Did not expect %d", unexpected)
	}
	expected := 4943923
	if input != expected {
		t.Errorf("Expected %d, got %d", expected, input)
	}
}

func TestFuelRequired(t *testing.T) {
	input := int64(100756)
	expected := int64(50346)
	if output := FuelRequired(input); output != expected {
		t.Errorf("Expected %d, got %d", expected, output)
	}
}
