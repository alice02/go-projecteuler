package main

import (
	"fmt"
	"math"
)

func SplitDigits(num int) []int {
	array := []int{}
	for ; num > 0; num /= 10 {
		array = append([]int{num % 10}, array...)
	}
	return array
}

func Problem030() int {
	m := []int{}
	for i := 0; i < 10; i++ {
		m = append(m, int(math.Pow(float64(i), 5.0)))
	}
	max := int(math.Pow(9, 5))*6 + 1
	answer := 0
	for i := 2; i < max; i++ {
		digits := SplitDigits(i)
		sum := 0
		for _, d := range digits {
			sum += m[d]
		}
		if sum == i {
			answer += i
		}
	}
	return answer
}

func main() {
	fmt.Println(Problem030())
}
