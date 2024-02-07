// Build Tower
package main

import "fmt"

func TowerBuilder(nFloors int) []string {
	tower := []string{}
	for i := 0; i < nFloors; i++ {
		towerStack := ""
		for n := 1; n <= (nFloors*2 - 1); n++ {
			if n >= nFloors-i && n <= nFloors+i {
				towerStack += "*"
			} else {
				towerStack += " "
			}
		}
		tower = append(tower, towerStack)
	}
	return tower
}

func main() {
	fmt.Println(TowerBuilder(3))
}
