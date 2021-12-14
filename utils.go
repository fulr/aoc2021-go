package aoc2021

import (
	"log"
	"os"
	"strconv"
	"strings"
)

func readFile(file string) []string {
	f, err := os.ReadFile(file)
	if err != nil {
		log.Fatal(err)
	}
	s := string(f)

	return strings.Split(strings.ReplaceAll(s, "\r", ""), "\n")
}

func readInt(file string) []int {
	lines := readFile(file)
	nrs := make([]int, len(lines))
	for i, v := range lines {
		a, err := strconv.Atoi(v)
		if err != nil {
			log.Fatal("no i", err)
		}
		nrs[i] = a
	}
	return nrs
}
