package main

import (
	"fmt"
	"math"
)

func InitializeFlag(number int) []bool {
	array := make([]bool, number)
	for i := range array {
		array[i] = true
	}
	array[0] = false
	array[1] = false
	return array
}

func EratosthenesSieve(number int) []int {
	primes := []int{}
	flags := InitializeFlag(number)
	for i := 2; i <= int(math.Sqrt(float64(number))); i++ {
		if flags[i] {
			for j := i * i; j < number; j += i {
				flags[j] = false
			}
		}
	}
	for i := 0; i < number; i++ {
		if flags[i] {
			primes = append(primes, i)
		}
	}
	return primes
}

func CalculateArraySum(array []int) int {
	sum := 0
	for _, i := range array {
		sum += i
	}
	return sum
}

func Problem010() int {
	primes := EratosthenesSieve(2000000)
	return CalculateArraySum(primes)
}

func main() {
	fmt.Println(Problem010())
}
