package main

import (
	"fmt"
	"strconv"
	"strings"
)

var debug bool

func advent5A(test string) (int, error) {
	var s []int
	output := 0
	for _, temp := range strings.Split(test, ",") {
		i, _ := strconv.Atoi(temp)
		s = append(s, i)
	}

	cursor := 0
	for s[cursor] != 99 {
		var p1 int
		var p2 int
		if s[cursor+1] < len(s) && s[cursor+1] >= 0 {
			p1 = s[s[cursor+1]]
		}
		if s[cursor+2] < len(s) && s[cursor+2] >= 0 {
			p2 = s[s[cursor+2]]
		}
		if s[cursor] > 99 {
			// multi mode
			t := int(s[cursor] / 100)
			if int((t%100)/10) == 1 {
				p2 = s[cursor+2]
			}
			if t%10 == 1 {
				p1 = s[cursor+1]
			}
		}
		if s[cursor]%100 == 1 {
			// add
			s[s[cursor+3]] = p1 + p2
			println(fmt.Sprintf("\t%d %d+%d->%d (%d)", s[cursor], p1, p2, s[cursor+3], s[s[cursor+3]]))
			cursor = cursor + 4
		} else if s[cursor]%100 == 2 {
			// multiply
			s[s[cursor+3]] = p1 * p2
			println(fmt.Sprintf("\t%d %d*%d->%d (%d)", s[cursor], p1, p2, s[cursor+3], s[s[cursor+3]]))
			cursor = cursor + 4
		} else if s[cursor]%100 == 3 {
			// input to location
			println(fmt.Sprintf("\t%d<-%d", s[cursor], s[cursor+1]))
			s[s[cursor+1]] = 1 // for AC unit
			cursor = cursor + 2
		} else if s[cursor]%100 == 4 {
			// output from location
			println(fmt.Sprintf("\t%d->%d", s[cursor], p1))
			println(fmt.Sprintf("!%d", p1))
			output = p1
			cursor = cursor + 2
		}
	}

	return output, nil
}

func advent5B(test string) (int, error) {
	var s []int
	for _, temp := range strings.Split(test, ",") {
		i, _ := strconv.Atoi(temp)
		s = append(s, i)
	}

	return runOpcode(s, 0, 5), nil
}

func runOpcode(opcodes []int, input int, phase int) int {
	debug = true
	s := make([]int, len(opcodes))
	_ = copy(s, opcodes)
	output := 0
	cursor := 0
	relative := 0
	phaseSetting := true
	if phase == -1 {
		phaseSetting = false
	}
	for s[cursor] != 99 {
		var p1 int
		var p2 int
		var p3 int
		if s[cursor+1] < len(s) && s[cursor+1] >= 0 {
			p1 = s[s[cursor+1]]
		}
		if s[cursor+2] < len(s) && s[cursor+2] >= 0 {
			p2 = s[s[cursor+2]]
		}
		if s[cursor+3] < len(s) && s[cursor+3] >= 0 {
			p3 = s[cursor+3]
		}
		if s[cursor] > 99 {
			// multi mode
			t := int(s[cursor] / 100)
			// position mode
			if int((t%100)/10) == 1 {
				p2 = s[cursor+2]
			}
			if t%10 == 1 {
				p1 = s[cursor+1]
			}
			// relative mode
			if int((t%1000)/100) == 2 {
				p3 = relative + s[cursor+2]
			}
			if int((t%100)/10) == 2 {
				p2 = s[relative+s[cursor+2]]
			}
			if t%10 == 2 {
				p1 = s[relative+s[cursor+1]]
			}
		}
		if s[cursor]%100 == 1 {
			// add
			s[p3] = p1 + p2
			if debug {
				println(fmt.Sprintf("\t%d %d+%d->%d (%d)", s[cursor], p1, p2, p3, s[p3]))
			}
			cursor = cursor + 4
		} else if s[cursor]%100 == 2 {
			// multiply
			s[p3] = p1 * p2
			if debug {
				println(fmt.Sprintf("\t%d %d*%d->%d (%d)", s[cursor], p1, p2, p3, s[p3]))
			}
			cursor = cursor + 4
		} else if s[cursor]%100 == 3 {
			// input to location
			if debug {
				println(fmt.Sprintf("\t%d<-%d", s[cursor], p1))
			}
			if phaseSetting {
				s[p1] = phase
				phaseSetting = false
			} else {
				// todo: fix
				s[1000] = input
			}
			cursor = cursor + 2
		} else if s[cursor]%100 == 4 {
			// output from location
			if debug {
				println(fmt.Sprintf("\t%d->%d", s[cursor], p1))
				println(fmt.Sprintf("%d", p1))
			}
			output = p1
			cursor = cursor + 2
		} else if s[cursor]%100 == 5 {
			// jump if true
			if debug {
				println(fmt.Sprintf("\t%d %dT>%d", s[cursor], p1, p2))
			}
			if p1 != 0 {
				cursor = p2
			} else {
				cursor = cursor + 3
			}
		} else if s[cursor]%100 == 6 {
			// jump if false
			if debug {
				println(fmt.Sprintf("\t%d %dF>%d", s[cursor], p1, p2))
			}
			if p1 == 0 {
				cursor = p2
			} else {
				cursor = cursor + 3
			}
		} else if s[cursor]%100 == 7 {
			// less than
			if debug {
				println(fmt.Sprintf("\t%d %d<%d->%d", s[cursor], p1, p2, p3))
			}
			if p1 < p2 {
				s[p3] = 1
			} else {
				s[p3] = 0
			}
			cursor = cursor + 4
		} else if s[cursor]%100 == 8 {
			// equal to
			if debug {
				println(fmt.Sprintf("\t%d %d>%d->%d", s[cursor], p1, p2, p3))
			}
			if p1 == p2 {
				s[p3] = 1
			} else {
				s[p3] = 0
			}
			cursor = cursor + 4
		} else if s[cursor]%100 == 9 {
			// set relative
			relative = relative + p1
			if debug {
				println(fmt.Sprintf("\t%d %dR>%d", s[cursor], p1, relative))
			}
			cursor = cursor + 2
		}
	}
	return output
}
