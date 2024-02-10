// Number of trailing zeros of N!
package main

import "fmt"

func Zeros(n int) int {
	if n < 5 {
		return 0
	}
	counter := 0
	temp := n

	for {
		if temp < 5 {
			break
		}
		temp = temp / 5
		counter = counter + temp
	}

	return counter
}

func main() {
	// fmt.Println("527515743", Zeros(527515743))
	// fmt.Println("5", Zeros(5))
	// fmt.Println("12", Zeros(12))
	// fmt.Println("21", Zeros(21))
	// fmt.Println("30", Zeros(30))
	fmt.Println("126", Zeros(126))
}
