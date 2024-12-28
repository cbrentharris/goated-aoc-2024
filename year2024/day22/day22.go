package day22

import (
	"goated-aoc-2024/year2024"
	"strconv"
	"strings"
	"sync"
)

func SumOfNthSecrets(input string, n int) int {
	lines := strings.Split(input, "\n")
	total := 0
	for _, line := range lines {
		num, _ := strconv.Atoi(line)
		for i := 0; i < n; i++ {
			num = CalculateNextSecret(num)
		}
		total += num
	}
	return total
}

func MaxBananas(input string, n int) int {
	lines := strings.Split(input, "\n")
	sequenceValues := make(map[string]int)
	maxSequence := 0
	var mu sync.Mutex
	var wg sync.WaitGroup
	tasks := make(chan int, len(lines)*2)
	worker := func() {
		defer wg.Done()
		for num := range tasks {
			queue := year2024.NewCircularQueue[int](4)
			seen := year2024.NewHashSet[string]()
			secret := num
			for i := 0; i < n; i++ {
				previousDigit := secret % 10
				secret = CalculateNextSecret(secret)
				digit := year2024.Mod(secret, 10)
				queue.Add(digit - previousDigit)
				if i >= 3 {
					sequenceKey := queue.ToString()
					if seen.Contains(sequenceKey) {
						continue
					}
					seen.Add(sequenceKey)
					mu.Lock()
					sequenceValues[sequenceKey] = sequenceValues[sequenceKey] + digit
					if sequenceValues[sequenceKey] > maxSequence {
						maxSequence = sequenceValues[sequenceKey]
					}
					mu.Unlock()
				}
			}

		}
	}
	wg.Add(8)
	for i := 0; i < 8; i++ {
		go worker()
	}
	for _, line := range lines {
		num, _ := strconv.Atoi(line)
		tasks <- num
	}
	close(tasks)
	wg.Wait()
	return maxSequence
}

func CalculateNextSecret(initialSecret int) int {
	multiplied := initialSecret << 6
	mixed := mix(initialSecret, multiplied)
	stage1 := prune(mixed)
	divided := stage1 >> 5
	mixed2 := mix(stage1, divided)
	stage2 := prune(mixed2)
	multiplied2 := stage2 << 11
	mixed3 := mix(stage2, multiplied2)
	stage3 := prune(mixed3)
	return stage3
}

func prune(a int) int {
	return a % 16777216
}

func mix(a, b int) int {
	return a ^ b
}
