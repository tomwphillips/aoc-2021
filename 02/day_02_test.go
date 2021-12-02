package main

import (
	"reflect"
	"testing"
)


func TestParseInstruction(t *testing.T) {
	var tests = []struct{
		i string
		p position
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
