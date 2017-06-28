package main

import (
	"fmt"
	"math"
)

func makeRange(min, max int) []int {
	a := make([]int, max-min+1)
	for i := range a {
		a[i] = min + i
	}
	return a
}

func sumSquares(numbers []int) float64 {
	var sum float64
	for _, n := range numbers {
		sum += math.Pow(float64(n), 2.0)
	}
	return sum
}

func squareSum(numbers []int) float64 {
	var sum float64
	for _, n := range numbers {
		sum += float64(n)
	}
	return math.Pow(sum, 2.0)
}

func squareDifference(numbers []int) float64 {
	return squareSum(numbers) - sumSquares(numbers)
}

func main() {
	input := makeRange(1, 100)
	fmt.Println(squareDifference(input))
}

/*
PROBLEM 6

The sum of the squares of the first ten natural numbers is,

12 + 22 + ... + 102 = 385

The square of the sum of the first ten natural numbers is,

(1 + 2 + ... + 10)2 = 552 = 3025

Hence the difference between the sum of the squares of the first ten natural numbers
and the square of the sum is 3025 − 385 = 2640.

Find the difference between the sum of the squares
of the first one hundred natural numbers and the square of the sum.

*/
