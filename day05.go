package aoc2021

import (
	"fmt"
	"log"
	"strconv"
	"strings"
)

func Day05Part1() {
	inputLines := readFile("inputs/day05.txt")

	lines := make([]PointLine, len(inputLines))
	for i, l := range inputLines {
		lines[i] = parseD5Line(l)
	}

	hits := make(map[Point]int)

	for _, l := range lines {
		if l.From.X == l.To.X {
			inc := 1
			if l.From.Y > l.To.Y {
				inc = -1
			}
			fmt.Println(l)
			for i := l.From.Y; i != l.To.Y; i += inc {
				hits[Point{X: l.From.X, Y: i}]++
			}
			hits[l.To]++
		}
		if l.From.Y == l.To.Y {
			inc := 1
			if l.From.X > l.To.X {
				inc = -1
			}
			fmt.Println(l)
			for i := l.From.X; i != l.To.X; i += inc {
				hits[Point{Y: l.From.Y, X: i}]++
			}
			hits[l.To]++
		}
	}

	count := 0
	for _, v := range hits {
		if v > 1 {
			count++
		}
	}

	result := count
	fmt.Printf("part1 %v\n", result)
}

func Day05Part2() {
	inputLines := readFile("inputs/day05.txt")

	lines := make([]PointLine, len(inputLines))
	for i, l := range inputLines {
		lines[i] = parseD5Line(l)
	}

	hits := make(map[Point]int)

	for _, l := range lines {
		incX := cmp(l.From.X, l.To.X)
		incY := cmp(l.From.Y, l.To.Y)
		for x, y := l.From.X, l.From.Y; x != l.To.X || y != l.To.Y; x, y = x+incX, y+incY {
			hits[Point{X: x, Y: y}]++
		}
		hits[l.To]++
	}

	count := 0
	for _, v := range hits {
		if v > 1 {
			count++
		}
	}

	result := count
	fmt.Printf("part2 %v\n", result)
}

func cmp(a int, b int) int {
	if a < b {
		return 1
	}
	if b < a {
		return -1
	}
	return 0
}

func parseD5Line(s string) PointLine {
	ps := strings.Split(s, " -> ")
	return PointLine{
		From: parsePoint(ps[0]),
		To:   parsePoint(ps[1]),
	}
}

func parsePoint(s string) Point {
	xy := strings.Split(s, ",")
	x, err := strconv.Atoi(xy[0])
	if err != nil {
		log.Fatal(err)
	}
	y, err := strconv.Atoi(xy[1])
	if err != nil {
		log.Fatal(err)
	}
	return Point{X: x, Y: y}
}
