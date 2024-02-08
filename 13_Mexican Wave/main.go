// Mexican Wave

package main

import (
	"fmt"
	"strings"
)

func wave(words string) []string {
	result := []string{}
	for i := 0; i < len(words); i++ {
		w := ""
		for n, c := range words {
			if n == i {
				w = w + strings.ToUpper(string(c))
			} else {
				w = w + string(c)
			}
		}

		if string(words[i]) != " " {
			result = append(result, w)
		}

	}

	return result
}

func main() {
	// fmt.Println(wave("abc"))
	fmt.Println(wave(" x yz"))
	// fmt.Println(wave("a a a a a"))
}
