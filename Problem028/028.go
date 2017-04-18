package main

import "fmt"

func Problem028() int {
	sum := 1
	target := 1
	for i := 3; i <= 1001; i += 2 {
		d := i - 1
		for j := 0; j < 4; j++ {
			target += d
			sum += target
		}
	}
	return sum
}

func main() {
	fmt.Println(Problem028())
}
