package main

import (
	"bufio"
	"fmt"
	"io"
	"log"
	"os"
	"strconv"
	"strings"
)

type Line struct {
	x_start int
	y_start int
	x_end   int
	y_end   int
}

func NewLine(s string) Line {
	coordinates := strings.Split(s, " -> ")
	start := strings.Split(coordinates[0], ",")
	end := strings.Split(coordinates[1], ",")

	x_start, err := strconv.Atoi(start[0])
	if err != nil {
		log.Fatal(err)
	}

	y_start, err := strconv.Atoi(start[1])
	if err != nil {
		log.Fatal(err)
	}

	x_end, err := strconv.Atoi(end[0])
	if err != nil {
		log.Fatal(err)
	}

	y_end, err := strconv.Atoi(end[1])
	if err != nil {
		log.Fatal(err)
	}

	return Line{x_start, y_start, x_end, y_end}
}

func IsVerticalOrHorizontal(l Line) bool {
	return l.x_start == l.x_end || l.y_start == l.y_end
}

func NewGrid(s int) *[][]int {
	g := make([][]int, s)
	for i := range g {
		g[i] = make([]int, s)
	}
	return &g
}

func MarkPoints(g *[][]int, l Line) {
	var x_start int
	var x_end int
	var y_start int
	var y_end int

	if l.x_start <= l.x_end {
		x_start = l.x_start
		x_end = l.x_end
	} else {
		x_start = l.x_end
		x_end = l.x_start
	}

	if l.y_start <= l.y_end {
		y_start = l.y_start
		y_end = l.y_end
	} else {
		y_start = l.y_end
		y_end = l.y_start
	}

	for i := x_start; i <= x_end; i++ {
		for j := y_start; j <= y_end; j++ {
			(*g)[j][i] += 1
		}
	}
}

func ReadInput(reader io.Reader) []Line {
	scanner := bufio.NewScanner(reader)
	var ls []Line
	for scanner.Scan() {
		l := NewLine(scanner.Text())
		ls = append(ls, l)
	}
	return ls
}

func CountPointsAtLeast(g *[][]int, n int) int {
	var count int
	for _, row := range *g {
		for i := range row {
			if row[i] >= n {
				count += 1
			}
		}
	}
	return count
}

func GetSize(ls []Line) int {
	var size int
	for _, l := range ls {
		if l.x_start > size {
			size = l.x_start
		}
		if l.y_start > size {
			size = l.y_start
		}
		if l.x_end > size {
			size = l.x_end
		}
		if l.y_end > size {
			size = l.y_end
		}
	}
	return size + 1
}

func main() {
	ls := ReadInput(os.Stdin)
	size := GetSize(ls)
	g := NewGrid(size)

	for _, l := range ls {
		if IsVerticalOrHorizontal(l) {
			MarkPoints(g, l)
		}
	}
	count := CountPointsAtLeast(g, 2)
	fmt.Println(count)
}
