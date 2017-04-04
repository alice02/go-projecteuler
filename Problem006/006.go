package main

import (
	"fmt"
	"math"
)

const max = 100

func Problem006() int {
	sumSquare := int(math.Pow(((max * (1 + max)) / 2), 2))
	squareSum := max * (max + 1) * (2*max + 1) / 6
	return sumSquare - squareSum
}

func main() {
	fmt.Println(Problem006())
}
