package main

import "fmt"

func main() {
	var i, num float64 = 0, 1
	for i = 40; i > 20; i-- {
		num *= i / (i - 20)
	}
	fmt.Println(num)
}
