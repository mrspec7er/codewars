// Pete, the baker
package main

import "fmt"

func Cakes(recipe, available map[string]int) int {
	maxIteration := 99999
	for i, r := range recipe {
		isAvailable := false
		for n, a := range available {
			if i == n {
				isAvailable = true
				if a/r < maxIteration {
					maxIteration = a / r
				}
			}
		}
		if isAvailable == false {
			return 0
		}
	}
	return maxIteration
}

func main() {
	fmt.Println(Cakes(map[string]int{"apples": 3, "flour": 300, "sugar": 150, "milk": 100, "oil": 100}, map[string]int{"sugar": 500, "flour": 2000, "milk": 2000}))
}
