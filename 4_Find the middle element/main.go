// Find the middle element
package main

import (
	"sort"
)

func Gimme(array [3]int) int {
	sorted := array
	sort.Ints(sorted[:])

	if array[0] == sorted[1] {
		return 0
	}

	if array[1] == sorted[1] {
		return 1
	}

	return 2
}

func main() {
	Gimme([3]int{-15, -17, 14})
}
