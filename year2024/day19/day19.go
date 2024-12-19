package day19

import (
	"goated-aoc-2024/year2024"
	"strings"
	"sync"
)

func CountPossibleDesigns(input string) int {
	patterns, trie := parseInput(input)
	memo := make(map[string]int)
	total := 0
	for _, pattern := range patterns {
		if countPossiblePatterns([]rune(pattern), trie, &memo) > 0 {
			total += 1
		}
	}
	return total
}

func CountPossibleDistinctDesigns(input string) int {
	patterns, trie := parseInput(input)
	total := 0
	var wg sync.WaitGroup
	result := make(chan int)
	for _, pattern := range patterns {
		wg.Add(1)
		go func(p string) {
			defer wg.Done()
			memo := make(map[string]int)
			result <- countPossiblePatterns([]rune(p), trie, &memo)
		}(pattern)
	}

	go func() {
		wg.Wait()
		close(result)
	}()

	for r := range result {
		total += r
	}

	return total
}

func countPossiblePatterns(pattern []rune, towels *year2024.Trie, memo *map[string]int) int {
	if len(pattern) == 0 {
		return 1
	}
	memoKey := string(pattern)
	count, exists := (*memo)[memoKey]
	if exists {
		return count
	}
	total := 0
	for index := range pattern {
		if towels.SearchRunes(pattern[:index+1]) {
			subSearch := countPossiblePatterns(pattern[index+1:], towels, memo)
			total += subSearch
		}
	}
	(*memo)[memoKey] = total
	return total
}

func parseInput(input string) ([]string, *year2024.Trie) {
	lines := strings.Split(input, "\n")
	towels := strings.Split(lines[0], ", ")
	trie := year2024.NewTrie()
	for _, towel := range towels {
		trie.Insert(towel)
	}
	return lines[2:], trie
}
