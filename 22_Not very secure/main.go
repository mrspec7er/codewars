// Not very secure
package main

import "fmt"

func alphanumeric(str string) bool {
	if len(str) == 0 {
		return false
	}

	for _, c := range str {
		if c == ' ' || c == '_' {
			return false
		}

		if c < 48 || c > 57 && c < 65 || c > 90 && c < 97 || c > 122 {
			return false
		}
	}

	return true
}

func main() {
	fmt.Println(alphanumeric("asd"))
	fmt.Println(alphanumeric("as_d"))
	fmt.Println(alphanumeric("asdc112}ASD"))
}
