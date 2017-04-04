package main

import "fmt"

func Gcd(a, b int) int {
	if b == 0 {
		return a
	} else {
		return Gcd(b, a%b)
	}
}

func Lcm(a, b int) int {
	return a * b / Gcd(a, b)
}

func Problem005() int {
	num := 1
	for i := 2; i <= 20; i++ {
		num = Lcm(num, i)
	}
	return num
}

func main() {
	fmt.Println(Problem005())
}
