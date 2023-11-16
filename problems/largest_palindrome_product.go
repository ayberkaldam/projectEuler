package problems

import "fmt"

func FindLargestPalindrome() int {
	var result int
	for i := 999; i > 99; i-- {
		for j := 999; j > 99; j-- {
			multiply := i * j
			if fmt.Sprint(multiply) == Reverse(fmt.Sprint(multiply)) && multiply > result {
				result = multiply
			}
		}

	}
	return result
}

func Reverse(sentence string) string {
	rns := []rune(sentence)
	for i, j := 0, len(rns)-1; i < j; i, j = i+1, j-1 {
		rns[i], rns[j] = rns[j], rns[i]
	}
	return string(rns)
}
