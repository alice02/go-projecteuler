package main

import "fmt"

func gcd(a, b int) int {
	if b == 0 {
		return a
	} else {
		return gcd(b, a%b)
	}
}

func lcm(a, b int) int {
	return a * b / gcd(a, b)
}

func main() {
	num := 1
	for i := 2; i <= 20; i++ {
		num = lcm(num, i)
	}
	fmt.Println(num)
}
