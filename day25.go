package aoc2021

import (
	"fmt"
)

func Day25Part1() {
	lines := readFile("inputs/day25.txt")

	s := d25state(lines)
	result := 1
	for !s.move() {
		result++
	}

	fmt.Printf("day 25 part1 %v\n", result)
}

func Day25Part2() {
	lines := readFile("inputs/day25.txt")

	result := len(lines)
	fmt.Printf("day 25 part2 %v\n", result)
}

type d25state []string

func (s d25state) move() bool {
	done := true
	for y := range s {
		n := []rune(s[y])

		for x, a := range s[y] {
			if a != '>' {
				continue
			}
			nx := (x + 1) % len(s[y])
			if s[y][nx] == '.' {
				n[x], n[nx] = n[nx], n[x]
				done = false
			}
		}
		s[y] = string(n)
	}

	newstate := d25state(make([]string, len(s)))
	copy(newstate, s)
	for y := range s {
		n := (y + 1) % len(s)
		for x, a := range s[y] {
			if a != 'v' {
				continue
			}
			if s[n][x] == '.' {
				newstate[y] = newstate[y][:x] + "." + newstate[y][x+1:]
				newstate[n] = newstate[n][:x] + "v" + newstate[n][x+1:]
				done = false
			}
		}
	}
	copy(s, newstate)
	return done
}

func (s d25state) count() (int, int) {
	a := 0
	b := 0
	for _, y := range s {
		for _, x := range y {
			switch x {
			case '>':
				a++
			case 'v':
				b++
			}
		}
	}
	return a, b
}
