package main

import (
	"fmt"
	"math"
)

func Problem029() int {
	m := make(map[int]bool)
	for a := 2; a <= 100; a++ {
		for b := 2; b <= 100; b++ {
			m[int(math.Pow(float64(a), float64(b)))] = true
		}
	}
	return len(m)
}

func main() {
	fmt.Println(Problem029())
}
