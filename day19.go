package aoc2021

import (
	"fmt"
	"strings"
)

func Day19Part1() {
	lines := readFile("inputs/day19.txt")

	scanners := make([][]v3, 0, 20)
	for _, l := range lines {
		if l == "" {
			continue
		}
		if l[1] == '-' {
			scanners = append(scanners, make([]v3, 0))
			continue
		}
		s := strings.Split(l, ",")
		scanners[len(scanners)-1] = append(scanners[len(scanners)-1], v3{toint(s[0]), toint(s[1]), toint(s[2])})
	}
	fmt.Println(len(scanners))

	transformations := make([]transformation, 0)
	for i, a := range scanners {
		for j, b := range scanners {
			if i == j {
				continue
			}
			for _, m := range transformers {
				hits := make(map[v3]int)
				for _, x := range a {
					for _, y := range b {
						hits[y.diff(m.mult(x))]++
					}
				}
				for v, c := range hits {
					if c == 12 {
						transformations = append(transformations, transformation{m, v, i, j})
						break
					}
				}
			}
		}
	}

	fmt.Println("done")

	beacons := make(map[int]map[v3]bool)
	for s := range scanners {
		beacons[s] = make(map[v3]bool)
		for _, v := range scanners[s] {
			beacons[s][v] = true
		}
	}
outer:
	for {
		for _, t := range transformations {
			for v := range beacons[t.from] {
				beacons[t.to][t.transform(v)] = true
			}
		}
		l := len(beacons[0])
		for _, b := range beacons {
			if l != len(b) {
				break
			}
			break outer
		}
	}

	result := len(beacons[0])
	fmt.Printf("day 19 part1 %v\n", result)
}

func Day19Part2() {
	lines := readFile("inputs/day19.txt")
	scanners := make([][]v3, 0, 20)
	for _, l := range lines {
		if l == "" {
			continue
		}
		if l[1] == '-' {
			scanners = append(scanners, make([]v3, 0))
			continue
		}
		s := strings.Split(l, ",")
		scanners[len(scanners)-1] = append(scanners[len(scanners)-1], v3{toint(s[0]), toint(s[1]), toint(s[2])})
	}
	fmt.Println(len(scanners))

	transformations := make([]transformation, 0)
	for i, a := range scanners {
		for j, b := range scanners {
			if i == j {
				continue
			}
			for _, m := range transformers {
				hits := make(map[v3]int)
				for _, x := range a {
					for _, y := range b {
						hits[y.diff(m.mult(x))]++
					}
				}
				for v, c := range hits {
					if c == 12 {
						transformations = append(transformations, transformation{m, v, i, j})
						break
					}
				}
			}
		}
	}

	fmt.Println("done")

	beacons := make(map[int]map[v3]bool)
	for s := range scanners {
		beacons[s] = map[v3]bool{{0, 0, 0}: true}
	}
outer:
	for {
		for _, t := range transformations {
			for v := range beacons[t.from] {
				beacons[t.to][t.transform(v)] = true
			}
		}
		for _, b := range beacons {
			if len(scanners) != len(b) {
				break
			}
			break outer
		}
	}

	result := 0
	for a := range beacons[0] {
		for b := range beacons[0] {
			dist := a.manhattan(b)
			if dist > result {
				result = dist
			}
		}
	}

	fmt.Printf("day 19 part2 %v\n", result)
}

type v3 [3]int
type m3 [3][3]int

type transformation struct {
	matrix      m3
	translation v3
	from        int
	to          int
}

func (t transformation) transform(v v3) v3 {
	return t.matrix.mult(v).add(t.translation)
}

func (a v3) manhattan(b v3) int {
	dist := 0
	if a[0] > b[0] {
		dist += a[0] - b[0]
	} else {
		dist += b[0] - a[0]
	}
	if a[1] > b[1] {
		dist += a[1] - b[1]
	} else {
		dist += b[1] - a[1]
	}
	if a[2] > b[2] {
		dist += a[2] - b[2]
	} else {
		dist += b[2] - a[2]
	}
	return dist
}

func (a v3) add(b v3) v3 {
	return v3{
		a[0] + b[0],
		a[1] + b[1],
		a[2] + b[2],
	}
}

func (a v3) diff(b v3) v3 {
	return v3{
		a[0] - b[0],
		a[1] - b[1],
		a[2] - b[2],
	}
}

