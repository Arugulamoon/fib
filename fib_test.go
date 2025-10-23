package fib_test

import (
	"slices"
	"testing"

	"github.com/Arugulamoon/fib"
)

/*
1) Applied age-old technique of stepping away from computer.
Lo and behold, mind remained on the problem and found a path forward.
Back at the computer, re-read the Wiki paragraph below, and set some
constants/assumptions (starting sequence is {0, 1} and the example sequence):

"The Fibonacci sequence is a sequence in which each element
is the sum of the two elements that precede it.
Numbers that are part of the Fibonacci sequence are known as Fibonacci numbers.
Many writers begin the sequence with 0 and 1,
although some authors start it from 1 and 1 and some (as did Fibonacci) from 1 and 2.
Starting from 0 and 1, the sequence begins:
0, 1, 1, 2, 3, 5, 8, 13, 21, 34, 55, 89, 144, ..."
*/

/*
4) Wrote Table-Driven Test to validate, and in so doing,
realized/saw that `seq` with `want` appended, is the sequence growing.
Need to wire in `max` now to know when to stop.
*/
func TestNextFibNum(t *testing.T) {
	tests := []struct {
		seq  []int
		want int
	}{
		{seq: []int{0, 1}, want: 1},
		{seq: []int{0, 1, 1}, want: 2},
		{seq: []int{0, 1, 1, 2}, want: 3},
		{seq: []int{0, 1, 1, 2, 3}, want: 5},
		{seq: []int{0, 1, 1, 2, 3, 5}, want: 8},
	}

	for _, tt := range tests {
		got := fib.NextFibNum(tt.seq)
		if got != tt.want {
			t.Errorf("with %#v, got %d; want %d", tt.seq, got, tt.want)
		}
	}
}

// 6) And then added tests to validate, and it looks close.
// Needs the recursion now.
func TestFibSeq(t *testing.T) {
	tests := []struct {
		seq  []int
		want []int
	}{
		{seq: []int{0, 1}, want: []int{0, 1, 1}},
		{seq: []int{0, 1, 1}, want: []int{0, 1, 1, 2}},
		{seq: []int{0, 1, 1, 2}, want: []int{0, 1, 1, 2, 3}},
		{seq: []int{0, 1, 1, 2, 3}, want: []int{0, 1, 1, 2, 3, 5}},
		{seq: []int{0, 1, 1, 2, 3, 5}, want: []int{0, 1, 1, 2, 3, 5, 8}},
	}

	for _, tt := range tests {
		got := fib.FibSeq(tt.seq)
		if !slices.Equal(got, tt.want) {
			t.Errorf("with %#v, got %d; want %d", tt.seq, got, tt.want)
		}
	}
}

// 7) Adding tests to drive towards adding recursion
// by removing `seq` from 6) TestFibSeq test inputs
func TestFib(t *testing.T) {
	startingSequence := []int{0, 1}

	tests := []struct {
		seq  []int
		num  int
		want []int
	}{
		{num: 1, want: []int{0, 1, 1}},
		{num: 2, want: []int{0, 1, 1, 2}},
		{num: 3, want: []int{0, 1, 1, 2, 3}},
		{num: 5, want: []int{0, 1, 1, 2, 3, 5}},
		{num: 8, want: []int{0, 1, 1, 2, 3, 5, 8}},
		// 10) Adding test for Non-Fibonacci Number (next is 13, not 9)
		// assuming should return up to 8
		{num: 9, want: []int{0, 1, 1, 2, 3, 5, 8}},
		// 11) And for good measure, adding in a couple more tests
		{num: 55, want: []int{0, 1, 1, 2, 3, 5, 8, 13, 21, 34, 55}},
		{num: 144, want: []int{0, 1, 1, 2, 3, 5, 8, 13, 21, 34, 55, 89, 144}},
		// 12) And special cases
	}

	for _, tt := range tests {
		got := fib.Fib(startingSequence, tt.num)
		if !slices.Equal(got, tt.want) {
			t.Errorf("with %d, got %d; want %d", tt.num, got, tt.want)
		}
	}
}

// 12) And special cases, unsure if even valid since starting sequence
// said to be {0, 1} and num > 1 (assumptions)
// So, the `want` values are assumptions.
func TestFibSpecialCases(t *testing.T) {
	tests := []struct {
		seq              []int
		num              int
		startingSequence []int
		want             []int
	}{
		{num: 0, startingSequence: []int{0}, want: []int{0}},
		{num: -1, startingSequence: []int{}, want: nil},
	}

	for _, tt := range tests {
		got := fib.Fib(tt.startingSequence, tt.num)
		if !slices.Equal(got, tt.want) {
			t.Errorf("with %d, got %d; want %d", tt.num, got, tt.want)
		}
	}
}
