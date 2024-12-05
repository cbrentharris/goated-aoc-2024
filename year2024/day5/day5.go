package day5

import (
    "regexp"
    "strconv"
    "strings"
)

var (
    ruleRegex   = regexp.MustCompile("\\d+\\|\\d+")
    updateRegex = regexp.MustCompile("(\\d+,)*(\\d+)")
)

func MiddlePageCount(input string, countIncorrect bool) int {
    lines := strings.Split(input, "\n")
    pageDependencies := make(map[string]map[string]struct{})
    pageRequirements := make(map[string]map[string]struct{})
    total := 0

    for _, line := range lines {
        if ruleRegex.MatchString(line) {
            pageNumbers := strings.Split(line, "|")
            pageBefore := pageNumbers[0]
            pageAfter := pageNumbers[1]
            pagesThatHaveToComeAfter, exists := pageDependencies[pageBefore]
            if exists {
                pagesThatHaveToComeAfter[pageAfter] = struct{}{}
            } else {
                pagesThatHaveToComeAfter = make(map[string]struct{})
                pagesThatHaveToComeAfter[pageAfter] = struct{}{}
                pageDependencies[pageBefore] = pagesThatHaveToComeAfter
            }

            pagesThatHaveToComeBefore, exists := pageRequirements[pageAfter]
            if exists {
                pagesThatHaveToComeBefore[pageBefore] = struct{}{}
            } else {
                pagesThatHaveToComeBefore = make(map[string]struct{})
                pagesThatHaveToComeBefore[pageBefore] = struct{}{}
                pageRequirements[pageAfter] = pagesThatHaveToComeBefore
            }

        } else if updateRegex.MatchString(line) {
            pages := strings.Split(line, ",")
            seenPages := make(map[string]struct{})
            validUpdate := true
            for _, page := range pages {
                seenPages[page] = struct{}{}
                dependencies, hasPageDependencies := pageDependencies[page]
                if hasPageDependencies {
                    for dependency := range dependencies {
                        _, seenDependency := seenPages[dependency]
                        if seenDependency {
                            validUpdate = false
                        }
                    }
                }
            }

            if countIncorrect {
                if !validUpdate {
                    sorted := correctUpdate(pages, pageRequirements, seenPages)
                    middlePageNumber, _ := strconv.Atoi(sorted[(len(sorted)-1)/2])
                    total += middlePageNumber
                }
            } else {
                middlePageNumber, _ := strconv.Atoi(pages[(len(pages)-1)/2])
                if validUpdate {
                    total += middlePageNumber
                }
            }

        }
    }
    return total
}

func correctUpdate(pages []string, pageRequirements map[string]map[string]struct{}, pagesSeen map[string]struct{}) []string {
    correctedUpdate := make([]string, len(pages))

    for page := range pagesSeen {
        requirements, hasRequirements := pageRequirements[page]
        if !hasRequirements {
            index := 0
            for correctedUpdate[index] != "" {
                index++
            }
            correctedUpdate[index] = page
        } else {
            seenRequirements := 0
            for requirement := range requirements {
                _, requirementHasBeenSeen := pagesSeen[requirement]
                if requirementHasBeenSeen {
                    seenRequirements++
                }
            }
            for correctedUpdate[seenRequirements] != "" {
                seenRequirements++
            }
            correctedUpdate[seenRequirements] = page
        }
    }

    return correctedUpdate
}
