package main

import "testing"

const epsilon float64 = 0.00000000001

func Approximately(a, b float64) bool {
	r := a - b
	if r < 0 {
		r = -r
	}
	return r < epsilon
}

func DoTest(iterations int64, expected float64, t *testing.T) {
	var result float64 = pi(iterations)
	if !Approximately(result, expected) {
		t.Fatalf("Iterations: %v, Expected: %v, Got: %v", iterations, expected, result)
	}
}

func TestPi(t *testing.T) {
	DoTest(10, 3.0418396189294032, t)
	DoTest(100, 3.1315929035585537, t)
	DoTest(1000, 3.140592653839794, t)
	DoTest(10000, 3.1414926535900345, t)
	DoTest(100000, 3.1415826535897198, t)
	DoTest(1000000, 3.1415916535897743, t)
}
