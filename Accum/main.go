package main

import (
	"fmt"
	"math"
	"sort"
	"strconv"
	"strings"
)

func CreateRepetition(character rune, repetititon int) string {
	result := make([]rune, repetititon)
	capital := string(character)
	capital = strings.ToUpper(capital)
	capitalR := []rune(capital)
	for i := 0; i < repetititon; i++ {
		if i == 0 {
			result[i] = capitalR[0]
		} else {
			result[i] = character
		}
	}
	return string(result)
}

func Accum(s string) string {
	s2 := strings.ToLower(s)
	var result strings.Builder
	for i, v := range s2 {
		if i > 0 {
			result.WriteRune('-')
		}
		result.WriteString(CreateRepetition(v, i+1))
	}
	return result.String()
}

func ToJadenCase(str string) string {
	fields := strings.Fields(str)
	var result strings.Builder
	for i, s := range fields {
		if i > 0 {
			result.WriteRune(' ')
		}
		result.WriteString(strings.Title(s))
	}
	return result.String()
}

func AddCharacter(array []rune, character rune) []rune {
	for _, c := range array {
		if c == character {
			return array
		}
	}
	return append(array, character)
}

func TwoToOne(s1 string, s2 string) string {
	array := make([]rune, 0, 1)
	for _, character := range s1 {
		array = AddCharacter(array, character)
	}
	for _, character := range s2 {
		array = AddCharacter(array, character)
	}
	sort.Slice(array, func(i0, i1 int) bool {
		return array[i0] < array[i1]
	})
	return string(array)
}

func FindNextSquare(sq int64) int64 {
	var a float64 = float64(sq)
	a = math.Sqrt(a)
	a = math.Floor(a)
	var i int64 = int64(a)
	if i*i != sq {
		return -1
	}
	i = (i + 1) * (i + 1)
	return i
}

func PrinterError(s string) string {
	good := []rune{'a', 'b', 'c', 'd', 'e', 'f', 'g', 'h', 'i', 'j', 'k', 'l', 'm'}
	result := 0
	for _, character := range s {
		found := false
		for _, g := range good {
			if character == g {
				found = true
				break
			}
		}
		if !found {
			result++
		}
	}
	return strconv.Itoa(result) + "/" + strconv.Itoa(len(s))
}

func solution(str, ending string) bool {
	a1 := []rune(str)
	a2 := []rune(ending)
	strLen := len(str)
	endingLen := len(ending)
	for i := 0; i < endingLen; i++ {
		if a2[endingLen-i-1] != a1[strLen-i-1] {
			return false
		}
	}
	return true
}

func NbYear(p0 int, percent float64, aug int, p int) int {
	n := float64(p0)
	a := float64(aug)
	f := float64(p)
	d := percent * 0.01
	i := 0
	for n < f {
		n += math.Floor(a + n*d)
		i++
	}
	return i
}

func Multiple3And5(number int) int {
	var result int = 0
	for i := 3; i < number; i++ {
		if i%3 == 0 {
			result += i
			fmt.Println(i)
			continue
		}
		if i%5 == 0 {
			result += i
			fmt.Println(i)
		}
	}
	return result
}

func main() {
	test := Accum("ZpglnRxqenU")
	fmt.Println(test)
	test = "All the rules in this world were made by someone no smarter than you. So make your own."
	fmt.Println(ToJadenCase(test))

	a := "xyaabbbccccdefww"
	b := "xxxxyyyyabklmopq"
	fmt.Println(TwoToOne(a, b))

	var i int64 = 15241383936
	var j int64 = FindNextSquare(i)
	fmt.Println("i:", i, "j:", j)

	test = "kkkwwwaaaaaaaaaaaaaabbbbbbbbbbbbbbbbbbmmmmmmmmmmmmmmmmmmmxyzuuuuu"
	fmt.Println("PrinterError:", PrinterError(test))

	a = "banana"
	b = "ana"
	fmt.Println("solution:", solution(a, b))
	fmt.Println("NbYear", NbYear(1500000, 0.25, 1000, 2000000))
	fmt.Println("Multiple3And5", Multiple3And5(10))
}
