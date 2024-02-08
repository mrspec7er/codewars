// Your order, please
package main

import (
	"fmt"
	"strings"
)

func Order(sentence string) string {
	words := strings.Fields(sentence)
	orderedWords := [9]string{}
	for i := 0; i < len(words); i++ {
		for _, c := range words[i] {
			if c == '1' {
				orderedWords[0] = words[i]
			}
			if c == '2' {
				orderedWords[1] = words[i]
			}
			if c == '3' {
				orderedWords[2] = words[i]
			}
			if c == '4' {
				orderedWords[3] = words[i]
			}
			if c == '5' {
				orderedWords[4] = words[i]
			}
			if c == '6' {
				orderedWords[5] = words[i]
			}
			if c == '7' {
				orderedWords[6] = words[i]
			}
			if c == '8' {
				orderedWords[7] = words[i]
			}
			if c == '9' {
				orderedWords[8] = words[i]
			}
		}
	}

	result := ""

	for i := 0; i < len(words); i++ {
		result = result + orderedWords[i]
		if i != len(words)-1 {
			result = result + " "
		}
	}

	return result
}

func main() {
	fmt.Println(Order("4of Fo1r pe6ople g3ood th5e the2"))
}
