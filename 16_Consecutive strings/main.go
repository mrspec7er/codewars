// Consecutive strings
package main

import "fmt"

func LongestConsec(strarr []string, k int) string {
	lStararr := ""
	for i := 0; i < len(strarr)-(k-1); i++ {
		tempStarrar := ""
		for n := i; n < i+k; n++ {
			tempStarrar = tempStarrar + strarr[n]
		}

		if len(tempStarrar) > len(lStararr) {
			lStararr = tempStarrar
		}
	}

	return lStararr
}

func main() {
	fmt.Println(LongestConsec([]string{"zone", "abigail", "theta", "form", "libe", "zas"}, 2))
}
