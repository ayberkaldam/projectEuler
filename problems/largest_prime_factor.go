package problems

import "math"

func FindLargestPrimeNumber(number int) int {
	largest := 0
	var counter int = 2
	for counter <= int(number) {
		if number%counter == 0 {
			if counter > largest {
				largest = counter
			}
			number = number / counter
		} else {
			counter++
			for !IsPrime(int64(counter)) {
				counter++
			}
		}
	}
	return largest
}

func IsPrime(number int64) bool {
	if number < 2 {
		return false
	}
	maxDivisor := int64(math.Sqrt(float64(number)))
	for i := int64(2); i <= maxDivisor; i++ {
		if number%i == 0 {
			return false
		}
	}
	return true
}
