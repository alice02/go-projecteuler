package main

import "fmt"

const (
	max = 28123
)

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

func CheckAbundantNumber(num int) bool {
	if num < 1 {
		return false
	}
	return num < CalcArraySum(CalcActualDivisor(num))
}

func main() {
	abundantNumbers := []int{}
	answer := []int{0}
	for i := 1; i <= max; i++ {
		answer = append(answer, i)
		if CheckAbundantNumber(i) {
			abundantNumbers = append(abundantNumbers, i)
		}
	}

	for _, i := range abundantNumbers {
		for _, j := range abundantNumbers {
			if i+j > max {
				break
			}
			answer[i+j] = 0
		}
	}
	fmt.Println(CalcArraySum(answer))
}
