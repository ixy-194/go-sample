package fizzbuzz_test

import (
	"go-sample/fizzbuzz"
	"testing"
)

// テストの定義
func Test_1を入力したら1が返ってくること(t *testing.T) {
	result := fizzbuzz.FizzBuzz(1)
	if result != "1" {
		t.Errorf("FizzBuzz(1) = %s; want 1", result)
	}
}

func Test_3を入力したらFizzが返ってくること(t *testing.T) {
	result := fizzbuzz.FizzBuzz(3)
	if result != "Fizz" {
		t.Errorf("FizzBuzz(3) = %s; want Fizz", result)
	}
}

func Test_5を入力したらBuzzが返ってくること(t *testing.T) {
	result := fizzbuzz.FizzBuzz(5)
	if result != "Buzz" {
		t.Errorf("FizzBuzz(5) = %s; want Buzz", result)
	}
}

func Test_15を入力したらFizzBuzzが返ってくること(t *testing.T) {
	result := fizzbuzz.FizzBuzz(15)
	if result != "FizzBuzz" {
		t.Errorf("FizzBuzz(15) = %s; want FizzBuzz", result)
	}
}
