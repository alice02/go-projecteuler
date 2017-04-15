package main

import (
	"fmt"
	"math/big"
)

const limit = 1000

func CalcRepeatendLength(num int) int {
	for num%2 == 0 {
		num /= 2
	}
	for num%5 == 0 {
		num /= 5
	}
	if num == 1 {
		return 0
	}

	c := 1
	m := big.NewInt(int64(num))
	n := big.NewInt(10)
	s := new(big.Int)
	ratio := big.NewInt(10)
	for {
		_, s := big.NewInt(10).DivMod(n, m, s)
		if s.Cmp(big.NewInt(1)) == 0 {
			break
		}
		n.Mul(n, ratio)
		c++
	}
	return c
}

func Problem026() int {
	maxSequence, denominator := 0, 0
	for i := limit - 1; i > 2; i-- {
		length := CalcRepeatendLength(i)
		if maxSequence < length {
			maxSequence = length
			denominator = i
		}
	}
	return denominator
}

func main() {
	fmt.Println(Problem026())
}
