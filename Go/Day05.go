package main

import AH "./adventhelper"

import (
	"strconv"
	"strings"
)

type Pos struct {
	x, y int
}

type Line struct {
	start, end Pos
}

func parseLine(s string) Line {
	ps := strings.Split(s, " -> ")
	lhs := strings.Split(ps[0], ",")
	rhs := strings.Split(ps[1], ",")

	x0, _ := strconv.Atoi(lhs[0])
	y0, _ := strconv.Atoi(lhs[1])
	x1, _ := strconv.Atoi(rhs[0])
	y1, _ := strconv.Atoi(rhs[1])

	return Line{start:Pos{x:x0, y:y0}, end:Pos{x:x1, y:y1}}
}

func addLine(m map[Pos]int, l Line, no_d bool) {

	if no_d {
		if (l.start.x != l.end.x) && (l.start.y != l.end.y) {
			return
		}
	}

	v := l.start
	if (l.start.x == l.end.x) { // y varies
		step := AH.Sign(l.end.y - l.start.y)
		for ; v.y != l.end.y; v.y += step {
			m[v] += 1
		}
	} else if (l.start.y == l.end.y) { // x varies
		step := AH.Sign(l.end.x - l.start.x)
		for ; v.x != l.end.x; v.x += step {
			m[v] += 1
		}
	} else { // x and y vary
		stepx, stepy := AH.Sign(l.end.x - l.start.x), AH.Sign(l.end.y - l.start.y)
		for ; (v.x != l.end.x) || (v.y != l.end.y) ;  {
			m[v] += 1
			v.y += stepy
			v.x += stepx
		}
	}
	m[v] += 1

	return
}

func addLines(ls []Line) (p1 int, p2 int) {
	space1 := make(map[Pos]int)
	space2 := make(map[Pos]int)
	for _, l := range ls {
		addLine(space1, l, true)
		addLine(space2, l, false)
	}

	for _, v := range space1 {
		if v > 1 {
			p1 +=1
		}
	}

	for _, v := range space2 {
		if v > 1 {
			p2 +=1
		}
	}

	return
}

func main() {
	ss, _ := AH.ReadStrFile("../input/input05.txt")
	lines := []Line{}
	for _, s := range ss {
		lines = append(lines, parseLine(s))
	}

	p1, p2 := addLines(lines)	
	AH.PrintSoln(5, p1, p2)

	return
}
