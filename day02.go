package aoc2021

import (
	"fmt"
	"log"
	"strconv"
	"strings"
)

type Line struct {
	Cmd   string
	Count int
}

func NewLine(s string) *Line {
	parts := strings.Split(s, " ")
	i, err := strconv.Atoi(parts[1])
	if err != nil {
		log.Fatal("no conv")
	}
	return &Line{Cmd: parts[0], Count: i}
}

func Day02Part1() {
	lines := readFile("inputs/day02.txt")
	h := 0
	v := 0
	for _, l := range lines {
		c := NewLine(l)
		switch c.Cmd {
		case "forward":
			h += c.Count
		case "up":
			v -= c.Count
		case "down":
			v += c.Count
		default:
			log.Fatal("error")
		}
	}
	fmt.Printf("day 02 part1 h %v v %v mult %v\n", h, v, h*v)
}

func Day02Part2() {
	lines := readFile("inputs/day02.txt")
	h := 0
	v := 0
	aim := 0
	for _, l := range lines {
		c := NewLine(l)
		switch c.Cmd {
		case "forward":
			h += c.Count
			v += aim * c.Count
		case "up":
			aim -= c.Count
		case "down":
			aim += c.Count
		default:
			log.Fatal("error")
		}
	}
	fmt.Printf("day 02 part 2 h %v v %v mult %v\n", h, v, h*v)
}
