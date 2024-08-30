package main

import (
	"fmt"
	"strconv"
	"strings"
)

func GetCount(str string) (count int) {
	vowels := []rune{'a', 'e', 'i', 'o', 'u'}
	count = 0
	for _, a := range str {
		for _, b := range vowels {
			if a == b {
				count++
			}
		}
	}
	return
}

func IsVowel(r rune) bool {
	vowels := []rune{'a', 'e', 'i', 'o', 'u'}
	for _, c := range vowels {
		if c == r {
			return true
		}
	}
	return false
}

func Disemvowel(comment string) string {
	runes := []rune(comment)
	runesLow := []rune(strings.ToLower(comment))
	var sb strings.Builder
	for i := 0; i < len(runes); i++ {
		if !IsVowel(runesLow[i]) {
			sb.WriteRune(runes[i])
		}
	}
	return sb.String()
}

func HighAndLow(in string) string {
	fields := strings.Fields(in)
	var max int
	var min int
	for i, v := range fields {
		num, _ := strconv.Atoi(v)
		if i == 0 {
			max = num
			min = num
		}
		if max < num {
			max = num
		}
		if min > num {
			min = num
		}
	}
	return fmt.Sprintf("%d %d", max, min)
}

func main() {
	s := "Hello, World, LOL!"
	fmt.Println("GetCount(s):", GetCount(s))
	fmt.Println("Disemvowel(s):", Disemvowel(s))
	s2 := "8 3 -5 42 -1 0 0 -9 4 7 4 -4"
	fmt.Println("HighAndLow(s2):", HighAndLow(s2))

	var a int = 1
	fmt.Println("a:", a)
}
