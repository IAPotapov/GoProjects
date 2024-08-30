package main

import (
	"fmt"
	"math"
	"strings"
)

func GetOppositeMap() map[string]string {
	m := make(map[string]string)
	m["EAST"] = "WEST"
	m["WEST"] = "EAST"
	m["NORTH"] = "SOUTH"
	m["SOUTH"] = "NORTH"
	return m
}

func FindOpposite(arr []string, oppositeMap map[string]string) (index int, found bool) {
	index = 0
	found = false
	for i := len(arr) - 1; i > 0; i-- {
		s := strings.ToUpper(arr[i])
		opp := oppositeMap[s]
		if arr[i-1] == opp {
			index = i - 1
			found = true
			break
		}
	}
	return
}

func DirReduc(arr []string) []string {
	result := make([]string, len(arr))
	copy(result, arr)
	oppositeMap := GetOppositeMap()

	for len(result) > 0 {
		index, found := FindOpposite(result, oppositeMap)
		if !found {
			break
		}
		r1 := result[:index]
		r2 := result[index+2:]
		r1 = append(r1, r2...)
		result = r1
	}

	return result
}

func CalculateSubarraySum(numbers []int, index int, length int) int {
	var result int = 0
	for i := 0; i < length; i++ {
		result += numbers[index+i]
	}
	return result
}

func MaximumSubarraySumIndex(numbers []int, index int) int {
	max := numbers[index]
	for i := 2; index+i <= len(numbers); i++ {
		s := CalculateSubarraySum(numbers, index, i)
		if s > max {
			max = s
		}
	}
	return max
}

func MaximumSubarraySum(numbers []int) int {
	max := 0
	for i := 0; i < len(numbers); i++ {
		s := MaximumSubarraySumIndex(numbers, i)
		if s > max {
			max = s
		}
	}
	return max
}

func Cakes(recipe, available map[string]int) int {
	var min int = 0
	var initFlag bool = false
	for name, recipeValue := range recipe {
		availableValue, found := available[name]
		if !found {
			min = 0
			break
		}
		n := availableValue / recipeValue
		if !initFlag {
			min = n
			initFlag = true
		}
		if min > n {
			min = n
		}
	}
	return min
}

func ProductFibGo(prod uint64) (Fi, FiPlus uint64, correct bool) {
	Fi = 0
	FiPlus = 1
	p := Fi * FiPlus
	correct = false
	for FiPlus < prod {
		if p == prod {
			correct = true
			break
		}
		FiPlusPlus := Fi + FiPlus
		Fi = FiPlus
		FiPlus = FiPlusPlus
		p = Fi * FiPlus
		if p > prod {
			correct = false
			break
		}
	}
	return
}

func ProductFib(prod uint64) [3]uint64 {
	Fi, FiPlus, correct := ProductFibGo(prod)
	var result [3]uint64
	result[0] = Fi
	result[1] = FiPlus
	if correct {
		result[2] = 1
	} else {
		result[2] = 0
	}
	return result
}

func PrimeFactorsI(n *int, value int) int {
	result := 0
	for *n > 1 {
		if *n%value > 0 {
			break
		}
		*n /= value
		result++
	}
	return result
}

func AppendResult(result string, prime, power int) string {
	var s string
	if power == 1 {
		s = fmt.Sprintf("(%d)", prime)
	} else {
		s = fmt.Sprintf("(%d**%d)", prime, power)
	}
	return result + s
}

func PrimeFactors(n int) string {
	result := ""
	p := PrimeFactorsI(&n, 2)
	if p > 0 {
		result = AppendResult(result, 2, p)
	}
	for i := 3; i < n/2; i += 2 {
		p = PrimeFactorsI(&n, i)
		if p > 0 {
			result = AppendResult(result, i, p)
		}
	}
	if n > 1 {
		result = AppendResult(result, n, 1)
	}
	return result
}

func Abs(value int) int {
	if value > 0 {
		return value
	} else {
		return -value
	}
}

func MaxAbs(lst []int) int {
	result := Abs(lst[0])
	for i := 0; i < len(lst); i++ {
		value := Abs(lst[i])
		if result < value {
			result = value
		}
	}
	return result
}

