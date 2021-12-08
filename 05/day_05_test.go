package main

import (
	"bytes"
	"reflect"
	"testing"
)

func TestNewLines(t *testing.T) {
	tests := []struct {
		in   string
		want Line
	}{
		{"0,9 -> 5,9", Line{0, 9, 5, 9}},
		{"593,10 -> 593,98", Line{593, 10, 593, 98}},
	}

	for _, test := range tests {
		got := NewLine(test.in)
		if got != test.want {
			t.Errorf("NewLine(%v) = %v, want %v", test.in, got, test.want)
		}
	}
}

func TestIsVerticalOrHorizontal(t *testing.T) {
	tests := []struct {
		in   Line
		want bool
	}{
		{Line{0, 9, 5, 9}, true},
		{Line{9, 4, 3, 4}, true},
		{Line{6, 4, 2, 0}, false},
	}
	for _, test := range tests {
		if got := IsVerticalOrHorizontal(test.in); got != test.want {
			t.Errorf("IsVerticalOrHorizontal(%v) = %t, want %t", test.in, got, test.want)
		}
	}
}

func TestNewGrid(t *testing.T) {
	sizes := []int{10, 100}
	for _, size := range sizes {
		g := NewGrid(size)
		if len(*g) != size {
			t.Fatalf("len(g) = %d, want %d", len(*g), size)
		}
		for i, row := range *g {
			if len(row) != size {
				t.Errorf("len(g[%d]) = %d, want %d", i, len(row), size)
			}
		}
	}
}

func TestMarkPoints(t *testing.T) {
	ls := []Line{
		{0, 9, 5, 9},
		{9, 4, 3, 4},
		{2, 2, 2, 1},
		{7, 0, 7, 4},
		{0, 9, 2, 9},
		{3, 4, 1, 4},
	}
	g := NewGrid(10)
	want := [][]int{
		{0, 0, 0, 0, 0, 0, 0, 1, 0, 0},
		{0, 0, 1, 0, 0, 0, 0, 1, 0, 0},
		{0, 0, 1, 0, 0, 0, 0, 1, 0, 0},
		{0, 0, 0, 0, 0, 0, 0, 1, 0, 0},
		{0, 1, 1, 2, 1, 1, 1, 2, 1, 1},
		{0, 0, 0, 0, 0, 0, 0, 0, 0, 0},
		{0, 0, 0, 0, 0, 0, 0, 0, 0, 0},
		{0, 0, 0, 0, 0, 0, 0, 0, 0, 0},
		{0, 0, 0, 0, 0, 0, 0, 0, 0, 0},
		{2, 2, 2, 1, 1, 1, 0, 0, 0, 0},
	}
	for _, l := range ls {
		MarkPoints(g, l)
	}
	for i := range *g {
		if got := (*g)[i]; !reflect.DeepEqual(got, want[i]) {
			t.Errorf("g[%d] = %v, want %v", i, got, want[i])
		}
	}

}

func TestReadInput(t *testing.T) {
	in := []string{"10,20 -> 30,40", "1,2 -> 3,4"}
	want := []Line{{10, 20, 30, 40}, {1, 2, 3, 4}}

	var buffer bytes.Buffer
	for _, line := range in {
		buffer.WriteString(line + "\n")
	}

	got := ReadInput(&buffer)
	if !reflect.DeepEqual(got, want) {
		t.Errorf("ReadInput(%v) = %v, want %v", in, got, want)
	}
}

func TestCountPointsAtLeast(t *testing.T) {
	in := make([][]int, 3)
	in[0] = []int{1, 0, 0}
	in[1] = []int{2, 0, 0}
	in[2] = []int{2, 0, 0}

	want := 2
	got := CountPointsAtLeast(&in, 2)
	if got != want {
		t.Errorf("CountPointsAtLeast(%v, 2) = %v, want %v", in, got, want)
	}
}

func TestGetSize(t *testing.T) {
	in := []Line{
		{1, 2, 3, 4},
		{0, 0, 1, 1},
		{10, 0, 0, 0},
	}
	want := 11
	got := GetSize(in)
	if got != want {
		t.Errorf("GetSize(in) = %d, want %d", got, want)
	}
}
