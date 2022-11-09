package fizzbuzz

func Fizzbuzz(i int) string {
	if i%5 == 0 {
		return "Buzz"
	}
	return "Fizz"
}
