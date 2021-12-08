package main

import (
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

func ParseInput(in []byte) []int {
	s := strings.Split(strings.Trim(string(in), "\n"), ",")
	x := make([]int, len(s))
	var err error
	for i := range s {
		x[i], err = strconv.Atoi(s[i])
		if err != nil {
			log.Fatal(err)
		}
	}
	return x
}

func StepFish(timer int) (int, bool) {
	var new_timer int
	if timer > 0 {
		new_timer = timer - 1
	} else {
		new_timer = 6
	}
	return new_timer, timer == 0
}

func StepFishes(timers []int) []int {
	var reproduced bool
	var new_fishes int
	for i := range timers {
		timers[i], reproduced = StepFish(timers[i])
		if reproduced {
			new_fishes++
		}
	}

	new_timers := make([]int, new_fishes)
	for i := range new_timers {
		new_timers[i] = 8
	}
	timers = append(timers, new_timers...)

	return timers
}

func main() {
	input, err := os.ReadFile(os.Args[1])
	if err != nil {
		log.Fatalf("Error reading file: %v", err)
	}

	timers := ParseInput(input)

	for i := 0; i < 80; i++ {
		timers = StepFishes(timers)
	}

	fmt.Printf("Number of fish: %d\n", len(timers))
}
