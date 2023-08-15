package fizzbuzz

import (
	"fmt"
	"testing"
)

var cases = []struct {
	input int
	want  string
}{
	{1, "1"},
	{2, "2"},
	{3, "Fizz"},
	{4, "4"},
	{5, "Buzz"},
	{6, "Fizz"},
}

func TestFizzBuzz(t *testing.T) {

	for _, c := range cases {
		name := fmt.Sprintf("should return %d when input %s", c.input, c.want)
		t.Run(name, func(t *testing.T) {

			got := FizzBuzz(c.input)

			if got != c.want {
				t.Errorf("want %s but got %s", c.want, got)
			}
		})
	}
}
