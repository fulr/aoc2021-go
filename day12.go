package aoc2021

import (
	"fmt"
	"strings"
	"unicode"
)

func Day12Part1() {
	lines := readFile("inputs/day12.txt")
	g := NewGraph(lines)
	paths := []Path{NewPath("start")}
	for {
		done := true
		newp := make([]Path, 0)
		for _, p := range paths {
			last := p[len(p)-1]
			if last == "end" {
				newp = append(newp, p)
				continue
			}
		neighbors:
			for _, nextCave := range g.Find(last) {
				newPath := p.Add(nextCave)
				if unicode.IsLower([]rune(nextCave)[0]) {
					for _, cave := range p {
						if cave == nextCave {
							continue neighbors
						}
					}
				}
				for _, t := range paths {
					if newPath.Equal(t) {
						continue neighbors
					}
				}
				newp = append(newp, newPath)
				done = false
			}
		}
		if done {
			break
		}
		paths = newp
	}

	result := 0
	for _, p := range paths {
		if p[len(p)-1] == "end" {
			result++
		}
	}

	fmt.Printf("day 12 part1 %v\n", result)
}

func Day12Part2() {
	lines := readFile("inputs/day12.txt")
	g := NewGraph(lines)
	paths := []Path{NewPath("start")}
	for {
		done := true
		newp := make([]Path, 0)
		for _, p := range paths {
			last := p[len(p)-1]
			if last == "end" {
				newp = append(newp, p)
				continue
			}
		neighbors:
			for _, nextCave := range g.Find(last) {
				if nextCave == "start" {
					continue
				}
				newPath := p.Add(nextCave)
				if unicode.IsLower([]rune(nextCave)[0]) {
					counts := make(map[string]int)
					for _, x := range newPath {
						if unicode.IsLower([]rune(x)[0]) {
							counts[x]++
						}
					}
					once := false
					for _, c := range counts {
						if c == 2 {
							if once {
								continue neighbors
							} else {
								once = true
							}
						}
						if c > 2 {
							continue neighbors
						}
					}
				}
				for _, t := range paths {
					if newPath.Equal(t) {
						continue neighbors
					}
				}
				newp = append(newp, newPath)
				done = false
			}
		}
		if done {
			break
		}
		paths = newp
	}

	result := 0
	for _, p := range paths {
		if p[len(p)-1] == "end" {
			result++
		}
	}
	fmt.Printf("day 12 part2 %v\n", result)
}

type Graph struct {
	g [][2]string
}

func NewGraph(lines []string) Graph {
	g := make([][2]string, 0)
	for _, l := range lines {
		s := strings.Split(l, "-")
		g = append(g, [2]string{s[0], s[1]})
	}
	return Graph{
		g: g,
	}
}

func (g Graph) Find(s string) []string {
	result := make([]string, 0)
	for _, x := range g.g {
		if x[0] == s {
			result = append(result, x[1])
		}
		if x[1] == s {
			result = append(result, x[0])
		}
	}
	return result
}

type Path []string

func NewPath(s string) Path {
	return Path{s}
}

func (p Path) Add(s string) Path {
	return append(append([]string{}, p...), s)
}

func (a Path) Equal(b Path) bool {
	if len(a) != len(b) {
		return false
	}
	for i, x := range a {
		if x != b[i] {
			return false
		}
	}
	return true
}
