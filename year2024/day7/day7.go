package day7

import (
	"goated-aoc-2024/year2024"
	"strconv"
	"strings"
)

func GuessCalibrationEquations(input string, useConcat bool) int64 {
	equationLines := strings.Split(input, "\n")
	total := int64(0)
	for _, equationline := range equationLines {
		splitOnColon := strings.Split(equationline, ":")
		result, _ := strconv.Atoi(splitOnColon[0])
		parameters := year2024.ToIntSlice(strings.Fields(splitOnColon[1]))

		if canReach(result, 1, parameters[0], parameters, useConcat) {
			total += int64(result)
		}
	}
	return total
}

func canReach(target int, currentIndex int, intermediateResult int, values []int, useConcatOperator bool) bool {
	if intermediateResult > target {
		return false
	}

	if currentIndex == len(values) {
		if target == intermediateResult {
			return true
		} else {
			return false
		}
	}
	value := values[currentIndex]
	currentIndex++
	added := intermediateResult + value
	multiplied := intermediateResult * value
	canReachAddition := canReach(target, currentIndex, added, values, useConcatOperator)
	if canReachAddition {
		return true
	}
	canReachMultiplication := canReach(target, currentIndex, multiplied, values, useConcatOperator)
	if canReachMultiplication {
		return true
	}

	if useConcatOperator {
		concatted, _ := strconv.Atoi(strconv.Itoa(intermediateResult) + strconv.Itoa(value))
		return canReach(target, currentIndex, concatted, values, useConcatOperator)

	}
	return false
}
