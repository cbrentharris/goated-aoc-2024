package day1

import (
	"os"
	"strconv"
	"strings"
)

func PartOne() (int, error) {
	totalFuelRequired := 0
	lines, _ := ReadInput()
	for _, line := range lines {
		mass, _ := strconv.ParseInt(line, 10, 64)
		fuelRequired := mass/3 - 2
		totalFuelRequired += int(fuelRequired)
	}
	return totalFuelRequired, nil
}

func PartTwo() (int, error) {
	totalFuelRequired := 0
	lines, _ := ReadInput()
	for _, line := range lines {
		mass, _ := strconv.ParseInt(line, 10, 64)
		fuelRequired := FuelRequired(mass)
		totalFuelRequired += int(fuelRequired)
	}
	return totalFuelRequired, nil
}

func ReadInput() ([]string, error) {
	file, err := os.ReadFile("input.txt")
	if err != nil {
		return nil, err
	}
	contents := string(file)
	lines := strings.Split(contents, "\n")
	return lines, nil
}

func FuelRequired(mass int64) int64 {
	fuelRequired := mass/3 - 2
	if fuelRequired <= 0 {
		return 0
	}
	return fuelRequired + FuelRequired(fuelRequired)
}
