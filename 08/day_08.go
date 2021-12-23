package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"sort"
	"strings"
)

type Entry struct {
	Patterns []string
	Output   []string
}

func ParseLine(l string) Entry {
	l_arr := strings.Split(l, "|")
	return Entry{
		strings.Fields(l_arr[0]),
		strings.Fields(l_arr[1]),
	}
}

func CountEasyDigits(e Entry) int {
	var count int
	for _, output := range e.Output {
		if len(output) == 2 || len(output) == 4 || len(output) == 3 || len(output) == 7 {
			count++
		}
	}
	return count
}

func ReadInput(filename string) []Entry {
	file, err := os.Open(filename)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	var entries []Entry
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		entries = append(entries, ParseLine(scanner.Text()))
	}
	return entries
}

func overlaps(a string, b string) int {
	m := make(map[rune]bool)

	for _, r := range a {
		m[r] = true
	}

	var c int

	for _, r := range b {
		_, found := m[r]
		if found {
			c++
		}
	}
	return c
}

func sortString(s string) string {
	ss := strings.Split(s, "")
	sort.Strings(ss)
	return strings.Join(ss, "")
}

func Decode(e Entry) int {
	encoding := make(map[int]string)

	// easy patterns
	ps := make(map[string]bool)
	for _, p := range e.Patterns {
		switch len(p) {
		case 2:
			encoding[1] = sortString(p)
		case 4:
			encoding[4] = sortString(p)
		case 3:
			encoding[7] = sortString(p)
		case 7:
			encoding[8] = sortString(p)
		default:
			ps[p] = true
		}
	}

	// remaining patterns
	for p := range ps {
		switch {
		case len(p) == 5 && overlaps(p, encoding[1]) == 2:
			encoding[3] = sortString(p)
		case len(p) == 5 && overlaps(p, encoding[4]) == 3:
			encoding[5] = sortString(p)
		case len(p) == 5:
			encoding[2] = sortString(p)
		case len(p) == 6 && overlaps(p, encoding[4]) == 4:
			encoding[9] = sortString(p)
		case len(p) == 6 && overlaps(p, encoding[7]) == 3:
			encoding[0] = sortString(p)
		case len(p) == 6:
			encoding[6] = sortString(p)
		}
	}

	decoding := make(map[string]int)
	for digit, pattern := range encoding {
		decoding[pattern] = digit
	}

	x := make([]int, 0, 4)
	for _, o := range e.Output {
		x = append(x, decoding[sortString(o)])
	}

	return x[0]*1000 + x[1]*100 + x[2]*10 + x[3]
}

func Part1(entries []Entry) {
	var total int
	for _, entry := range entries {
		total += CountEasyDigits(entry)
	}
	fmt.Printf("Part 1: %d\n", total)
}

func Part2(entries []Entry) {
	var total int
	for _, entry := range entries {
		total += Decode(entry)
	}
	fmt.Printf("Part 2: %d\n", total)
}

func main() {
	entries := ReadInput(os.Args[1])

	Part1(entries)
	Part2(entries)
}
