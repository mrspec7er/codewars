// Are they the "same"?
package main

import "fmt"

func Comp(array1 []int, array2 []int) bool {

	if len(array2) == 0 {
		return false
	}
	for _, a2 := range array2 {
		isSame := false
		for n, a1 := range array1 {
			if a1 == 0 {
				continue
			}
			if a2 == (a1 * a1) {
				array1[n] = 0
				isSame = true
				break
			}
		}

		if !isSame {
			return isSame
		}
	}

	return true
}

func main() {
	// var a1 = []int{121, 144, 19, 161, 19, 144, 19, 11}
	// var a2 = []int{11 * 11, 121 * 121, 144 * 144, 19 * 19, 161 * 161, 19 * 19, 144 * 144, 19 * 19}

	// a1 := []int{121, 144, 19, 161, 19, 144, 19, 11}
	// a2 := []int{11 * 21, 121 * 121, 144 * 144, 19 * 19, 161 * 161, 19 * 19, 144 * 144, 19 * 19}

	// var a1 []int
	// a1 = nil
	// a2 := []int{11 * 11, 121 * 121, 144 * 144, 19 * 19, 161 * 161, 19 * 19, 144 * 144, 19 * 19}

	a1 := []int{2, 3, 3}
	a2 := []int{}

	// a1 := []int{-65, -79, 44, -3, 24, 66, 96, -3, -19, -1}
	// a2 := []int{1, 9, 9, 361, 576, 1936, 4225, 4356, 6241, 9216}
	fmt.Println(Comp(a1, a2))
}
