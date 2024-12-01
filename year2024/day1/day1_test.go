package day1

import "testing"

func TestDayOneExample(t *testing.T) {
	input := `3   4
4   3
2   5
1   3
3   9
3   3`
	output := DayOne(input)

	if output != 11 {
		t.Errorf("Expected day 1 example to be 11, instead got %d", output)
	}
}

func TestDayOne(t *testing.T) {
	output := DayOne(ReadInput())
	if output != 2264607 {
		t.Errorf("Expected day 1 run to be 2264607, instead got %d", output)
	}
}

func TestDayOnePartTwoExample(t *testing.T) {
	input := `3   4
4   3
2   5
1   3
3   9
3   3`
	output := DayOnePartTwo(input)

	if output != 31 {
		t.Errorf("Expected day one part two example to be 31, got %d", output)
	}
}

func TestDayOnePartTwo(t *testing.T) {
	output := DayOnePartTwo(ReadInput())
	if output != 19457120 {
		t.Errorf("Expected day 1 part two run to be 19457120, instead got %d", output)
	}
}
