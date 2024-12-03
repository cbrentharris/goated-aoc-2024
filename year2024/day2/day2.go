package day2

import (
	"strconv"
	"strings"
)

type Direction int

const (
	Increasing Direction = iota
	Decreasing
	Undecided
)

func ReportAnalyzer(input string, dampener bool) int {
	reports := strings.Split(input, "\n")
	var totalSafeReports int
	for _, report := range reports {
		levels := parseLevels(report)
		if isSafeReport(levels, dampener) {
			totalSafeReports++
		}
	}
	return totalSafeReports
}

func parseLevels(report string) []int {
	stringLevels := strings.Fields(report)
	levels := make([]int, len(stringLevels))
	for index, stringLevel := range stringLevels {
		level, err := strconv.Atoi(stringLevel)
		if err != nil {
			panic(err)
		}
		levels[index] = level
	}
	return levels
}

func isSafeReport(levels []int, dampener bool) bool {
	startIndex, endIndex := getDampenerRange(len(levels), dampener)
	for indexToSkip := startIndex; indexToSkip < endIndex; indexToSkip++ {
		if isSafeWithSkip(levels, indexToSkip) {
			return true
		}
	}
	return false
}

func getDampenerRange(levelSize int, dampener bool) (int, int) {
	// Without a dampener, we will not skip any index (as it is set to -1, and we only perform one loop)
	if dampener {
		return 0, levelSize
	}
	return -1, 0
}

func isSafeWithSkip(levels []int, skipIndex int) bool {
	direction := Undecided
	for i, j := 0, 1; j < len(levels); {
		if i == skipIndex {
			i++
			if i == j {
				j++
			}
			continue
		}
		if j == skipIndex {
			j++
			continue
		}

		if !isTransitionSafe(levels[i], levels[j], direction) {
			return false
		}

		direction = getDirection(levels[i], levels[j])
		i++
		j++
	}
	return true
}

func isTransitionSafe(a, b int, currentDirection Direction) bool {
	return isWithinDelta(a, b) &&
		(currentDirection == Undecided || currentDirection == getDirection(a, b))
}

func delta(a int, b int) int {
	if a >= b {
		return a - b
	} else {
		return b - a
	}
}

func isWithinDelta(a int, b int) bool {
	delta := delta(a, b)
	return delta > 0 && delta < 4
}

func getDirection(a int, b int) Direction {
	if a > b {
		return Decreasing
	} else if a < b {
		return Increasing
	} else {
		return Undecided
	}
}
