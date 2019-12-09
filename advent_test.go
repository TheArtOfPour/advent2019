package main

import (
	"testing"

	"github.com/google/go-cmp/cmp"
)

func TestDay1(t *testing.T) {
	input := `12`
	out, _ := advent1A(input)
	expected := 2
	if !cmp.Equal(out, expected) {
		t.Errorf("Didn't match %s", cmp.Diff(out, expected))
	}
	input = `14`
	out, _ = advent1A(input)
	expected = 2
	if !cmp.Equal(out, expected) {
		t.Errorf("Didn't match %s", cmp.Diff(out, expected))
	}
	input = `1969`
	out, _ = advent1A(input)
	expected = 654
	if !cmp.Equal(out, expected) {
		t.Errorf("Didn't match %s", cmp.Diff(out, expected))
	}
	input = `100756`
	out, _ = advent1A(input)
	expected = 33583
	if !cmp.Equal(out, expected) {
		t.Errorf("Didn't match %s", cmp.Diff(out, expected))
	}
	input = `14`
	out, _ = advent1B(input)
	expected = 2
	if !cmp.Equal(out, expected) {
		t.Errorf("Didn't match %s", cmp.Diff(out, expected))
	}
	input = `1969`
	out, _ = advent1B(input)
	expected = 966
	if !cmp.Equal(out, expected) {
		t.Errorf("Didn't match %s", cmp.Diff(out, expected))
	}
	input = `100756`
	out, _ = advent1B(input)
	expected = 50346
	if !cmp.Equal(out, expected) {
		t.Errorf("Didn't match %s", cmp.Diff(out, expected))
	}
}

func TestDay2(t *testing.T) {
	input := `1,9,10,3,2,3,11,0,99,30,40,50`
	out, _ := advent2A(input)
	expected := 3500
	if !cmp.Equal(out, expected) {
		t.Errorf("Didn't match %s", cmp.Diff(out, expected))
	}
}

// should be 731
func TestDay3(t *testing.T) {
	input := `R20,U5,L17,D3
U7,R6,D4,L4`
	out, _ := advent3A(input)
	expected := 6
	if !cmp.Equal(out, expected) {
		t.Errorf("Didn't match %s", cmp.Diff(out, expected))
	}
	input = `L17,U5,R20,D3
U7,R6,D20,L4,U16`
	out, _ = advent3A(input)
	expected = 5
	if !cmp.Equal(out, expected) {
		t.Errorf("Didn't match %s", cmp.Diff(out, expected))
	}
	input = `R75,D30,R83,U83,L12,D49,R71,U7,L72
U62,R66,U55,R34,D71,R55,D58,R83`
	out, _ = advent3A(input)
	expected = 159
	if !cmp.Equal(out, expected) {
		t.Errorf("Didn't match %s", cmp.Diff(out, expected))
	}
	input = `R98,U47,R26,D63,R33,U87,L62,D20,R33,U53,R51
U98,R91,D20,R16,D67,R40,U7,R15,U6,R7`
	out, _ = advent3A(input)
	expected = 135
	if !cmp.Equal(out, expected) {
		t.Errorf("Didn't match %s", cmp.Diff(out, expected))
	}
	input = `R75,D30,R83,U83,L12,D49,R71,U7,L72
U62,R66,U55,R34,D71,R55,D58,R83`
	out, _ = advent3B(input)
	expected = 610
	if !cmp.Equal(out, expected) {
		t.Errorf("Didn't match %s", cmp.Diff(out, expected))
	}
	input = `R98,U47,R26,D63,R33,U87,L62,D20,R33,U53,R51
U98,R91,D20,R16,D67,R40,U7,R15,U6,R7`
	out, _ = advent3B(input)
	expected = 410
	if !cmp.Equal(out, expected) {
		t.Errorf("Didn't match %s", cmp.Diff(out, expected))
	}
}

