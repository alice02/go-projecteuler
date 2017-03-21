package main

import "fmt"

func main() {
	n := 1000
	largetst := 0
	for a := 0; a < 1000; a++ {
		b := (2*a*n - n*n) / (2*a - 2*n)
		c := n - a - b
		if a*a+b*b == c*c {
			if num := a * b * c; largetst < num {
				largetst = num
			}
		}
	}
	fmt.Println(largetst)
}
