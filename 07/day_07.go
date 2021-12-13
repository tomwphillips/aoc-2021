package main

import (
	"fmt"
	"log"
	"os"
	"sort"
	"strconv"
	"strings"
)

type Crab struct {
	Position int
}

func ParseInput(input []byte) []Crab {
	positions := strings.Split(strings.Trim(string(input), "\n"), ",")
	crabs := make([]Crab, len(positions))
	for i := range positions {
		p, err := strconv.Atoi(positions[i])
		if err != nil {
			log.Fatal(err)
		}
		crabs[i] = Crab{p}
	}
	return crabs
}

// Align returns fuel required to align crabs at position
func Align(crabs []Crab, pos int) int {
	var total_fuel int
	for i := range crabs {
		fuel := crabs[i].Position - pos
		if fuel < 0 {
			fuel = fuel * -1
		}
		total_fuel += fuel
	}
	return total_fuel
}

func MostFuelEfficientPosition(crabs []Crab) int {
	pos := make([]int, len(crabs))
	for i := range crabs {
		pos[i] = crabs[i].Position
	}
	return median(pos)
}

func median(x []int) int {
	sort.Ints(x)
	if len(x)%2 == 0 {
		return x[len(x)/2]
	} else {
		return x[len(x)/2+1]
	}
}

func Part1(crabs []Crab) {
	pos := MostFuelEfficientPosition(crabs)
	fuel := Align(crabs, pos)
	fmt.Printf("Fuel at position %d: %d\n", pos, fuel)
}

func main() {
	input, err := os.ReadFile(os.Args[1])
	if err != nil {
		log.Fatalf("Error reading file: %v", err)
	}

	crabs := ParseInput(input)
	Part1(crabs)
}