func TestDay4(t *testing.T) {
	input := ``
	out, _ := advent4A(input)
	expected := 0
	if !cmp.Equal(out, expected) {
		t.Errorf("Didn't match %s", cmp.Diff(out, expected))
	}
	out, _ = advent4B(input)
	expected = 0
	if !cmp.Equal(out, expected) {
		t.Errorf("Didn't match %s", cmp.Diff(out, expected))
	}
}

func TestDay5(t *testing.T) {
	input := `3,21,1008,21,8,20,1005,20,22,107,8,21,20,1006,20,31,1106,0,36,98,0,0,1002,21,125,20,4,20,1105,1,46,104,999,1105,1,46,1101,1000,1,20,4,20,1105,1,46,98,99`
	out, _ := advent5A(input)
	expected := 0
	if !cmp.Equal(out, expected) {
		t.Errorf("Didn't match %s", cmp.Diff(out, expected))
	}
	out, _ = advent5B(input)
	expected = 0
	if !cmp.Equal(out, expected) {
		t.Errorf("Didn't match %s", cmp.Diff(out, expected))
	}
}

func TestDay6(t *testing.T) {
	input := `COM)B
B)C
C)D
D)E
E)F
B)G
G)H
D)I
E)J
J)K
K)L`
	out, _ := advent6A(input)
	expected := 42
	if !cmp.Equal(out, expected) {
		t.Errorf("Didn't match %s", cmp.Diff(out, expected))
	}
	input = `COM)B
B)C
C)D
D)E
E)F
B)G
G)H
D)I
E)J
J)K
K)L
K)YOU
I)SAN`
	out, _ = advent6B(input)
	expected = 4
	if !cmp.Equal(out, expected) {
		t.Errorf("Didn't match %s", cmp.Diff(out, expected))
	}
}

func TestDay7(t *testing.T) {
	input := `3,15,3,16,1002,16,10,16,1,16,15,15,4,15,99,0,0`
	out, _ := advent7A(input)
	expected := 43210
	if !cmp.Equal(out, expected) {
		t.Errorf("Didn't match %s", cmp.Diff(out, expected))
	}
	input = `3,23,3,24,1002,24,10,24,1002,23,-1,23,101,5,23,23,1,24,23,23,4,23,99,0,0`
	out, _ = advent7A(input)
	expected = 54321
	if !cmp.Equal(out, expected) {
		t.Errorf("Didn't match %s", cmp.Diff(out, expected))
	}
	input = `3,31,3,32,1002,32,10,32,1001,31,-2,31,1007,31,0,33,1002,33,7,33,1,33,31,31,1,32,31,31,4,31,99,0,0,0`
	out, _ = advent7A(input)
	expected = 65210
	if !cmp.Equal(out, expected) {
		t.Errorf("Didn't match %s", cmp.Diff(out, expected))
	}
	input = `3,26,1001,26,-4,26,3,27,1002,27,2,27,1,27,26,27,4,27,1001,28,-1,28,1005,28,6,99,0,0,5`
	out, _ = advent7B(input)
	expected = 139629729
	if !cmp.Equal(out, expected) {
		t.Errorf("Didn't match %s", cmp.Diff(out, expected))
	}
	input = `3,52,1001,52,-5,52,3,53,1,52,56,54,1007,54,5,55,1005,55,26,1001,54,-5,54,1105,1,12,1,53,54,53,1008,54,0,55,1001,55,1,55,2,53,55,53,4,53,1001,56,-1,56,1005,56,6,99,0,0,0,0,10`
	out, _ = advent7B(input)
	expected = 18216
	if !cmp.Equal(out, expected) {
		t.Errorf("Didn't match %s", cmp.Diff(out, expected))
	}
}

func TestDay8(t *testing.T) {
	input := ``
	out, _ := advent8A(input)
	expected := 0
	if !cmp.Equal(out, expected) {
		t.Errorf("Didn't match %s", cmp.Diff(out, expected))
	}
	input = `0222112222120000`
	out, _ = advent8B(input)
	expected = 1
	if !cmp.Equal(out, expected) {
		t.Errorf("Didn't match %s", cmp.Diff(out, expected))
	}
}

