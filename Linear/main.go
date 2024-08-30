package main

import (
	"fmt"
	"strings"

	"github.com/gSpera/morse"
)

var MORSE_CODE map[string]string = map[string]string{}

func IsLinear(x int) bool {

	if x == 1 {
		return true
	}
	if (x-1)%2 == 0 {
		if IsLinear(x / 2) {
			return true
		}
	}
	if (x-1)%3 == 0 {
		if IsLinear(x / 3) {
			return true
		}
	}
	return false
}

func DblLinear(n int) int {
	value := 1
	index := 0
	for index < n {
		index++
		value++
		for !IsLinear(value) {
			value++
		}

	}
	return value
}

func GetEnd(bits string, begin int) int {
	for j := begin + 1; j < len(bits); j++ {
		if bits[j] != bits[begin] {
			return j
		}
	}
	return len(bits)
}

func GetSize(bits string) int {
	min := len(bits)
	for i := 0; i < len(bits); {
		end := GetEnd(bits, i)
		n := end - i
		if min > n {
			min = n
		}
		i = end
	}
	return min
}

func DecodeBits(bits string) string {
	for bits[0] == '0' {
		bits = bits[1:]
	}
	for bits[len(bits)-1] == '0' {
		bits = bits[:len(bits)-1]
	}
	unit := GetSize(bits)
	var sb strings.Builder

	for i := 0; i < len(bits); {
		end := GetEnd(bits, i)
		n := (end - i) / unit
		if bits[i] == '1' && n == 1 {
			sb.WriteRune('.')
		}
		if bits[i] == '1' && n == 3 {
			sb.WriteRune('-')
		}
		if bits[i] == '0' && n == 3 {
			sb.WriteRune(' ')
		}
		if bits[i] == '0' && n == 7 {
			sb.WriteString("   ")
		}

		i = end
	}
	return sb.String()
}

func DecodeMorse(morseCode string) string {
	var sb strings.Builder
	words := strings.Split(strings.TrimSpace(morseCode), "  ")
	for _, w := range words {
		if sb.Len() > 0 {
			sb.WriteString(" ")
		}
		letters := strings.Split(w, " ")
		for _, c := range letters {
			sb.WriteString(MORSE_CODE[c])
		}
	}
	return sb.String()
}

func main() {
	var r rune
	var s string
	n := 1000
	s = fmt.Sprintf("DblLinear(%d): %d", n, DblLinear(n))
	fmt.Println(s)

	str := "--."
	ch := morse.RuneToText(str)

	fmt.Printf("The morse code %s converts to: %c\n", str, ch)

	for i := int('A'); i <= int('Z'); i++ {
		r = rune(i)
		s = string(r)
		MORSE_CODE[morse.ToMorse(s)] = string(r)
	}
	r = '.'
	s = string(r)
	MORSE_CODE[morse.ToMorse(s)] = string(r)

	//s = ".... . -.--   .--- ..- -.. ."
	//fmt.Printf("DecodeMorse(%s): %s\n", s, DecodeMorse(s))

	//s = "1100110011001100000011000000111111001100111111001111110000000000000011001111110011111100111111000000110011001111110000001111110011001100000011"
	s = "11111100111111"
	fmt.Printf("GetSize(s): %d\n", GetSize(s))
	s2 := DecodeBits(s)
	fmt.Printf("DecodeBits(s): %s\n", s2)
	fmt.Printf("DecodeMorse(s2): %s\n", DecodeMorse(s2))
}
