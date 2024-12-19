package day19

import (
	"goated-aoc-2024/year2024"
	"strings"
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
	memo := make(map[string]int)
	total := 0
	for _, pattern := range patterns {
		total += countPossiblePatterns([]rune(pattern), trie, &memo)
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
