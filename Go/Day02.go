package main

import AH "./adventhelper"

import (
	"strings"
	"strconv"
)

type Instruction struct {
	Dir   string
	Value int
}

type Pos struct {
	V int
	H int
	Aim int
}

func parseLine(s string) Instruction {
	parts := strings.Split(s, " ")
	value, _ := strconv.Atoi(parts[1])

	return Instruction{Dir:parts[0], Value:value}
}

func (p *Pos) apply(i Instruction) () {
	switch i.Dir {
		case "forward": p.H += i.Value
		case "up":      p.V -= i.Value
		case "down":    p.V += i.Value

		default: // do nothing
	}	

	return
}

func (p *Pos) applyWithAim(i Instruction) () {
	switch i.Dir {
		case "forward":
			p.H += i.Value
			p.V += p.Aim * i.Value
		case "up": p.Aim -= i.Value
		case "down": p.Aim += i.Value

		default: // do nothing
	}	

	return
}

func main() {
	ss, _ := AH.ReadStrFile("../input/input02.txt")

	p1 := Pos{H:0, V:0}
	p2 := Pos{H:0, V:0}
	for _, s := range ss {
		i := parseLine(s)
		p1.apply(i)
		p2.applyWithAim(i)
	}

	AH.PrintSoln(2, p1.H * p1.V, p2.H * p2.V)

	return
}
