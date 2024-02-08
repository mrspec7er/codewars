// Find the missing letter
package main

import (
	"fmt"
)

func FindMissingLetter(chars []rune) rune {
	n := 0
	index := int(chars[0])
	for i := index; i < i+len(chars); i++ {

		if int(chars[n]) != i {
			return chars[n] - 1
		}
		n++
	}
	return 'c'
}

func main() {
	fmt.Println(FindMissingLetter([]rune{'a', 'b', 'c', 'd', 'f'}))
	fmt.Println(FindMissingLetter([]rune{'O', 'Q', 'R', 'S'}))
}
