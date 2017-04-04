package main

import "fmt"

func CalcActualDivisor(num int) []int {
	divisors := []int{}
	for i := 1; i < num; i++ {
		if num%i == 0 {
			divisors = append(divisors, i)
		}
	}
	return divisors
}

func CalcArraySum(array []int) int {
	sum := 0
	for _, i := range array {
		sum += i
	}
	return sum
}

func main() {
	sum := 0
	max := 10000
	for i := 1; i < max; i++ {
		x := CalcArraySum(CalcActualDivisor(i))
		if i != x && i < x && CalcArraySum(CalcActualDivisor(x)) == i {
			sum += i + x
		}
	}
	fmt.Println(sum)
}
