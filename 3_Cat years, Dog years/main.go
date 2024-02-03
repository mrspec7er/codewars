// Cat years, Dog years
package main

import "fmt"

func CalculateYears(years int) (result [3]int) {
	if years == 1 {
		return [3]int{1, 15, 15}
	}

	if years == 2 {
		return [3]int{2, 24, 24}
	}

	catYear := 24
	dogYear := 24

	for i := 2; i < years; i++ {
		catYear = catYear + 4
		dogYear = dogYear + 5
	}

	return [3]int{years, catYear, dogYear}
}

func main() {
	fmt.Println(CalculateYears(10))
}
