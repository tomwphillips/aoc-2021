package main

import (
	"reflect"
	"testing"
)

func TestParseInstruction(t *testing.T) {
	var tests = []struct {
		i    string
		p    position
		want position
	}{
		{"forward 1", position{0, 0}, position{1, 0}},
		{"forward 1", position{1, 0}, position{2, 0}},
		{"down 5", position{0, 0}, position{0, 5}},
		{"up 3", position{0, 0}, position{0, -3}},
	}

	for _, test := range tests {
		parseInstruction(test.i, &test.p)
		if !reflect.DeepEqual(test.want, test.p) {
			t.Errorf("parseInstruction, got %v, want %v", test.p, test.want)
		}
	}
}

func TestCalculatePart2Position(t *testing.T) {
	tests := []struct {
		instruction string
		position    Part2Position
	}{
		{"forward 5", Part2Position{5, 0, 0}},
		{"down 5", Part2Position{5, 0, 5}},
		{"forward 8", Part2Position{13, 40, 5}},
		{"up 3", Part2Position{13, 40, 2}},
		{"down 8", Part2Position{13, 40, 10}},
		{"forward 2", Part2Position{15, 60, 10}},
	}
	for i := 1; i <= len(tests); i++ {
		in := make([]string, i)
		for j := range in {
			in[j] = tests[j].instruction
		}

		want := tests[i-1].position
		got := CalculatePart2Position(in)
		if want != got {
			t.Errorf("CalculatePart2Position(%v) = %v, want %v", in, want, got)
		}
	}
}
