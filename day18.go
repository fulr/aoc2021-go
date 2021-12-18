package aoc2021

import (
	"fmt"
	"log"
)

func Day18Part1() {
	lines := readFile("inputs/day18.txt")

	var r *snailfishNode
	for _, l := range lines {
		s, _ := parseSnailfish(l, 0)
		s.reduce()
		r = addSnailfish(r, s)
		r.reduce()
	}

	result := r.magnitude()
	fmt.Printf("day 18 part1 %v\n", result)
}

func Day18Part2() {
	lines := readFile("inputs/day18.txt")

	result := len(lines)
	for _, a := range lines {
		for _, b := range lines {
			x, _ := parseSnailfish(a, 0)
			y, _ := parseSnailfish(b, 0)
			x.reduce()
			y.reduce()
			r := addSnailfish(x, y)
			r.reduce()
			mag := r.magnitude()
			if mag > result {
				result = mag
			}
		}
	}

	fmt.Printf("day 18 part2 %v\n", result)
}

type snailfishNode struct {
	left  *snailfishNode
	right *snailfishNode
	leaf  bool
	value int
}

func parseSnailfish(s string, i int) (*snailfishNode, int) {
	switch s[i] {
	case '[':
		left, lefti := parseSnailfish(s, i+1)
		if s[lefti] != ',' {
			log.Fatal(", expected")
		}
		right, righti := parseSnailfish(s, lefti+1)
		if s[righti] != ']' {
			log.Fatal("] expected")
		}
		return &snailfishNode{
			left:  left,
			right: right,
			leaf:  false,
			value: 0,
		}, righti + 1
	case ']':
		fallthrough
	case ',':
		log.Fatal("not expected")
		return nil, 0
	default:
		return &snailfishNode{
			nil,
			nil,
			true,
			int(s[i]) - '0',
		}, i + 1
	}
}

func (s snailfishNode) String() string {
	if s.leaf {
		return fmt.Sprint(s.value)
	} else {
		return fmt.Sprintf("[%v,%v]", s.left, s.right)
	}
}

func (s *snailfishNode) magnitude() int {
	if s.leaf {
		return s.value
	} else {
		return 3*s.left.magnitude() + 2*s.right.magnitude()
	}
}

func addSnailfish(a *snailfishNode, b *snailfishNode) *snailfishNode {
	if a == nil {
		return b
	}
	return &snailfishNode{
		left:  a,
		right: b,
		leaf:  false,
	}
}

func (s *snailfishNode) reduce() {
	for s.explode() || s.split() {
	}
}

func (s *snailfishNode) explode() bool {
	_, _, exploded := s.explodeIntern(1)
	return exploded
}

func (s *snailfishNode) explodeIntern(level int) (int, int, bool) {
	if !s.leaf {
		if level > 4 {
			s.leaf = true
			s.value = 0
			left := s.left.value
			right := s.right.value
			s.left = nil
			s.right = nil
			return left, right, true
		} else {
			left, right, exploded := s.left.explodeIntern(level + 1)
			if exploded {
				s.right.addValue(right, true)
				return left, 0, true
			}
			left, right, exploded = s.right.explodeIntern(level + 1)
			if exploded {
				s.left.addValue(left, false)
				return 0, right, true
			}
		}
	}
	return 0, 0, false
}

func (s *snailfishNode) addValue(value int, goLeft bool) {
	if s.leaf {
		s.value += value
	} else {
		if goLeft {
			s.left.addValue(value, goLeft)
		} else {
			s.right.addValue(value, goLeft)
		}
	}
}

func (s *snailfishNode) split() bool {
	if s.leaf {
		if s.value > 9 {
			s.left = &snailfishNode{nil, nil, true, s.value / 2}
			if s.value%2 == 0 {
				s.right = &snailfishNode{nil, nil, true, s.value / 2}
			} else {
				s.right = &snailfishNode{nil, nil, true, s.value/2 + 1}
			}
			s.leaf = false
			s.value = 0
			return true
		}
	} else {
		return s.left.split() || s.right.split()
	}
	return false
}
