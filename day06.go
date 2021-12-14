package aoc2021

import (
	"fmt"
)

func Day06Part1() {
	fishes := make([]int, 0, len(inputDay06))
	fishes = append(fishes, inputDay06[:]...)
	for generations := 0; generations < 80; generations++ {
		next := make([]int, len(fishes))
		for i, f := range fishes {
			next[i] = f - 1
			if f == 0 {
				next[i] = 6
				next = append(next, 8)
			}
		}
		fishes = next
	}

	result := len(fishes)
	fmt.Printf("part1 %v\n", result)
}

func Day06Part2() {
	var fishes [9]int
	for _, f := range inputDay06 {
		fishes[f]++
	}
	for generations := 0; generations < 256; generations++ {
		zero := fishes[0]
		fishes[0] = fishes[1]
		fishes[1] = fishes[2]
		fishes[2] = fishes[3]
		fishes[3] = fishes[4]
		fishes[4] = fishes[5]
		fishes[5] = fishes[6]
		fishes[6] = fishes[7] + zero
		fishes[7] = fishes[8]
		fishes[8] = zero
	}

	result := 0
	for _, f := range fishes {
		result += f
	}

	fmt.Printf("part2 %v\n", result)
}

var testInputDay06 = [...]int{3, 4, 3, 1, 2}
var inputDay06 = [...]int{
	2, 1, 1, 4, 4, 1, 3, 4, 2, 4, 2, 1, 1, 4, 3, 5, 1, 1, 5, 1, 1, 5, 4,
	5, 4, 1, 5, 1, 3, 1, 4, 2, 3, 2, 1, 2, 5, 5, 2, 3, 1, 2, 3, 3, 1, 4,
	3, 1, 1, 1, 1, 5, 2, 1, 1, 1, 5, 3, 3, 2, 1, 4, 1, 1, 1, 3, 1, 1, 5,
	5, 1, 4, 4, 4, 4, 5, 1, 5, 1, 1, 5, 5, 2, 2, 5, 4, 1, 5, 4, 1, 4, 1,
	1, 1, 1, 5, 3, 2, 4, 1, 1, 1, 4, 4, 1, 2, 1, 1, 5, 2, 1, 1, 1, 4, 4,
	4, 4, 3, 3, 1, 1, 5, 1, 5, 2, 1, 4, 1, 2, 4, 4, 4, 4, 2, 2, 2, 4, 4,
	4, 2, 1, 5, 5, 2, 1, 1, 1, 4, 4, 1, 4, 2, 3, 3, 3, 3, 3, 5, 4, 1, 5,
	1, 4, 5, 5, 1, 1, 1, 4, 1, 2, 4, 4, 1, 2, 3, 3, 3, 3, 5, 1, 4, 2, 5,
	5, 2, 1, 1, 1, 1, 3, 3, 1, 1, 2, 3, 2, 5, 4, 2, 1, 1, 2, 2, 2, 1, 3,
	1, 5, 4, 1, 1, 5, 3, 3, 2, 2, 3, 1, 1, 1, 1, 2, 4, 2, 2, 5, 1, 2, 4,
	2, 1, 1, 3, 2, 5, 5, 3, 1, 3, 3, 1, 4, 1, 1, 5, 5, 1, 5, 4, 1, 1, 1,
	1, 2, 3, 3, 1, 2, 3, 1, 5, 1, 3, 1, 1, 3, 1, 1, 1, 1, 1, 1, 5, 1, 1,
	5, 5, 2, 1, 1, 5, 2, 4, 5, 5, 1, 1, 5, 1, 5, 5, 1, 1, 3, 3, 1, 1, 3, 1}
