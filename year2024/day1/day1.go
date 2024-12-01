package day1

import (
	"fmt"
	"math"
	"os"
	"sort"
	"strconv"
	"strings"
)

func DayOne(input string) int {
	firstListOfIds, secondListOfIds, err := parseInput(input)
	if err != nil {
		panic("Unable to parse input correctly")
	}
	sort.Ints(firstListOfIds)
	sort.Ints(secondListOfIds)
	var total int
	for i := 0; i < len(firstListOfIds); i++ {
		firstId := firstListOfIds[i]
		secondId := secondListOfIds[i]
		delta := int(math.Abs(float64(firstId) - float64(secondId)))
		total += delta
	}
	return total
}

func ReadInput() string {
	file, err := os.ReadFile("input.txt")
	if err != nil {
		panic("Error trying to read input.txt")
	}
	contents := string(file)
	return contents
}

func DayOnePartTwo(input string) int {
	firstListOfIds, secondListOfIds, err := parseInput(input)
	if err != nil {
		panic("Unable to parse input correctly")
	}
	secondFrequencyList := make(map[int]int)

	for _, k := range secondListOfIds {
		secondFrequencyList[k]++
	}

	var score int
	for _, k := range firstListOfIds {
		freq, exists := secondFrequencyList[k]
		if exists {
			score += k * freq
		}
	}
	return score
}

func parseInput(input string) ([]int, []int, error) {
	lines := strings.Split(strings.TrimSpace(input), "\n")
	var firstIdList []int
	var secondIdList []int
	for _, line := range lines {
		words := strings.Fields(line)
		if len(words) != 2 {
			return nil, nil, fmt.Errorf("improper input line -- expected two tokens and found %d", len(words))
		}
		firstId, firstIdParseError := strconv.Atoi(words[0])
		if firstIdParseError != nil {
			return nil, nil, firstIdParseError
		}
		firstIdList = append(firstIdList, firstId)
		secondId, secondIdParseError := strconv.Atoi(words[1])
		if secondIdParseError != nil {
			return nil, nil, secondIdParseError
		}
		secondIdList = append(secondIdList, secondId)
	}
	return firstIdList, secondIdList, nil
}
