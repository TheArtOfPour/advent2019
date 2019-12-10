package main

import (
	"fmt"
	"strings"
)

func getSeen(b [][]rune, i int, j int) int {
	debug := false
	// scan in lines
	a := make([][]rune, len(b))
	for i := 0; i < len(b); i++ {
		for j := 0; j < len(b[i]); j++ {
			a[i] = append(a[i], b[i][j])
		}
	}
	//_ = copy(a, b)
	seen := 0
	spiral := 'u'
	steps := 0
	done := false
	x := j
	y := i
	for !done {
		if steps > 2*len(a) && steps > 2*len(a[i]) {
			done = true
		}
		switch spiral {
		case 'u':
			steps = steps + 1
			for s := 0; s < steps; s++ {
				y = y - 1
				oob := false
				blocked := false
				ct := 0
				relX := x - j
				relY := y - i
				for !oob {
					ct = ct + 1
					newX := j + (ct * relX)
					newY := i + (ct * relY)
					if newX < 0 || newX >= len(a[i]) || newY < 0 || newY >= len(a) {
						oob = true
						break
					}
					if blocked {
						a[newY][newX] = '0'
					} else if a[newY][newX] == '#' {
						seen = seen + 1
						blocked = true
						a[newY][newX] = 'X'
					}
					if newX == 0 || newY == 0 {
						break
					}
				}
			}
			spiral = 'r'
			break
		case 'r':
			spiral = 'd'
			for s := 0; s < steps; s++ {
				x = x + 1
				oob := false
				blocked := false
				ct := 0
				relX := x - j
				relY := y - i
				for !oob {
					ct = ct + 1
					newX := j + (ct * relX)
					newY := i + (ct * relY)
					if newX < 0 || newX >= len(a[i]) || newY < 0 || newY >= len(a) {
						oob = true
						break
					}
					if blocked {
						a[newY][newX] = '0'
					} else if a[newY][newX] == '#' {
						seen = seen + 1
						blocked = true
						a[newY][newX] = 'X'
					}
					if newX == 0 || newY == 0 {
						break
					}
				}
			}
			break
		case 'd':
			steps = steps + 1
			spiral = 'l'
			for s := 0; s < steps; s++ {
				y = y + 1
				oob := false
				blocked := false
				ct := 0
				relX := x - j
				relY := y - i
				for !oob {
					ct = ct + 1
					newX := j + (ct * relX)
					newY := i + (ct * relY)
					if newX < 0 || newX >= len(a[i]) || newY < 0 || newY >= len(a) {
						oob = true
						break
					}
					if blocked {
						a[newY][newX] = '0'
					} else if a[newY][newX] == '#' {
						seen = seen + 1
						blocked = true
						a[newY][newX] = 'X'
					}
					if newX == 0 || newY == 0 {
						break
					}
				}
			}
			break
		case 'l':
			spiral = 'u'
			for s := 0; s < steps; s++ {
				x = x - 1
				oob := false
				blocked := false
				ct := 0
				relX := x - j
				relY := y - i
				for !oob {
					ct = ct + 1
					newX := j + (ct * relX)
					newY := i + (ct * relY)
					if newX < 0 || newX >= len(a[i]) || newY < 0 || newY >= len(a) {
						oob = true
						break
					}
					if blocked {
						a[newY][newX] = '0'
					} else if a[newY][newX] == '#' {
						seen = seen + 1
						blocked = true
						a[newY][newX] = 'X'
					}
					if newX == 0 || newY == 0 {
						break
					}
				}
			}
			break
		}
	}
	if debug {
		printAsteroids(a)
	}
	return seen
}

func getBestBase(a [][]rune) int {
	debug := false
	total := 0
	best := 0
	for i := 0; i < len(a); i++ {
		for j := 0; j < len(a[i]); j++ {
			if a[i][j] == '#' {
				total = total + 1
				curr := getSeen(a, i, j)
				if curr > best {
					best = curr
					if debug {
						println(fmt.Sprintf("New Best %v @ (%v,%v)", best, j, i))
					}
				} else if debug {
					println(fmt.Sprintf("         %v @ (%v,%v)", curr, j, i))
				}
			}
		}
	}
	println(fmt.Sprintf("%v/%v", best, total))
	return best
}

func printAsteroids(a [][]rune) {
	println()
	for i := 0; i < len(a); i++ {
		for j := 0; j < len(a[i]); j++ {
			print(string(a[i][j]))
		}
		println()
	}
	println()
}

func get200th(b [][]rune, i int, j int) {
	debug := false
	// scan in lines
	a := make([][]rune, len(b))
	for i := 0; i < len(b); i++ {
		for j := 0; j < len(b[i]); j++ {
			a[i] = append(a[i], b[i][j])
		}
	}
	//_ = copy(a, b)
	seen := 0
	spiral := 'u'
	steps := 0
	done := false
	x := j
	y := i
	for !done {
		if steps > 2*len(a) && steps > 2*len(a[i]) {
			done = true
		}
		switch spiral {
		case 'u':
			steps = steps + 1
			for s := 0; s < steps; s++ {
				y = y - 1
				if x < 0 || x >= len(a[i]) || y < 0 || y >= len(a) {
					continue
				}
				if a[y][x] == '#' {
					seen = seen + 1
				}
				if seen == 200 {
					println(fmt.Sprintf("200th: %v", x*100+y))
					done = true
					break
				}
			}
			spiral = 'r'
			break
		case 'r':
			spiral = 'd'
			for s := 0; s < steps; s++ {
				x = x + 1
				if x < 0 || x >= len(a[i]) || y < 0 || y >= len(a) {
					continue
				}
				if a[y][x] == '#' {
					seen = seen + 1
				}
				if seen == 200 {
					println(fmt.Sprintf("200th: %v", x*100+y))
					done = true
					break
				}
			}
			break
		case 'd':
			steps = steps + 1
			spiral = 'l'
			for s := 0; s < steps; s++ {
				y = y + 1
				if x < 0 || x >= len(a[i]) || y < 0 || y >= len(a) {
					continue
				}
				if a[y][x] == '#' {
					seen = seen + 1
				}
				if seen == 200 {
					println(fmt.Sprintf("200th: %v", x*100+y))
					done = true
					break
				}
			}
			break
		case 'l':
			spiral = 'u'
			for s := 0; s < steps; s++ {
				x = x - 1
				if x < 0 || x >= len(a[i]) || y < 0 || y >= len(a) {
					continue
				}
				if a[y][x] == '#' {
					seen = seen + 1
				}
				if seen == 200 {
					println(fmt.Sprintf("200th: %v", x*100+y))
					done = true
					break
				}
			}
			break
		}
	}
	if debug {
		printAsteroids(a)
	}
}

func advent10A(test string) int {
	asteroids := [][]rune{}
	for _, temp := range strings.Split(test, "\n") {
		temp = strings.TrimSpace(temp)
		asteroids = append(asteroids, []rune(temp))
	}
	printAsteroids(asteroids)
	return getBestBase(asteroids)
}

func advent10B(test string) int {
	//New Best 288 @ (17,22)
	asteroids := [][]rune{}
	for _, temp := range strings.Split(test, "\n") {
		temp = strings.TrimSpace(temp)
		asteroids = append(asteroids, []rune(temp))
	}
	printAsteroids(asteroids)
	get200th(asteroids, 17, 22)
	//1230 too high
	return 0
}
