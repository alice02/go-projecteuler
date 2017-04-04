package main

import "fmt"

var (
	length = map[int]int{
		0:  0,
		1:  3, // one
		2:  3, // two
		3:  5, // three
		4:  4, // four
		5:  4, // five
		6:  3, // six
		7:  5, // seven
		8:  5, // eight
		9:  4, // nine
		10: 3, // ten
		11: 6, // eleven
		12: 6, // twelve
		13: 8, // thirteen
		14: 8, // fourteen
		15: 7, // fifteen
		16: 7, // sixteen
		17: 9, // seventeen
		18: 8, // eighteen
		19: 8, // nineteen
		20: 6, // twenty
		30: 6, // thirty
		40: 5, // forty
		50: 5, // fifty
		60: 5, // sixty
		70: 7, // seventy
		80: 6, // eighty
		90: 6, // ninety
	}
)

func CalcNumberLength(num int) int {
	if num <= 0 || num >= 1000 {
		panic("num must in range 1-999")
	}
	if num < 20 {
		return length[num]
	}
	if num >= 20 && num < 100 {
		return length[int(num/10)*10] + length[num%10]
	}
	if num%100 == 0 {
		return length[num/100] + len("hundred")
	}
	return length[num/100] + len("hundred") + len("and") + CalcNumberLength(num%100)
}

func main() {
	sum := 0
	for i := 1; i <= 999; i++ {
		sum += CalcNumberLength(i)
	}
	fmt.Println(sum + len("onethousand"))
}
