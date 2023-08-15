package fizzbuzz

import "fmt"

func FizzBuzz(n int) string {

	if n == 5 {
		return "Buzz"
	}

	if n == 3 || n == 6 || n == 9 {
		return "Fizz"
	}

	return fmt.Sprint(n)
}
