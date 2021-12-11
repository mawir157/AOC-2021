package main

import AH "./adventhelper"

import (
	"strconv"
)

type Pos struct {
	X, Y int
}

type Octopus struct {
	Level int
	Flashed bool
}

func nbrs(p Pos) (ns []Pos) {
	ns  = []Pos{ {X:p.X - 1, Y:p.Y - 1}, {X:p.X - 1, Y:p.Y}, {X:p.X - 1, Y:p.Y + 1},
	             {X:p.X,     Y:p.Y - 1},                     {X:p.X,     Y:p.Y + 1},
	             {X:p.X + 1, Y:p.Y - 1}, {X:p.X + 1, Y:p.Y}, {X:p.X + 1, Y:p.Y + 1}}

	return 
}

func parseToOctopus(ss []string) (ocks map[Pos]*Octopus) {
	ocks = make(map[Pos]*Octopus)
	for i, s := range ss {
		for j, c := range s {
			v, _ := strconv.Atoi(string(c))
			ocks[Pos{X:i, Y:j}] = &Octopus{v, false}
		}
	}

	return
}

func octoTick(ocks map[Pos]*Octopus) (total int) {
	// step 1 increment all ocks
	for p, o := range ocks {
		o.Level += 1;
		ocks[p] = o
	}

	// flash all octopus
	flashCount := 1;
	for ; flashCount != 0; {
		flashCount = 0
		for p, o := range ocks {
			if (o.Level > 9) && (!o.Flashed) {
				// this octopus has flashed
				flashCount += 1;
				o.Flashed = true
				// find neighbours...
				ns := nbrs(p)
				// ...and increment them
				for _, npos := range ns {
					if nOck, ok := ocks[npos]; ok {
						nOck.Level += 1;
					}
				}
			} 
		}
		
		total += flashCount
	}

	// all the octopus have finished flashing
	for _, o := range ocks {
		if o.Level > 9 {
			o.Level = 0;
		}
		o.Flashed = false
	}
	return
}

func main() {
	js, _ := AH.ReadStrFile("../input/input11.txt")
	ocks:= parseToOctopus(js)

	part1 := 0
	part2 := 0
	for part2 = 0; part2 < 100; part2++ {
		part1 += octoTick(ocks)
	}

	for ; true ; {
		part2++;
		if octoTick(ocks) == len(ocks) {
			break
		}
	}

	AH.PrintSoln(11, part1, part2)

	return
}
