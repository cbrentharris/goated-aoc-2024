package day3

import (
    "regexp"
    "strconv"
)

var (
    // Every valid mul instruction must be mul(num,num)
    mulOperatorRegex = regexp.MustCompile("mul\\(\\d+,\\d+\\)")
    numberRegex      = regexp.MustCompile("\\d+")
    crazyRegex       = regexp.MustCompile("(?<=^|[^a-zA-Z0-9])(?!.*\\bdon't\\(\\))(?!.*\\bdo\\(\\)[^a-zA-Z0-9]*mul\\().*\\bmul\\([^\\)]*\\)")
)

func CorruptedProgramExecutor(program string) int {
    matches := mulOperatorRegex.FindAllString(program, -1)
    total := 0
    for _, match := range matches {
        lhs, rhs := parseLhsAndRhs(match)
        total += lhs * rhs
    }
    return total
}
func CorruptedProgramExecutorV2(program string) int {
    matches := crazyRegex.FindAllString(program, -1)
    total := 0
    for _, match := range matches {
        lhs, rhs := parseLhsAndRhs(match)
        total += lhs * rhs
    }
    return total
}

func parseLhsAndRhs(mul string) (int, int) {
    stringNums := numberRegex.FindAllString(mul, -1)
    lhs, lhserr := strconv.Atoi(stringNums[0])
    rhs, rhserr := strconv.Atoi(stringNums[1])
    if lhserr != nil {
        panic(lhserr)
    }

    if rhserr != nil {
        panic(rhserr)
    }
    return lhs, rhs
}
