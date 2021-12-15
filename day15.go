package aoc2021

import (
	"fmt"
	"math"
)

func Day15Part1() {
	lines := readFile("inputs/day15.txt")

	set := NewPointSet(lines)

	result := findPath(set, Point{0, 0}, Point{99, 99})

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

	result := findPath(newSet, Point{0, 0}, Point{499, 499})
	fmt.Printf("day 15 part2 %v\n", result)
}

func findPath(set PointSet, from Point, to Point) int {
	todo := []Point{from}
	totalRisk := make(map[Point]int)
	totalRisk[todo[0]] = 0
	visited := make(map[Point]bool)
	prev := make(map[Point]Point)
	for len(todo) > 0 {
		min := math.MaxInt
		var minPoint Point
		for _, p := range todo {
			tr, trOk := totalRisk[p]
			if !visited[p] && trOk && tr < min {
				min = tr
				minPoint = p
			}
		}

		newTodo := make([]Point, 0)
		for _, p := range todo {
			if p == minPoint {
				continue
			}
			newTodo = append(newTodo, p)
		}
		todo = newTodo

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
			todo = append(todo, neighbor)
		}
	}

	return totalRisk[to]
}
