package main

import (
	"strconv"
	"strings"
)

func advent9A(test string) int {
	s := make([]int, 999999)
	for j, temp := range strings.Split(test, ",") {
		i, _ := strconv.Atoi(temp)
		s[j] = i
	}

	return runOpcode(s, 1, -1)
}

func advent9B(test string) int {
	return 0
}
