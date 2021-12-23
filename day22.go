package aoc2021

import (
	"fmt"
	"math"
	"strings"
)

func Day22Part1() {
	lines := readFile("inputs/day22.txt")

	set := make(map[v3]bool)

	for _, l := range lines {
		d := parseD22(l)
		if d.min[0] < -50 || d.max[0] > 50 || d.min[1] < -50 || d.max[1] > 50 || d.min[2] < -50 || d.max[2] > 50 {
			continue
		}
		for x := d.min[0]; x <= d.max[0]; x++ {
			for y := d.min[1]; y <= d.max[1]; y++ {
				for z := d.min[2]; z <= d.max[2]; z++ {
					set[v3{x, y, z}] = d.cmd == "on"
				}
			}
		}
	}

	result := 0
	for _, b := range set {
		if b {
			result++
		}
	}
	fmt.Printf("day 22 part1 %v\n", result)
}

func Day22Part2() {
	lines := readFile("inputs/day22.txt")

	list := make([]d22line, 0)
	for _, l := range lines {
		list = append(list, parseD22(l))
	}

	scales := [3]nlScale{newScale(), newScale(), newScale()}
	for _, l := range list {
		scales[0].addMax(l.max[0])
		scales[0].addMin(l.min[0])
		scales[1].addMax(l.max[1])
		scales[1].addMin(l.min[1])
		scales[2].addMax(l.max[2])
		scales[2].addMin(l.min[2])
	}
	fmt.Println(len(scales[0]), len(scales[1]), len(scales[2]))

	set := make(map[v3]bool)
	for i, l := range list {
		fmt.Println(i, l)
		if l.cmd == "on" {
			for x, xs := range scales[0] {
				for y, ys := range scales[1] {
					for z, zs := range scales[2] {
						if l.min[0] <= xs[1] && l.max[0] >= xs[0] &&
							l.min[1] <= ys[1] && l.max[1] >= ys[0] &&
							l.min[2] <= zs[1] && l.max[2] >= zs[0] {
							set[v3{x, y, z}] = true
						}
					}
				}
			}
		} else {
			for x, xs := range scales[0] {
				for y, ys := range scales[1] {
					for z, zs := range scales[2] {
						if l.min[0] <= xs[1] && l.max[0] >= xs[0] &&
							l.min[1] <= ys[1] && l.max[1] >= ys[0] &&
							l.min[2] <= zs[1] && l.max[2] >= zs[0] {
							delete(set, v3{x, y, z})
						}
					}
				}
			}
		}
	}

	result := 0
	for k, v := range set {
		if v {
			result += (scales[0][k[0]][1] - scales[0][k[0]][0] + 1) *
				(scales[1][k[1]][1] - scales[1][k[1]][0] + 1) *
				(scales[2][k[2]][1] - scales[2][k[2]][0] + 1)
		}
	}

	fmt.Printf("day 22 part2 %v\n", result)
}

type d22line struct {
	cmd string
	min v3
	max v3
}

func parseD22(s string) d22line {
	parts := strings.Split(s, " ")
	coords := strings.Split(parts[1], ",")

	result := d22line{cmd: parts[0]}

	result.min[0], result.max[0] = parseCoord(coords[0])
	result.min[1], result.max[1] = parseCoord(coords[1])
	result.min[2], result.max[2] = parseCoord(coords[2])

	return result
}

func parseCoord(s string) (int, int) {
	p := strings.Split(s[2:], "..")
	return toint(p[0]), toint(p[1])
}

type nlScale [][2]int

func newScale() nlScale {
	return nlScale{{math.MinInt, math.MaxInt}}
}

func (s *nlScale) addMax(max int) {
	i := 0
	for !((*s)[i][0] <= max && max <= (*s)[i][1]) {
		i++
	}
	newScale := make([][2]int, 0)
	for x, mm := range *s {
		if x == i {
			newScale = append(newScale, [2]int{mm[0], max}, [2]int{max + 1, mm[1]})
		} else {
			newScale = append(newScale, mm)
		}
	}
	*s = newScale
}

func (s *nlScale) addMin(min int) {
	i := 0
	for !((*s)[i][0] <= min && min <= (*s)[i][1]) {
		i++
	}
	newScale := make([][2]int, 0)
	for x, mm := range *s {
		if x == i {
			newScale = append(newScale, [2]int{mm[0], min - 1}, [2]int{min, mm[1]})
		} else {
			newScale = append(newScale, mm)
		}
	}
	*s = newScale
}
