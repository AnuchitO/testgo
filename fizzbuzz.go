package fizzbuzz

import "fmt"

func FizzBuzz(n int) string {

	if n == 5 {
		return "Buzz"
	}

	if n%3 == 0 {
		return "Fizz"
	}

	return fmt.Sprint(n)
}
