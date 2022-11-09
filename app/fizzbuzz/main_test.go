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
	// 7の時は7になるか
}
