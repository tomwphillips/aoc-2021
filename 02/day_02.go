package main

import (
	"bufio"
	"fmt"
	"log"
	"strings"
	"os"
	"strconv"
)

type position struct {
	horizontal_position int
	depth int
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

func main() {
	p := position{horizontal_position: 0, depth: 0}
	scanner := bufio.NewScanner(os.Stdin)
	for scanner.Scan() {
		parseInstruction(scanner.Text(), &p)
	}
	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}
	fmt.Println(p.horizontal_position * p.depth)
}
