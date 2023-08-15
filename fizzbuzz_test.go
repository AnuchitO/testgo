package fizzbuzz

import "testing"

func TestFizzBuzzShouldReturn1WhenInput1(t *testing.T) {
	input := 1
	want := "1"

	got := FizzBuzz(input)

	if got != want {
		t.Errorf("want %s but got %s", want, got)
	}
}

func TestFizzBuzzShouldReturn2WhenInput2(t *testing.T) {
	input := 2
	want := "2"

	got := FizzBuzz(input)

	if got != want {
		t.Errorf("want %s but got %s", want, got)
	}
}

func TestFizzBuzzShouldReturn3WhenInputFizz(t *testing.T) {
	input := 3
	want := "Fizz"

	got := FizzBuzz(input)

	if got != want {
		t.Errorf("want %s but got %s", want, got)
	}
}

func TestFizzBuzzShouldReturn4WhenInput4(t *testing.T) {
	input := 4
	want := "4"

	got := FizzBuzz(input)

	if got != want {
		t.Errorf("want %s but got %s", want, got)
	}
}
