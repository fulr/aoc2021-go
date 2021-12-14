package aoc2021

import (
	"fmt"
	"sort"
)

func Day10Part1() {
	lines := readFile("inputs/day10.txt")

	end := map[rune]rune{
		')': '(',
		']': '[',
		'}': '{',
		'>': '<',
	}
	points := map[rune]int{
		')': 3,
		']': 57,
		'}': 1197,
		'>': 25137,
	}

	result := 0
	for _, l := range lines {
		stack := make([]rune, 0)
	line:
		for _, r := range l {
			if x, ok := end[r]; ok {
				if stack[len(stack)-1] == x {
					stack = stack[:len(stack)-1]
					continue
				}
				result += points[r]
				break line
			}
			stack = append(stack, r)
		}
	}
	fmt.Printf("day 10 part1 %v\n", result)
}

func Day10Part2() {
	lines := readFile("inputs/day10.txt")

	end := map[rune]rune{
		')': '(',
		']': '[',
		'}': '{',
		'>': '<',
	}
	points := map[rune]int{
		'(': 1,
		'[': 2,
		'{': 3,
		'<': 4,
	}

	results := make([]int, 0)
line:
	for _, l := range lines {
		stack := make([]rune, 0)
		for _, r := range l {
			if x, ok := end[r]; ok {
				if stack[len(stack)-1] == x {
					stack = stack[:len(stack)-1]
					continue
				}
				continue line
			}
			stack = append(stack, r)
		}
		result := 0
		for i := len(stack) - 1; i >= 0; i-- {
			result = result*5 + points[stack[i]]
		}
		results = append(results, result)
	}

	sort.Ints(results)

	result := results[len(results)/2]

	fmt.Printf("day 10 part2 %v\n", result)
}
