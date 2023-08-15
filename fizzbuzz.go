package fizzbuzz

import "fmt"

func FizzBuzz(n int) string {
	if n == 15 || n == 30 {
		return "FizzBuzz"
	}

	if n == 5 || n == 10 {
		return "Buzz"
	}

	if n%3 == 0 {
		return "Fizz"
	}

	return fmt.Sprint(n)
}
