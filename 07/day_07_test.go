package main

import (
	"reflect"
	"testing"
)

func TestParseInput(t *testing.T) {
	in := []byte("16,1,2\n")
	got := ParseInput(in)
	want := []Crab{
		{16}, {1}, {2},
	}
	if !reflect.DeepEqual(got, want) {
		t.Errorf("ParseInput(%v) = %v, want %v", in, got, want)
	}
}

func TestAlign(t *testing.T) {
	crabs := ParseInput([]byte("16,1,2,0,4,2,7,1,2,14"))
	tests := []struct {
		in   int
		want int
	}{
		{2, 37},
		{1, 41},
		{3, 39},
		{10, 71},
	}
	for _, test := range tests {
		got := Align(crabs, test.in)
		if got != test.want {
			t.Errorf("Align(%d) = %d, want %d", test.in, got, test.want)
		}
	}
}

func TestMostFuelEfficientPosition(t *testing.T) {
	in := ParseInput([]byte("16,1,2,0,4,2,7,1,2,14"))
	got := MostFuelEfficientPosition(in)
	want := 2
	if got != want {
		t.Errorf("MostFuelEfficientPosition(%v) = %d, want %d", in, got, want)
	}
}
