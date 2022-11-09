package fizzbuzz

import "testing"

func TestFizzbuzz(t *testing.T) {
	// 3の時はFizzになるか
	v := Fizzbuzz(3)
	if v != "Fizz" {
		t.Error("Expected Fizz", v)
	}

	// 5の時はBuzzになるか
	v = Fizzbuzz(5)
	if v != "Buzz" {
		t.Error("Expected Buzz", v)
	}
	// 15の時はFizzBuzzになるか
	v = Fizzbuzz(15)
	if v != "FizzBuzz" {
		t.Error("Expected FizzBuzz", v)
	}

	// 7の時は7になるか
	v = Fizzbuzz(7)
	if v != "7" {
		t.Error("Expected 7", v)
	}
}
