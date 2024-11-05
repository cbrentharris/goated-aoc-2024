package day2

import "testing"

func TestPartOne(t *testing.T) {
	input, err := PartOne()
	if err != nil {
		t.Error(err)
		return
	}
	// Assert input is expected value
	expected := 11590668
	if input != expected {
		t.Errorf("Expected %d, got %d", expected, input)
	}
}

func TestPartTwo(t *testing.T) {
	input, err := PartTwo()
	if err != nil {
		t.Error(err)
		return
	}
	expected := 2254
	if input != expected {
		t.Errorf("Expected %d, got %d", expected, input)
	}
}
