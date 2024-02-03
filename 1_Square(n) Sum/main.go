package main

func SquareSum(numbers []int) int {
	// your code here

	accumulator := 0
	for _, n := range numbers {
		accumulator = accumulator + (n * n)
	}

	return accumulator
}
