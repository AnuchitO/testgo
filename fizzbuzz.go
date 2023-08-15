package fizzbuzz

import (
	"fmt"
)

func FizzBuzz(n int) string {
	if isFizzBuzz(n) {
		return "FizzBuzz"
	}

	if isBuzz(n) {
		return "Buzz"
	}

	if isFizz(n) {
		return "Fizz"
	}

	return fmt.Sprint(n)
}

func isFizzBuzz(n int) bool {
	return n%15 == 0
}

func isBuzz(n int) bool {
	return n%5 == 0
}

func isFizz(n int) bool {
	return n%3 == 0
}
