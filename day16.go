package aoc2021

import (
	"fmt"
	"log"
	"math"
	"strconv"
	"strings"
)

var toBits = map[rune]string{
	'0': "0000",
	'1': "0001",
	'2': "0010",
	'3': "0011",
	'4': "0100",
	'5': "0101",
	'6': "0110",
	'7': "0111",
	'8': "1000",
	'9': "1001",
	'A': "1010",
	'B': "1011",
	'C': "1100",
	'D': "1101",
	'E': "1110",
	'F': "1111",
}

func Day16Part1() {
	lines := readFile("inputs/day16.txt")

	bits := hexToBits(lines[0])

	p, _ := parse(bits, math.MaxInt)

	result := 0
	for _, s := range p {
		result += s.versionSum()
	}

	fmt.Printf("day 16 part1 %v\n", result)
}

func Day16Part2() {
	lines := readFile("inputs/day16.txt")

	bits := hexToBits(lines[0])

	p, _ := parse(bits, math.MaxInt)

	fmt.Println(p)

	result := p[0].eval()

	fmt.Printf("day 16 part2 %v\n", result)
}

func hexToBits(input string) string {
	var s strings.Builder

	for _, r := range input {
		s.WriteString(toBits[r])
	}

	return s.String()
}

func parse(s string, maxPackets int) ([]packet, int) {
	i := 0
	result := make([]packet, 0)
	for i < len(s)-8 && len(result) < maxPackets {
		current := s[i:]
		packet := packet{
			version: toint(current[:3]),
			typeId:  toint(current[3:6]),
		}

		switch packet.typeId {
		case 4:
			parseLiteral(current[6:], &packet)
		default:
			parseOperator(current[6:], &packet)
		}
		i += 6 + packet.totalLength
		result = append(result, packet)
	}
	return result, i
}

func parseOperator(s string, packet *packet) {
	if s[0] == '0' {
		length := toint(s[1:16])
		packet.totalLength = length + 16
		packet.subpackets, _ = parse(s[16:packet.totalLength], math.MaxInt)
	} else {
		length := toint(s[1:12])
		packet.subpackets, packet.totalLength = parse(s[12:], length)
		packet.totalLength += 12
	}
}

func parseLiteral(s string, packet *packet) {
	packet.literal = 0
	i := 0
	for {
		if s[i] == '0' {
			packet.literal = packet.literal*16 + toint(s[i+1:i+5])
			packet.totalLength = i + 5
			break
		} else {
			packet.literal = packet.literal*16 + toint(s[i+1:i+5])
		}
		i += 5
	}
}

func toint(s string) int {
	result, err := strconv.ParseInt(s, 2, 64)
	if err != nil {
		log.Fatal(err)
	}
	return int(result)
}

type packet struct {
	version     int
	typeId      int
	totalLength int
	literal     int
	subpackets  []packet
}

func (p packet) versionSum() int {
	sum := 0
	for _, s := range p.subpackets {
		sum += s.versionSum()
	}
	return p.version + sum
}

func (p packet) eval() int {
	result := 0
	switch p.typeId {
	case 0:
		for _, s := range p.subpackets {
			result += s.eval()
		}
	case 1:
		result = 1
		for _, s := range p.subpackets {
			result *= s.eval()
		}
	case 2:
		result = math.MaxInt
		for _, s := range p.subpackets {
			e := s.eval()
			if e < result {
				result = e
			}
		}
	case 3:
		result = 0
		for _, s := range p.subpackets {
			e := s.eval()
			if e > result {
				result = e
			}
		}
	case 4:
		result = p.literal
	case 5:
		if p.subpackets[0].eval() > p.subpackets[1].eval() {
			result = 1
		}
	case 6:
		if p.subpackets[0].eval() < p.subpackets[1].eval() {
			result = 1
		}
	case 7:
		if p.subpackets[0].eval() == p.subpackets[1].eval() {
			result = 1
		}
	default:
		log.Fatal("uh")
	}
	return result
}
