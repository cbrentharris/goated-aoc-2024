package day14

import (
	"context"
	"fmt"
	"goated-aoc-2024/year2024"
	"regexp"
	"strings"
	"sync"
)

var numberRegex = regexp.MustCompile("-?\\d+")

type Velocity struct {
	X, Y int
}

func CalculateSafetyScore(input string, width, height int, seconds int) int {
	quadrantScores := make(map[int]int)
	for _, line := range strings.Split(input, "\n") {
		nums := year2024.ToIntSlice(numberRegex.FindAllString(line, 4))
		positionX := nums[0]
		positionY := nums[1]
		velocityX := nums[2]
		velocityY := nums[3]
		finalX := mod(positionX+velocityX*seconds, width)
		finalY := mod(positionY+velocityY*seconds, height)
		quadrant, err := determineQuadrant(year2024.Coordinate{X: finalX, Y: finalY}, width, height)
		if err == nil {
			quadrantScores[quadrant]++
		}
	}
	total := 1
	for _, v := range quadrantScores {
		total *= v
	}
	return total
}

type Robot struct {
	Coordinate year2024.Coordinate
	V          Velocity
}

func FindChristmasTree(input string, width, height int) int {
	initialRobotCoordinates := year2024.NewHashSet[Robot]()
	for _, line := range strings.Split(input, "\n") {
		nums := year2024.ToIntSlice(numberRegex.FindAllString(line, 4))
		positionX := nums[0]
		positionY := nums[1]
		velocityX := nums[2]
		velocityY := nums[3]
		initialRobotCoordinates.Add(Robot{Coordinate: year2024.Coordinate{X: positionX, Y: positionY}, V: Velocity{X: velocityX, Y: velocityY}})
	}
	numWorkers := 8
	upperLimit := 10000000
	result := make(chan int, 1)

	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	var wg sync.WaitGroup

	worker := func(ctx context.Context, seconds <-chan int) {
		defer wg.Done()
		for {
			select {
			case <-ctx.Done():
				return
			case second, ok := <-seconds:
				if !ok {
					return
				}

				robotCoordinates := year2024.NewHashSet[year2024.Coordinate]()
				for robot := range initialRobotCoordinates.Iterator() {
					finalX := mod(robot.Coordinate.X+robot.V.X*second, width)
					finalY := mod(robot.Coordinate.Y+robot.V.Y*second, height)
					robotCoordinates.Add(year2024.Coordinate{X: finalX, Y: finalY})
				}
				// Improve with goroutines
				for c := range robotCoordinates.Iterator() {
					// Find the first triangle like
					//        *
					//       ***
					//      *****
					containsNextThree := robotCoordinates.Contains(year2024.Coordinate{X: c.X - 1, Y: c.Y + 1}) && robotCoordinates.Contains(year2024.Coordinate{X: c.X + 1, Y: c.Y + 1}) && robotCoordinates.Contains(year2024.Coordinate{X: c.X, Y: c.Y + 1})
					containsNextFive := robotCoordinates.Contains(year2024.Coordinate{X: c.X - 2, Y: c.Y + 2}) && robotCoordinates.Contains(year2024.Coordinate{X: c.X + 2, Y: c.Y + 2}) && robotCoordinates.Contains(year2024.Coordinate{X: c.X - 1, Y: c.Y + 2}) && robotCoordinates.Contains(year2024.Coordinate{X: c.X + 1, Y: c.Y + 2}) && robotCoordinates.Contains(year2024.Coordinate{X: c.X, Y: c.Y + 2})
					if containsNextThree && containsNextFive {
						for y := 0; y < height; y++ {
							for x := 0; x < width; x++ {
								if robotCoordinates.Contains(year2024.Coordinate{X: x, Y: y}) {
									fmt.Print("*")
								} else {
									fmt.Print(".")
								}
							}
							fmt.Println("")
						}
						select {
						case result <- second:
							cancel()
						default:
						}
						return
					}
				}
			}
		}
	}
	seconds := make(chan int, numWorkers*10)
	for i := 0; i < numWorkers; i++ {
		wg.Add(1)
		go worker(ctx, seconds)
	}

	go func() {
		for s := 0; s < upperLimit; s++ {
			select {
			case <-ctx.Done():
				break
			case seconds <- s:
			}

		}
		close(seconds)
	}()

	go func() {
		wg.Wait()
		close(result)
	}()

	for r := range result {
		return r
	}
	return 0
}

func determineQuadrant(coordinate year2024.Coordinate, width, height int) (int, error) {
	middleX := width / 2
	middleY := height / 2
	onMiddleX := coordinate.X == middleX
	onMiddleY := coordinate.Y == middleY

	if onMiddleX || onMiddleY {
		return 0, fmt.Errorf("on middle line")
	}

	switch {
	case coordinate.X < middleX && coordinate.Y < middleY:
		return 1, nil
	case coordinate.X > middleX && coordinate.Y < middleY:
		return 2, nil
	case coordinate.X < middleX && coordinate.Y > middleY:
		return 3, nil
	}
	return 4, nil
}

func mod(a, b int) int {
	return ((a % b) + b) % b
}
