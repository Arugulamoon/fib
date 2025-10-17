package fib

const START_NUM = 2

/*
2) The insight from time away was that num1 and num2 were just the
last 2 elements of the slice, which then also meant it had to be passed
along to the next function, in the function, instead of being a struct with
method. Maybe it could still be done with it, but the path seemed clearer without.
So, deleted all that code.
*/

// 3) Added an explicit function to pop last and second last elements off slice.
func NextFibNum(seq []int) int {
	return seq[len(seq)-1] + seq[len(seq)-2]
}

// 5) Took `seq` and `want` from 4) TestNextFibNum to make this function.
// The `want` is just the value of `NextFibNum(seq)`.
func FibSeq(seq []int) []int {
	return append(seq, NextFibNum(seq))
}

// 8) And back to reimplementing Fib() by manually iterating...
// func Fib(num int) []int {
// 	startingSequence := []int{0, 1}
// 	first := FibSeq(startingSequence)
// 	if reachedNum(num, first) {
// 		return first
// 	} else {
// 		second := FibSeq(first)
// 		if reachedNum(num, second) {
// 			return second
// 		} else {
// 			third := FibSeq(second)
// 			return third
// 		}
// 	}
// }

// 9) And Recursion!
func Fib(seq []int, num int) []int {
	// 13) Added to handle out of bounds special cases
	if len(seq) == 0 {
		return nil
	} else if len(seq) < 2 {
		return seq
	}

	next := NextFibNum(seq)

	// num isn't a Fibonacci Number
	if next > num {
		return seq
	}

	newSeq := FibSeq(seq)

	// reached num
	if next == num {
		return newSeq
	}

	// have not reached num yet, keep going
	return Fib(newSeq, num)
}
