package aoc2021

import (
	"math"
	"testing"
)

func TestDay16(t *testing.T) {
	Day16Part1()
	Day16Part2()
}

func TestDay16Tests(t *testing.T) {
	table := []struct {
		input  string
		output int
	}{
		{"D2FE28", 2021},
		{"C200B40A82", 3},
		{"04005AC33890", 54},
		{"880086C3E88112", 7},
		{"CE00C43D881120", 9},
	}
	for _, te := range table {
		p, _ := parse(hexToBits(te.input), math.MaxInt)
		output := p[0].eval()
		if output != te.output {
			t.Fail()
		}
	}
}
