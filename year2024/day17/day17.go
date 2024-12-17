package day17

import (
	"goated-aoc-2024/year2024"
	"math"
	"regexp"
	"strconv"
	"strings"
)

var (
	aRegister = 'A'
	bRegister = 'B'
	cRegister = 'C'
)

type OpCode int

const (
	DivisionA = iota
	BitwiseXOrB
	Modulo
	Jump
	BitwiseXOrBAndC
	Output
	DivisionB
	DivisionC
)

func ExecuteProgram(input string) string {
	registers, instructionsAndOperands := initializeRegistersAndProgram(input)
	return executeProgram(instructionsAndOperands, registers)
}

func FindCorrectRegisterValue(input string) int {
	registers, instructionsAndOperands := initializeRegistersAndProgram(input)
	totalPositionOutputs := len(instructionsAndOperands)
	total := 0
	programString := year2024.Join(instructionsAndOperands, ",")
	for i := 0; i < totalPositionOutputs; i++ {
		endVal := int(math.Pow(8, float64(i+1)))
		for j := 0; j <= endVal; j++ {
			iter := total + j
			registers[aRegister] = iter
			result := executeProgram(instructionsAndOperands, registers)
			neededStr := year2024.Join(instructionsAndOperands[len(instructionsAndOperands)-1-i:], ",")
			if result == programString {
				return iter
			}
			if neededStr == result {
				total += j
				total *= 8
				break
			}
		}
	}
	return total
}

func executeProgram(instructionsAndOperands []int, registers map[rune]int) string {
	var output []int
	instructionPointer := 0
	jumped := false
	pointer := 0
	for instructionPointer < len(instructionsAndOperands)-1 {
		pointer, jumped, output = execute(OpCode(instructionsAndOperands[instructionPointer]), instructionsAndOperands[instructionPointer+1], registers, output)
		if jumped {
			instructionPointer = pointer
		} else {
			instructionPointer += 2
		}
	}
	result := make([]string, len(output))
	for i, value := range output {
		result[i] = strconv.Itoa(value)
	}
	return strings.Join(result, ",")
}

func comboOperand(value int, registers map[rune]int) int {
	switch value {
	case 0:
		return value
	case 1:
		return value
	case 2:
		return value
	case 3:
		return value
	case 4:
		return registers[aRegister]
	case 5:
		return registers[bRegister]
	case 6:
		return registers[cRegister]
	case 7:
		panic("reserved combo operand")
	}
	panic("unknown combo operand")
}

func execute(opCode OpCode, operand int, registers map[rune]int, output []int) (int, bool, []int) {
	switch opCode {
	case DivisionA:
		numerator := registers[aRegister]
		denominator := math.Pow(2, float64(comboOperand(operand, registers)))
		division := numerator / int(denominator)
		registers[aRegister] = division
	case BitwiseXOrB:
		bRegisterValue := registers[bRegister]
		result := bRegisterValue ^ operand
		registers[bRegister] = result
	case Modulo:
		registers[bRegister] = year2024.Mod(comboOperand(operand, registers), 8)
	case Jump:
		aRegisterValue := registers[aRegister]
		if aRegisterValue > 0 {
			return operand, true, output
		}
	case BitwiseXOrBAndC:
		bRegisterValue := registers[bRegister]
		cRegistervalue := registers[cRegister]
		registers[bRegister] = bRegisterValue ^ cRegistervalue
	case Output:
		output = append(output, year2024.Mod(comboOperand(operand, registers), 8))
	case DivisionB:
		numerator := registers[aRegister]
		denominator := math.Pow(2, float64(comboOperand(operand, registers)))
		division := numerator / int(denominator)
		registers[bRegister] = division
	case DivisionC:
		numerator := registers[aRegister]
		denominator := math.Pow(2, float64(comboOperand(operand, registers)))
		division := numerator / int(denominator)
		registers[cRegister] = division
	}
	return 0, false, output
}

func initializeRegistersAndProgram(input string) (map[rune]int, []int) {
	registerRegex := regexp.MustCompile("Register (\\w): (\\d+)")
	registers := make(map[rune]int)
	//instructionPointer := 0
	lines := strings.Split(input, "\n")
	for i := 0; i < 3; i++ {
		elements := registerRegex.FindStringSubmatch(lines[i])
		r := rune(elements[1][0])
		num, _ := strconv.Atoi(elements[2])
		registers[r] = num
	}
	programRegex := regexp.MustCompile("Program: (\\d(,\\d)*)")
	programString := programRegex.FindStringSubmatch(lines[4])[1]
	instructionsAndOperands := year2024.ToIntSlice(strings.Split(programString, ","))
	return registers, instructionsAndOperands
}
