package aoc2021

import (
	"fmt"
)

func Day21Part1() {
	// p1Pos := 4
	// p2Pos := 8
	p1Pos := 5
	p2Pos := 6
	p1Points := 0
	p2Points := 0
	die := 1
	dieCount := 0

	result := 0
	for {
		roll := 0
		roll, die = rollDie(0, die)
		roll, die = rollDie(roll, die)
		roll, die = rollDie(roll, die)
		dieCount += 3

		p1Points, p1Pos = pointsPos(p1Points, p1Pos, roll)

		if p1Points >= 1000 {
			result = dieCount * p2Points
			break
		}

		roll, die = rollDie(0, die)
		roll, die = rollDie(roll, die)
		roll, die = rollDie(roll, die)
		dieCount += 3

		p2Points, p2Pos = pointsPos(p2Points, p2Pos, roll)

		if p2Points >= 1000 {
			result = dieCount * p1Points
			break
		}
	}

	fmt.Printf("day 21 part1 %v\n", result)
}

func Day21Part2() {
	dirac := buildOutcomes()
	states := make(map[d21state]int)
	states[d21state{0, 5, 0, 6}] = 1

	p1 := 0
	p2 := 0
	for {
		done := true
		newState := make(map[d21state]int)
		for state, count := range states {
			if count == 0 {
				continue
			}
			if state.p2Points >= 21 {
				p2 += count
				continue
			}
			done = false
			for roll, nrOfUniverses := range dirac {
				p1Points, p1Pos := pointsPos(state.p1Points, state.p1Pos, roll)
				newState[d21state{p1Points, p1Pos, state.p2Points, state.p2Pos}] += nrOfUniverses * count
			}
		}
		states = newState
		newState = make(map[d21state]int)
		for state, count := range states {
			if count == 0 {
				continue
			}
			if state.p1Points >= 21 {
				p1 += count
				continue
			}
			done = false
			for roll, nrOfUniverses := range dirac {
				p2Points, p2Pos := pointsPos(state.p2Points, state.p2Pos, roll)
				newState[d21state{state.p1Points, state.p1Pos, p2Points, p2Pos}] += nrOfUniverses * count

			}
		}
		states = newState
		if done {
			break
		}
	}

	fmt.Printf("day 21 part2 p1 %v p2 %v\n", p1, p2)
}

func rollDie(rollIn int, diceIn int) (roll int, die int) {
	roll = rollIn + diceIn
	die = diceIn + 1
	if die == 101 {
		die = 1
	}
	return
}

func pointsPos(points, pos, roll int) (int, int) {
	pos += roll
	pos = (pos-1)%10 + 1
	points += pos
	return points, pos
}

type d21state struct {
	p1Points int
	p1Pos    int
	p2Points int
	p2Pos    int
}

func buildOutcomes() map[int]int {
	result := make(map[int]int)
	for a := 1; a <= 3; a++ {
		for b := 1; b <= 3; b++ {
			for c := 1; c <= 3; c++ {
				result[a+b+c]++
			}
		}
	}
	return result
}
