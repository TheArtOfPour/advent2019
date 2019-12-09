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

// Abs 6pack
func Abs(x int) int {
	if x < 0 {
		return -x
	}
	return x
}

func main() {
	fileContents, _ := ioutil.ReadFile("./inputs/day10.txt")
	input := string(fileContents)
	day10A := advent10A(input)
	fmt.Printf("10A: %v\n", day10A)
	day10B := advent10B(input)
	fmt.Printf("10B: %v\n\n", day10B)
}

func runAll() {
	fileContents, _ := ioutil.ReadFile("./inputs/day1.txt")
	input := string(fileContents)
	day1A, _ := advent1A(input)
	fmt.Printf("1A: %v\n", day1A)
	day1B, _ := advent1B(input)
	fmt.Printf("1B: %v\n\n", day1B)

	fileContents, _ = ioutil.ReadFile("./inputs/day2.txt")
	input = string(fileContents)
	day2A, _ := advent2A(input)
	fmt.Printf("2A: %v\n", day2A)
	day2B, _ := advent2B(input)
	fmt.Printf("2B: %v\n\n", day2B)

	fileContents, _ = ioutil.ReadFile("./inputs/day3.txt")
	input = string(fileContents)
	day3A, _ := advent3A(input)
	fmt.Printf("3A: %v\n", day3A)
	day3B, _ := advent3B(input)
	fmt.Printf("3B: %v\n\n", day3B)

	fileContents, _ = ioutil.ReadFile("./inputs/day4.txt")
	input = string(fileContents)
	day4A, _ := advent4A(input)
	fmt.Printf("4A: %v\n", day4A)
	day4B, _ := advent4B(input)
	fmt.Printf("4B: %v\n\n", day4B)

	fileContents, _ = ioutil.ReadFile("./inputs/day5.txt")
	input = string(fileContents)
	day5A, _ := advent5A(input)
	fmt.Printf("5A: %v\n", day5A)
	day5B, _ := advent5B(input)
	fmt.Printf("5B: %v\n\n", day5B)

	fileContents, _ = ioutil.ReadFile("./inputs/day6.txt")
	input = string(fileContents)
	day6A, _ := advent6A(input)
	fmt.Printf("6A: %v\n", day6A)
	day6B, _ := advent6B(input)
	fmt.Printf("6B: %v\n\n", day6B)

	fileContents, _ = ioutil.ReadFile("./inputs/day7.txt")
	input = string(fileContents)
	day7A, _ := advent7A(input)
	fmt.Printf("7A: %v\n", day7A)
	day7B, _ := advent7B(input)
	fmt.Printf("7B: %v\n\n", day7B)

	fileContents, _ = ioutil.ReadFile("./inputs/day8.txt")
	input = string(fileContents)
	day8A, _ := advent8A(input)
	fmt.Printf("8A: %v\n", day8A)
	day8B, _ := advent8B(input)
	fmt.Printf("8B: %v\n\n", day8B)

	fileContents, _ = ioutil.ReadFile("./inputs/day9.txt")
	input = string(fileContents)
	day9A := advent9A2(input)
	fmt.Printf("9A: %v\n", day9A)
	day9B := advent9B2(input)
	fmt.Printf("9B: %v\n\n", day9B)
}
