package main

import (
	"strconv"
	"strings"
)

func advent2A(test string) (int, error) {
	var s []int
	for _, temp := range strings.Split(test, ",") {
		i, _ := strconv.Atoi(temp)
		s = append(s, i)
	}

	s[1] = 12
	s[2] = 2
	cursor := 0
	for s[cursor] != 99 {
		if s[cursor] == 1 {
			s[s[cursor+3]] = s[s[cursor+1]] + s[s[cursor+2]]
		} else if s[cursor] == 2 {
			s[s[cursor+3]] = s[s[cursor+1]] * s[s[cursor+2]]
		}
		cursor = cursor + 4
	}

	return s[0], nil
}

func advent2B(test string) (int, error) {
	target := 19690720
	var s []int
	for _, temp := range strings.Split(test, ",") {
		i, _ := strconv.Atoi(temp)
		s = append(s, i)
	}

	orig := make([]int, len(s))
	copy(orig, s)
	found := false

	for i := 0; i < 100; i++ {
		for j := 0; j < 100; j++ {
			copy(s, orig)
			s[1] = i
			s[2] = j
			cursor := 0
			for s[cursor] != 99 {
				if s[cursor] == 1 {
					s[s[cursor+3]] = s[s[cursor+1]] + s[s[cursor+2]]
				} else if s[cursor] == 2 {
					s[s[cursor+3]] = s[s[cursor+1]] * s[s[cursor+2]]
				}
				cursor = cursor + 4
			}
			if s[0] == target {
				found = true
				break
			}
		}
		if found {
			break
		}
	}

	return 100*s[1] + s[2], nil
}
