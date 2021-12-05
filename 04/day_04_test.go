package main

import (
	"bytes"
	"reflect"
	"testing"
)

func TestNewBoard(t *testing.T) {
	in := []string{
		" 1  2  3",
		" 4  5  6",
		" 7  8  9",
	}
	got := NewBoard(in)
	want := map[int]Position{
		1: {0, 0},
		2: {0, 1},
		3: {0, 2},
		4: {1, 0},
		5: {1, 1},
		6: {1, 2},
		7: {2, 0},
		8: {2, 1},
		9: {2, 2},
	}
	if !reflect.DeepEqual(got.numbers, want) {
		t.Errorf("NewBoard(%v) = %v, want %v", in, got, want)
	}

	if got.size != 3 {
		t.Errorf("Got size %d, not 3", got.size)
	}
}

func TestCheckBoard(t *testing.T) {
	b := NewBoard([]string{
		"1 2 3",
		"4 5 6",
		"7 8 9",
	})

	ns := []int{1, 5, 9, 7}

	for _, n := range ns {
		if CheckBoard(b, n) {
			t.Fatalf("CheckBoard(b, %d) returned true, expected false", n)
		}
	}

	if !CheckBoard(b, 4) {
		t.Fatal("CheckBoard(b, 4) returned false, expected true")
	}
}

func TestComputeScore(t *testing.T) {
	b := NewBoard([]string{
		"14 21 17 24  4",
		"10 16 15  9 19",
		"18  8 23 26 20",
		"22 11 13  6  5",
		" 2  0 12  3  7",
	})
	ns := []int{7, 4, 9, 5, 11, 17, 23, 2, 0, 14, 21, 24}
	got := ComputeScore(b, ns)
	want := 4512
	if got != want {
		t.Errorf("ComputeScore = %d, want %d", got, want)
	}
}

func TestReadInput(t *testing.T) {
	tests := []struct {
		in     []string
		ns     []int
		num_bs int
		size   int
	}{
		{[]string{
			"10,20,30,40",
			"",
			" 1  2",
			" 3  4",
			"",
			"10 20",
			"30 40",
		},
			[]int{10, 20, 30, 40},
			2,
			2,
		},
		{[]string{
			"1,2,3,4,5",
			"",
			"10 20 30",
			"40 50 60",
			"70 80 90",
			"",
			"01 02 03",
			"04 05 06",
			"07 08 09",
			"",
			"30 40 50",
			"60 70 80",
			"90 10 20",
		},
			[]int{1, 2, 3, 4, 5},
			3,
			3,
		},
	}

	for _, test := range tests {
		var buffer bytes.Buffer
		for _, line := range test.in {
			buffer.WriteString(line + "\n")
		}

		ns, bs := ReadInput(&buffer)

		if !reflect.DeepEqual(ns, test.ns) {
			t.Errorf("ReadInput ns = %v, want %v", ns, test.ns)
		}

		if len(bs) != test.num_bs {
			t.Errorf("Got %d boards, want %d", len(bs), test.num_bs)
		}

		for _, b := range bs {
			if b.size != test.size {
				t.Errorf("Size %d, want %d", b.size, test.size)
			}
		}

	}

}
