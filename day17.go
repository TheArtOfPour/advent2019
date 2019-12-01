package main

import (
	"fmt"
	"strconv"
	"strings"
)

func getCoords(s string) []coord {
	coords := make([]coord, 0)
	stringSlice := strings.Split(s, "\n")
	for _, s := range stringSlice {
		c := coord{X: 0, Y: 0}
		s = strings.TrimSuffix(s, "\r")
		parts := strings.Split(s, ", ")
		lockedCoord := parts[0]
		lockedParts := strings.Split(lockedCoord, "=")
		lockedValue, _ := strconv.Atoi(lockedParts[1])
		if lockedParts[0] == "x" {
			c.X = lockedValue
		} else {
			c.Y = lockedValue
		}

		freeCoord := parts[1]
		freeParts := strings.Split(freeCoord, "=")
		freeValueParts := strings.Split(freeParts[1], "..")
		freeLowerBound, _ := strconv.Atoi(freeValueParts[0])
		freeUpperBound, _ := strconv.Atoi(freeValueParts[1])
		for i := freeLowerBound; i <= freeUpperBound; i++ {
			if freeParts[0] == "x" {
				c.X = i
			} else {
				c.Y = i
			}
			coords = append(coords, c)
		}
	}
	return coords
}

func getBounds(coords []coord) (int, int, int, int) {
	minX := 1000000
	maxX := 0
	minY := 1000000
	maxY := 0
	for _, c := range coords {
		if c.X < minX {
			minX = c.X
		}
		if c.X > maxX {
			maxX = c.X
		}
		if c.Y < minY {
			minY = c.Y
		}
		if c.Y > maxY {
			maxY = c.Y
		}
	}

	return minX, maxX, minY, maxY
}

func buildScan(s string) ([][]rune, int, int, int) {

	coords := getCoords(s)

	minX, maxX, minY, maxY := getBounds(coords)

	scan := make([][]rune, maxY+1)
	for i := 0; i < len(scan); i++ {
		scan[i] = make([]rune, maxX-minX+3)
		for j := 0; j < len(scan[i]); j++ {
			scan[i][j] = '.'
		}
	}

	for _, c := range coords {
		scan[c.Y][c.X-minX+1] = '#'
	}

	startX := 500 - minX + 1

	scan[0][startX] = '+'

	return scan, minY, maxY, startX
}

func printScan(scan [][]rune) {
	fmt.Print("\n")
	for i := 0; i < len(scan); i++ {
		for j := 0; j < len(scan[i]); j++ {
			fmt.Printf("%s", string(scan[i][j]))
		}
		fmt.Print("\n")
	}
	fmt.Print("\n")
}

func runWater(scan *[][]rune, maxY int, startX int, startY int) bool {
	// look directly below start x, y
	// . -> | recurse with startY + 1
	// # -> ~
	// ~ -> nothing
	// detect if bounded?
	// current := scan[startY][startX]
	if startY == maxY {
		return false
	}

	below := (*scan)[startY+1][startX]
	if below == '.' {
		(*scan)[startY+1][startX] = '|'
		runWater(scan, maxY, startX, startY+1)
	} else if below == '~' {
		// check if bounded, if bounded proceed
	} else if below == '#' {
		(*scan)[startY][startX] = '~'

		// look left and right
		offset := 1
		yOffset := 0
		boundedLeft := false
		boundedRight := false
		pouringLeft := false
		pouringRight := false
		for true {
			// bound or pouring
			if (boundedLeft || pouringLeft) && (boundedRight || pouringRight) {
				// bound
				if boundedLeft && boundedRight {
					// reset one level up
					offset = 1
					yOffset++
					boundedLeft = false
					boundedRight = false
				} else {
					break
				}
			}

			if !boundedLeft && !pouringLeft {
				left := (*scan)[startY-yOffset][startX-offset]
				if left == '#' {
					boundedLeft = true
				} else if left == '.' || left == '|' {
					(*scan)[startY-yOffset][startX-offset] = '~'
					leftBelow := (*scan)[startY-yOffset+1][startX-offset]
					if leftBelow == '.' {
						(*scan)[startY-yOffset+1][startX-offset] = '|'
						runWater(scan, maxY, startX-offset, startY-yOffset+1)
						pouringLeft = true
					}
				}
			}

			if !boundedRight && !pouringRight {
				// debug
				if startX+offset >= len((*scan)[0]) {
					printScan((*scan))
					break
				}
				right := (*scan)[startY-yOffset][startX+offset]
				if right == '#' {
					boundedRight = true
				} else if right == '.' || right == '|' {
					(*scan)[startY-yOffset][startX+offset] = '~'
					rightBelow := (*scan)[startY-yOffset+1][startX+offset]
					if rightBelow == '.' {
						(*scan)[startY-yOffset+1][startX+offset] = '|'
						runWater(scan, maxY, startX+offset, startY-yOffset+1)
						pouringRight = true
					}
				}
			}

			offset++
		}

		if boundedLeft && boundedRight {
			// water level increases
			(*scan)[startY-1][startX] = '~'
			runWater(scan, maxY, startX, startY-1)
		}
	}
	return false
}

func getTotal(scan [][]rune, minY int, maxY int) int {
	total := 0
	for i := minY; i <= maxY; i++ {
		for j := 0; j < len(scan[i]); j++ {
			if scan[i][j] == '|' || scan[i][j] == '~' {
				total++
			}
		}
	}
	return total
}

func advent17A(test string) int {
	scan, minY, maxY, startX := buildScan(test)
	//printScan(scan)
	runWater(&scan, maxY, startX, 0)
	printScan(scan)
	return getTotal(scan, minY, maxY)
}

func advent17B(test string) int {
	return 0
}
