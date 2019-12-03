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
	input := ``
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
	input := ``
	out, _ := advent6A(input)
	expected := 0
	if !cmp.Equal(out, expected) {
		t.Errorf("Didn't match %s", cmp.Diff(out, expected))
	}
	input = ``
	out, _ = advent6B(input)
	expected = 0
	if !cmp.Equal(out, expected) {
		t.Errorf("Didn't match %s", cmp.Diff(out, expected))
	}
}

func TestDay7(t *testing.T) {
	input := ``
	out, _ := advent7A(input)
	expected := ""
	if !cmp.Equal(out, expected) {
		t.Errorf("Didn't match %s", cmp.Diff(out, expected))
	}
	outInt, _ := advent7B(input)
	expectedInt := 0
	if !cmp.Equal(outInt, expectedInt) {
		t.Errorf("Didn't match %s", cmp.Diff(outInt, expectedInt))
	}
}

func TestDay8(t *testing.T) {
	input := ``
	out := advent8A(input)
	expected := 0
	if !cmp.Equal(out, expected) {
		t.Errorf("Didn't match %s", cmp.Diff(out, expected))
	}
	input = ``
	out = advent8B(input)
	expected = 0
	if !cmp.Equal(out, expected) {
		t.Errorf("Didn't match %s", cmp.Diff(out, expected))
	}
}

func TestDay9(t *testing.T) {
	input := ``
	out := advent9A(input)
	expected := 0
	if !cmp.Equal(out, expected) {
		t.Errorf("Didn't match %s", cmp.Diff(out, expected))
	}
	input = ``
	out = advent9A(input)
	expected = 0
	if !cmp.Equal(out, expected) {
		t.Errorf("Didn't match %s", cmp.Diff(out, expected))
	}
	input = ``
	out = advent9A(input)
	expected = 0
	if !cmp.Equal(out, expected) {
		t.Errorf("Didn't match %s", cmp.Diff(out, expected))
	}
	// out = advent9B(input)
	// expected = 0
	// if !cmp.Equal(out, expected) {
	// 	t.Errorf("Didn't match %s", cmp.Diff(out, expected))
	// }
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
