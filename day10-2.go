package main

import (
	"fmt"
	"math"
	"sort"
	"strings"
)

type Slope struct {
	dx       int
	dy       int
	negative bool
}

type Coord struct {
	x int
	y int
}

func Abs32(x int) int {
	if x < 0 {
		x = -x
	}
	return x
}

func Gcd(x, y int) int {
	for y != 0 {
		x, y = y, x%y
	}
	return x
}

func (p Coord) distance(q Coord) float64 {
	return math.Sqrt(math.Pow(float64(p.x-q.x), 2) + math.Pow(float64(p.y-q.y), 2))

}

func (s *Slope) shorten() {
	gcds := Abs32(Gcd(s.dx, s.dy))
	s.dx /= gcds
	s.dy /= gcds
}

func (s *Slope) toAngle() float64 {
	angle := math.Atan2(float64(s.dy), float64(s.dx))/(2*math.Pi)*360.0 + 90
	if angle < 0 {
		angle += 360
	}
	return angle
}

func makeGrid(input []string) [][]int {
	rows, cols := len(input), len(input)
	spaceMap := make([][]int, rows)
	for i := 0; i < rows; i++ {
		spaceMap[i] = make([]int, cols)
	}
	for y := 0; y < rows; y++ {
		for x := 0; x < cols; x++ {
			if input[y][x] == 35 {
				spaceMap[y][x] = 1
			}
		}
	}
	return spaceMap
}

func printGrid(input [][]int) {
	rows, cols := len(input), len(input[0])
	for y := 0; y < rows; y++ {
		for x := 0; x < cols; x++ {
			if input[y][x] == 1 {
				fmt.Print("#")
			} else {
				fmt.Print(".")
			}
		}
		fmt.Println()
	}
}

func countAsteroids(grid [][]int, posX, posY int) map[Slope][]Coord {
	rows, cols := len(grid), len(grid[0])
	dirs := make(map[Slope][]Coord)
	for y := 0; y < rows; y++ {
		for x := 0; x < cols; x++ {
			if grid[y][x] == 1 && !(posX == x && posY == y) {
				var dir Slope
				dir.dx = x - posX
				dir.dy = y - posY
				(&dir).shorten()
				newPos := Coord{x, y}
				dirs[dir] = append(dirs[dir], newPos)
			}
		}
	}
	return dirs
}

func part1(grid [][]int) (int, Coord) {
	maxAsteroids := 0
	var maxX, maxY int
	rows, cols := len(grid), len(grid[0])
	for y := 0; y < rows; y++ {
		for x := 0; x < cols; x++ {
			if grid[y][x] == 1 {
				currAsteroids := len(countAsteroids(grid, x, y))
				if currAsteroids > maxAsteroids {
					maxAsteroids = currAsteroids
					maxX = x
					maxY = y
				}
			}
		}
	}
	return maxAsteroids, Coord{maxX, maxY}
}

func advent10B2(test string) int {
	turret := Coord{x: 17, y: 22}
	target := 200
	grid := [][]int{}
	for i, temp := range strings.Split(test, "\n") {
		temp = strings.TrimSpace(temp)
		grid = append(grid, []int{})
		for j := 0; j < len(temp); j++ {
			if string(temp[j]) == "#" {
				grid[i] = append(grid[i], 1)
			} else {
				grid[i] = append(grid[i], 0)
			}
		}
	}
	slopeAsteroidsMap := countAsteroids(grid, turret.x, turret.y)
	anglesAsteroidsMap := make(map[float64][]Coord)
	angles := make([]float64, 0)
	for slope := range slopeAsteroidsMap {
		angle := (&slope).toAngle()
		anglesAsteroidsMap[angle] = slopeAsteroidsMap[slope]
		sort.Slice(anglesAsteroidsMap[angle], func(p, q int) bool {
			return anglesAsteroidsMap[angle][p].distance(turret) < anglesAsteroidsMap[angle][q].distance(turret)
		})
		if len(anglesAsteroidsMap[angle]) >= 1 {
			angles = append(angles, angle)
		}
		sort.Float64s(angles)
	}
	i := 0
	for i < target {
		for _, a := range angles {
			if len(anglesAsteroidsMap[a]) > 0 {
				i++
				if i == target {
					fmt.Println(anglesAsteroidsMap[a][0])
					return anglesAsteroidsMap[a][0].x*100 + anglesAsteroidsMap[a][0].y
				}
				anglesAsteroidsMap[a] = anglesAsteroidsMap[a][1:]
			}
		}
	}
	return 0
}
