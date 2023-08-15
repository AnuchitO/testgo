package fizzbuzz

import "testing"

func TestFizzBuzzShouldReturn1WhenInput1(t *testing.T) {
	input := 1

	got := FizzBuzz(input)

	if got != "1" {
		t.Errorf("want 1 but got %s", got)
	}
}

func TestFizzBuzzShouldReturn2WhenInput2(t *testing.T) {
	input := 2

	got := FizzBuzz(input)

	if got != "2" {
		t.Errorf("want 2 but got %s", got)
	}
}

func TestFizzBuzzShouldReturn3WhenInputFizz(t *testing.T) {
	input := 3

	got := FizzBuzz(input)

	if got != "Fizz" {
		t.Errorf("want 3 but got %s", got)
	}
}
