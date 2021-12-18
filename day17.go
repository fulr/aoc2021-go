package aoc2021

import (
	"fmt"
)

// target area: x=265..287, y=-103..-58

func Day17Part1() {
	// targetYMin := -10
	// targetYMax := -5
	// targetXMin := 20
	// targetXMax := 30
	targetYMin := -103
	targetYMax := -58
	targetXMin := 265
	targetXMax := 287

	maxMaxY := 0
	for ivx := 2; ivx < 300; ivx++ {
		for ivy := 2; ivy < 300; ivy++ {
			maxY := 0
			x := 0
			y := 0
			vx := ivx
			vy := ivy
			hit := false
			for {
				x += vx
				y += vy
				vx += cmp(vx, 0)
				vy--
				if y > maxY {
					maxY = y
				}
				if x >= targetXMin && x <= targetXMax && y >= targetYMin && y <= targetYMax {
					hit = true
					break
				}
				if y < targetYMin {
					break
				}
			}
			if hit {
				if maxY > maxMaxY {
					maxMaxY = maxY
				}
			}
		}
	}

	result := maxMaxY
	fmt.Printf("day 17 part1 %v\n", result)
}

func Day17Part2() {
	// test data
	// targetYMin := -10
	// targetYMax := -5
	// targetXMin := 20
	// targetXMax := 30

	// real data
	targetYMin := -103
	targetYMax := -58
	targetXMin := 265
	targetXMax := 287

	count := 0
	for ivx := 2; ivx < 300; ivx++ {
		for ivy := -200; ivy < 500; ivy++ {
			x := 0
			y := 0
			vx := ivx
			vy := ivy
			hit := false
			for {
				x += vx
				y += vy
				vx += cmp(vx, 0)
				vy--
				if x >= targetXMin && x <= targetXMax && y >= targetYMin && y <= targetYMax {
					hit = true
					break
				}
				if y < targetYMin {
					break
				}
			}
			if hit {
				count++
			}
		}
	}

	result := count
	fmt.Printf("day 17 part2 %v\n", result)
}
