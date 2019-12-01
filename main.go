package main

import (
	"fmt"
	"io/ioutil"
	"sort"
)

type sortRunes []rune

func (s sortRunes) Less(i, j int) bool {
	return s[i] < s[j]
}

func (s sortRunes) Swap(i, j int) {
	s[i], s[j] = s[j], s[i]
}

func (s sortRunes) Len() int {
	return len(s)
}

// SortString sorts a string
func SortString(s string) string {
	r := []rune(s)
	sort.Sort(sortRunes(r))
	return string(r)
}

func main() {
	fileContents, _ := ioutil.ReadFile("./inputs/day2.txt")
	input := string(fileContents)
	out, _ := advent2A(input)
	fmt.Printf("Result %v\n", out)
}

func runAll() {
	fileContents, _ := ioutil.ReadFile("./inputs/day1.txt")
	input := string(fileContents)
	fmt.Printf("Day1A: %v\n", advent1A(input))
	fmt.Printf("Day1B: %v\n", advent1B(input))
}