var primes []int
var sums map[int]int

func IsPrime(value int) bool {
	max := int(math.Sqrt(float64(value)))
	for i := 0; i < len(primes); i++ {
		p := primes[i]
		if p > max {
			break
		}
		if value%p == 0 {
			return false
		}
	}
	return true
}

func ProcessPrime(lst []int, value int) {
	sum := 0
	count := 0
	for i := 0; i < len(lst); i++ {
		if Abs(lst[i])%value == 0 {
			sum += lst[i]
			count++
		}
	}
	if count > 0 {
		primes = append(primes, value)
		sums[value] = sum
	}
}

func SumOfDivided(lst []int) string {
	sums = make(map[int]int)
	primes = make([]int, 0)
	ProcessPrime(lst, 2)
	n := MaxAbs(lst)
	for i := 3; i <= n; i += 2 {
		if IsPrime(i) {
			ProcessPrime(lst, i)
		}
	}
	var sb strings.Builder
	for i := 0; i < len(primes); i++ {
		s := fmt.Sprintf("(%d %d)", primes[i], sums[primes[i]])
		sb.WriteString(s)
	}
	return sb.String()
}

func SnailPerimeter(snaipMap [][]int, i int) []int {
	var x, y int
	result := make([]int, 0)
	n := len(snaipMap)
	min := i
	max := n - i + 1
	y = min
	for x = min; x <= max; x++ {
		result = append(result, snaipMap[y-1][x-1])
	}
	x = max
	for y = min + 1; y <= max; y++ {
		result = append(result, snaipMap[y-1][x-1])
	}
	y = max
	for x = max - 1; x >= min; x-- {
		result = append(result, snaipMap[y-1][x-1])
	}
	x = min
	for y = max - 1; y > min; y-- {
		result = append(result, snaipMap[y-1][x-1])
	}
	return result
}

func Snail(snaipMap [][]int) []int {
	result := make([]int, 0)
	n := len(snaipMap)
	iterations := n / 2
	if n%2 > 0 {
		iterations++
	}
	for i := 1; i <= iterations; i++ {
		ri := SnailPerimeter(snaipMap, i)
		result = append(result, ri...)
	}
	return result
}

func main() {
	//var test = []string{"NORTH", "SOUTH", "EAST", "WEST"}
	var test = []string{"NORTH", "SOUTH", "SOUTH", "EAST", "WEST", "NORTH", "NORTH"}
	fmt.Println("DirReduc:", DirReduc(test))
	var numbers = []int{1, 2}
	//var numbers = []int{-2, -1, -3, -4, -1, -2, -1, 100, 500}
	fmt.Println("MaximumSubarraySum:", MaximumSubarraySum(numbers))
	// Expect(Cakes(map[string]int{"flour": 500, "sugar": 200, "eggs": 1},map[string]int{"flour": 1200, "sugar": 1200, "eggs": 5, "milk": 200})).To(Equal(2))
	recipe := map[string]int{"flour": 500, "sugar": 200, "eggs": 1}
	available := map[string]int{"flour": 1200, "sugar": 1200, "eggs": 5, "milk": 200}
	fmt.Println("Cakes:", Cakes(recipe, available))

	// dotest(4895, [3]uint64{55, 89, 1})
	// fmt.Println("ProductFib:", ProductFib(4895))

	// fmt.Println("PrimeFactors:", PrimeFactors(7775460))

	// snailMap := [][]int{{1, 2, 3}, {4, 5, 6}, {7, 8, 9}}
	// fmt.Println("Snail:", Snail(snailMap))
	/*
			lst1 := []int{12, 15}
		        dotest(lst1, "(2 12)(3 27)(5 15)")

		        lst2 := []int{15,21,24,30,45}
		        dotest(lst2, "(2 54)(3 135)(5 90)(7 21)")
	*/
	lst1 := []int{12, 15}
	lst2 := []int{15, 21, 24, 30, 45}
	fmt.Println("SumOfDivided(lst1):", SumOfDivided(lst1))
	fmt.Println("SumOfDivided(lst2):", SumOfDivided(lst2))
}
