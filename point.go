package aoc2021

import (
	"log"
	"strconv"
	"strings"
)

type Point struct {
	X int
	Y int
}

func NewPoint(s string) Point {
	xy := strings.Split(s, ",")
	x, err := strconv.Atoi(xy[0])
	if err != nil {
		log.Fatal(err)
	}
	y, err := strconv.Atoi(xy[1])
	if err != nil {
		log.Fatal(err)
	}
	return Point{X: x, Y: y}
}

func (p Point) Neighbors8() []Point {
	return []Point{
		{p.X + 1, p.Y + 1},
		{p.X + 1, p.Y},
		{p.X + 1, p.Y - 1},
		{p.X - 1, p.Y + 1},
		{p.X - 1, p.Y},
		{p.X - 1, p.Y - 1},
		{p.X, p.Y + 1},
		{p.X, p.Y - 1},
	}
}

func (p Point) Neighbors4() []Point {
	return []Point{
		{p.X + 1, p.Y},
		{p.X - 1, p.Y},
		{p.X, p.Y + 1},
		{p.X, p.Y - 1},
	}
}
