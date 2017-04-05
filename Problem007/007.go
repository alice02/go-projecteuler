package main

import (
	"fmt"
	"math"
)

const max = 10001

func IsPrime(number int) bool {
	if number < 0 {
		return false
	}
	if number == 0 || number == 1 {
		return false
	}
	if number == 2 {
		return true
	}
	if number%2 == 0 {
		return false
	}
	for i := 3; i <= int(math.Sqrt(float64(number))); i += 2 {
		if number%i == 0 {
			return false
		}
	}
	return true
}

func Problem007() int {
	num, count := 0, 0
	for {
		if IsPrime(num) {
			count += 1
			if count == max {
				break
			}
		}
		num++
	}
	return num
}

func main() {
	fmt.Println(Problem007())
}
