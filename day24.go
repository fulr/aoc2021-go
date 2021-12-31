package aoc2021

import (
	"fmt"
)

func Day24Part1() {
	params := [][]int{ // w1+c1+b2==w2
		{1, 13, 6},   // w1=3
		{1, 15, 7},   // w2=9
		{1, 15, 10},  // w3=4
		{1, 11, 2},   // w4 = 9
		{26, -7, 15}, // w4+2-7=w5 -> w4-5=w5 -> w4=9 -> w5=4
		{1, 10, 8},   // w6=1
		{1, 10, 1},   // w7 = 9
		{26, -5, 10}, // w7+1-5=w8 -> w7-4=w8 -> w7=9 -> w8=5
		{1, 15, 5},   // w9 = 7
		{26, -3, 3},  // w10=w9+5-3 -> w10=w9+2 -> w10=9
		{26, 0, 5},   // w11=w6+8+0 -> w11=9
		{26, -5, 11}, // w12=w3+10-5 -> w12=w3+5 -> w12=9
		{26, -9, 12}, // w13=w2+7-9 -> w13=w2-2 -> w13=7
		{26, 0, 10},  // w14=w1+6 -> w14=9
	}
	// 39494195799979

	fmt.Printf("day 24 part1 %v\n", params)
}

func Day24Part2() {
	params := [][]int{ // w1+c1+b2==w2
		{1, 13, 6},   // w1=1
		{1, 15, 7},   // w2=3
		{1, 15, 10},  // w3=1
		{1, 11, 2},   // w4 = 6
		{26, -7, 15}, // w5=w4+2-7 -> w5=w4-5 -> w4=6 -> w5=1
		{1, 10, 8},   // w6=1
		{1, 10, 1},   // w7 = 5
		{26, -5, 10}, // w8=w7+1-5 -> w8=w7-4 -> w7=5 -> w8=1
		{1, 15, 5},   // w9 = 1
		{26, -3, 3},  // w10=w9+5-3 -> w10=w9+2 -> w10=3
		{26, 0, 5},   // w11=w6+8+0 -> w11=9
		{26, -5, 11}, // w12=w3+10-5 -> w12=w3+5 -> w12=6
		{26, -9, 12}, // w13=w2+7-9 -> w13=w2-2 -> w13=1
		{26, 0, 10},  // w14=w1+6 -> w14=7
	}
	// 13161151139617

	fmt.Printf("day 24 part2 %v\n", params)
}

func d24recc(params [][]int, zIn, wIn int) (int, bool) {
	if len(params) == 0 {
		return wIn, zIn == 0
	}
	for w := 9; w > 0; w-- {
		z := check(params[0][0], params[0][1], params[0][2], w, zIn)
		ww := 10*wIn + w
		if len(params) > 7 {
			fmt.Println(ww)
		}
		r, ok := d24recc(params[1:], z, ww)
		if ok {
			return r, true
		}
	}
	return 0, false
}

func check(a, b, c, w, z int) int {
	// inp w
	// mul x 0
	// add x z
	// mod x 26
	// div z 1 = a
	// add x 13 = b
	x := (z % 26) + b // w1+c1+b2==w2
	z /= a
	// eql x w
	// eql x 0
	if x != w {
		// mul y 0
		// add y 25
		// mul y x
		// add y 1
		// mul z y
		z *= 26
		// mul y 0
		// add y w
		// add y 6 = c
		// mul y x
		// add z y
		z += w + c
	}
	return z
}
