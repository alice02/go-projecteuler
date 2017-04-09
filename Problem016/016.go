package main

import (
	"fmt"
	"math/big"
	"strconv"
)

func DigitSum(strNumber string) int {
	sum := 0
	for _, d := range strNumber {
		i, _ := strconv.Atoi(string(d))
		sum += i
	}
	return sum
}

func Problem016() int {
	num := new(big.Int).Exp(big.NewInt(2), big.NewInt(1000), nil).String()
	return DigitSum(num)
}

func main() {
	fmt.Println(Problem016())
}
