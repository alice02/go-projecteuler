package main

import "fmt"

func main() {
	value := 600851475143
	factors := []int{}
	i := 2
	for value > 1 {
		if value%i == 0 {
			factors = append(factors, i)
			value = value / i
		} else {
			i += 1
		}
	}
	fmt.Println(factors[len(factors)-1])
}
