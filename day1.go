package main

import (
	"strconv"
	"strings"
)

func advent1A(test string) (int, error) {
	total := 0
	stringSlice := strings.Split(test, "\n")
	for _, s := range stringSlice {
		s = strings.TrimSuffix(s, "\r")
		i, err := strconv.Atoi(s)
		if err != nil {
			return 0, err
		}
		total += calcFuel(i)
	}
	return total, nil
}

func advent1B(test string) (int, error) {
	total := 0
	stringSlice := strings.Split(test, "\n")
	for _, s := range stringSlice {
		s = strings.TrimSuffix(s, "\r")
		i, err := strconv.Atoi(s)
		if err != nil {
			return 0, err
		}
		total += calcFuelRecursive(i)
	}
	return total, nil
}

func calcFuel(i int) int {
	fuel := int(i/3) - 2
	if fuel < 0 {
		return 0
	}
	return fuel
}

func calcFuelRecursive(i int) int {
	fuel := int(i/3) - 2
	if fuel < 0 {
		return 0
	}
	return fuel + calcFuelRecursive(fuel)
}
