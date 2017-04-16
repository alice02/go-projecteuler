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
	if number <= 1 {
		return []int{}
	}
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

func IsPrime(number int, primes []int) bool {
	n := int(math.Abs(float64(number)))
	for _, p := range primes {
		if p == n {
			return true
		}
	}
	return false
}

func CountPrimeSequence(a, b int) int {
	n := 0
	primes := EratosthenesSieve(1000)
	for {
		if n > b {
			break
		}
		x := n*n + a*n + b
		if IsPrime(x, primes) {
			n += 1
		} else {
			break
		}
	}
	return n
}

func Problem027() int {
	max := 0
	answer := 0
	primes := EratosthenesSieve(1000)
	for b := -999; b < 999; b++ {
		if b%2 == 0 {
			continue
		}
		if !IsPrime(b, primes) {
			continue
		}
		for a := -b + 1; a < 999; a++ {
			n := CountPrimeSequence(a, b)
			if max < n {
				max = n
				answer = a * b
			}
		}
	}
	return answer
}

func main() {
	fmt.Println(Problem027())
}
