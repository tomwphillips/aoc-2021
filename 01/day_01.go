package main

import (
	"bufio"
	"fmt"
	"io"
	"log"
	"os"
	"strconv"
)

func ReadMeasurements(reader io.Reader) []int {
	scanner := bufio.NewScanner(reader)
	var ms []int
	for scanner.Scan() {
		m, err := strconv.Atoi(scanner.Text())
		if err != nil {
			log.Fatal(err)
		}
		ms = append(ms, m)
	}
	return ms
}

func Part1(ms []int) int {
	var increases int
	for i := 1; i < len(ms); i++ {
		if ms[i] > ms[i-1] {
			increases++
		}
	}
	return increases
}

func Part2(ms []int) int {
	var increases int
	for i := 1; i < len(ms)-2; i++ {
		if ms[i]+ms[i+1]+ms[i+2] > ms[i-1]+ms[i]+ms[i+1] {
			increases++
		}
	}
	return increases
}

func main() {
	ms := ReadMeasurements(os.Stdin)
	fmt.Printf("Part 1: %d\n", Part1(ms))
	fmt.Printf("Part 2: %d\n", Part2(ms))
}
