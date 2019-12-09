package main

import "strings"

func advent10A(test string) int {
	for _, temp := range strings.Split(test, "\n") {
		print(temp)
	}
	return 0
}

func advent10B(test string) int {
	return 0
}
