// Sum of odd numbers
package main

import "fmt"

func RowSumOddNumbers(n int) int {
	start := 1

	for i := 0; i < n; i++ {
		start = start + (2 * i)
	}

	accumulate := start

	for i := 1; i < n; i++ {
		accumulate = accumulate + start + (i * 2)
	}

	return accumulate
}

func main() {
	fmt.Println(RowSumOddNumbers(1))
	fmt.Println(RowSumOddNumbers(2))
	fmt.Println(RowSumOddNumbers(13))
	fmt.Println(RowSumOddNumbers(19))
	fmt.Println(RowSumOddNumbers(41))
}
