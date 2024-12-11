package day11

import (
	"goated-aoc-2024/year2024"
	"math"
	"strings"
	"sync"
)

// Rule 1: 0 -> 1
// Rule 2: evenDigits -> split in half
// Rule 3: multiply by 2024

type BlinkState struct {
	StoneId    int
	BlinkCount int
}

func CountStones(input string, blinkCount int) int {
	memo := make(map[BlinkState]int)
	total := 0
	wg := sync.WaitGroup{}
	intermediateChannel := make(chan int)
	var mu sync.Mutex
	for _, stone := range year2024.ToIntSlice(strings.Split(input, " ")) {
		wg.Add(1)
		go func(stone int) {
			defer wg.Done()
			// Each goroutine computes the result for a stone
			result := countEventualStones(BlinkState{StoneId: stone, BlinkCount: 0}, &memo, &mu, blinkCount)
			intermediateChannel <- result // Send the result to the total channel
		}(stone)
	}
	go func() {
		wg.Wait()
		close(intermediateChannel)
	}()
	for intermediateTotal := range intermediateChannel {
		total += intermediateTotal
	}
	return total
}

func countEventualStones(blinkState BlinkState, memo *map[BlinkState]int, mu *sync.Mutex, maxBlinks int) int {
	mu.Lock()
	count, exists := (*memo)[blinkState]
	mu.Unlock()
	if exists {
		return count
	}
	if blinkState.BlinkCount == maxBlinks {
		return 1
	}

	if blinkState.StoneId == 0 {
		child := countEventualStones(BlinkState{StoneId: 1, BlinkCount: blinkState.BlinkCount + 1}, memo, mu, maxBlinks)
		mu.Lock()
		(*memo)[blinkState] = child
		mu.Unlock()
		return child
	} else if hasEvenDigits(blinkState.StoneId) {
		left, right := split(blinkState.StoneId)
		leftChild := countEventualStones(BlinkState{StoneId: left, BlinkCount: blinkState.BlinkCount + 1}, memo, mu, maxBlinks)
		rightChild := countEventualStones(BlinkState{StoneId: right, BlinkCount: blinkState.BlinkCount + 1}, memo, mu, maxBlinks)
		mu.Lock()
		(*memo)[blinkState] = leftChild + rightChild
		mu.Unlock()
		return leftChild + rightChild
	} else {
		child := countEventualStones(BlinkState{StoneId: blinkState.StoneId * 2024, BlinkCount: blinkState.BlinkCount + 1}, memo, mu, maxBlinks)
		mu.Lock()
		(*memo)[blinkState] = child
		mu.Unlock()
		return child
	}
}

func split(n int) (int, int) {
	numDigits := countDigits(n)
	mid := numDigits / 2
	divisor := int(math.Pow10(mid))
	left := n / divisor
	right := n % divisor
	return left, right
}

func hasEvenDigits(n int) bool {
	return countDigits(n)%2 == 0
}

func countDigits(n int) int {
	numDigits := int(math.Log10(float64(n))) + 1
	return numDigits
}
