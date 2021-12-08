package main

import (
	"bytes"
	"reflect"
	"testing"
)

func TestReadMeasurements(t *testing.T) {
	input := []string{
		"199",
		"200",
		"208",
		"210",
		"200",
		"207",
		"240",
		"269",
		"260",
		"263",
	}
	want := []int{199,
		200,
		208,
		210,
		200,
		207,
		240,
		269,
		260,
		263,
	}

	var buffer bytes.Buffer
	for _, line := range input {
		buffer.WriteString(line + "\n")
	}

	got := ReadMeasurements(&buffer)
	if !reflect.DeepEqual(got, want) {
		t.Errorf("ReadMeasurements(input) = %v, want %v", got, want)
	}
}

func TestPart1(t *testing.T) {
	in := []int{199,
		200,
		208,
		210,
		200,
		207,
		240,
		269,
		260,
		263,
	}
	got := Part1(in)
	want := 7
	if got != want {
		t.Errorf("Part1(in) = %d, want %d", got, want)
	}
}

func TestPart2(t *testing.T) {
	in := []int{199,
		200,
		208,
		210,
		200,
		207,
		240,
		269,
		260,
		263,
	}
	got := Part2(in)
	want := 5
	if got != want {
		t.Errorf("Part2(ms) = %d, want %d", got, want)
	}
}