func (m m3) mult(v v3) v3 {
	return v3{
		m[0][0]*v[0] + m[0][1]*v[1] + m[0][2]*v[2],
		m[1][0]*v[0] + m[1][1]*v[1] + m[1][2]*v[2],
		m[2][0]*v[0] + m[2][1]*v[1] + m[2][2]*v[2],
	}
}

var transformers = []m3{
	{{0, 0, 1}, {0, 1, 0}, {1, 0, 0}},
	{{0, 0, 1}, {1, 0, 0}, {0, 1, 0}},
	{{0, 1, 0}, {0, 0, 1}, {1, 0, 0}},
	{{0, 1, 0}, {1, 0, 0}, {0, 0, 1}},
	{{1, 0, 0}, {0, 0, 1}, {0, 1, 0}},
	{{1, 0, 0}, {0, 1, 0}, {0, 0, 1}},

	{{0, 0, 1}, {0, 1, 0}, {-1, 0, 0}},
	{{0, 0, 1}, {-1, 0, 0}, {0, 1, 0}},
	{{0, 1, 0}, {0, 0, 1}, {-1, 0, 0}},
	{{0, 1, 0}, {-1, 0, 0}, {0, 0, 1}},
	{{-1, 0, 0}, {0, 0, 1}, {0, 1, 0}},
	{{-1, 0, 0}, {0, 1, 0}, {0, 0, 1}},

	{{0, 0, 1}, {0, -1, 0}, {1, 0, 0}},
	{{0, 0, 1}, {1, 0, 0}, {0, -1, 0}},
	{{0, -1, 0}, {0, 0, 1}, {1, 0, 0}},
	{{0, -1, 0}, {1, 0, 0}, {0, 0, 1}},
	{{1, 0, 0}, {0, 0, 1}, {0, -1, 0}},
	{{1, 0, 0}, {0, -1, 0}, {0, 0, 1}},

	{{0, 0, -1}, {0, 1, 0}, {1, 0, 0}},
	{{0, 0, -1}, {1, 0, 0}, {0, 1, 0}},
	{{0, 1, 0}, {0, 0, -1}, {1, 0, 0}},
	{{0, 1, 0}, {1, 0, 0}, {0, 0, -1}},
	{{1, 0, 0}, {0, 0, -1}, {0, 1, 0}},
	{{1, 0, 0}, {0, 1, 0}, {0, 0, -1}},

	{{0, 0, 1}, {0, -1, 0}, {-1, 0, 0}},
	{{0, 0, 1}, {-1, 0, 0}, {0, -1, 0}},
	{{0, -1, 0}, {0, 0, 1}, {-1, 0, 0}},
	{{0, -1, 0}, {-1, 0, 0}, {0, 0, 1}},
	{{-1, 0, 0}, {0, 0, 1}, {0, -1, 0}},
	{{-1, 0, 0}, {0, -1, 0}, {0, 0, 1}},

	{{0, 0, -1}, {0, -1, 0}, {1, 0, 0}},
	{{0, 0, -1}, {1, 0, 0}, {0, -1, 0}},
	{{0, -1, 0}, {0, 0, -1}, {1, 0, 0}},
	{{0, -1, 0}, {1, 0, 0}, {0, 0, -1}},
	{{1, 0, 0}, {0, 0, -1}, {0, -1, 0}},
	{{1, 0, 0}, {0, -1, 0}, {0, 0, -1}},

	{{0, 0, -1}, {0, 1, 0}, {-1, 0, 0}},
	{{0, 0, -1}, {-1, 0, 0}, {0, 1, 0}},
	{{0, 1, 0}, {0, 0, -1}, {-1, 0, 0}},
	{{0, 1, 0}, {-1, 0, 0}, {0, 0, -1}},
	{{-1, 0, 0}, {0, 0, -1}, {0, 1, 0}},
	{{-1, 0, 0}, {0, 1, 0}, {0, 0, -1}},

	{{0, 0, -1}, {0, -1, 0}, {-1, 0, 0}},
	{{0, 0, -1}, {-1, 0, 0}, {0, -1, 0}},
	{{0, -1, 0}, {0, 0, -1}, {-1, 0, 0}},
	{{0, -1, 0}, {-1, 0, 0}, {0, 0, -1}},
	{{-1, 0, 0}, {0, 0, -1}, {0, -1, 0}},
	{{-1, 0, 0}, {0, -1, 0}, {0, 0, -1}},
}
