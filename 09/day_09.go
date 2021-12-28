package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"strings"
	"unicode/utf8"
)

type point struct {
	x int
	y int
}

func ParseLine(l string) []int {
	a := make([]int, 0, utf8.RuneCountInString(l))
	for _, r := range strings.Trim(l, "\n") {
		a = append(a, int(r)-'0')
	}
	return a
}

func ReadInput(reader io.Reader) *[][]int {
	scanner := bufio.NewScanner(reader)
	v := make([][]int, 0)
	for scanner.Scan() {
		v = append(v, ParseLine(scanner.Text()))
	}
	return &v
}

func adjacentPoints(p point) []point {
	return []point{
		{p.x - 1, p.y}, // left
		{p.x, p.y - 1}, // up
		{p.x + 1, p.y}, // right
		{p.x, p.y + 1}, // down
	}
}

func validPoints(ps []point, x_size int, y_size int) []point {
	vps := make([]point, 0)
	for _, p := range ps {
		if 0 <= p.x && p.x < x_size && 0 <= p.y && p.y < y_size {
			vps = append(vps, p)
		}
	}
	return vps
}

func SumRiskLevels(heightmap *[][]int) int {
	y_size := len(*heightmap)
	x_size := len((*heightmap)[0])

	var sum int
	for y := range *heightmap {
		for x := range (*heightmap)[y] {
			ps := adjacentPoints(point{x, y})
			ps = validPoints(ps, x_size, y_size)

			low_point := true
			for _, p := range ps {
				if (*heightmap)[y][x] >= (*heightmap)[p.y][p.x] {
					low_point = false
					break
				}
			}
			if low_point {
				sum += 1 + (*heightmap)[y][x]
			}
		}
	}
	return sum
}

func main() {
	heightmap := ReadInput(os.Stdin)
	fmt.Printf("Part 1: %d\n", SumRiskLevels(heightmap))
}
