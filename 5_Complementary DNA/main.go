// Complementary DNA
package main

import "fmt"

func DNAStrand(dna string) string {
	result := []rune{}
	for i, v := range dna {
		if v == 'A' {
			result = append(result, 'T')
			continue
		}

		if v == 'T' {
			result = append(result, 'A')
			continue
		}

		if v == 'C' {
			result = append(result, 'G')
			continue
		}

		if v == 'G' {
			result = append(result, 'C')
			continue
		}

		result = append(result, result[i])
	}

	return string(result)
}

func main() {
	fmt.Println(DNAStrand("AAAA"))
	fmt.Println(DNAStrand("ATTGC"))
	fmt.Println(DNAStrand("GTAT"))
}
