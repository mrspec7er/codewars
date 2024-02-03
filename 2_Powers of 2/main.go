// Powers of 2
package main

import (
	"fmt"
	"math"
)

func PowersOfTwo(n int) []uint64 {
	result := []uint64{}

	for i := 0; i <= n; i++ {
		current := math.Pow(2, float64(i))

		result = append(result, uint64(current))
	}

	return result
}

func main() {
	fmt.Println(PowersOfTwo(4))
}
