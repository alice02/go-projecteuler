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

func main() {
	num := new(big.Int).MulRange(1, 100).String()
	fmt.Println(DigitSum(num))
}
