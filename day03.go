package aoc2021

import (
	"fmt"
	"log"
	"strconv"
)

func Day03Part1() {
	lines := readFile("inputs/day03.txt")

	count := make([]int, len(lines[0]))
	for _, l := range lines {
		for i, c := range l {
			if c == '1' {
				count[i]++
			}
		}
	}

	gammaRate := ""
	epsilonRate := ""
	for _, c := range count {
		if c > len(lines)/2 {
			gammaRate += "1"
			epsilonRate += "0"
		} else {
			gammaRate += "0"
			epsilonRate += "1"
		}
	}

	fmt.Printf("gamma %v epsilon %v\n", gammaRate, epsilonRate)

	g, err := strconv.ParseInt(gammaRate, 2, 0)
	if err != nil {
		log.Fatal(err)
	}
	e, err := strconv.ParseInt(epsilonRate, 2, 0)
	if err != nil {
		log.Fatal(err)
	}

	result := g * e
	fmt.Printf("part1 %v\n", result)
}

func findCO2OrO2(lines []string, pos int, co2 bool) string {
	if len(lines) == 1 {
		return lines[0]
	}
	one := make([]string, 0)
	zero := make([]string, 0)
	for _, l := range lines {
		if l[pos] == '1' {
			one = append(one, l)
		} else {
			zero = append(zero, l)
		}
	}
	if (len(one) >= len(zero) || co2) && !(len(one) >= len(zero) && co2) { //xor
		return findCO2OrO2(one, pos+1, co2)
	} else {
		return findCO2OrO2(zero, pos+1, co2)

	}
}

func Day03Part2() {
	lines := readFile("inputs/day03.txt")

	o2, err := strconv.ParseInt(findCO2OrO2(lines, 0, false), 2, 0)
	if err != nil {
		log.Fatal(err)
	}
	co2, err := strconv.ParseInt(findCO2OrO2(lines, 0, true), 2, 0)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Printf("o2 %v co2 %v\n", o2, co2)

	result := o2 * co2
	fmt.Printf("part2 %v\n", result)
}
