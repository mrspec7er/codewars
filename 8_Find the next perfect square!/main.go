// Find the next perfect square!
package main

import (
	"fmt"
	"math"
)

func FindNextSquare(sq int64) int64 {

	sqDefault := math.Sqrt(float64(sq))

	_, fractional := math.Modf(sqDefault)

	if fractional != 0 {
		return -1
	}

	return int64(sqDefault+1) * int64(sqDefault+1)

}

func main() {
	fmt.Println(FindNextSquare(625))
}
