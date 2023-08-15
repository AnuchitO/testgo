package fizzbuzz

import "testing"

func TestFizzBuzzShouldReturn1WhenInput1(t *testing.T) {
	input := 1

	want := FizzBuzz(input)

	if want != "1" {
		t.Errorf("want 1 but got %s", want)
	}
}
