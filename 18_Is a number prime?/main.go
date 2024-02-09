// Is a number prime?
package main

import (
	"fmt"
	"math"
)

func IsPrime(n int) bool {
	if n < 2 {
		return false
	}

	if n == 3 || n == 2 {
		return true
	}

	if n%2 == 0 {
		return false
	}

	sq_root := int(math.Sqrt(float64(n)))
	for i := 3; i <= sq_root; i += 2 {
		if n%i == 0 {
			return false
		}
	}
	return true
}

func main() {
	fmt.Println(IsPrime(64))
}
