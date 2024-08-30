package main

import (
	"fmt"
)

func GetSubMatrix(matrix [][]int, i int) [][]int {
	n := len(matrix)
	result := make([][]int, n-1)
	for j := 0; j < n-1; j++ {
		result[j] = make([]int, n-1)
	}
	for row := 1; row < n; row++ {
		for col := 0; col < n; col++ {
			if col != i {
				c2 := col
				if c2 > i {
					c2--
				}
				result[row-1][c2] = matrix[row][col]
			}
		}
	}
	return result
}

func Determinant(matrix [][]int) int {
	n := len(matrix)
	if n == 1 {
		return matrix[0][0]
	}
	result := 0
	for i := 0; i < n; i++ {
		// get sub-matrix for i
		sm := GetSubMatrix(matrix, i)
		// get the sign
		sign := 1
		if i%2 > 0 {
			sign = -1
		}
		// sum the part
		result += sign * matrix[0][i] * Determinant(sm)
	}
	return result
}

// func GetSubPyramid(pyramid [][]int, i int) [][]int {
// 	n := len(pyramid)
// 	result := make([][]int, n-1)
// 	for row := 0; row < n-1; row++ {
// 		result[row] = make([]int, row+1)
// 		for col := 0; col <= row; col++ {
// 			r0 := row + 1
// 			c0 := col + i
// 			result[row][col] = pyramid[r0][c0]
// 		}
// 	}
// 	return result
// }

var cashe map[int]int

func GetHash(row, col int) int {
	return row*10000 + col
}

func GetSubSum(pyramid [][]int, row, col int) int {
	n := len(pyramid)
	if row == n-1 {
		return pyramid[row][col]
	}
	hash := GetHash(row, col)
	v, ok := cashe[hash]
	if ok {
		return v
	}
	v0 := pyramid[row][col]
	v1 := GetSubSum(pyramid, row+1, col+0)
	v2 := GetSubSum(pyramid, row+1, col+1)
	result := v0 + v1
	if v2 > v1 {
		result = v0 + v2
	}
	cashe[hash] = result
	return result
}

func LongestSlideDown(pyramid [][]int) int {
	cashe = make(map[int]int)
	v0 := pyramid[0][0]
	v1 := GetSubSum(pyramid, 1, 0)
	v2 := GetSubSum(pyramid, 1, 1)
	if v1 > v2 {
		return v0 + v1
	} else {
		return v0 + v2
	}
}

func main() {

	// Expect(Determinant([][]int{{2, 5, 3}, {1, -2, -1}, {1, 3, 4}})).To(Equal(-20))
	// test := [][]int{{2, 5, 3}, {1, -2, -1}, {1, 3, 4}}
	// result := Determinant(test)
	// fmt.Println("Determinant(test):", result)

	// Expect(LongestSlideDown([][]int{{3}, {7, 4}, {2, 4, 6}, {8, 5, 9, 3}})).To(Equal(23))
	pyramid := [][]int{{3}, {7, 4}, {2, 4, 6}, {8, 5, 9, 3}}
	//fmt.Println("GetSubPyramid(pyramid, 0):", GetSubPyramid(pyramid, 0))
	//fmt.Println("GetSubPyramid(pyramid, 1):", GetSubPyramid(pyramid, 1))
	result := LongestSlideDown(pyramid)
	fmt.Println("LongestSlideDown:", result)

	/*
			Expect(LongestSlideDown([][]int{{75}, {95, 64}, {17, 47, 82}, {18, 35, 87, 10},
		                                       {20, 4, 82, 47, 65}, {19, 1, 23, 75, 3, 34},
		                                       {88, 2, 77, 73, 7, 63, 67}, {99, 65, 4, 28, 6, 16, 70, 92},
		                                       {41, 41, 26, 56, 83, 40, 80, 70, 33},
		                                       {41, 48, 72, 33, 47, 32, 37, 16, 94, 29},
		                                       {53, 71, 44, 65, 25, 43, 91, 52, 97, 51, 14},
		                                       {70, 11, 33, 28, 77, 73, 17, 78, 39, 68, 17, 57},
		                                       {91, 71, 52, 38, 17, 14, 91, 43, 58, 50, 27, 29, 48},
		                                       {63, 66, 4, 68, 89, 53, 67, 30, 73, 16, 69, 87, 40, 31},
		                                       {4, 62, 98, 27, 23, 9, 70, 98, 73, 93, 38, 53, 60, 4, 23}})).To(Equal(1074))

	*/
	//pyramid2 := [][]int{{75}, {95, 64}, {17, 47, 82}, {18, 35, 87, 10},{20, 4, 82, 47, 65}, {19, 1, 23, 75, 3, 34},{88, 2, 77, 73, 7, 63, 67}, {99, 65, 4, 28, 6, 16, 70, 92},{41, 41, 26, 56, 83, 40, 80, 70, 33},{41, 48, 72, 33, 47, 32, 37, 16, 94, 29},{53, 71, 44, 65, 25, 43, 91, 52, 97, 51, 14},{70, 11, 33, 28, 77, 73, 17, 78, 39, 68, 17, 57},{91, 71, 52, 38, 17, 14, 91, 43, 58, 50, 27, 29, 48},{63, 66, 4, 68, 89, 53, 67, 30, 73, 16, 69, 87, 40, 31},{4, 62, 98, 27, 23, 9, 70, 98, 73, 93, 38, 53, 60, 4, 23}}
	pyramid = [][]int{{75}, {95, 64}, {17, 47, 82}, {18, 35, 87, 10},
		{20, 4, 82, 47, 65}, {19, 1, 23, 75, 3, 34},
		{88, 2, 77, 73, 7, 63, 67}, {99, 65, 4, 28, 6, 16, 70, 92},
		{41, 41, 26, 56, 83, 40, 80, 70, 33},
		{41, 48, 72, 33, 47, 32, 37, 16, 94, 29},
		{53, 71, 44, 65, 25, 43, 91, 52, 97, 51, 14},
		{70, 11, 33, 28, 77, 73, 17, 78, 39, 68, 17, 57},
		{91, 71, 52, 38, 17, 14, 91, 43, 58, 50, 27, 29, 48},
		{63, 66, 4, 68, 89, 53, 67, 30, 73, 16, 69, 87, 40, 31},
		{4, 62, 98, 27, 23, 9, 70, 98, 73, 93, 38, 53, 60, 4, 23}}

	result = LongestSlideDown(pyramid)
	fmt.Println("LongestSlideDown:", result)
}
