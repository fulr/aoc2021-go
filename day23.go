package aoc2021

import (
	"fmt"
)

func Day23Part1() {
	from := d23state{
		hallway: "...........",
		rooms:   [4]string{"BC", "BC", "DA", "DA"},
	}

	// test
	// from := d23state{
	// 	hallway: "...........",
	// 	rooms:   [4]string{"BA", "CD", "BC", "DA"},
	// }

	to := d23state{
		hallway: "...........",
		rooms:   [4]string{"AA", "BB", "CC", "DD"},
	}
	result := d23findPathHeap(from, to)
	fmt.Printf("day 23 part1 %v\n", result)
}

func Day23Part2() {
	from := d23state{
		hallway: "...........",
		rooms:   [4]string{"BDDC", "BCBC", "DBAA", "DACA"},
	}

	// test
	// from := d23state{
	// 	hallway: "...........",
	// 	rooms:   [4]string{"BDDA", "CCBD", "BBAC", "DACA"},
	// }

	to := d23state{
		hallway: "...........",
		rooms:   [4]string{"AAAA", "BBBB", "CCCC", "DDDD"},
	}
	result := d23findPathHeap(from, to)
	fmt.Printf("day 23 part2 %v\n", result)
}

type d23state struct {
	hallway string
	rooms   [4]string
}

type d23stateWithEnergy struct {
	d23state
	energy int
}

func (s d23state) next() []d23stateWithEnergy {
	result := []d23stateWithEnergy{}
	// move out of the room
	for k, x := range roomHallwayPos {
		roomPos := 0
		for roomPos < len(s.rooms[k]) {
			if s.rooms[k][roomPos] != '.' {
				break
			}
			roomPos++
		}
		if roomPos == len(s.rooms[k]) {
			continue
		}
	outer:
		for nx := x + 1; nx < len(s.hallway); nx++ {
			for _, xx := range roomHallwayPos {
				if xx == nx {
					continue outer
				}
			}
			if s.hallway[nx] != '.' {
				break
			}

			h := s.hallway[:nx] + s.rooms[k][roomPos:roomPos+1] + s.hallway[nx+1:]
			r := [4]string{}
			copy(r[:], s.rooms[:])
			r[k] = r[k][0:roomPos] + "." + r[k][roomPos+1:]
			result = append(result, d23stateWithEnergy{d23state{h, r}, (1 + nx - x + roomPos) * energyCosts[s.rooms[k][roomPos]]})
		}
	outer2:
		for nx := x - 1; nx >= 0; nx-- {
			for _, xx := range roomHallwayPos {
				if xx == nx {
					continue outer2
				}
			}
			if s.hallway[nx] != '.' {
				break
			}

			h := s.hallway[:nx] + s.rooms[k][roomPos:roomPos+1] + s.hallway[nx+1:]
			r := [4]string{}
			copy(r[:], s.rooms[:])
			r[k] = r[k][0:roomPos] + "." + r[k][roomPos+1:]
			result = append(result, d23stateWithEnergy{d23state{h, r}, (1 + x - nx + roomPos) * energyCosts[s.rooms[k][roomPos]]})
		}
	}

	// move from hallway to room
outer3:
	for x, a := range s.hallway {
		if a == '.' {
			continue
		}
		dest := destRoom[s.hallway[x]]
		destX := roomHallwayPos[dest]
		for rr, ar := range s.rooms[dest] {
			if ar != '.' && dest != destRoom[s.rooms[dest][rr]] {
				continue outer3
			}
		}
		if s.rooms[dest][0] != '.' {
			continue
		}
		inc := cmp(x, destX)
		// check for free way
		for nx := x + inc; nx != destX; nx += inc {
			if s.hallway[nx] != '.' {
				continue outer3
			}
		}
		rooms := [4]string{}
		copy(rooms[:], s.rooms[:])
		roomPos := 0
		for roomPos < len(rooms[dest])-1 && rooms[dest][roomPos+1] == '.' {
			roomPos++
		}
		rooms[dest] = rooms[dest][0:roomPos] + string(a) + rooms[dest][roomPos+1:]
		h := s.hallway[:x] + "." + s.hallway[x+1:]
		result = append(result, d23stateWithEnergy{d23state{h, rooms}, (roomPos + delta(destX, x) + 1) * energyCosts[s.hallway[x]]})
	}

	return result
}

func delta(a, b int) int {
	if a > b {
		return a - b
	} else {
		return b - a
	}
}

func sswap(s string, a, b int) string {
	r := []rune(s)
	r[a], r[b] = r[b], r[a]
	return string(r)
}

var energyCosts = map[byte]int{
	'A': 1,
	'B': 10,
	'C': 100,
	'D': 1000,
}

var destRoom = map[byte]int{
	'A': 0,
	'B': 1,
	'C': 2,
	'D': 3,
}

var roomHallwayPos = map[int]int{
	0: 2,
	1: 4,
	2: 6,
	3: 8,
}

func d23findPathHeap(from, to d23state) int {
	todo := new(d23heap)
	todo.insert(0, from)
	total := make(map[d23state]int)
	total[from] = 0
	visited := make(map[d23state]bool)
	prev := make(map[d23state]d23state)
	for len(*todo) > 0 {
		min, minState := todo.pop()
		if visited[minState] {
			continue
		}
		visited[minState] = true

		if minState == to {
			break
		}

		for _, neighbor := range minState.next() {
			tr, trOk := total[neighbor.d23state]
			if trOk && tr < min+neighbor.energy {
				continue
			}
			total[neighbor.d23state] = min + neighbor.energy
			prev[neighbor.d23state] = minState
			todo.insert(min+neighbor.energy, neighbor.d23state)
		}
	}

	current := to
	for {
		fmt.Println(total[current], current)
		c, ok := prev[current]
		if ok {
			current = c
		} else {
			break
		}
	}

	return total[to]
}

type d23heapItem struct {
	key   int
	state d23state
}

type d23heap []d23heapItem

func (h d23heap) swap(a, b int) {
	h[a], h[b] = h[b], h[a]
}

func (h d23heap) decrease(i int) {
	for i > 0 && h[i].key < h[i/2].key {
		h.swap(i, i/2)
		i /= 2
	}
}

func (h *d23heap) insert(key int, s d23state) {
	*h = append(*h, d23heapItem{key, s})
	h.decrease(len(*h) - 1)
}

func (h d23heap) heapify(i int) {
	left := 2*i + 1
	right := 2*i + 2
	min := i
	if left < len(h) && h[left].key < h[min].key {
		min = left
	}
	if right < len(h) && h[right].key < h[min].key {
		min = right
	}
	if min != i {
		h.swap(i, min)
		h.heapify(min)
	}
}

func (h *d23heap) pop() (int, d23state) {
	result := (*h)[0]
	h.swap(0, len(*h)-1)
	*h = (*h)[:len(*h)-1]
	h.heapify(0)
	return result.key, result.state
}
