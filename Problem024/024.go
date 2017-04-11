package main

import (
	"fmt"
	"math"
	"math/big"
)

const target = 1000000

func PermutationSize(num int) int {
	x := new(big.Int)
	return int(x.MulRange(2, int64(num)).Int64())
}

func RemoveElement(s []int, i int) []int {
	s = append(s[:i], s[i+1:]...)
	n := make([]int, len(s))
	copy(n, s)
	return n
}

func ElementsToInt(elements []int) int {
	value := 0
	for i := 0; i < len(elements); i++ {
		value += elements[i] * int(math.Pow(10, float64(len(elements)-i-1)))
	}
	return value
}

func Problem024() int {
	elements := []int{0, 1, 2, 3, 4, 5, 6, 7, 8, 9}
	answer := []int{}
	num := target - 1
	for i := len(elements); i > 0; i-- {
		size := PermutationSize(i - 1)
		index := num / size
		num = num % size
		answer = append(answer, elements[index])
		elements = RemoveElement(elements, index)
	}
	return ElementsToInt(answer)
}

func main() {
	fmt.Println(Problem024())
}
