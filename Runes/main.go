package main

import (
	"fmt"
	"strconv"
	//"strings"
)

func FeastUnicode(beast []rune, dish []rune) bool {
	i := len(beast) - 1
	j := len(dish) - 1
	if beast[0] != dish[0] {
		return false
	} else if beast[i] != dish[j] {
		return false
	}
	return true
}

func Feast(beast string, dish string) bool {
	// convert to rune slices for Unicode support
	beastU := []rune(beast)
	dishU := []rune(dish)
	return FeastUnicode(beastU, dishU)
}

func DNAtoRNA(dna string) string {
	rna := make([]rune, len(dna))
	for i, c := range dna {
		if string(c) == "T" {
			rna[i] = 'U'
		} else {
			rna[i] = c
		}
	}
	return string(rna)
}

func DNAStrand(dna string) string {
	result := make([]rune, len(dna))
	for i, character := range dna {
		switch character {
		case 'A':
			result[i] = 'T'
		case 'T':
			result[i] = 'A'
		case 'C':
			result[i] = 'G'
		case 'G':
			result[i] = 'C'
		default:
			result[i] = character
		}
	}
	return string(result)
}

func Parse(s string) (int, int) {
	xs := s[:1]
	ys := s[2:]
	x, _ := strconv.Atoi(xs)
	y, _ := strconv.Atoi(ys)
	return x, y
}

func Points(games []string) int {
	result := 0
	var point int
	for _, s := range games {
		x, y := Parse(s)
		if x > y {
			point = 3
		}
		if x < y {
			point = 0
		}
		if x == y {
			point = 1
		}
		result += point
	}
	return result
}

func main() {
	a := "test"
	d := "torn boat"
	b := Feast(a, d)
	fmt.Println("Correct:", b)
	dna := "GCAT"
	rna := DNAtoRNA(dna)
	fmt.Println("dna:", dna)
	fmt.Println("rna:", rna)

	log := []string{"1:0", "2:0", "3:0", "4:0", "2:1", "3:1", "4:1", "3:2", "4:2", "4:3"}
	fmt.Println("Points(log):", Points(log))

	test := "GTAT"
	fmt.Println("DNAStrand:", DNAStrand(test))
}
