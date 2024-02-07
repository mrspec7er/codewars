// Persistent Bugger
package main

import (
	"fmt"
)

func Persistence(n int) int {
	counter := 0

	if n < 10 {
		return counter
	}

	return test(n, &counter)
}

func test(n int, counter *int) int {
	slc := []int{}
	for n > 0 {
		slc = append(slc, n%10)
		n = n / 10
	}

	acm := 1
	for i := 0; i < len(slc); i++ {
		acm = acm * slc[i]
	}

	*counter++
	if acm > 9 {
		test(acm, counter)
	}

	return *counter
}

func main() {
	fmt.Println(Persistence(25))
}
