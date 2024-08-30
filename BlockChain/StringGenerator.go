package main

import "strings"

type StringGenerator struct {
	StringValue [11]byte
	value       [11]byte
	baseMap     map[byte]byte
}

func (sgen *StringGenerator) Initialize(seed uint64) {
	sgen.baseMap = make(map[byte]byte)

	var i byte = 0
	for c := byte('A'); c <= byte('Z'); c++ {
		sgen.baseMap[i] = c
		i++
	}
	for c := byte('a'); c <= byte('z'); c++ {
		sgen.baseMap[i] = c
		i++
	}
	for c := byte('0'); c <= byte('9'); c++ {
		sgen.baseMap[i] = c
		i++
	}
	sgen.baseMap[i] = byte('+')
	i++
	sgen.baseMap[i] = byte('/')

	// 11 byte bufer is enough for 64 bits value, when we store 6 bits per byte
	for j := 0; j < 11; j++ {
		sgen.value[j] = byte(seed % 64)
		sgen.StringValue[j] = sgen.baseMap[sgen.value[j]]
		seed /= 64
	}
}

func (sgen *StringGenerator) Next() {
	for i := 0; i < 11; i++ {
		if sgen.value[i] < 63 {
			sgen.value[i]++
			sgen.StringValue[i] = sgen.baseMap[sgen.value[i]]
			return
		}

		sgen.value[i] = 0
		sgen.StringValue[i] = sgen.baseMap[0]
	}
}

func (sgen *StringGenerator) GetString() string {
	var sb strings.Builder
	for i := 10; i >= 0; i-- {
		sb.WriteByte(sgen.StringValue[i])
	}
	return sb.String()
}
