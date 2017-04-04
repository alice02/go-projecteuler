package main

import "fmt"
import "strconv"

func IsPalindrome(num int) bool {
	s := strconv.Itoa(num)
	runes := []rune(s)
	for i, j := 0, len(runes)-1; i < j; i, j = i+1, j-1 {
		runes[i], runes[j] = runes[j], runes[i]
	}
	y, _ := strconv.Atoi(string(runes))
	if num == y {
		return true
	}
	return false

}

func Problem004() int {
	largest := 0
	for i := 1; i < 1000; i++ {
		for j := 1; j < 1000; j++ {
			num := i * j
			if IsPalindrome(num) {
				if num > largest {
					largest = num
				}
			}
		}
	}
	return largest
}

func main() {
	fmt.Println(Problem004())
}
