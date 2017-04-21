package main

import (
	"fmt"
	"math/big"
)

func Problem029() int {
	m := make(map[interface{}]bool)
	for a := 2; a <= 100; a++ {
		for b := 2; b <= 100; b++ {
			i, e := big.NewInt(int64(a)), big.NewInt(int64(b))
			i.Exp(i, e, nil)
			m[i] = true
		}
	}
	return len(m)
}

func main() {
	fmt.Println(Problem029())
}
