package day13

import (
	"goated-aoc-2024/year2024"
	"math"
	"regexp"
	"strconv"
	"strings"
)

type Button struct {
	XOffset int
	YOffset int
}

type ClawMachine struct {
	ButtonA         Button
	ButtonB         Button
	PrizeCoordinate year2024.Coordinate
}

var (
	numberRegex      = regexp.MustCompile("\\d+")
	buttonATokenCost = 3
	buttonBTokenCost = 1
	calibration      = 10000000000000
)

func TokensRequiredToWinAllPrizes(input string, calibrate bool) int {
	lines := strings.Split(input, "\n")

	total := 0
	for i := 0; i < len(lines); i += 4 {
		firstButtonLine := lines[i]
		secondButtonLine := lines[i+1]
		prizeLine := lines[i+2]
		firstButtonX, firstButtonY := findTwoNumbers(firstButtonLine)
		secondButtonX, secondButtonY := findTwoNumbers(secondButtonLine)
		prizeX, prizeY := findTwoNumbers(prizeLine)
		if calibrate {
			prizeX += calibration
			prizeY += calibration
		}
		clawMachine := ClawMachine{ButtonA: Button{XOffset: firstButtonX, YOffset: firstButtonY}, ButtonB: Button{XOffset: secondButtonX, YOffset: secondButtonY}, PrizeCoordinate: year2024.Coordinate{X: prizeX, Y: prizeY}}
		var cost int
		var reached bool
		if calibrate {
			// Hit stacktrace length issues, need to solve with math
			cost, reached = solveSystemOfEquations(clawMachine)
		} else {
			// Can use the DP solution without calibration
			cost, reached = calculateMinimalTokensRequired(clawMachine)
		}
		if reached {
			total += cost
		}

	}
	return total
}

func findTwoNumbers(line string) (int, int) {
	numbersAsString := numberRegex.FindAllString(line, 2)
	first, _ := strconv.Atoi(numbersAsString[0])
	second, _ := strconv.Atoi(numbersAsString[1])
	return first, second
}

func calculateMinimalTokensRequired(clawMachine ClawMachine) (int, bool) {
	memo := make(map[MemoKey]MemoValue)
	return calculate(clawMachine, year2024.Coordinate{X: 0, Y: 0}, 0, &memo)
}

type MemoKey struct {
	Coordinate year2024.Coordinate
	TokensUsed int
}

type MemoValue struct {
	TokensUsed int
	Reached    bool
}

func calculate(clawMachine ClawMachine, current year2024.Coordinate, currentTokensUsed int, memo *map[MemoKey]MemoValue) (int, bool) {
	memoKey := MemoKey{TokensUsed: currentTokensUsed, Coordinate: current}
	memoValue, exists := (*memo)[memoKey]

	if exists {
		return memoValue.TokensUsed, memoValue.Reached
	}

	if current.X > clawMachine.PrizeCoordinate.X || current.Y > clawMachine.PrizeCoordinate.Y {
		memoValue = MemoValue{Reached: false, TokensUsed: 0}
		(*memo)[memoKey] = memoValue
		return 0, false
	}

	if current == clawMachine.PrizeCoordinate {
		memoValue = MemoValue{Reached: true, TokensUsed: currentTokensUsed}
		(*memo)[memoKey] = memoValue
		return currentTokensUsed, true

	}

	buttonACost, buttonAReached := calculate(clawMachine, year2024.Coordinate{X: current.X + clawMachine.ButtonA.XOffset, Y: current.Y + clawMachine.ButtonA.YOffset}, currentTokensUsed+buttonATokenCost, memo)
	buttonBCost, buttonBReached := calculate(clawMachine, year2024.Coordinate{X: current.X + clawMachine.ButtonB.XOffset, Y: current.Y + clawMachine.ButtonB.YOffset}, currentTokensUsed+buttonBTokenCost, memo)

	if !buttonAReached && !buttonBReached {
		memoValue = MemoValue{Reached: false, TokensUsed: 0}
		(*memo)[memoKey] = memoValue
		return 0, false
	}

	if !buttonAReached {
		memoValue = MemoValue{Reached: true, TokensUsed: buttonBCost}
		(*memo)[memoKey] = memoValue
		return buttonBCost, true
	}

	if !buttonBReached {
		memoValue = MemoValue{Reached: true, TokensUsed: buttonACost}
		(*memo)[memoKey] = memoValue
		return buttonACost, true
	}

	minimumCost := min(buttonACost, buttonBCost)
	memoValue = MemoValue{Reached: true, TokensUsed: minimumCost}
	(*memo)[memoKey] = memoValue
	return minimumCost, true
}

func solveSystemOfEquations(clawMachine ClawMachine) (int, bool) {
	// 94 22 -- Xs (a first)
	// 34 67
	// Use Cramer's rule https://en.wikipedia.org/wiki/Cramer%27s_rule
	// since unknowns = equations
	determinant := calculateDeterminant(clawMachine.ButtonA.XOffset, clawMachine.ButtonB.XOffset, clawMachine.ButtonA.YOffset, clawMachine.ButtonB.YOffset)
	// Infinite solutions
	if determinant == 0 {
		return 0, false
	}
	buttonAMoves := float64(calculateDeterminant(clawMachine.PrizeCoordinate.X, clawMachine.ButtonB.XOffset, clawMachine.PrizeCoordinate.Y, clawMachine.ButtonB.YOffset)) / float64(determinant)
	buttonBMoves := float64(calculateDeterminant(clawMachine.ButtonA.XOffset, clawMachine.PrizeCoordinate.X, clawMachine.ButtonA.YOffset, clawMachine.PrizeCoordinate.Y)) / float64(determinant)

	// Must have whole button presses
	if isInteger(buttonAMoves) && isInteger(buttonBMoves) {
		return int(buttonAMoves)*buttonATokenCost + int(buttonBMoves)*buttonBTokenCost, true
	}

	return 0, false
}

func calculateDeterminant(a, b, c, d int) int {
	return a*d - b*c
}

func isInteger(float float64) bool {
	return math.Mod(float, 1) == 0
}
