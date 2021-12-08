package main

import (
	"reflect"
	"testing"
)

func TestParseInput(t *testing.T) {
	in := []byte("3,4,3,1,2\n")
	want := []int{3, 4, 3, 1, 2}
	got := ParseInput(in)
	if !reflect.DeepEqual(got, want) {
		t.Errorf("ParseInput(%v) = %v, want %v", in, got, want)
	}
}

func TestStepFish(t *testing.T) {
	tests := []struct {
		in              int
		want_state      int
		want_reproduced bool
	}{
		{3, 2, false},
		{2, 1, false},
		{1, 0, false},
		{0, 6, true},
		{6, 5, false},
	}

	for _, test := range tests {
		got_state, got_reproduced := StepFish(test.in)
		if got_state != test.want_state || got_reproduced != test.want_reproduced {
			t.Errorf("StepFish(%d) = (%d, %t), want (%d, %t)", test.in, got_state, got_reproduced, test.want_state, test.want_reproduced)
		}
	}
}

func TestStepFishes(t *testing.T) {
	states := [][]int{
		{3, 4, 3, 1, 2},
		{2, 3, 2, 0, 1},
		{1, 2, 1, 6, 0, 8},
		{0, 1, 0, 5, 6, 7, 8},
		{6, 0, 6, 4, 5, 6, 7, 8, 8},
		{5, 6, 5, 3, 4, 5, 6, 7, 7, 8},
		{4, 5, 4, 2, 3, 4, 5, 6, 6, 7},
		{3, 4, 3, 1, 2, 3, 4, 5, 5, 6},
		{2, 3, 2, 0, 1, 2, 3, 4, 4, 5},
		{1, 2, 1, 6, 0, 1, 2, 3, 3, 4, 8},
		{0, 1, 0, 5, 6, 0, 1, 2, 2, 3, 7, 8},
		{6, 0, 6, 4, 5, 6, 0, 1, 1, 2, 6, 7, 8, 8, 8},
		{5, 6, 5, 3, 4, 5, 6, 0, 0, 1, 5, 6, 7, 7, 7, 8, 8},
		{4, 5, 4, 2, 3, 4, 5, 6, 6, 0, 4, 5, 6, 6, 6, 7, 7, 8, 8},
		{3, 4, 3, 1, 2, 3, 4, 5, 5, 6, 3, 4, 5, 5, 5, 6, 6, 7, 7, 8},
		{2, 3, 2, 0, 1, 2, 3, 4, 4, 5, 2, 3, 4, 4, 4, 5, 5, 6, 6, 7},
		{1, 2, 1, 6, 0, 1, 2, 3, 3, 4, 1, 2, 3, 3, 3, 4, 4, 5, 5, 6, 8},
		{0, 1, 0, 5, 6, 0, 1, 2, 2, 3, 0, 1, 2, 2, 2, 3, 3, 4, 4, 5, 7, 8},
		{6, 0, 6, 4, 5, 6, 0, 1, 1, 2, 6, 0, 1, 1, 1, 2, 2, 3, 3, 4, 6, 7, 8, 8, 8, 8},
	}

	for i := range states[:len(states)-1] {
		in := states[i]
		want := states[i+1]
		got := StepFishes(in)
		if !reflect.DeepEqual(got, want) {
			t.Errorf("StepFishes(%v) = %v, want %v", in, got, want)
		}
	}
}
