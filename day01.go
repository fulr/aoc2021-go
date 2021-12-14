package aoc2021

import (
	"fmt"
)

func Day01Part1() {
	nrs := readInt("inputs/day01.txt")

	resultPart1 := 0
	for i := range nrs[:len(nrs)-1] {
		if nrs[i] < nrs[i+1] {
			resultPart1++
		}
	}
	fmt.Printf("day 01 part1 %v\n", resultPart1)
}

func Day01Part2() {
	nrs := readInt("inputs/day01.txt")
	resultPart2 := 0
	for i := range nrs[:len(nrs)-3] {
		if nrs[i]+nrs[i+1]+nrs[i+2] < nrs[i+1]+nrs[i+2]+nrs[i+3] {
			resultPart2++
		}
	}
	fmt.Printf("day 01 part2 %v\n", resultPart2)
}
