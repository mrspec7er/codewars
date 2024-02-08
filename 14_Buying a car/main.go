// Buying a car

package main

import (
	"fmt"
	"math"
)

func NbMonths(startPriceOld, startPriceNew, savingperMonth int, percentLossByMonth float64) [2]int {
	cashflow := (float64(startPriceOld) - (float64(startPriceNew)))
	counter := 0
	deg := percentLossByMonth
	cPriceOld := float64(startPriceOld)
	cPriceNew := float64(startPriceNew)

	for cashflow < 0 {

		counter++
		if counter%2 == 0 {
			deg = deg + 0.5
		}
		cPriceOld = cPriceOld - (float64(cPriceOld) * deg / 100)
		cPriceNew = cPriceNew - (float64(cPriceNew) * deg / 100)

		cashflow = (float64(savingperMonth*counter) + cPriceOld) - cPriceNew

	}

	resCashflow := math.Round(cashflow)

	return [2]int{counter, int(resCashflow)}
}

func main() {
	fmt.Println(NbMonths(7500, 32000, 300, 1.55))
}
