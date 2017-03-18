package main

import "fmt"

func main() {
	square_sum, sum_square := 0, 0

	for i := 1; i <= 100; i++ {
		square_sum += i * i
		sum_square += i
	}
	sum_square = sum_square * sum_square

	fmt.Println(sum_square - square_sum)
}
