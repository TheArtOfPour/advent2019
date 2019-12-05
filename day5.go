package main

import (
	"fmt"
	"strconv"
	"strings"
)

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
			s[s[cursor+1]] = 5 // for heat
			cursor = cursor + 2
		} else if s[cursor]%100 == 4 {
			// output from location
			println(fmt.Sprintf("\t%d->%d", s[cursor], p1))
			println(fmt.Sprintf("!%d", p1))
			output = p1
			cursor = cursor + 2
		} else if s[cursor]%100 == 5 {
			// jump if true
			println(fmt.Sprintf("\t%d %dT>%d", s[cursor], p1, p2))
			if p1 != 0 {
				cursor = p2
			} else {
				cursor = cursor + 3
			}
		} else if s[cursor]%100 == 6 {
			// jump if false
			println(fmt.Sprintf("\t%d %dF>%d", s[cursor], p1, p2))
			if p1 == 0 {
				cursor = p2
			} else {
				cursor = cursor + 3
			}
		} else if s[cursor]%100 == 7 {
			// less than
			println(fmt.Sprintf("\t%d %d<%d->%d", s[cursor], p1, p2, s[cursor+3]))
			if p1 < p2 {
				s[s[cursor+3]] = 1
			} else {
				s[s[cursor+3]] = 0
			}
			cursor = cursor + 4
		} else if s[cursor]%100 == 8 {
			// greater than
			println(fmt.Sprintf("\t%d %d>%d->%d", s[cursor], p1, p2, s[cursor+3]))
			if p1 == p2 {
				s[s[cursor+3]] = 1
			} else {
				s[s[cursor+3]] = 0
			}
			cursor = cursor + 4
		}
	}

	// 1433567 too low
	return output, nil
}
