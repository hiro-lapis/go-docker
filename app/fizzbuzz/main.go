package fizzbuzz

func Fizzbuzz(i int) string {
	if i%5 == 0 {
		return "Buzz"
	}
	if i%15 == 0 {
		return "FizzBuzz"
	}
	return "Fizz"
}
