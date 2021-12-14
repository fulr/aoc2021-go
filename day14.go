package aoc2021

import (
	"fmt"
	"math"
)

func Day14Part1() {
	lines := readFile("inputs/day14.txt")

	rules := make([]rule, 0)
	for _, l := range lines[2:] {
		r := []rune(l)
		rules = append(rules, rule{r[0], r[1], r[6]})
	}

	var head *atom
	r := []rune(lines[0])
	for i := range r {
		head = &atom{r[len(r)-1-i], head}
	}

	for i := 0; i < 10; i++ {
		for current := head; current.next != nil; current = current.next {
			for _, rule := range rules {
				if rule.first == current.element && rule.second == current.next.element {
					current.next = &atom{rule.addon, current.next}
					current = current.next
					break
				}
			}
		}
	}

	counts := make(map[rune]int)
	for current := head; current != nil; current = current.next {
		counts[current.element]++
	}

	min := math.MaxInt
	max := 0
	for _, v := range counts {
		if v < min {
			min = v
		}
		if v > max {
			max = v
		}
	}

	fmt.Println(counts)

	result := max - min

	fmt.Printf("day 14 part1 %v\n", result)
}

func Day14Part2() {
	lines := readFile("inputs/day14.txt")
	rules := make([]rule, 0)
	for _, l := range lines[2:] {
		r := []rune(l)
		rules = append(rules, rule{r[0], r[1], r[6]})
	}

	r := []rune(lines[0])
	ruleCounts := make(map[rule]int)
	elementCounts := make(map[rune]int)
	for i := range r[:len(r)-1] {
		elementCounts[r[i]]++
		for _, rule := range rules {
			if r[i] == rule.first && r[i+1] == rule.second {
				ruleCounts[rule]++
			}
		}
	}
	elementCounts[r[len(r)-1]]++

	for i := 0; i < 40; i++ {
		newCounts := make(map[rule]int)
		for rule, count := range ruleCounts {
			for _, nextRule := range rules {
				if nextRule.first == rule.first && nextRule.second == rule.addon {
					newCounts[nextRule] += count
				}
				if nextRule.first == rule.addon && nextRule.second == rule.second {
					newCounts[nextRule] += count
				}
			}
			elementCounts[rule.addon] += count
		}
		ruleCounts = newCounts
	}

	min := math.MaxInt
	max := 0
	for _, v := range elementCounts {
		if v < min {
			min = v
		}
		if v > max {
			max = v
		}
	}

	result := max - min

	fmt.Printf("day 14 part2 %v\n", result)
}

type atom struct {
	element rune
	next    *atom
}

type rule struct {
	first  rune
	second rune
	addon  rune
}

func (r rule) String() string {
	return fmt.Sprintf("%c%c -> %c", r.first, r.second, r.addon)
}