func TestDay9(t *testing.T) {
	input := `109,1,204,-1,1001,100,1,100,1008,100,16,101,1006,101,0,99`
	out := advent9A(input)
	expected := 99
	if !cmp.Equal(out, expected) {
		t.Errorf("Didn't match %s", cmp.Diff(out, expected))
	}
	input = `1102,34915192,34915192,7,4,7,99,0`
	out = advent9A(input)
	expected = 1219070632396864
	if !cmp.Equal(out, expected) {
		t.Errorf("Didn't match %s", cmp.Diff(out, expected))
	}
	input = `104,1125899906842624,99`
	out = advent9A(input)
	expected = 1125899906842624
	if !cmp.Equal(out, expected) {
		t.Errorf("Didn't match %s", cmp.Diff(out, expected))
	}
	input = ``
	out = advent9B(input)
	expected = 0
	if !cmp.Equal(out, expected) {
		t.Errorf("Didn't match %s", cmp.Diff(out, expected))
	}
}

func TestDay10(t *testing.T) {
	input := ``
	out := advent10A(input)
	expected := 0
	if !cmp.Equal(out, expected) {
		t.Errorf("Didn't match %s", cmp.Diff(out, expected))
	}
	// out = advent10B(input)
	// expected = 0
	// if !cmp.Equal(out, expected) {
	// 	t.Errorf("Didn't match %s", cmp.Diff(out, expected))
	// }
}

func TestDay11(t *testing.T) {
	unitTest := calculatePower(3, 5, 8)
	unitExpected := 4
	if !cmp.Equal(unitTest, unitExpected) {
		t.Errorf("Didn't match %s", cmp.Diff(unitTest, unitExpected))
	}
	unitTest = calculatePower(122, 79, 57)
	unitExpected = -5
	if !cmp.Equal(unitTest, unitExpected) {
		t.Errorf("Didn't match %s", cmp.Diff(unitTest, unitExpected))
	}
	unitTest = calculatePower(217, 196, 39)
	unitExpected = 0
	if !cmp.Equal(unitTest, unitExpected) {
		t.Errorf("Didn't match %s", cmp.Diff(unitTest, unitExpected))
	}
	unitTest = calculatePower(101, 153, 71)
	unitExpected = 4
	if !cmp.Equal(unitTest, unitExpected) {
		t.Errorf("Didn't match %s", cmp.Diff(unitTest, unitExpected))
	}
	// input := `18`
	// out := advent11A(input)
	// expected := 138
	// if !cmp.Equal(out, expected) {
	// 	t.Errorf("Didn't match %s", cmp.Diff(out, expected))
	// }
	// out = advent11B(input)
	// expected = 66
	// if !cmp.Equal(out, expected) {
	// 	t.Errorf("Didn't match %s", cmp.Diff(out, expected))
	// }
}

func TestDay12(t *testing.T) {
	input := ``
	out := advent12A(input)
	expected := 0
	if !cmp.Equal(out, expected) {
		t.Errorf("Didn't match %s", cmp.Diff(out, expected))
	}
	// out = advent12B(input)
	// expected = 66
	// if !cmp.Equal(out, expected) {
	// 	t.Errorf("Didn't match %s", cmp.Diff(out, expected))
	// }
}

func TestDay13(t *testing.T) {
	input := ``
	out := advent13A(input)
	expected := 0
	if !cmp.Equal(out, expected) {
		t.Errorf("Didn't match %s", cmp.Diff(out, expected))
	}
	input = ``
	out = advent13B(input)
	expected = 0
	if !cmp.Equal(out, expected) {
		t.Errorf("Didn't match %s", cmp.Diff(out, expected))
	}
}
func TestDay17(t *testing.T) {
	input := ``
	out := advent17A(input)
	expected := 0
	if !cmp.Equal(out, expected) {
		t.Errorf("Didn't match %s", cmp.Diff(out, expected))
	}
	// input = ``
	// out = advent17B(input)
	// expected = 1
	// if !cmp.Equal(out, expected) {
	// 	t.Errorf("Didn't match %s", cmp.Diff(out, expected))
	// }
}

func TestAdvent(t *testing.T) {
}
