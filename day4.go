package main

import (
	"strconv"
	"strings"
)

func advent4A(test string) (int, error) {
	result := 0
	s := strings.Split(test, "-")
	start, _ := strconv.Atoi(s[0])
	end, _ := strconv.Atoi(s[1])
	for i := start; i <= end; i++ {
		adjacent := false
		increase := false
		p1 := int(i / 100000)
		p2 := int((i % 100000) / 10000)
		p3 := int((i % 10000) / 1000)
		p4 := int((i % 1000) / 100)
		p5 := int((i % 100) / 10)
		p6 := i % 10
		//println(fmt.Sprintf("%d %d %d %d %d %d %d", i, p1, p2, p3, p4, p5, p6))
		if p6 >= p5 && p5 >= p4 && p4 >= p3 && p3 >= p2 && p2 >= p1 {
			increase = true
		}
		if p1 == p2 || p2 == p3 || p3 == p4 || p4 == p5 || p5 == p6 {
			adjacent = true
		}
		if adjacent && increase {
			result = result + 1
		}
	}
	return result, nil
}

func advent4B(test string) (int, error) {
	result := 0
	s := strings.Split(test, "-")
	start, _ := strconv.Atoi(s[0])
	end, _ := strconv.Atoi(s[1])
	for i := start; i <= end; i++ {
		adjacent := false
		increase := false
		exclusive := false
		p1 := int(i / 100000)
		p2 := int((i % 100000) / 10000)
		p3 := int((i % 10000) / 1000)
		p4 := int((i % 1000) / 100)
		p5 := int((i % 100) / 10)
		p6 := i % 10
		//println(fmt.Sprintf("%d %d %d %d %d %d %d", i, p1, p2, p3, p4, p5, p6))
		if p6 >= p5 && p5 >= p4 && p4 >= p3 && p3 >= p2 && p2 >= p1 {
			increase = true
		}
		if p1 == p2 || p2 == p3 || p3 == p4 || p4 == p5 || p5 == p6 {
			adjacent = true
		}
		if (p1 == p2 && p2 != p3) ||
			(p2 == p3 && p3 != p4 && p1 != p2) ||
			(p3 == p4 && p4 != p5 && p2 != p3) ||
			(p4 == p5 && p5 != p6 && p3 != p4) ||
			(p5 == p6 && p4 != p5) {
			exclusive = true
		}
		if adjacent && increase && exclusive {
			result = result + 1
		}
	}
	return result, nil
}
