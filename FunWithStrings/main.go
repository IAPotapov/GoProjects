package main

import (
	"fmt"
	"sort"
	"strings"
)

type Data struct {
	Character                          rune
	CountS1, CountS2, CountMax, Leader int
	Report                             string
}

var result []Data

func NewReport(r rune, c1, c2, cm int) string {
	var sb strings.Builder
	if c1 == c2 {
		sb.WriteString("=:")
	} else if c1 == cm {
		sb.WriteString("1:")
	} else {
		sb.WriteString("2:")
	}
	sb.WriteString(strings.Repeat(string(r), cm))
	return sb.String()
}

func NewData(s1, s2 string, r rune) Data {
	c1 := strings.Count(s1, string(r))
	c2 := strings.Count(s2, string(r))
	cm := c1
	if c2 > c1 {
		cm = c2
	}
	var i int
	if c1 > c2 {
		i = 1
	}
	if c2 > c1 {
		i = 2
	}
	if c1 == c2 {
		i = 3
	}
	return Data{
		Character: r,
		CountS1:   c1,
		CountS2:   c2,
		CountMax:  cm,
		Leader:    i,
		Report:    NewReport(r, c1, c2, cm),
	}
}

func Mix(s1, s2 string) string {
	result = make([]Data, 0)
	for _, r := range "abcdefghijklmnopqrstuvwxyz" {
		result = append(result, NewData(s1, s2, r))
	}
	sort.Slice(result, func(i0, i1 int) bool {
		if result[i0].CountMax != result[i1].CountMax {
			return result[i0].CountMax > result[i1].CountMax
		}
		if result[i0].Leader != result[i1].Leader {
			return result[i0].Leader < result[i1].Leader
		}
		return result[i0].Character < result[i1].Character
	})
	var sb strings.Builder
	for i, v := range result {
		if v.CountMax < 2 {
			continue
		}
		if i > 0 {
			sb.WriteRune('/')
		}
		sb.WriteString(v.Report)
	}
	return sb.String()
}

func main() {

	//s1 := "my&friend&Paul has heavy hats! &"
	//s2 := "my friend John has many many friends &"
	// "2:nnnnn/1:aaaa/1:hhh/2:mmm/2:yyy/2:dd/2:ff/2:ii/2:rr/=:ee/=:ss"
	//fmt.Println("Mix:", Mix(s1, s2))

	// dotest("Are they here", "yes, they are here", "2:eeeee/2:yy/=:hh/=:rr")
	//fmt.Println("Mix:", Mix("Are they here", "yes, they are here"))
	//dotest("looping is fun but dangerous", "less dangerous than coding",  "1:ooo/1:uuu/2:sss/=:nnn/1:ii/2:aa/2:dd/2:ee/=:gg")
	s1 := "looping is fun but dangerous"
	s2 := "less dangerous than coding"
	fmt.Println("Mix:", Mix(s1, s2))
}
