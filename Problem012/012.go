package main

import (
	"fmt"
	"math"
)

func GenerateTriangularNumber(n int) int {
	return n * (n + 1) / 2
}

func CalcDivisorLength(n int) int {
	if n == 1 {
		return 1
	}
	m := int(math.Sqrt(float64(n)))
	result := 2
	i := 2
	for i <= m {
		if n%i == 0 {
			result += 2
		}
		i += 1
	}
	if i-1 == m {
		result -= 1
	}
	return result

}
func main() {
	for n := 0; n < 100000000; n++ {
		x := GenerateTriangularNumber(n)
		num := CalcDivisorLength(x)
		if num > 500 {
			fmt.Println(x)
			break
		}
	}
}
