package main

import "fmt"

func Count(target int, coins []int) int {
	if len(coins) == 0 {
		return 1
	}

	sum := 0
	coin := coins[0]
	q := (target / coin) + 1
	for i := 0; i < q; i++ {
		sum += Count(target-coin*i, coins[1:])
	}
	return sum
}

func Problem031() int {
	coins := []int{200, 100, 50, 20, 10, 5, 2}
	target := 200
	return Count(target, coins)
}

func main() {
	fmt.Println(Problem031())
}
