package problems

func EvenFiboNumTotals() int {
	result := 0
	fibonacciNumFirst := 1
	fibonacciNumSecond := 2
	for fibonacciNumSecond < 4000000 {
		if fibonacciNumSecond%2 == 0 {
			result = result + fibonacciNumSecond

		}
		temp := fibonacciNumFirst
		fibonacciNumFirst = fibonacciNumSecond
		fibonacciNumSecond = fibonacciNumSecond + temp

	}
	return result
}
