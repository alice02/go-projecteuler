package main

import (
	"fmt"
	"math/big"
)

const limit = 1000

func Fibonacci() func() *big.Int {
	a, b := big.NewInt(0), big.NewInt(1)
	return func() *big.Int {
		a, b = b, new(big.Int).Add(a, b)
		return a
	}
}

func Problem025() int {
	f := Fibonacci()
	count := 1
	for ; ; count++ {
		num := f()
		if len(num.String()) == limit {
			break
		}
	}
	return count
}

func main() {
	fmt.Println(Problem025())
}
