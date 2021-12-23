package main

import (
	"reflect"
	"testing"
)

func TestParseLine(t *testing.T) {
	in := "be cfbegad cbdgef fgaecd cgeb fdcge agebfd fecdb fabcd edb |	fdgacbe cefdb cefbgd gcbe"
	got := ParseLine(in)
	want := Entry{
		[]string{"be", "cfbegad", "cbdgef", "fgaecd", "cgeb", "fdcge", "agebfd", "fecdb", "fabcd", "edb"},
		[]string{"fdgacbe", "cefdb", "cefbgd", "gcbe"},
	}
	if !reflect.DeepEqual(got, want) {
		t.Errorf("ParseLine(%s) = %v, want %v", in, got, want)
	}
}

func TestCountEasyDigits(t *testing.T) {
	tests := []struct {
		in   Entry
		want int
	}{
		{Entry{[]string{}, []string{"fdgacbe", "cefdb", "cefbgd", "gcbe"}}, 2},
		{Entry{[]string{}, []string{"fcgedb", "cgb", "dgebacf", "gc"}}, 3},
		{Entry{[]string{}, []string{"efabcd", "cedba", "gadfec", "cb"}}, 1},
	}
	for _, test := range tests {
		got := CountEasyDigits(test.in)
		if got != test.want {
			t.Errorf("CountEasyDigits(%v) = %d, want %d", test.in, got, test.want)
		}
	}
}

func TestOverlaps(t *testing.T) {
	tests := []struct {
		a    string
		b    string
		want int
	}{
		{"abc", "abc", 3},
		{"abc", "c", 1},
		{"abc", "d", 0},
	}
	for _, test := range tests {
		got := overlaps(test.a, test.b)
		if got != test.want {
			t.Errorf("overlaps(%v, %v) = %v, want %v", test.a, test.b, got, test.want)
		}
	}

}

func TestDecode(t *testing.T) {
	in := Entry{
		[]string{"acedgfb", "cdfbe", "gcdfa", "fbcad", "dab", "cefabd", "cdfgeb", "eafb", "cagedb", "ab"},
		[]string{"cdfeb", "fcadb", "cdfeb", "cdbaf"},
	}
	want := 5353
	got := Decode(in)
	if got != want {
		t.Errorf("Decode(%v) = %d, want %d", in, got, want)
	}
}
