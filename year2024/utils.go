package year2024

import (
	"os"
	"strconv"
)

func ReadInput(location string) string {
	file, err := os.ReadFile(location)
	if err != nil {
		panic(err)
	}
	return string(file)
}

type Coordinate struct {
	X int
	Y int
}

func ToIntSlice(strings []string) []int {
	ints := make([]int, len(strings))
	for i, s := range strings {
		n, err := strconv.Atoi(s)
		if err != nil {
			panic(err)
		}
		ints[i] = n
	}
	return ints
}

func ToRuneSlice(strings []string) [][]rune {
	runes := make([][]rune, len(strings))
	for i, s := range strings {
		runes[i] = []rune(s)
	}
	return runes
}

func OffTheMap(coordinate Coordinate, grid [][]rune) bool {
	return coordinate.Y >= len(grid) || coordinate.Y < 0 || coordinate.X >= len(grid[0]) || coordinate.X < 0
}

func Abs(a int, b int) int {
	if a >= b {
		return a - b
	} else {
		return b - a
	}
}

type Deque[T any] struct {
	elements []T
}

func (q *Deque[T]) Enqueue(value T) {
	q.elements = append(q.elements, value)
}

func (q *Deque[T]) RemoveFirst() (T, bool) {
	var zeroValue T
	if len(q.elements) == 0 {
		return zeroValue, false
	}
	front := q.elements[0]
	q.elements = q.elements[1:]
	return front, true
}
func (q *Deque[T]) RemoveLast() (T, bool) {
	var zeroValue T
	if len(q.elements) == 0 {
		return zeroValue, false
	}
	tail := q.elements[len(q.elements)-1]
	q.elements = q.elements[:len(q.elements)-1]
	return tail, true
}

func (q *Deque[T]) IsEmpty() bool {
	return len(q.elements) == 0
}

type MinHeap struct {
	elements []int
}

func (h *MinHeap) Offer(element int) {
	h.elements = append(h.elements, element)
	h.bubbleUp(len(h.elements) - 1)
}

func (h *MinHeap) Remove() (int, bool) {
	if len(h.elements) == 0 {
		return 0, false
	}
	root := h.elements[0]

	h.elements[0] = h.elements[len(h.elements)-1]
	h.elements = h.elements[:len(h.elements)-1]

	h.bubbleDown(0)
	return root, true
}

func (h *MinHeap) bubbleUp(index int) {
	parentIndex := (index - 1) / 2
	for index > 0 && h.elements[index] < h.elements[parentIndex] {
		h.elements[index], h.elements[parentIndex] = h.elements[parentIndex], h.elements[index]
		index = parentIndex
		parentIndex = (index - 1) / 2
	}
}

func (h *MinHeap) bubbleDown(index int) {
	leftChild := 2*index + 1
	rightChild := 2*index + 2
	smallest := index

	if leftChild < len(h.elements) && h.elements[leftChild] < h.elements[smallest] {
		smallest = leftChild
	}

	if rightChild < len(h.elements) && h.elements[rightChild] < h.elements[smallest] {
		smallest = rightChild
	}

	if smallest != index {
		h.elements[index], h.elements[smallest] = h.elements[smallest], h.elements[index]
		h.bubbleDown(smallest)
	}
}

func (h *MinHeap) IsEmpty() bool {
	return len(h.elements) == 0
}

func (h *MinHeap) Peek() (int, bool) {
	if len(h.elements) == 0 {
		return 0, false
	}
	return h.elements[0], true
}
