package aoc2021

import (
	"fmt"
)

func Day01() {
	nrs := readInt("inputs/day01.txt")

	resultPart1 := 0
	for i := range nrs[:len(nrs)-1] {
		if nrs[i] < nrs[i+1] {
			resultPart1++
		}
	}
	fmt.Printf("part1 %v\n", resultPart1)

	resultPart2 := 0
	for i := range nrs[:len(nrs)-3] {
		if nrs[i]+nrs[i+1]+nrs[i+2] < nrs[i+1]+nrs[i+2]+nrs[i+3] {
			resultPart2++
		}
	}
	fmt.Printf("part2 %v\n", resultPart2)
}
