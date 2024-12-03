package day2

import (
	"testing"
	"year2024"
)

func TestDayTwoPartOneExample(t *testing.T) {
	exampleInput := `7 6 4 2 1
1 2 7 8 9
9 7 6 2 1
1 3 2 4 5
8 6 4 4 1
1 3 6 7 9`
	totalSafeReports := ReportAnalyzer(exampleInput, false)
	if totalSafeReports != 2 {
		t.Errorf("Expected two safe reports, found %d", totalSafeReports)
	}
}

func TestDayTwoPartOne(t *testing.T) {
	totalSafeReports := ReportAnalyzer(year2024.ReadInput("input.txt"), false)

	if totalSafeReports != 359 {
		t.Errorf("Total safe reports: %d", totalSafeReports)
	}
}

func TestDayTwoPartTwoExample(t *testing.T) {
	exampleInput := `7 6 4 2 1
1 2 7 8 9
9 7 6 2 1
1 3 2 4 5
8 6 4 4 1
1 3 6 7 9`
	totalSafeReports := ReportAnalyzer(exampleInput, true)
	if totalSafeReports != 4 {
		t.Errorf("Expected four safe reports, found %d", totalSafeReports)
	}
}

func TestDayTwoPartTwo(t *testing.T) {
	totalSafeReports := ReportAnalyzer(year2024.ReadInput("input.txt"), true)
	if totalSafeReports != 418 {
		t.Errorf("Total safe reports: %d", totalSafeReports)
	}
}
