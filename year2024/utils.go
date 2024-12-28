package year2024

import (
	"fmt"
	"os"
	"strconv"
	"strings"
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

func OffTheMap[T any](coordinate Coordinate, grid [][]T) bool {
	return coordinate.Y >= len(grid) || coordinate.Y < 0 || coordinate.X >= len(grid[0]) || coordinate.X < 0
}

func OffTheMap2[T any](coordinate Coordinate, grid *[][]T) bool {
	return coordinate.Y >= len(*grid) || coordinate.Y < 0 || coordinate.X >= len((*grid)[0]) || coordinate.X < 0
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

func (q *Deque[T]) Size() int {
	return len(q.elements)
}

type MinHeap[T comparable] struct {
	elements []T
	less     func(a, b T) bool
}

func NewMinHeap[T comparable](less func(a, b T) bool) *MinHeap[T] {
	return &MinHeap[T]{
		less: less,
	}
}

func (h *MinHeap[T]) Offer(element T) {
	h.elements = append(h.elements, element)
	h.bubbleUp(len(h.elements) - 1)
}

func (h *MinHeap[T]) Remove() (T, bool) {
	var zeroValue T
	if len(h.elements) == 0 {
		return zeroValue, false
	}
	root := h.elements[0]

	h.elements[0] = h.elements[len(h.elements)-1]
	h.elements = h.elements[:len(h.elements)-1]

	h.bubbleDown(0)
	return root, true
}

func (h *MinHeap[T]) bubbleUp(index int) {
	parentIndex := (index - 1) / 2
	for index > 0 && h.less(h.elements[index], h.elements[parentIndex]) {
		h.elements[index], h.elements[parentIndex] = h.elements[parentIndex], h.elements[index]
		index = parentIndex
		parentIndex = (index - 1) / 2
	}
}

func (h *MinHeap[T]) bubbleDown(index int) {
	leftChild := 2*index + 1
	rightChild := 2*index + 2
	smallest := index

	if leftChild < len(h.elements) && h.less(h.elements[leftChild], h.elements[smallest]) {
		smallest = leftChild
	}

	if rightChild < len(h.elements) && h.less(h.elements[rightChild], h.elements[smallest]) {
		smallest = rightChild
	}

	if smallest != index {
		h.elements[index], h.elements[smallest] = h.elements[smallest], h.elements[index]
		h.bubbleDown(smallest)
	}
}

func (h *MinHeap[T]) IsEmpty() bool {
	return len(h.elements) == 0
}

func (h *MinHeap[T]) Peek() (T, bool) {
	var zeroValue T
	if len(h.elements) == 0 {
		return zeroValue, false
	}
	return h.elements[0], true
}

func (h *MinHeap[T]) Size() int {
	return len(h.elements)
}

func AdjacentCoordinates[T any](coordinate Coordinate, grid [][]T) []Coordinate {
	var adjacencies []Coordinate
	up := Coordinate{X: coordinate.X, Y: coordinate.Y - 1}
	down := Coordinate{X: coordinate.X, Y: coordinate.Y + 1}
	left := Coordinate{X: coordinate.X - 1, Y: coordinate.Y}
	right := Coordinate{X: coordinate.X + 1, Y: coordinate.Y}
	for _, c := range []Coordinate{up, down, left, right} {
		if !OffTheMap(c, grid) {
			adjacencies = append(adjacencies, c)
		}
	}
	return adjacencies
}

func AdjacentCoordinatesIncludingOffTheMap(coordinate Coordinate) []Coordinate {
	up := Coordinate{X: coordinate.X, Y: coordinate.Y - 1}
	down := Coordinate{X: coordinate.X, Y: coordinate.Y + 1}
	left := Coordinate{X: coordinate.X - 1, Y: coordinate.Y}
	right := Coordinate{X: coordinate.X + 1, Y: coordinate.Y}
	return []Coordinate{up, down, left, right}
}

func FullAdjacenciesIncludingOffTheMap(coordinate Coordinate) []Coordinate {
	up := Coordinate{X: coordinate.X, Y: coordinate.Y - 1}
	down := Coordinate{X: coordinate.X, Y: coordinate.Y + 1}
	left := Coordinate{X: coordinate.X - 1, Y: coordinate.Y}
	right := Coordinate{X: coordinate.X + 1, Y: coordinate.Y}
	upLeft := Coordinate{X: coordinate.X - 1, Y: coordinate.Y - 1}
	upRight := Coordinate{X: coordinate.X + 1, Y: coordinate.Y - 1}
	downLeft := Coordinate{X: coordinate.X - 1, Y: coordinate.Y + 1}
	downRight := Coordinate{X: coordinate.X + 1, Y: coordinate.Y + 1}
	return []Coordinate{up, down, left, right, upLeft, upRight, downLeft, downRight}
}

type HashSet[T comparable] struct {
	elements map[T]struct{}
}

func NewHashSet[T comparable]() *HashSet[T] {
	return &HashSet[T]{elements: make(map[T]struct{})}
}

func (h *HashSet[T]) Add(element T) {
	if h.elements == nil {
		h.elements = make(map[T]struct{})
	}
	h.elements[element] = struct{}{}
}

func (h *HashSet[T]) Remove(element T) {
	delete(h.elements, element)
}

func (h *HashSet[T]) Size() int {
	return len(h.elements)
}

func (h *HashSet[T]) Iterator() <-chan T {
	ch := make(chan T)
	go func() {
		defer close(ch)
		for key := range h.elements {
			ch <- key
		}
	}()
	return ch
}

func (h *HashSet[T]) Contains(key T) bool {
	_, exists := h.elements[key]
	return exists
}

func (h *HashSet[T]) Values() []T {
	values := make([]T, len(h.elements))
	index := 0
	for k := range h.elements {
		values[index] = k
		index++
	}
	return values
}

func (h *HashSet[T]) Clone() *HashSet[T] {
	newHashSet := NewHashSet[T]()
	for obj := range h.Iterator() {
		newHashSet.Add(obj)
	}
	return newHashSet
}

func Mod(a, b int) int {
	return ((a % b) + b) % b
}

func Join(a []int, sep string) string {
	s := make([]string, len(a))
	for index, value := range a {
		s[index] = strconv.Itoa(value)
	}
	return strings.Join(s, sep)
}

type TrieNode struct {
	children map[rune]*TrieNode
	isEnd    bool
}

type Trie struct {
	root *TrieNode
}

func NewTrie() *Trie {
	return &Trie{
		root: &TrieNode{children: make(map[rune]*TrieNode)},
	}
}

func (t *Trie) Insert(word string) {
	currentNode := t.root
	for _, char := range word {
		if _, exists := currentNode.children[char]; !exists {
			currentNode.children[char] = &TrieNode{children: make(map[rune]*TrieNode)}
		}
		currentNode = currentNode.children[char]
	}
	currentNode.isEnd = true
}

func (t *Trie) Search(word string) bool {
	currentNode := t.root
	for _, char := range word {
		if _, exists := currentNode.children[char]; !exists {
			return false
		}
		currentNode = currentNode.children[char]
	}
	return currentNode.isEnd
}

func (t *Trie) SearchRunes(input []rune) bool {
	currentNode := t.root
	for _, char := range input {
		if _, exists := currentNode.children[char]; !exists {
			return false
		}
		currentNode = currentNode.children[char]
	}
	return currentNode.isEnd
}

func ManhattanDistance(origin Coordinate, point Coordinate) int {
	return Abs(origin.X, point.X) + Abs(origin.Y, point.Y)
}

type CircularQueue[T comparable] struct {
	elements []T
	size     int
	head     int
	tail     int
}

func NewCircularQueue[T comparable](size int) *CircularQueue[T] {
	elements := make([]T, size+1)
	return &CircularQueue[T]{elements: elements}
}

func (h *CircularQueue[T]) IsFull() bool {
	if h.tail > h.head {
		return h.head == 0 && h.tail == len(h.elements)-1
	} else {
		return h.head-h.tail == 1
	}
}

func (h *CircularQueue[T]) Add(element T) {
	if h.IsFull() {
		h.head = Mod(h.head+1, len(h.elements))
		h.tail = Mod(h.tail+1, len(h.elements))
		h.elements[Mod(h.tail-1, len(h.elements))] = element
	} else {
		h.tail = Mod(h.tail+1, len(h.elements))
		h.elements[Mod(h.tail-1, len(h.elements))] = element
	}
}

func (h *CircularQueue[T]) IsEmpty() bool {
	return h.tail == h.head
}

func (h *CircularQueue[T]) ToString() string {
	var buffer string
	for i := 0; i < len(h.elements)-1; i++ {
		index := Mod(h.head+i, len(h.elements))
		if index == h.tail {
			break
		}
		buffer += fmt.Sprintf("%v", h.elements[index])
		buffer += ","
	}
	return buffer
}
