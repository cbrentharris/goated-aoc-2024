package year2024

import "os"

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
