package aoc2021

import (
	"fmt"
	"log"
	"math"
	"strconv"
	"strings"
)

type fold struct {
	value int
	x     bool
}

func Day13Part1() {
	lines := readFile("inputs/day13.txt")

	paper := make(map[Point]bool)
	folds := make([]fold, 0)
	for _, l := range lines {
		if l == "" {
			continue
		}
		if strings.HasPrefix(l, "fold along x=") {
			x, err := strconv.Atoi(l[13:])
			if err != nil {
				log.Fatal("fold x", err)
			}
			folds = append(folds, fold{x, true})
			continue
		}
		if strings.HasPrefix(l, "fold along y=") {
			x, err := strconv.Atoi(l[13:])
			if err != nil {
				log.Fatal(err)
			}
			folds = append(folds, fold{x, false})
			continue
		}
		p := NewPoint(l)
		paper[p] = true
	}

	f := folds[0]
	foldedPaper := make(map[Point]bool)
	if f.x {
		for p := range paper {
			foldedPaper[Point{int(math.Abs(float64(p.X - f.value))), p.Y}] = true
		}
	} else {
		for p := range paper {
			foldedPaper[Point{p.X, int(math.Abs(float64(p.Y - f.value)))}] = true
		}
	}

	result := len(foldedPaper)
	fmt.Printf("day 13 part1 %v\n", result)
}

func Day13Part2() {
	lines := readFile("inputs/day13.txt")

	paper := make(map[Point]bool)
	folds := make([]fold, 0)
	for _, l := range lines {
		if l == "" {
			continue
		}
		if strings.HasPrefix(l, "fold along x=") {
			x, err := strconv.Atoi(l[13:])
			if err != nil {
				log.Fatal("fold x", err)
			}
			folds = append(folds, fold{x, true})
			continue
		}
		if strings.HasPrefix(l, "fold along y=") {
			x, err := strconv.Atoi(l[13:])
			if err != nil {
				log.Fatal(err)
			}
			folds = append(folds, fold{x, false})
			continue
		}
		p := NewPoint(l)
		paper[p] = true
	}

	for _, f := range folds {
		foldedPaper := make(map[Point]bool)
		if f.x {
			for p := range paper {
				if p.X > f.value {
					foldedPaper[Point{2*f.value - p.X, p.Y}] = true
				} else {
					foldedPaper[Point{p.X, p.Y}] = true
				}
			}
		} else {
			for p := range paper {
				if p.Y > f.value {
					foldedPaper[Point{p.X, 2*f.value - p.Y}] = true
				} else {
					foldedPaper[Point{p.X, p.Y}] = true
				}
			}
		}
		paper = foldedPaper
	}

	minX := 10000000
	minY := 10000000
	maxX := 0
	maxY := 0
	for p := range paper {
		if p.X < minX {
			minX = p.X
		}
		if p.X > maxX {
			maxX = p.X
		}
		if p.Y < minY {
			minY = p.Y
		}
		if p.Y > maxY {
			maxY = p.Y
		}
	}

	screen := make([][]rune, maxY-minY+1)

	for y := range screen {
		screen[y] = make([]rune, maxX-minX+1)
		for x := range screen[y] {
			screen[y][x] = ' '
		}
	}

	for p, v := range paper {
		if v {
			screen[p.Y-minY][p.X-minX] = '#'
		}
	}

	for _, l := range screen {
		fmt.Println(string(l))
	}

	result := len(paper)
	fmt.Printf("day 13 part2 %v\n", result)
}
