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

type Position struct {
	row int
	col int
}

type Board struct {
	size       int
	numbers    map[int]Position
	row_counts []int
	col_counts []int
}

func NewBoard(s []string) *Board {
	b := Board{}
	b.numbers = make(map[int]Position)

	for i, row := range s {
		cols := strings.Fields(row)

		if i == 0 {
			b.size = len(cols)
		}

		for j, col := range cols {
			n, err := strconv.Atoi(col)
			if err != nil {
				log.Fatal(err)
			}
			b.numbers[n] = Position{i, j}
		}
	}

	b.row_counts = make([]int, b.size)
	b.col_counts = make([]int, b.size)

	return &b
}

func CheckBoard(b *Board, n int) bool {
	p, ok := b.numbers[n]
	if ok {
		b.row_counts[p.row] += 1
		b.col_counts[p.col] += 1

		for _, row_count := range b.row_counts {
			if row_count == b.size {
				return true
			}
		}

		for _, col_count := range b.col_counts {
			if col_count == b.size {
				return true
			}
		}
	}
	return false
}

func ComputeScore(b *Board, ns []int) int {
	for _, n := range ns {
		delete(b.numbers, n)
	}

	var sum int
	for n := range b.numbers {
		sum += n
	}

	return sum * ns[len(ns)-1]
}

func ReadInput(reader io.Reader) ([]int, []*Board) {
	scanner := bufio.NewScanner(reader)

	scanner.Scan()
	str_ns := strings.Split(scanner.Text(), ",")
	ns := make([]int, len(str_ns))
	for i, str_n := range str_ns {
		n, err := strconv.Atoi(str_n)
		if err != nil {
			log.Fatal(err)
		}
		ns[i] = n
	}

	scanner.Scan()

	var bs []*Board
	var lines []string
	for scanner.Scan() {
		if line := scanner.Text(); line != "" {
			lines = append(lines, line)
		} else {
			bs = append(bs, NewBoard(lines))
			lines = []string{}
		}
	}
	bs = append(bs, NewBoard(lines))
	return ns, bs
}

func main() {
	ns, bs := ReadInput(os.Stdin)
	for i, n := range ns {
		for j := range bs {
			if CheckBoard(bs[j], n) {
				score := ComputeScore(bs[j], ns[:i+1])
				fmt.Printf("Score: %d\n", score)
				os.Exit(0)
			}
		}
	}
}
