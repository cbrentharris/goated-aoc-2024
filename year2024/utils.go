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
