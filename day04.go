package aoc2021

import (
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

func Day04Part1() {
	bingo := readBingo("inputs/day04.txt")

	winBoard := 0
	winNumber := 0
all:
	for _, randomNr := range bingo.RandomNumbers {
		for boardNr, board := range bingo.Boards {
			if board.Check(randomNr) {
				winBoard = boardNr
				winNumber = randomNr
				break all
			}
		}
	}

	sum := bingo.Boards[winBoard].Sum()

	result := sum * winNumber
	fmt.Printf("day 04 part1 %v\n", result)
}

func Day04Part2() {
	bingo := readBingo("inputs/day04.txt")

	countOfWonBoards := 0
	winBoard := 0
	winNumber := 0
all:
	for _, randomNr := range bingo.RandomNumbers {
		for boardNr, board := range bingo.Boards {
			if board.Check(randomNr) {
				countOfWonBoards++
			}
			if countOfWonBoards == len(bingo.Boards) {
				winBoard = boardNr
				winNumber = randomNr
				break all
			}
		}
	}

	sum := bingo.Boards[winBoard].Sum()

	result := sum * winNumber
	fmt.Printf("day 04 part2 %v\n", result)
}

func readBingo(file string) Bingo {
	f, err := os.ReadFile(file)
	if err != nil {
		log.Fatal(err)
	}
	s := string(f)

	segments := strings.Split(strings.ReplaceAll(s, "\r", ""), "\n\n")

	rngString := strings.Split(segments[0], ",")

	rn := make([]int, len(rngString))

	for i, r := range rngString {
		a, err := strconv.Atoi(r)
		if err != nil {
			log.Fatal(err)
		}
		rn[i] = a
	}

	boards := make([]*Board, len(segments)-1)
	hits := make([][]int, len(boards))
	for bnr, board := range segments[1:] {
		boards[bnr] = parseBoard(board)
		hits[bnr] = make([]int, 10)
	}

	return Bingo{
		RandomNumbers: rn,
		Boards:        boards,
		Hits:          hits,
	}
}

func parseBoard(b string) *Board {
	result := make([][]int, 5)
	hits := make([][]bool, 5)
	for y, line := range strings.Split(b, "\n") {
		result[y] = make([]int, 5)
		for x := 0; x < 5; x++ {
			a, err := strconv.Atoi(strings.TrimSpace(line[x*3 : (x*3)+2]))
			if err != nil {
				log.Fatal(err)
			}
			result[y][x] = a
		}
		hits[y] = make([]bool, 5)
	}
	return &Board{
		Numbers:  result,
		Hits:     hits,
		HitCount: make([]int, 10),
	}
}

type Board struct {
	Numbers  [][]int
	Hits     [][]bool
	HitCount []int
	Won      bool
}

func (board *Board) Check(randomNr int) bool {
all:
	for y, line := range board.Numbers {
		for x, nr := range line {
			if nr == randomNr {
				board.HitCount[x]++
				board.HitCount[y+5]++
				board.Hits[y][x] = true
				if (board.HitCount[x] == 5 || board.HitCount[y+5] == 5) && !board.Won {
					board.Won = true
					return true
				}
				break all
			}
		}
	}
	return false
}

func (board Board) Sum() int {
	sum := 0
	for y, line := range board.Numbers {
		for x, nr := range line {
			if !board.Hits[y][x] {
				sum += nr
			}
		}
	}
	return sum
}

type Bingo struct {
	RandomNumbers []int
	Boards        []*Board
	Hits          [][]int
}
