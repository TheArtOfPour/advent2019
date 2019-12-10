package main

import (
	"fmt"
	"strings"
)

func advent11A(test string) int {
	for _, temp := range strings.Split(test, "\n") {
		println(fmt.Sprintf("%v", temp))
	}
	return 0
}

func advent11B(test string) int {
	return 0
}
