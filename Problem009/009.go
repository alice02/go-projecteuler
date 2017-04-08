package main

import "fmt"

func Problem009() int {
	n := 1000
	largest := 0
	for a := 0; a < n/2; a++ {
		b := (2*a*n - n*n) / (2*a - 2*n)
		c := n - a - b
		if a*a+b*b == c*c {
			if num := a * b * c; largest < num {
				largest = num
			}
		}
	}
	return largest
}

func main() {
	fmt.Println(Problem009())
}
