package main

import (
	"fmt"
	"sort"
	"strconv"
	"strings"
)

func IsTriangle(a, b, c int) bool {
	sides := []int{a, b, c}
	sort.Slice(sides, func(i0, i1 int) bool {
		return sides[i0] < sides[i1]
	})
	if sides[0] <= 0 {
		return false
	}
	return (sides[0] + sides[1]) == sides[2]
}

func SpinWord(str string) string {
	runes := []rune(str)
	result := make([]rune, len(runes))
	for i := 0; i < len(str); i++ {
		result[i] = runes[len(runes)-1-i]
	}
	return string(result)
}

func SpinWords(str string) string {
	fields := strings.Fields(str)
	var result strings.Builder
	for i, s := range fields {
		if i > 0 {
			result.WriteRune(' ')
		}
		if len(s) < 5 {
			result.WriteString(s)
		} else {
			result.WriteString(SpinWord(s))
		}
	}
	return result.String()
}

func MoveZeros(arr []int) []int {
	result := make([]int, len(arr))
	var i int = 0
	for j := 0; j < len(arr); j++ {
		if arr[j] != 0 {
			result[i] = arr[j]
			i++
		}
	}
	return result
}

func Sprintf(value int) string {
	result := strconv.Itoa(value)
	if len(result) < 2 {
		result = "0" + result
	}
	return result
}

func HumanReadableTime(seconds int) string {
	var s int = seconds % 60
	seconds /= 60
	var m int = seconds % 60
	var h int = seconds / 60
	return Sprintf(h) + ":" + Sprintf(m) + ":" + Sprintf(s)
}

func NibbleString(value int) (result string) {
	switch {
	case value < 10:
		result = strconv.Itoa(value)
	case value == 10:
		result = "A"
	case value == 11:
		result = "B"
	case value == 12:
		result = "C"
	case value == 13:
		result = "D"
	case value == 14:
		result = "E"
	case value == 15:
		result = "F"
	}
	return
}

func HexString(value int) (result string) {
	if value > 255 {
		value = 255
	}
	if value < 0 {
		value = 0
	}
	low := value & 0x0f
	high := (value & 0xf0) >> 4
	result = NibbleString(high) + NibbleString(low)
	return
}

func RGB(r, g, b int) string {
	return HexString(r) + HexString(g) + HexString(b)
}

func main() {

	fmt.Println("IsTriangle:", IsTriangle(4, 2, 3))
	fmt.Println("SpinWords:", SpinWords("Hey fellow warriors"))
	fmt.Println(MoveZeros([]int{1, 2, 0, 1, 0, 1, 0, 3, 0, 1}))
	fmt.Println("Sprintf:", Sprintf(0))
	fmt.Println("HumanReadableTime:", HumanReadableTime(61))
	fmt.Println("RGB:", RGB(-1, 277, 3))
}
