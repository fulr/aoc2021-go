package aoc2021

import (
	"fmt"
	"testing"
)

func TestDay23P1(t *testing.T) {
	Day23Part1()
}

func TestDay23P2(t *testing.T) {
	Day23Part2()
}

func TestDay23T1(t *testing.T) {
	// from := d23state{
	// 	hallway: "...........",
	// 	rooms:   [4]string{"AA", "CD", "BC", "DA"},
	// }
	to := d23state{
		hallway: "..........D",
		rooms:   [4]string{".BCD", "..CD", "...A", "...."},
	}
	// r := from.next()
	// for _, x := range r {
	// 	fmt.Println(x)
	// 	fmt.Println(x.d23state == to)
	// }
	fmt.Println(to.next())
}
