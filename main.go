package main

import (
	"fmt"
	problems "projectEuler/problems"
)

func main() {
	problem1Result := problems.MultipleThreeOrFive()
	fmt.Println(problem1Result)
	problem2Result := problems.EvenFiboNumTotals()
	fmt.Println(problem2Result)
	problem3Result := problems.FindLargestPrimeNumber(600851475143)
	fmt.Println(problem3Result)
	problem4Result := problems.FindLargestPalindrome()
	fmt.Println(problem4Result)
	problem5Result := problems.FindSmallestMultipy()
	fmt.Println(problem5Result)
}
