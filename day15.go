package aoc2021

import (
	"fmt"
)

func Day15Part1() {
	lines := readFile("inputs/day15.txt")

	set := NewPointSet(lines)

	result := findPathHeap(set, Point{0, 0}, Point{99, 99})

	fmt.Printf("day 15 part1 %v\n", result)
}

func limitAdd(a int, b int) int {
	if a+b > 9 {
		return a + b - 9
	}
	return a + b
}

func Day15Part2() {
	lines := readFile("inputs/day15.txt")

	set := NewPointSet(lines)

	newSet := PointSet{}
	for i := 0; i < 5; i++ {
		for j := 0; j < 5; j++ {
			for p, r := range set {
				newSet[Point{p.X + i*100, p.Y + j*100}] = limitAdd(r, i+j)
			}
		}
	}

	result := findPathHeap(newSet, Point{0, 0}, Point{499, 499})
	fmt.Printf("day 15 part2 %v\n", result)
}

func findPathHeap(set PointSet, from Point, to Point) int {
	todo := new(heap)
	todo.insert(0, from)
	totalRisk := make(map[Point]int)
	totalRisk[from] = 0
	visited := make(map[Point]bool)
	prev := make(map[Point]Point)
	for len(*todo) > 0 {
		min, minPoint := todo.pop()
		if visited[minPoint] {
			continue
		}
		visited[minPoint] = true

		for _, neighbor := range minPoint.Neighbors4() {
			risk, ok := set[neighbor]
			if !ok {
				continue
			}
			tr, trOk := totalRisk[neighbor]
			if trOk && tr < min+risk {
				continue
			}
			totalRisk[neighbor] = min + risk
			prev[neighbor] = minPoint
			todo.insert(min+risk, neighbor)
		}
	}

	return totalRisk[to]
}

type heapItem struct {
	key   int
	point Point
}

type heap []heapItem

func (h heap) swap(a, b int) {
	h[a], h[b] = h[b], h[a]
}

func (h heap) decrease(i int) {
	for i > 0 && h[i].key < h[i/2].key {
		h.swap(i, i/2)
		i /= 2
	}
}

func (h *heap) insert(key int, point Point) {
	*h = append(*h, heapItem{key, point})
	h.decrease(len(*h) - 1)
}

func (h heap) heapify(i int) {
	left := 2*i + 1
	right := 2*i + 2
	min := i
	if left < len(h) && h[left].key < h[min].key {
		min = left
	}
	if right < len(h) && h[right].key < h[min].key {
		min = right
	}
	if min != i {
		h.swap(i, min)
		h.heapify(min)
	}
}

func (h *heap) pop() (int, Point) {
	result := (*h)[0]
	h.swap(0, len(*h)-1)
	*h = (*h)[:len(*h)-1]
	h.heapify(0)
	return result.key, result.point
}
