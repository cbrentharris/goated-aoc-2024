package year2024

import (
	"os"
	"strconv"
)

func ReadInput(location string) string {
	file, err := os.ReadFile(location)
	if err != nil {
		panic(err)
	}
	return string(file)
}

type Coordinate struct {
	X int
	Y int
}

func ToIntSlice(strings []string) []int {
	ints := make([]int, len(strings))
	for i, s := range strings {
		n, err := strconv.Atoi(s)
		if err != nil {
			panic(err)
		}
		ints[i] = n
	}
	return ints
}

func ToRuneSlice(strings []string) [][]rune {
	runes := make([][]rune, len(strings))
	for i, s := range strings {
		runes[i] = []rune(s)
	}
	return runes
}

func OffTheMap(coordinate Coordinate, grid [][]rune) bool {
	return coordinate.Y >= len(grid) || coordinate.Y < 0 || coordinate.X >= len(grid[0]) || coordinate.X < 0
}

func Abs(a int, b int) int {
	if a >= b {
		return a - b
	} else {
		return b - a
	}
}
