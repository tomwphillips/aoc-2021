package main

import "testing"

func TestComputeGammaRate(t *testing.T) {
	tests := []struct {
		in   []string
		want string
	}{
		{[]string{
			"00100",
			"11110",
			"10110",
			"10111",
			"10101",
			"01111",
			"00111",
			"11100",
			"10000",
			"11001",
			"00010",
			"01010",
		}, "10110"},
		{[]string{
			"100",
			"100",
			"110",
		}, "100"},
	}

	for _, test := range tests {
		got := computeGammaRate(test.in)
		if got != test.want {
			t.Errorf("computeGrammaRate = %v, want %v", got, test.want)
		}
	}

}

func TestComputeEpsilonRate(t *testing.T) {
	in := "10110"
	got := computeEpsilonRate(in)
	want := "01001"
	if got != want {
		t.Errorf("computeGrammaRate = %v, want %v", got, want)
	}
}

func TestBinaryToInt(t *testing.T) {
	tests := []struct {
		in   string
		want int
	}{
		{"10110", 22},
		{"01001", 9},
	}

	for _, test := range tests {
		got := binaryToInt(test.in)
		if got != test.want {
			t.Errorf("binaryToInt(%v) = %d, want %d", test.in, got, test.want)
		}
	}
}
