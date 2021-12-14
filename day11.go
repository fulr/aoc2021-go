package aoc2021

import (
	"fmt"
)

func Day11Part1() {
	lines := readFile("inputs/day11.txt")

	set := NewPointSet(lines)

	result := 0
	for i := 0; i < 100; i++ {
		flashes := make(map[Point]bool)
		for p := range set {
			set[p]++
		}
		for {
			flashed := false
			for p, l := range set {
				if flashes[p] {
					continue
				}
				if l > 9 {
					flashed = true
					set.IncNeighbors(p)
					flashes[p] = true
				}
			}
			if !flashed {
				break
			}
		}

		for p, f := range flashes {
			if f {
				result++
				set[p] = 0
			}
		}
	}

	fmt.Printf("day 11 part1 %v\n", result)
}

func Day11Part2() {
	lines := readFile("inputs/day11.txt")

	set := NewPointSet(lines)

	result := 0
	for {
		result++
		flashes := make(map[Point]bool)
		for p := range set {
			set[p]++
		}
		for {
			flashed := false
			for p, l := range set {
				if flashes[p] {
					continue
				}
				if l > 9 {
					flashed = true
					set.IncNeighbors(p)
					flashes[p] = true
				}
			}
			if !flashed {
				break
			}
		}
		count := 0
		for p, f := range flashes {
			if f {
				set[p] = 0
				count++
			}
		}
		if count == len(set) {
			break
		}
	}

	fmt.Printf("day 11 part2 %v\n", result)
}

type PointSet map[Point]int

func NewPointSet(lines []string) PointSet {
	set := make(map[Point]int)
	for y, l := range lines {
		for x, r := range l {
			set[Point{x, y}] = int(r) - '0'
		}
	}
	return set
}

func (s PointSet) IncNeighbors(p Point) {
	for _, n := range p.Nieghbors8() {
		if _, ok := s[n]; ok {
			s[n]++
		}
	}
}
