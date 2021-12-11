package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

type position struct {
	horizontal_position int
	depth               int
}

func parseInstruction(i string, p *position) {
	s := strings.Split(i, " ")
	direction := s[0]
	magnitude, err := strconv.Atoi(s[1])
	if err != nil {
		log.Fatal(err)
	}
	switch direction {
	case "forward":
		p.horizontal_position += magnitude
	case "down":
		p.depth += magnitude
	case "up":
		p.depth -= magnitude
	}
}

func Part1(instructions []string) int {
	p := position{horizontal_position: 0, depth: 0}
	for i := range instructions {
		parseInstruction(instructions[i], &p)
	}
	return p.horizontal_position * p.depth
}

type Part2Position struct {
	horizontal int
	depth      int
	aim        int
}

func CalculatePart2Position(instructions []string) Part2Position {
	p := Part2Position{0, 0, 0}
	for _, instruction := range instructions {
		s := strings.Split(instruction, " ")
		direction := s[0]
		magnitude, err := strconv.Atoi(s[1])
		if err != nil {
			log.Fatal(err)
		}
		switch direction {
		case "forward":
			p.horizontal += magnitude
			p.depth += p.aim * magnitude
		case "down":
			p.aim += magnitude
		case "up":
			p.aim -= magnitude
		}
	}
	return p
}

func Part2(instructions []string) int {
	p := CalculatePart2Position(instructions)
	return p.horizontal * p.depth
}

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	var instructions []string
	for scanner.Scan() {
		instructions = append(instructions, scanner.Text())
	}
	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}

	fmt.Printf("Part 1: %d\n", Part1(instructions))
	fmt.Printf("Part 2: %d\n", Part2(instructions))
}
