package day2

import (
	"os"
	"strconv"
	"strings"
)

const add int = 1
const mul int = 2
const halt int = 99

func PartOne() (int, error) {
	lines, _ := ReadInput()
	program := InitializeProgram(lines, 12, 2)
	for i := 0; i < len(program); i += 4 {
		opcode := program[i]
		if opcode == halt {
			return program[0], nil
		}
		value1Position := program[i+1]
		value2Position := program[i+2]
		value1 := program[value1Position]
		value2 := program[value2Position]
		storePosition := program[i+3]

		if opcode == add {
			program[storePosition] = value1 + value2
		} else if opcode == mul {
			program[storePosition] = value1 * value2
		} else {
			return 0, nil
		}
	}
	return 0, nil
}

func PartTwo() (int, error) {
	lines, _ := ReadInput()
	target := 19690720
	for noun := 0; noun < 100; noun++ {
		for verb := 0; verb < 100; verb++ {
			program := InitializeProgram(lines, noun, verb)

			for i := 0; i < len(program); i += 4 {
				opcode := program[i]
				if opcode == halt {
					if program[0] == target {
						return 100*noun + verb, nil
					} else {
						break
					}
				}
				value1Position := program[i+1]
				value2Position := program[i+2]
				value1 := program[value1Position]
				value2 := program[value2Position]
				storePosition := program[i+3]

				if opcode == add {
					program[storePosition] = value1 + value2
				} else if opcode == mul {
					program[storePosition] = value1 * value2
				} else {
					return 0, nil
				}
			}

		}
	}
	return 0, nil
}

func ReadInput() ([]string, error) {
	file, err := os.ReadFile("input.txt")
	if err != nil {
		return nil, err
	}
	contents := string(file)
	lines := strings.Split(contents, "\n")
	return strings.Split(lines[0], ","), nil
}

func InitializeProgram(input []string, noun, verb int) []int {
	program := make([]int, len(input))
	for i, line := range input {
		code, _ := strconv.Atoi(line)
		if i == 1 {
			code = noun
		} else if i == 2 {
			code = verb
		}
		program[i] = code
	}
	return program
}
