package day3

import (
    "regexp"
    "strconv"
)

var (
    // Every valid mul instruction must be mul(num,num)
    mulOperatorRegex = regexp.MustCompile("mul\\(\\d+,\\d+\\)")
    v2Regex          = regexp.MustCompile("(mul\\(\\d+,\\d+\\)|do\\(\\)|don't\\(\\))")
    numberRegex      = regexp.MustCompile("\\d+")
    doRegex          = regexp.MustCompile("do\\(\\)")
    dontRegex        = regexp.MustCompile("don't\\(\\)")
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
    matches := v2Regex.FindAllString(program, -1)
    total := 0
    multiplicationEnabled := true
    for _, match := range matches {
        switch {
        case doRegex.MatchString(match):
            multiplicationEnabled = true
        case dontRegex.MatchString(match):
            multiplicationEnabled = false
        case multiplicationEnabled && mulOperatorRegex.MatchString(match):
            lhs, rhs := parseLhsAndRhs(match)
            total += lhs * rhs
        }
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
