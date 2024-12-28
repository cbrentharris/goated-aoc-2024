package day23

import (
	"goated-aoc-2024/year2024"
	"sort"
	"strings"
)

type PotentialHistorian struct {
	Last    string
	Visited *year2024.HashSet[string]
}

func NumTConnected(input string) int {
	connections := make(map[string]*year2024.HashSet[string])
	lines := strings.Split(input, "\n")
	ts := year2024.NewHashSet[string]()
	for _, connection := range lines {
		split := strings.Split(connection, "-")
		left := split[0]
		right := split[1]
		rightConnections, rightExists := connections[right]

		if !rightExists {
			connections[right] = year2024.NewHashSet[string]()
			rightConnections = connections[right]
		}

		leftConnections, leftExists := connections[left]

		if !leftExists {
			connections[left] = year2024.NewHashSet[string]()
			leftConnections = connections[left]
		}
		leftConnections.Add(right)
		rightConnections.Add(left)
		if strings.HasPrefix(left, "t") {
			ts.Add(left)
		}
		if strings.HasPrefix(right, "t") {
			ts.Add(right)
		}
	}
	queue := year2024.Deque[PotentialHistorian]{}
	for t := range ts.Iterator() {
		visited := year2024.NewHashSet[string]()
		visited.Add(t)
		queue.Enqueue(PotentialHistorian{Last: t, Visited: visited})
	}
	uniquePaths := year2024.NewHashSet[string]()
	for !queue.IsEmpty() {
		next, _ := queue.RemoveFirst()
		if next.Visited.Size() == 3 {
			path := []string{}
			for e := range next.Visited.Iterator() {
				path = append(path, e)
			}
			sort.Strings(path)
			uniquePaths.Add(strings.Join(path, "|"))
			continue
		}

		neighbors := connections[next.Last]
		for n := range neighbors.Iterator() {
			if next.Visited.Contains(n) {
				continue
			}
			allConnected := true
			for v := range next.Visited.Iterator() {
				if !connections[n].Contains(v) {
					allConnected = false
				}
			}
			if !allConnected {
				continue
			}
			cloned := next.Visited.Clone()
			cloned.Add(n)
			queue.Enqueue(PotentialHistorian{Last: n, Visited: cloned})
		}
	}
	return uniquePaths.Size()
}
func NumTConnectedV2(input string) string {
	connections := make(map[string]*year2024.HashSet[string])
	lines := strings.Split(input, "\n")
	ts := year2024.NewHashSet[string]()
	for _, connection := range lines {
		split := strings.Split(connection, "-")
		left := split[0]
		right := split[1]
		rightConnections, rightExists := connections[right]

		if !rightExists {
			connections[right] = year2024.NewHashSet[string]()
			rightConnections = connections[right]
		}

		leftConnections, leftExists := connections[left]

		if !leftExists {
			connections[left] = year2024.NewHashSet[string]()
			leftConnections = connections[left]
		}
		leftConnections.Add(right)
		rightConnections.Add(left)
		if strings.HasPrefix(left, "t") {
			ts.Add(left)
		}
		if strings.HasPrefix(right, "t") {
			ts.Add(right)
		}
	}
	queue := year2024.NewMinHeap[PotentialHistorian](func(a, b PotentialHistorian) bool {
		return a.Visited.Size() > b.Visited.Size()
	})
	for k, _ := range connections {
		visited := year2024.NewHashSet[string]()
		visited.Add(k)
		queue.Offer(PotentialHistorian{Last: k, Visited: visited})
	}
	maxLength := 0
	sequence := ""
	seen := year2024.NewHashSet[string]()
	for !queue.IsEmpty() {
		next, _ := queue.Remove()
		var path []string
		for e := range next.Visited.Iterator() {
			path = append(path, e)
		}
		sort.Strings(path)
		key := strings.Join(path, ",")
		if seen.Contains(key) {
			continue
		}
		seen.Add(key)
		if next.Visited.Size() > maxLength {
			maxLength = next.Visited.Size()
			sequence = key
		}

		neighbors := connections[next.Last]
		tooSmall := false
		for v := range next.Visited.Iterator() {
			if connections[v].Size() <= maxLength {
				tooSmall = true
			}
		}
		if tooSmall {
			continue
		}
		for n := range neighbors.Iterator() {
			if next.Visited.Contains(n) {
				continue
			}
			allConnected := true
			if connections[n].Size() < maxLength || connections[n].Size() < next.Visited.Size() {
				continue
			}
			for v := range next.Visited.Iterator() {
				if !connections[n].Contains(v) {
					allConnected = false
				}
			}
			if !allConnected {
				continue
			}
			cloned := next.Visited.Clone()
			cloned.Add(n)
			queue.Offer(PotentialHistorian{Last: n, Visited: cloned})
		}
	}
	return sequence
}
