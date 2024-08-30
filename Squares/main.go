package main

import (
	"fmt"
	"math"
)

func IsCorrect(list []int64) bool {
	for i := 1; i < len(list); i++ {
		if list[i-1] >= list[i] {
			return false
		}
	}
	return true
}

func GetSquares(sqn int64) (result []int64, ok bool) {
	a := math.Sqrt(float64(sqn))
	a = math.Floor(a)
	b := int64(a)
	for i := b; i > 0; i-- {
		sqn2 := sqn - i*i
		if sqn2 == 0 {
			result = []int64{i}
			ok = true
			return
		}
		resulti, oki := GetSquares(sqn2)
		if oki {
			resulti = append(resulti, i)
			if IsCorrect(resulti) {
				result = resulti
				ok = true
				return
			}
		}
	}
	ok = false
	return
}

func Decompose(n int64) []int64 {
	for i := n - 1; i > 0; i-- {
		result, ok := GetSquares(n*n - i*i)
		if ok {
			result = append(result, i)
			if IsCorrect(result) {
				return result
			}
		}
	}
	result := make([]int64, 0)
	return result
}

/*
Hamming Numbers
The sequence of Hamming numbers 1, 2, 3, 4, 5, ...
consists of all numbers of the form Pow(2, i) * Pow(3, j) * Pow(5, k) where i, j and k are non-negative integers.
Axiom 1: The value 1 is in the sequence.
Axiom 2: If x is in the sequence, so are 2 * x, 3 * x, and 5 * x.
Axiom 3: The sequence contains no other values other than those that belong to it on account of Axioms 1 and 2.
*/

/*

var stopFlag bool = false



func Hammer(n int) uint {
	// var y float64 = math.Pow(float64(n), 0.3)
	// y = math.Floor(y)
	// iy := uint(y) + 2
	// fmt.Println("iy:", iy)
	var y float64 = math.Log2(float64(n))
	y = math.Ceil(y)
	iy := uint(y)
	count1 := iy / 2
	count2 := iy - count1

	list := make([]uint, 0)
	outputChan1 := make(chan uint)
	go GetHamming(0, count1, outputChan1)
	outputChan2 := make(chan uint)
	go GetHamming(count1, count2, outputChan2)

	// for {
	// 	if x, ok := <-outputChan1; ok {
	// 		list = append(list, x)
	// 	} else {
	// 		break
	// 	}
	// }
	done1 := false
	done2 := false
	for !done1 || !done2 {
		select {
		case x, ok := <-outputChan1:
			if ok {
				list = append(list, x)
			} else {
				done1 = true
			}
		case x, ok := <-outputChan2:
			if ok {
				list = append(list, x)
			} else {
				done2 = true
			}
		}
	}

	less := func(i, j int) bool {
		return list[i] < list[j]
	}
	sort.Slice(list, less)
	fmt.Println(list)
	return list[n-1]
}*/

// func GetHamming(n uint) (a, b, c uint) {
// 	a = n * 2
// 	b = n * 3
// 	c = n * 5
// 	return
// }

/*func Hammer(n int) uint {
	hmap := make(map[uint]bool)
	var i uint = 1
	hmap[i] = true
	for len(hmap) < n+1 {
		a, b, c := i*2, i*3, i*5
		hmap[a] = true
		hmap[b] = true
		hmap[c] = true
		for {
			i++
			if _, ok := hmap[i]; ok {
				break
			}
		}
	}
	list := make([]uint, 0)
	for i := range hmap {
		list = append(list, i)
	}
	less := func(i, j int) bool {
		return list[i] < list[j]
	}
	sort.Slice(list, less)
	fmt.Println(list)
	return list[n-1]
}*/

/*func IsHammer(value uint) bool {
	for value > 1 {
		found := false
		if value%2 == 0 {
			value /= 2
			found = true
		}
		if value%3 == 0 {
			value /= 3
			found = true
		}
		if value%5 == 0 {
			value /= 5
			found = true
		}
		if !found {
			break
		}
	}
	return value == 1
}

func Hammer(n int) uint {
	var index int = 1
	var value uint = 1
	for index < n {
		value++
		if IsHammer(value) {
			index++
		}
	}
	return value
}*/

// func GetHammingIJK(i, j, k uint) uint {
// 	var a, result uint
// 	result = 1
// 	for a = 1; a <= i; a++ {
// 		result *= 2
// 	}
// 	for a = 1; a <= j; a++ {
// 		result *= 3
// 	}
// 	for a = 1; a <= k; a++ {
// 		result *= 5
// 	}
// 	return result
// }

// func GetHamming(n int) uint {
// 	max := uint(math.Log2(float64(n))) + 1
// 	hmap := make(map[uint]bool)

// 	var i, j, k uint

// 	for k = 0; k <= max; k++ {
// 		for j = 0; j <= max; j++ {
// 			for i = 0; i <= max; i++ {
// 				hmap[GetHammingIJK(i, j, k)] = true
// 			}
// 		}
// 	}

// 	list := make([]uint, 0)
// 	for i, _ := range hmap {
// 		list = append(list, i)
// 	}
// 	less := func(i, j int) bool {
// 		return list[i] < list[j]
// 	}
// 	sort.Slice(list, less)
// 	//fmt.Println(list)
// 	return list[n-1]
// }

func Hammer(n int) uint {
	primes := []uint{2, 3, 5}
	values := []uint{2, 3, 5}
	indexes := []uint{0, 1, 2}
	results := make([]uint, n)
	results[0] = 1
	for i := 1; i < n; i++ {
		results[i] = values[0]
		for p := 1; p < len(primes); p++ {
			if results[i] > values[p] {
				results[i] = values[p]
			}
		}
		for p := 0; p < len(primes); p++ {
			if results[i] == values[p] {
				indexes[p] += 1
				values[p] = primes[p] * results[indexes[p]]
			}
		}
	}
	return results[n-1]
}

func main() {
	// var n int64 = 7
	// result := Decompose(n)
	// fmt.Println(result)

	for i := 1; i <= 19; i++ {
		fmt.Println("i:", i, "Hammer:", Hammer(i))
	}
	fmt.Println("i:", 10000, "Hammer:", Hammer(10000))
}
