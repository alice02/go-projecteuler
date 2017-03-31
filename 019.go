package main

import (
	"fmt"
	"time"
)

func main() {
	count := 0
	for year := 1901; year <= 2000; year++ {
		for month := 1; month <= 12; month++ {
			date := time.Date(year, time.Month(month), 1, 0, 0, 0, 0, time.Local)
			if date.Weekday().String() == "Sunday" {
				count += 1
			}
		}
	}
	fmt.Println(count)
}
