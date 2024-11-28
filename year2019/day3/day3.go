package day3

import (
	"math"
	"os"
	"strconv"
	"strings"
)

func ReadInput() ([]string, error) {
	file, err := os.ReadFile("input.txt")
	if err != nil {
		return nil, err
	}
	contents := string(file)
	lines := strings.Split(contents, "\n")
	return lines, nil
}

func DayOne() (int, error) {
	lines, _ := ReadInput()
	lineOne := strings.Split(lines[0], ",")
	lineTwo := strings.Split(lines[1], ",")
	intersectionPoints := findIntersectionPoints(lineOne, lineTwo)
	minimumDistance := Min(findManhattanDistances(Point{0, 0}, intersectionPoints))
	return minimumDistance, nil
}

func DayTwo() (int, error) {
	lines, _ := ReadInput()
	lineOne := strings.Split(lines[0], ",")
	lineTwo := strings.Split(lines[1], ",")
	intersectionPoints := findIntersectionPoints(lineOne, lineTwo)
	minimumDistance := Min(findStepDistances(lineOne, lineTwo, intersectionPoints))
	return minimumDistance, nil
}

func findStepDistances(wireOne []string, wireTwo []string, intersectionPoints []Point) []int {
	stepsOne := findSteps(wireOne, intersectionPoints)
	stepsTwo := findSteps(wireTwo, intersectionPoints)
	var combinedDistances []int
	for i := 0; i < len(stepsOne); i++ {
		combinedDistances = append(combinedDistances, stepsOne[i]+stepsTwo[i])
	}
	return combinedDistances
}

func findSteps(wire []string, points []Point) []int {
	segments := getSegments(wire)
	var steps []int
	for _, point := range points {
		steps = append(steps, findStepsToPoint(point, segments))
	}
	return steps
}

func findStepsToPoint(point Point, segments []Segment) int {
	total := 0.0
	for _, segment := range segments {
		if liesWithin(point, segment) {
			total += math.Abs(segment.start.y-point.y) + math.Abs(segment.start.x-point.x)
			return int(total)
		}
		total += math.Abs(segment.end.y-segment.start.y) + math.Abs(segment.end.x-segment.start.x)
	}
	return int(total)
}

func liesWithin(point Point, segment Segment) bool {
	if isHorizontalSlope(segment) {
		return point.y == segment.start.y && min(segment.start.x, segment.end.x) < point.x && max(segment.start.x, segment.end.x) > point.x
	} else {
		return point.x == segment.start.x && min(segment.start.y, segment.end.y) < point.y && max(segment.start.y, segment.end.y) > point.y
	}
}

func Min(values []int) int {
	if len(values) == 0 {
		panic("Cannot find minimum of empty list")
	}
	arrayMin := values[0]
	for _, value := range values {
		if value < arrayMin {
			arrayMin = value
		}
	}
	return arrayMin
}

func findManhattanDistances(origin Point, points []Point) []int {
	var distances []int
	for _, point := range points {
		distances = append(distances, manhattanDistance(origin, point))
	}
	return distances
}

func manhattanDistance(origin Point, point Point) int {
	return int(math.Abs(origin.x-point.x) + math.Abs(origin.y-point.y))
}

type Point struct {
	x float64
	y float64
}

type Segment struct {
	start Point
	end   Point
}

func findIntersectionPoints(one []string, two []string) []Point {
	segmentsOne := getSegments(one)
	segmentsTwo := getSegments(two)
	var intersectionPoints []Point
	for _, segmentOne := range segmentsOne {
		for _, segmentTwo := range segmentsTwo {
			if intersectionPoint := findIntersectionPoint(segmentOne, segmentTwo); intersectionPoint != nil {
				intersectionPoints = append(intersectionPoints, *intersectionPoint)
			}
		}
	}
	return intersectionPoints

}

func findIntersectionPoint(one Segment, two Segment) *Point {
	if !overlaps(one, two) {
		return nil
	}
	if isVerticalSlope(one) {
		return &Point{one.start.x, two.start.y}
	} else {
		return &Point{two.start.x, one.start.y}
	}
}

func overlaps(one Segment, two Segment) bool {
	if isVerticalSlope(one) && isHorizontalSlope(two) {
		return max(one.start.y, one.end.y) > two.start.y && two.start.y > min(one.start.y, one.end.y) &&
			max(two.start.x, two.end.x) > one.start.x && one.start.x > min(two.start.x, two.end.x)
	} else if isHorizontalSlope(one) && isVerticalSlope(two) {
		return max(two.start.y, two.end.y) > one.start.y && one.start.y > min(two.start.y, two.end.y) &&
			max(one.start.x, one.end.x) > two.start.x && two.start.x > min(one.start.x, one.end.x)
	} else {
		return false
	}
}

func isHorizontalSlope(segment Segment) bool {
	return segment.start.y == segment.end.y
}

func isVerticalSlope(segment Segment) bool {
	return segment.start.x == segment.end.x
}

func getSegments(one []string) []Segment {
	var initialPoint = Point{0, 0}
	segments := []Segment{}
	for _, instruction := range one {
		direction := instruction[0]
		distanceInt, _ := strconv.Atoi(instruction[1:])
		distance := float64(distanceInt)
		var endPoint Point
		switch direction {
		case 'U':
			endPoint = Point{initialPoint.x, initialPoint.y + distance}
		case 'D':
			endPoint = Point{initialPoint.x, initialPoint.y - distance}
		case 'L':
			endPoint = Point{initialPoint.x - distance, initialPoint.y}
		case 'R':
			endPoint = Point{initialPoint.x + distance, initialPoint.y}
		}
		segments = append(segments, Segment{initialPoint, endPoint})
		initialPoint = endPoint
	}
	return segments
}
