package main

import "fmt"

func MakeSequence(num int) []int {
	sequence := []int{num}
	for num != 1 {
		if num%2 == 0 {
			num = num / 2
		} else {
			num = 3*num + 1
		}
		sequence = append(sequence, num)
	}
	return sequence
}

func Problem014() int {
	max := 1000000
	longest := map[string]int{
		"length": 0,
		"value":  0,
	}
	for n := 1; n < max; n++ {
		sequence := MakeSequence(n)
		if len(sequence) > longest["length"] {
			longest["length"] = len(sequence)
			longest["value"] = n
		}
	}
	return longest["value"]
}

func main() {
	fmt.Println(Problem014())
}
