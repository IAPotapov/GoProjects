package main

import (
	"fmt"
	"strconv"

	//"math"
	"sort"
	"strings"
)

func SplitDuration(value int64) (y, d, h, m, s int) {
	s = int(value % 60)
	value /= 60
	m = int(value % 60)
	value /= 60
	h = int(value % 24)
	value /= 24
	d = int(value % 365)
	y = int(value / 365)
	return
}

func FormatValue(parts []string, value int, name string) []string {
	if value == 1 {
		parts = append(parts, "1 "+name)
	}
	if value > 1 {
		f := fmt.Sprintf("%d %ss", value, name)
		parts = append(parts, f)
	}
	return parts
}

func FormatDuration(seconds int64) string {
	y, d, h, m, s := SplitDuration(seconds)
	parts := make([]string, 0)
	parts = FormatValue(parts, y, "year")
	parts = FormatValue(parts, d, "day")
	parts = FormatValue(parts, h, "hour")
	parts = FormatValue(parts, m, "minute")
	parts = FormatValue(parts, s, "second")
	num := len(parts)
	if num == 0 {
		return "now"
	}
	if num == 2 {
		return parts[0] + " and " + parts[1]
	}
	var sb strings.Builder
	sb.WriteString(parts[0])
	for i := 1; i < num; i++ {
		if i == num-1 {
			sb.WriteString(" and ")
		} else {
			sb.WriteString(", ")
		}
		sb.WriteString(parts[i])
	}
	return sb.String()
}

type Element struct {
	First, Last int
	IsRange     bool
}

func Feed(elements []Element, value int) []Element {
	length := len(elements)
	e := Element{First: value, Last: value, IsRange: false}
	if length < 1 {
		elements = append(elements, e)
		return elements
	}
	last := elements[length-1]

	if last.IsRange {
		if value-last.Last == 1 {
			last.Last = value
			elements[length-1] = last
			return elements
		} else {
			elements = append(elements, e)
			return elements
		}
	}
	if length < 2 {
		elements = append(elements, e)
		return elements
	}
	previous := elements[length-2]
	if previous.IsRange || value-previous.Last != 2 {
		elements = append(elements, e)
		return elements
	}
	// remove last and previous
	elements = elements[:length-2]
	// convert to range
	e.IsRange = true
	e.First = previous.First
	elements = append(elements, e)

	return elements
}

func ConvertToString(elements []Element) string {
	var sb strings.Builder
	for i, e := range elements {
		if i > 0 {
			sb.WriteRune(',')
		}
		if e.IsRange {
			f := fmt.Sprintf("%d-%d", e.First, e.Last)
			sb.WriteString(f)
		} else {
			f := fmt.Sprintf("%d", e.First)
			sb.WriteString(f)
		}
	}
	return sb.String()
}

func Solution(list []int) string {
	elements := make([]Element, 0)
	for i := 0; i < len(list); i++ {
		elements = Feed(elements, list[i])
	}
	return ConvertToString(elements)
}

func ConvertToInt(B []int) int {
	result := 0
	for i, p := 0, 1; i < len(B); i, p = i+1, p*10 {
		result += p * B[len(B)-1-i]
	}
	return result
}

func GetCandidate(A []int, i int) int {
	Acopy := make([]int, len(A))
	copy(Acopy, A)
	B := Acopy[i+1:]
	sort.Ints(B)
	for j := 0; j < len(B); j++ {
		if Acopy[i] < B[j] {
			Acopy[i], B[j] = B[j], Acopy[i]
			return ConvertToInt(Acopy)
		}
	}
	return -1
}

func NextBigger(n int) int {
	s := strconv.Itoa(n)
	runes := []rune(s)
	length := len(runes)
	A := make([]int, length)
	for i, r := range runes {
		A[i], _ = strconv.Atoi(string(r))
	}
	for i := length - 2; i >= 0; i-- {
		candidate := GetCandidate(A, i)
		if candidate > 0 {
			return candidate
		}
	}
	return -1
}

func ConvertToInt2(B []int) int {
	result := 0
	for i, p := 0, 1; i < len(B); i, p = i+1, p*10 {
		result += p * B[i]
	}
	return result
}

func GetCandidate2(A []int, i int) int {
	Acopy := make([]int, len(A))
	copy(Acopy, A)
	B := Acopy[:i]
	sort.Ints(B)
	for j := len(B) - 1; j >= 0; j-- {
		if Acopy[i] > B[j] {
			if i == len(Acopy)-1 && B[j] == 0 {
				continue
			}
			Acopy[i], B[j] = B[j], Acopy[i]
			return ConvertToInt2(Acopy)
		}
	}
	return -1
}

func NextSmaller(n int) int {
	digits := make([]int, 0)
	n2 := n
	for n2 > 0 {
		digits = append(digits, n2%10)
		n2 /= 10
	}
	for i := 1; i < len(digits); i++ {
		candidate := GetCandidate2(digits, i)
		if candidate > 0 {
			return candidate
		}
	}
	return -1
}

func main() {
	/*var testD int64 = 3600
	fmt.Println("FormatDuration:", FormatDuration(testD))
	test2 := []int{-6, -3, -2, -1, 0, 1, 3, 4, 5, 7, 8, 9, 10, 11, 14, 15, 17, 18, 19, 20}
	fmt.Println("Solution:", Solution(test2))
	test3 := []int{59, 60, 61, 62, 68, 69, 70, 71, 80, 81, 82, 83, 84, 85, 88, 97}
	fmt.Println("Solution:", Solution(test3))
	//A := []int{1, 2, 3, 4}
	//h := make([][]int, 0)
	//Heaps(len(A), A)
	//fmt.Println("Heaps:", Heaps(len(A), A))*/
	// 59884848459853
	// 1234567890
	// n = 2081053450, should return 2081053405
	test := 2081053450
	fmt.Println("NextBigger:", NextBigger(test))
	fmt.Println("NextSmaller:", NextSmaller(test))
}
