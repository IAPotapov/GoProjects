package main

import "fmt"

type Data struct {
	list map[string]string
	A    []rune
}

var result = Data{}

// Heap's algorithm
// func Heaps(k int, d *Data) {
// 	if k == 1 {
// 		Test(d)
// 	} else {
// 		// Generate permutations with k-th unaltered
// 		Heaps(k-1, d)
// 		// Generate permutations for k-th swapped with each k-1 initial
// 		for i := 0; i < k-1; i++ {
// 			// Swap choice dependent on parity of k (even or odd)
// 			if k%2 == 0 {
// 				// zero-indexed, the k-th is at k-1
// 				d.A[i], d.A[k-1] = d.A[k-1], d.A[i]
// 			} else {
// 				d.A[0], d.A[k-1] = d.A[k-1], d.A[0]
// 			}
// 			Heaps(k-1, d)
// 		}
// 	}
// }

func Add() {
	var s string = string(result.A)
	result.list[s] = s
}

func Swap(index1, index2 int) {
	result.A[index1], result.A[index2] = result.A[index2], result.A[index1]
}

// Heap's algorithm
func Heaps(k int) {
	if k == 1 {
		Add()
	} else {
		Heaps(k - 1)
		for i := 0; i < k-1; i++ {
			if k%2 == 0 {
				Swap(i, k-1)
			} else {
				Swap(0, k-1)
			}
			Heaps(k - 1)
		}
	}
}

func Permutations(s string) []string {
	result.A = []rune(s)
	result.list = make(map[string]string)
	Heaps(len(result.A))
	p := make([]string, 0, len(result.list))
	for _, v := range result.list {
		p = append(p, v)
	}
	return p
}

type Interval struct {
	begin int
	end   int
}

func (i Interval) GetLength() int {
	return i.end - i.begin
}

func (i Interval) DoesOverlap(other Interval) bool {
	if i.end <= other.begin {
		return false
	}
	if other.end <= i.begin {
		return false
	}
	return true
}

func Merge(first, second Interval) Interval {
	result := Interval{begin: first.begin, end: first.end}
	if result.begin > second.begin {
		result.begin = second.begin
	}
	if result.end < second.end {
		result.end = second.end
	}
	return result
}

func Feed(list []Interval, value Interval) []Interval {
	var newValue Interval = value
	for i := 0; i < len(list); {
		if list[i].DoesOverlap(newValue) {
			newValue = Merge(newValue, list[i])
			// remove item with index i
			p1 := list[:i]
			p2 := list[i+1:]
			list = append(p1, p2...)
		} else {
			i++
		}
	}
	list = append(list, newValue)
	return list
}

func SumOfIntervals(intervals [][2]int) int {
	list := make([]Interval, 0)
	for _, v := range intervals {
		i := Interval{begin: v[0], end: v[1]}
		list = Feed(list, i)
	}
	var result int = 0
	for _, i := range list {
		result += i.GetLength()
	}
	return result
}

func main() {
	fmt.Println("Permutations:", Permutations("aabb"))
	// Expect(SumOfIntervals([][2]int{{1, 4}, {7, 10}, {3, 5}})).To(Equal(7))
	//test := [][2]int{{1, 4}, {7, 10}, {3, 5}}
	test := [][2]int{{10, 30}, {20, 30}, {1, 2}, {-1, 12}}
	fmt.Println("SumOfIntervals:", SumOfIntervals(test))
}
