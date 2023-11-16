package problems

func MultipleThreeOrFive() int {
	result := 0
	for i := 1; i < 1000; i++ {
		if i%3 == 0 || i%5 == 0 {
			result = result + i
		}
	}
	return result
}
