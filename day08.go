package aoc2021

import (
	"fmt"
	"strings"
)

func Day08Part1() {
	lines := readFile("inputs/day08.txt")

	hits := make(map[int]int)

	for _, l := range lines {
		parts := strings.Split(l, " | ")
		for _, w := range strings.Split(parts[1], " ") {
			hits[len(w)]++
		}
	}

	fmt.Println(hits)
	result := hits[2] + hits[3] + hits[4] + hits[7]
	fmt.Printf("day 08 part1 %v\n", result)
}

func Day08Part2() {
	lines := readFile("inputs/day08.txt")

	result := 0
	for _, l := range lines {
		parts := strings.Split(l, " | ")
		digits := make([]Set, 10)
		done := make([]bool, 10)
		numbers := make([]Set, 10)
		for i, d := range strings.Split(parts[0], " ") {
			digits[i] = NewSet(d)
		}
		for i, s := range digits {
			switch len(s) {
			case 2:
				numbers[1] = s
				done[i] = true
			case 3:
				numbers[7] = s
				done[i] = true
			case 4:
				numbers[4] = s
				done[i] = true
			case 7:
				numbers[8] = s
				done[i] = true
			}
		}
		for i, d := range digits {
			if done[i] {
				continue
			}
			if d.Contains(numbers[4]) {
				numbers[9] = d
				done[i] = true
				break
			}
		}
		for i, d := range digits {
			if done[i] {
				continue
			}
			if d.Contains(numbers[7]) && len(d) == 5 {
				numbers[3] = d
				done[i] = true
			}
			if d.Contains(numbers[7]) && len(d) == 6 {
				numbers[0] = d
				done[i] = true
			}
		}
		for i, d := range digits {
			if done[i] {
				continue
			}
			if len(d) == 6 {
				numbers[6] = d
				done[i] = true
				break
			}
		}
		for i, d := range digits {
			if done[i] {
				continue
			}
			if numbers[6].Contains(d) {
				numbers[5] = d
				done[i] = true
				break
			}
		}
		for i, d := range digits {
			if done[i] {
				continue
			}
			numbers[2] = d
			done[i] = true
			break
		}
		for i := range digits {
			if done[i] {
				continue
			}
			fmt.Println(done)
			fmt.Println(parts[0])
			panic("check")
		}

		nr := 0
		for _, w := range strings.Split(parts[1], " ") {
			s := NewSet(w)
			for i, n := range numbers {
				if n.Matches(s) {
					nr = nr*10 + i
					break
				}
			}
		}
		fmt.Println(nr)
		result += nr
	}

	fmt.Printf("day 08 part2 %v\n", result)
}

type Set map[rune]bool

func NewSet(s string) Set {
	r := make(Set)
	for _, c := range s {
		r[c] = true
	}
	return r
}

func (a Set) Diff(b Set) Set {
	r := make(Set)
	for k, v := range a {
		if v && !b[k] {
			r[k] = true
		}
	}
	return r
}

func (a Set) Contains(b Set) bool {
	for k, v := range b {
		if v && !a[k] {
			return false
		}
	}
	return true
}

func (a Set) Matches(b Set) bool {
	return a.Contains(b) && b.Contains(a)
}

func (a Set) Union(b Set) Set {
	r := make(Set)
	for k, v := range a {
		if v {
			r[k] = true
		}
	}
	for k, v := range b {
		if v {
			r[k] = true
		}
	}
	return r
}

func (a Set) Intersect(b Set) Set {
	r := make(Set)
	for k, v := range a {
		if v && b[k] {
			r[k] = true
		}
	}
	return r
}
