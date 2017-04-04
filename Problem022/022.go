package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
	"strings"
)

func LoadData(filename string) string {
	var fp *os.File
	var err error

	fp, err = os.Open(filename)
	if err != nil {
		panic(err)
	}
	defer fp.Close()

	var rawString string
	scanner := bufio.NewScanner(fp)
	for scanner.Scan() {
		rawString = scanner.Text()
	}
	return rawString
}

func ParseText(input string) []string {
	return strings.Split(input, ",")
}

func SortStringArray(array []string) []string {
	sort.Strings(array)
	return array
}

func CalcScore(word string, rank int) int {
	score := 0
	length := len(word)
	for _, c := range word[1 : length-1] {
		score += int(c - 'A' + 1)
	}
	return score * rank
}

func main() {
	data := ParseText(LoadData("names.txt"))
	sortedWords := SortStringArray(data)
	sum := 0
	for i, word := range sortedWords {
		sum += CalcScore(word, i+1)
	}
	fmt.Println(sum)
}
