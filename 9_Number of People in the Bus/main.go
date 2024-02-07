// Number of People in the Bus
package main

import "fmt"

func Number(busStops [][2]int) int {

	counter := 0
	for _, p := range busStops {
		fmt.Println(p[0], p[1])
		counter = counter + (p[0] - p[1])
	}
	return counter // Good Luck!
}

func main() {
	fmt.Println(Number([][2]int{{3, 0}, {9, 1}, {4, 10}, {12, 2}, {6, 1}, {7, 10}}))
}
