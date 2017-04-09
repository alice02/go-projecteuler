package main

import "fmt"

func Problem015() int {
	var i, num float64 = 0, 1
	for i = 40; i > 20; i-- {
		num *= i / (i - 20)
	}
	return int(num)
}

func main() {
	fmt.Println(Problem015())
}
