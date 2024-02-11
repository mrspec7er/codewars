// Weight for weight
package main

import (
	"fmt"
	"sort"
	"strconv"
	"strings"
)

func OrderWeight(strng string) string {
	initialList := strings.Split(strng, " ")
	resultList := []struct {
		str   string
		value int
	}{}

	for _, eachList := range initialList {
		value := 0
		for _, l := range eachList {
			c, _ := strconv.Atoi(string(l))
			value = value + c
		}
		resultList = append(resultList, struct {
			str   string
			value int
		}{str: eachList, value: value})
	}

	sort.Slice(resultList, func(i, j int) bool {
		if resultList[i].value == resultList[j].value {
			return resultList[i].str < resultList[j].str
		}
		return resultList[i].value < resultList[j].value
	})

	result := ""

	for i, r := range resultList {
		result = result + r.str
		if i != len(resultList)-1 {
			result = result + " "
		}
	}
	return result
}

func main() {
	fmt.Println(OrderWeight("2000 10003 1234000 44444444 9999 11 11 22 123"))
}
