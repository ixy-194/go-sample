package fizzbuzz_test

import (
	"go-sample/fizzbuzz"
	"testing"
)

func TestFizzBuzz(t *testing.T) {
	tests := []struct {
		input    int
		expected string
	}{
		{1, "1"},
		{3, "Fizz"},
		{5, "Buzz"},
		{15, "FizzBuzz"},
		{7, "7"},
	}

	for _, test := range tests {
		result := fizzbuzz.FizzBuzz(test.input)

		if result != test.expected {
			t.Errorf("FizzBuzz(%d) = %s; want %s", test.input, result, test.expected)
		}
	}
}
