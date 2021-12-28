package main

import (
	"bytes"
	"reflect"
	"testing"
)

func TestParseLine(t *testing.T) {
	in := "12345"
	want := []int{1, 2, 3, 4, 5}
	got := ParseLine(in)
	if !reflect.DeepEqual(got, want) {
		t.Errorf("ParseLine(%v) = %v, want %v", in, got, want)
	}
}

func TestReadInput(t *testing.T) {
	in := []string{"123", "456"}
	want := [][]int{
		{1, 2, 3},
		{4, 5, 6},
	}

	var buffer bytes.Buffer
	for _, line := range in {
		buffer.WriteString(line + "\n")
	}

	got := ReadInput(&buffer)
	if !reflect.DeepEqual(*got, want) {
		t.Errorf("ReadInput(%v) = %v, want %v", in, got, want)
	}
}

func TestAdjacentPoints(t *testing.T) {
	tests := []struct {
		in   point
		want []point
	}{
		{point{0, 0}, []point{{-1, 0}, {0, -1}, {1, 0}, {0, 1}}},
	}
	for _, test := range tests {
		got := adjacentPoints(test.in)
		if !reflect.DeepEqual(got, test.want) {
			t.Errorf("adjacentPoints(%v) = %v, want %v", test.in, got, test.want)
		}
	}
}

func TestValidPoints(t *testing.T) {
	tests := []struct {
		ps     []point
		x_size int
		y_size int
		want   []point
	}{
		{[]point{{-1, 0}, {0, -1}, {1, 0}, {0, 1}}, 5, 5, []point{{1, 0}, {0, 1}}},
		{[]point{{3, 4}, {4, 3}, {5, 4}, {4, 5}}, 5, 5, []point{{3, 4}, {4, 3}}},
	}
	for _, test := range tests {
		got := validPoints(test.ps, test.x_size, test.y_size)
		if !reflect.DeepEqual(got, test.want) {
			t.Errorf("validPoints(%v, %d, %d) = %v, want %v", test.ps, test.x_size, test.y_size, got, test.want)
		}
	}
}
