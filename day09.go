package aoc2021

import (
	"fmt"
	"sort"
)

func Day09Part1() {
	lines := readFile("inputs/day09.txt")

	set := make(map[Point]int)
	for y, row := range lines {
		for x, level := range row {
			set[Point{x, y}] = int(level) - '0' + 1
		}
	}

	result := 0
	for k, v := range set {
		if l, ok := set[Point{k.X - 1, k.Y}]; ok && l <= v {
			continue
		}
		if l, ok := set[Point{k.X + 1, k.Y}]; ok && l <= v {
			continue
		}
		if l, ok := set[Point{k.X, k.Y - 1}]; ok && l <= v {
			continue
		}
		if l, ok := set[Point{k.X, k.Y + 1}]; ok && l <= v {
			continue
		}
		result += v
	}

	fmt.Printf("day 09 part1 %v\n", result)
}

func Day09Part2() {
	lines := readFile("inputs/day09.txt")

	set := make(map[Point]int)
	for y, row := range lines {
		for x, level := range row {
			if level != '9' {
				set[Point{x, y}] = int(level) - '0'
			}
		}
	}

	done := make(map[Point]bool)
	todo := make([]Point, 0)
	counts := make([]int, 0)

	for len(done) < len(set) {
		for k := range set {
			if _, ok := done[k]; ok {
				continue
			}
			todo = append(todo, k)
			break
		}
		count := 0
		for len(todo) > 0 {
			k := todo[len(todo)-1]
			todo = todo[:len(todo)-1]
			if _, ok := set[k]; !ok {
				continue
			}
			if v, ok := done[k]; ok && v {
				continue
			}
			count++
			done[k] = true
			todo = append(todo,
				Point{k.X - 1, k.Y},
				Point{k.X + 1, k.Y},
				Point{k.X, k.Y - 1},
				Point{k.X, k.Y + 1})
		}
		counts = append(counts, count)
	}

	sort.Ints(counts)

	result := counts[len(counts)-1] * counts[len(counts)-2] * counts[len(counts)-3]

	fmt.Printf("day 09 part2 %v\n", result)
}
