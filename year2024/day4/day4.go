package day4

import "strings"

type Point struct {
    Row    int
    Column int
}

var (
    xRune = 'X'
    mRune = 'M'
    aRune = 'A'
    sRune = 'S'
)

func WordSearch(input string) int {
    lines := strings.Split(input, "\n")
    total := 0
    for lineIndex, line := range lines {
        for wordIndex, letter := range line {
            if !(letter == xRune) {
                continue
            }
            var indicesToCheck [][]Point

            canGoUp := lineIndex >= 3
            canGoDown := len(lines)-lineIndex-1 >= 3
            canGoLeft := wordIndex >= 3
            canGoRight := len(line)-wordIndex-1 >= 3

            if canGoUp {
                indicesToCheck = append(indicesToCheck, []Point{
                    {lineIndex - 1, wordIndex},
                    {lineIndex - 2, wordIndex},
                    {lineIndex - 3, wordIndex},
                })
            }
            if canGoLeft {
                indicesToCheck = append(indicesToCheck, []Point{
                    {lineIndex, wordIndex - 1},
                    {lineIndex, wordIndex - 2},
                    {lineIndex, wordIndex - 3},
                })
            }
            if canGoDown {
                indicesToCheck = append(indicesToCheck, []Point{
                    {lineIndex + 1, wordIndex},
                    {lineIndex + 2, wordIndex},
                    {lineIndex + 3, wordIndex},
                })
            }
            if canGoRight {
                indicesToCheck = append(indicesToCheck, []Point{
                    {lineIndex, wordIndex + 1},
                    {lineIndex, wordIndex + 2},
                    {lineIndex, wordIndex + 3},
                })
            }

            if canGoUp && canGoLeft {
                indicesToCheck = append(indicesToCheck, []Point{
                    {lineIndex - 1, wordIndex - 1},
                    {lineIndex - 2, wordIndex - 2},
                    {lineIndex - 3, wordIndex - 3},
                })
            }
            if canGoUp && canGoRight {
                indicesToCheck = append(indicesToCheck, []Point{
                    {lineIndex - 1, wordIndex + 1},
                    {lineIndex - 2, wordIndex + 2},
                    {lineIndex - 3, wordIndex + 3},
                })
            }
            if canGoDown && canGoLeft {
                indicesToCheck = append(indicesToCheck, []Point{
                    {lineIndex + 1, wordIndex - 1},
                    {lineIndex + 2, wordIndex - 2},
                    {lineIndex + 3, wordIndex - 3},
                })
            }
            if canGoDown && canGoRight {
                indicesToCheck = append(indicesToCheck, []Point{
                    {lineIndex + 1, wordIndex + 1},
                    {lineIndex + 2, wordIndex + 2},
                    {lineIndex + 3, wordIndex + 3},
                })
            }

            for _, points := range indicesToCheck {
                if lines[points[0].Row][points[0].Column] == 'M' &&
                    lines[points[1].Row][points[1].Column] == 'A' &&
                    lines[points[2].Row][points[2].Column] == 'S' {
                    total += 1
                }
            }
        }
    }
    return total
}
func WordSearchV2(input string) int {
    lines := strings.Split(input, "\n")
    total := 0
    for lineIndex, line := range lines {
        for wordIndex, letter := range line {
            if !(letter == aRune) {
                continue
            }

            canGoUp := lineIndex >= 1
            canGoDown := len(lines)-lineIndex-1 >= 1
            canGoLeft := wordIndex >= 1
            canGoRight := len(line)-wordIndex-1 >= 1

            if !(canGoUp && canGoDown && canGoLeft && canGoRight) {
                continue
            }

            firstRune := int32(lines[lineIndex-1][wordIndex-1])
            secondRune := int32(lines[lineIndex+1][wordIndex+1])
            thirdRune := int32(lines[lineIndex-1][wordIndex+1])
            fourthRune := int32(lines[lineIndex+1][wordIndex-1])

            firstPairMatches := (firstRune == sRune && secondRune == mRune) || (firstRune == mRune && secondRune == sRune)
            secondPairMatches := (thirdRune == sRune && fourthRune == mRune) || (thirdRune == mRune && fourthRune == sRune)

            allMatch := firstPairMatches && secondPairMatches

            if allMatch {
                total += 1
            }
        }
    }
    return total
}
