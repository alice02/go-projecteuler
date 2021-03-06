package main

import "fmt"

func Problem002() int {
	sum, num := 0, 0
	f := []int{1, 1}
	for num < 4000000 {
		num = f[0] + f[1]
		f[0] = f[1]
		f[1] = num
		if num%2 == 0 {
			sum += num
		}
	}
	return sum
}

func main() {
	fmt.Println(Problem002())
}
