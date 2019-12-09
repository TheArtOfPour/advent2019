package main

import (
	"strconv"
	"time"

	"github.com/inancgumus/screen"
)

func getCount(s string) int {
	w := 25
	h := 6
	layers := len(s) / (w * h)
	i := [][]int{}
	for m := 0; m < h; m++ {
		i = append(i, []int{})
	}
	result := 0
	zeroMin := 99999999
	for l := 0; l < layers; l++ {
		zeroCount := 0
		oneCount := 0
		twoCount := 0
		for j := l * w * h; j < (l*w*h + w*h); j++ {
			n, _ := strconv.Atoi(string(s[j]))
			if n == 0 {
				zeroCount++
			}
			if n == 1 {
				oneCount++
			}
			if n == 2 {
				twoCount++
			}
		}
		if zeroCount < zeroMin {
			zeroMin = zeroCount
			result = oneCount * twoCount
		}
	}

	return result
}

func advent8A(test string) (int, error) {
	return getCount(test), nil
}

func printImage(h int, w int, i []int) {
	screen.MoveTopLeft()
	for o := 0; o < h*w; o++ {
		if o%w == 0 {
			println()
		}
		if i[o] == 0 {
			print(string('░'))
		} else if i[o] == 1 {
			print(string('█'))
		} else {
			print(string(' '))
		}
	}
	println()
	time.Sleep(10 * time.Millisecond)
}

func advent8B(s string) (int, error) {
	w := 25
	h := 6
	d := w * h
	i := []int{}
	for k := 0; k < d; k++ {
		i = append(i, 2)
	}

	// for each layer
	screen.Clear()
	for j, v := range s {
		n, _ := strconv.Atoi(string(v))
		c := int(j % d)
		if i[c] == 2 && n != 2 {
			i[c] = n
			printImage(h, w, i)
		}
	}
	printImage(h, w, i)

	return 0, nil
}
