package aoc2021

import (
	"fmt"
	"math"
)

func Day20Part1() {
	lines := readFile("inputs/day20.txt")

	lookup := lines[0]

	m := make(map[Point]int)
	for y, l := range lines[2:] {
		for x, r := range l {
			if r == '#' {
				m[Point{x, y}] = 1
			}
		}
	}

	minX := math.MaxInt
	maxX := math.MinInt
	minY := math.MaxInt
	maxY := math.MinInt
	for p := range m {
		if p.X > maxX {
			maxX = p.X
		}
		if p.X < minX {
			minX = p.X
		}
		if p.Y > maxY {
			maxY = p.Y
		}
		if p.Y < minY {
			minY = p.Y
		}
	}

	fmt.Printf("x: %v - %v y: %v - %v\n", minX, maxX, minY, maxY)

	newMap := make(map[Point]int)
	for x := minX - 5; x <= maxX+5; x++ {
		for y := minY - 5; y <= maxY+5; y++ {
			p := Point{x, y}
			l := lookupNr(m, p)
			if lookup[l] == '#' {
				newMap[p] = 1
			}
		}
	}
	m = newMap
	newMap = make(map[Point]int)
	for x := minX - 5; x <= maxX+5; x++ {
		for y := minY - 5; y <= maxY+5; y++ {
			p := Point{x, y}
			l := lookupNr(m, p)
			if lookup[l] == '#' {
				newMap[p] = 1
			}
		}
	}
	m = newMap

	result := 0
	for x := minX - 2; x <= maxX+2; x++ {
		for y := minY - 2; y <= maxY+2; y++ {
			p := Point{x, y}
			if m[p] == 1 {
				result++
			}
		}
	}

	fmt.Printf("day 20 part1 %v\n", result)
}

func Day20Part2() {
	lines := readFile("inputs/day20.txt")

	lookup := lines[0]

	m := make(map[Point]int)
	for y, l := range lines[2:] {
		for x, r := range l {
			if r == '#' {
				m[Point{x, y}] = 1
			}
		}
	}

	minX := 0
	maxX := len(lines[2]) - 1
	minY := 0
	maxY := len(lines) - 3

	fmt.Printf("x: %v - %v y: %v - %v\n", minX, maxX, minY, maxY)

	iterations := 50

	for i := 0; i < iterations; i++ {
		newMap := make(map[Point]int)
		for x := minX - 2*iterations - 10; x <= maxX+2*iterations+10; x++ {
			for y := minY - 2*iterations - 10; y <= maxY+2*iterations+10; y++ {
				p := Point{x, y}
				l := lookupNr(m, p)
				if lookup[l] == '#' {
					newMap[p] = 1
				}
			}
		}
		m = newMap
	}

	result := 0
	for x := minX - iterations - 10; x <= maxX+iterations+10; x++ {
		for y := minY - iterations - 10; y <= maxY+iterations+10; y++ {
			p := Point{x, y}
			if m[p] == 1 {
				result++
			}
		}
	}

	fmt.Printf("day 20 part2 %v\n", result)
}

func lookupNr(m map[Point]int, p Point) int {
	return m[Point{p.X - 1, p.Y - 1}]*256 +
		m[Point{p.X, p.Y - 1}]*128 +
		m[Point{p.X + 1, p.Y - 1}]*64 +
		m[Point{p.X - 1, p.Y}]*32 +
		m[Point{p.X, p.Y}]*16 +
		m[Point{p.X + 1, p.Y}]*8 +
		m[Point{p.X - 1, p.Y + 1}]*4 +
		m[Point{p.X, p.Y + 1}]*2 +
		m[Point{p.X + 1, p.Y + 1}]
}
