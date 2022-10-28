package main

import AH "./adventhelper"


import (
	"reflect"
)

type Pos struct {
	X, Y int
}

type CELL int
const (
	HOR CELL = iota
	VER
)

type SEA map[Pos]CELL

func BuildSea (ss []string) (sea SEA, x_hi int, y_hi int) {
	sea = make(SEA)
	for i, s := range ss {
		for j, c := range s {
			if c == '>' {
				sea[Pos{X:i, Y:j}] = HOR
			} else if c == 'v'{
				sea[Pos{X:i, Y:j}] = VER
			}
		}
	}

	x_hi, y_hi = len(ss), len(ss[0])
	return
}

func horStep (sea SEA, x_hi int, y_hi int) (new_sea SEA) {
	new_sea = make(SEA)
	for x := 0; x < x_hi; x++ {
		for y := 0; y < y_hi; y++ {
			p1 := Pos{X:x, Y:y}
			if v, ok := sea[p1]; ok {
				if v == HOR {
					p2 := Pos{X:x, Y:((y + 1) % y_hi)}
					_, blocked := sea[p2]
					if !blocked {
						new_sea[p2] = HOR
					} else {
						new_sea[p1] = HOR
					}
				} else {
					new_sea[p1] = v
				}
			}
		}
	}
	return
}

func verStep (sea SEA, x_hi int, y_hi int) (new_sea SEA) {
	new_sea = make(SEA)
	for x := 0; x < x_hi; x++ {
		for y := 0; y < y_hi; y++ {
			p1 := Pos{X:x, Y:y}
			if v, ok := sea[p1]; ok {
				if v == VER {
					p2 := Pos{X:((x + 1) % x_hi), Y:y}
					_, blocked  := sea[p2]
					if !blocked  {
						new_sea[p2] = VER
					} else {
						new_sea[p1] = VER
					}
				} else {
					new_sea[p1] = v
				}
			}
		}
	}

	return
}

func step(sea SEA, x_hi int, y_hi int) (new_sea SEA, eq bool) {
	new_sea = horStep(sea, x_hi, y_hi)
	new_sea = verStep(new_sea, x_hi, y_hi)

	eq = reflect.DeepEqual(sea, new_sea)

	return
}

func steps (sea SEA, x_hi int, y_hi int) (SEA, int) {
	eq, i := false, 0
	for i = 0; !eq; i++ {
		sea, eq = step(sea, x_hi, y_hi)
	}

	return sea, i
}

func main() {
	js, _ := AH.ReadStrFile("../input/input25.txt")
	sea, x, y := BuildSea(js)
	_, end := steps(sea, x, y)

	AH.PrintSoln(25, end, "I love you. See you next year")

	return
}
