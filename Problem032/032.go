package main

import "fmt"

func IsPandigital(n int) bool {
	digits, count, tmp := 0, 0, 0

	for n > 0 {
		tmp = digits
		digits = digits | 1<<uint((n%10)-1)
		if tmp == digits {
			return false
		}
		count++
		n /= 10
	}
	return digits == ((1 << uint(count)) - 1)
}

func Concat(a, b int) int {
	c := b
	for c > 0 {
		a *= 10
		c /= 10
	}
	return a + b
}

func Contains(s []int, e int) bool {
	for _, a := range s {
		if a == e {
			return true
		}
	}
	return false
}

func Problem032() int {
	products := []int{}
	sum := 0
	var prod, compiled int

	for m := 2; m < 100; m++ {
		nbegin := 1234
		if m > 9 {
			nbegin = 123
		}
		nend := 10000/m + 1
		for n := nbegin; n < nend; n++ {
			prod = m * n
			compiled = Concat(Concat(prod, n), m)
			if compiled >= 1e+8 && compiled < 1e+9 && IsPandigital(compiled) {
				if !Contains(products, prod) {
					products = append(products, prod)
				}

			}
		}
	}

	for _, v := range products {
		sum += v
	}
	return sum
}

func main() {
	fmt.Println(Problem032())
}
