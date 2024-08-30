package main

import "fmt"

func pi(iterations int64) float64 {
	// use float64 for everithing, hope compiler use vector operations
	var result, i, x float64
	x = 1
	// We calculate two steps per iteration
	for i < float64(iterations) {
		result += (4.0 / x) - (4.0 / (x + 2.0))
		x += 4.0
		i += 2
	}
	return result
}

func main() {
	var n int64 = 1000000000
	fmt.Printf("Iterations: %v, Pi: %v"+"\n", n, pi(n))
}
