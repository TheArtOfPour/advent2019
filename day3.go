package main

import (
	"strconv"
	"strings"
)

type direction string

const (
	up    direction = "U"
	down  direction = "D"
	left  direction = "L"
	right direction = "R"
)

type move struct {
	dir direction
	mag int
}

type point struct {
	x int
	y int
}
type state struct {
	point
	steps int
}

type grid map[point]map[int]int

func advent3A(s string) (int, error) {

	data := [][]move{}

	for _, l := range strings.Split(s, "\r\n") {
		s := strings.Split(l, ",")
		moves := []move{}
		for _, v := range s {
			m := move{dir: direction(v[0])}
			mag, err := strconv.Atoi(v[1:])
			if err != nil {
				panic(err)
			}
			m.mag = mag
			moves = append(moves, m)
		}
		data = append(data, moves)
	}

	g := make(grid)
	for i := 0; i < len(data); i++ {
		s0 := state{}
		for j := 0; j < len(data[i]); j++ {
			s0 = g.apply(s0, data[i][j], i)
		}
	}

	minDist := 9999999999
	minDelay := minDist
	origin := state{}
	for k, v := range g {
		if len(v) > 1 {
			dst := dist(origin.point, k)
			if dst < minDist {
				minDist = dst
			}
			delay := v[0] + v[1]
			if delay < minDelay {
				minDelay = delay
			}
		}
	}

	return minDist, nil
}

func advent3B(s string) (int, error) {

	data := [][]move{}

	for _, l := range strings.Split(s, "\r\n") {
		s := strings.Split(l, ",")
		moves := []move{}
		for _, v := range s {
			m := move{dir: direction(v[0])}
			mag, err := strconv.Atoi(v[1:])
			if err != nil {
				panic(err)
			}
			m.mag = mag
			moves = append(moves, m)
		}
		data = append(data, moves)
	}

	g := make(grid)
	for i := 0; i < len(data); i++ {
		s0 := state{}
		for j := 0; j < len(data[i]); j++ {
			s0 = g.apply(s0, data[i][j], i)
		}
	}

	minDist := 9999999999
	minDelay := minDist
	origin := state{}
	for k, v := range g {
		if len(v) > 1 {
			dst := dist(origin.point, k)
			if dst < minDist {
				minDist = dst
			}
			delay := v[0] + v[1]
			if delay < minDelay {
				minDelay = delay
			}
		}
	}

	return minDelay, nil
}

func abs(a int) int {
	if a < 0 {
		return -1 * a
	}
	return a
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func dist(a, b point) int {
	return abs(a.x-b.x) + abs(a.y-b.y)
}

func (g grid) apply(s state, m move, wire int) state {
	switch m.dir {
	case up:
		for i := 0; i < m.mag; i++ {
			s.y++
			s.steps++
			g.update(s, wire)
		}
	case down:
		for i := 0; i < m.mag; i++ {
			s.y--
			s.steps++
			g.update(s, wire)
		}
	case left:
		for i := 0; i < m.mag; i++ {
			s.x--
			s.steps++
			g.update(s, wire)
		}
	case right:
		for i := 0; i < m.mag; i++ {
			s.x++
			s.steps++
			g.update(s, wire)
		}
	}
	return s
}

func (g grid) update(s state, wire int) {
	if _, ok := g[s.point]; !ok {
		g[s.point] = make(map[int]int)
	}
	if _, ok := g[s.point][wire]; !ok {
		g[s.point][wire] = 9999999999
	}
	g[s.point][wire] = min(g[s.point][wire], s.steps)
}
