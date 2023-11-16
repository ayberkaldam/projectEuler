package problems

func FindSmallestMultipy() int {
	num := 20
	var result bool = IsMultiplyZeroToTwenty(num)
	for !result {
		num++
		result = IsMultiplyZeroToTwenty(num)
	}
	return num
}

func IsMultiplyZeroToTwenty(number int) bool {
	var result bool = true
	for i := 20; i > 0; i-- {
		if number%i != 0 {
			result = false
			break
		}
	}
	return result
}
