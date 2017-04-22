package main

import (
	"fmt"
	"math/big"
)

func Problem029() int {
	m := make(map[string]bool)
	for a := 2; a <= 100; a++ {
		for b := 2; b <= 100; b++ {
			bigA := big.NewInt(int64(a))
			bigB := big.NewInt(int64(b))
			m[bigA.Exp(bigA, bigB, nil).String()] = true
		}
	}
	return len(m)
}

func main() {
	fmt.Println(Problem029())
}
