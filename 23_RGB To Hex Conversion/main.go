// RGB To Hex Conversion
package main

import (
	"fmt"
	"strconv"
)

func RGB(r, g, b int) string {

	rHex := RGBToHex(r)
	gHex := RGBToHex(g)
	bHex := RGBToHex(b)

	return rHex + gHex + bHex
}

func RGBToHex(v int) string {
	if v <= 0 {
		return "00"
	}
	if v >= 255 {
		return "FF"
	}
	firstDigit := v / 16
	firstHex := HexConverter(firstDigit)
	secondDigit := v % 16
	secondHex := HexConverter(secondDigit)

	return firstHex + secondHex
}

func HexConverter(d int) string {
	if d <= 0 {
		return "0"
	}

	if d == 10 {
		return "A"
	}

	if d == 11 {
		return "B"
	}

	if d == 12 {
		return "C"
	}

	if d == 13 {
		return "D"
	}

	if d == 14 {
		return "E"
	}

	if d == 15 {
		return "F"
	}

	return strconv.Itoa(d)
}

func main() {
	fmt.Println(RGB(254, 253, 252))
}
