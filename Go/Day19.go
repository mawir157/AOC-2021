package main

import AH "./adventhelper"

import (
	"strconv"
	"strings"
)

type Pos struct {
	x int
	y int
	z int
}

func diff(p1 Pos, p2 Pos) (Pos) {
	return Pos{x:(p1.x - p2.x), y:(p1.y - p2.y), z:(p1.z - p2.z)}
}

type Ori int
const (
	s00 Ori = iota
	s10
	s20
	s30
	s40
	s50
	s60
	s70
	s01
	s11
	s21
	s31
	s41
	s51
	s61
	s71
	s02
	s12
	s22
	s32
	s42
	s52
	s62
	s72
	ERR
)

var ORIENTATIONS = []Ori{
s00, s10, s20, s30, s40, s50, s60, s70,
s01, s11, s21, s31, s41, s51, s61, s71,
s02, s12, s22, s32, s42, s52, s62, s72}

type Space []Pos

func norm(s Space, d Pos) (t Space) {
	t = Space{}
	for _, p := range s {
		t = append(t, diff(p, d))
	}

	return
}

// rotate then norm
func norm2(s Space, r Ori, d Pos) (t Space) {
	t = Space{}
	for _, p := range s {
		t = append(t, diff(rot(p, r), d))
	}
	return
}

func rot(p Pos, r Ori) Pos {
	switch r {
	case s00:
		return Pos{p.x, p.y, p.z}
	case s10:
		return Pos{p.x, -p.z, p.y}
	case s20:
		return Pos{p.x, -p.y, -p.z}
	case s30:
		return Pos{p.x, p.z, -p.y}
	case s40:
		return Pos{-p.x, p.y, -p.z}
	case s50:
		return Pos{p.z, p.y, -p.x}
	case s60:
		return Pos{-p.x, -p.y, p.z}
	case s70:
		return Pos{-p.x, -p.z, -p.y}

	case s01:
		return Pos{p.y, p.z, p.x}
	case s11:
		return Pos{-p.z, p.y, p.x}
	case s21:
		return Pos{-p.y, -p.z, p.x}
	case s31:
		return Pos{p.z, -p.y, p.x}
	case s41:
		return Pos{p.y, -p.z, -p.x}
	case s51:
		return Pos{p.y, -p.x, p.z}
	case s61:
		return Pos{-p.y, p.z, -p.x}
	case s71:
		return Pos{-p.z, -p.y, -p.x}

	case s02:
		return Pos{p.z, p.x, p.y}
	case s12:
		return Pos{p.y, p.x, -p.z}
	case s22:
		return Pos{-p.z, p.x, -p.y}
	case s32:
		return Pos{-p.y, p.x, p.z}
	case s42:
		return Pos{-p.z, -p.x, p.y}
	case s52:
		return Pos{-p.x, p.z, p.y}
	case s62:
		return Pos{p.z, -p.x, -p.y}
	case s72:
		return Pos{-p.y, -p.x, -p.z}

	default:
		return Pos{4, 0, 4}
	}
}

func rotate(s Space, r Ori) (t Space) {
	t = Space{}
	for _, p := range s {
		t = append(t, rot(p, r))
	}
	return
}

func overlap(s1 Space, s2 Space) int {
	set := make(map[Pos]bool)
	for _, p := range s1 {
		set[p] = true
	}
	for _, p := range s2 {
		set[p] = true
	}

	return len(s1) + len(s2) - len(set)
}

func matchPair(scan1 Space, scan2 Space) (Ori, Pos) {
	for _, o := range ORIENTATIONS {
		scan2_rot := rotate(scan2, o)
		for _, p2 := range scan2_rot {
			for _, p1 := range scan1 {
				dv := diff(p2, p1)
				scan2_rot_normed := norm(scan2_rot, dv)

				count := overlap(scan1, scan2_rot_normed)
				if count >= 12 {
					return o, diff(p2, p1)
				}
			}
		}
	}

	return ERR, Pos{4,0,4}
}


func overlap2(s1 Space, s2 Space, r Ori, d Pos) int {
	set := make(map[Pos]bool)
	for _, p := range s1 {
		set[p] = true
	}
	for _, p := range s2 {
		pp := diff(rot(p, r), d)
		set[pp] = true
	}

	return len(s1) + len(s2) - len(set)
}

func matchPair2(scan1 Space, scan2 Space) (Ori, Pos) {
	for _, o := range ORIENTATIONS {
		// scan2_rot := rotate(scan2, o)
		for _, p2 := range scan2 {
			for _, p1 := range scan1 {
				dv := diff(rot(p2, o), p1)
				// scan2_rot_normed := norm(scan2_rot, dv)

				count := overlap2(scan1, scan2, o, dv)
				if count >= 12 {
					return o, diff(p2, p1)
				}
			}
		}
	}

	return ERR, Pos{4,0,4}
}

func mrg(s Space, t Space) (st Space) {
	set := make(map[Pos]bool)
	st = Space{}

	for _, p := range s {
		set[p] = true
	}
	for _, p := range t {
		set[p] = true
	}	

	for k, _ := range set {
		st = append(st, k)
	}

	return
}

func reduce_once(ss []Space) (new []Space, done bool, loc Pos) { 
	done = false
	merged := Space{}
	new = []Space{}
	skip := 0
	s := ss[0]

	for j, t := range ss {
		if 0 != j {
			ori, pos := matchPair2(s, t)
			if ori != ERR {
				merged = mrg(s, norm2(t, ori, pos))
				loc = pos
				done = true
			}
		}
		if done { skip = j; break }
	}

	if done {
		new = append(new, merged)
		for q, s := range ss {
			if q == skip || q == 0 {
				continue
			} else {
				new = append(new, s)
			}
		}
	}

	return
}

func reduce(ss []Space) (Space, Space) {
	locs := Space{Pos{0,0,0}}
	loc := Pos{0,0,0}
	for ; len(ss) != 1;  {
		ss, _, loc = reduce_once(ss)

		locs = append(locs, loc)
	}

	return ss[0], locs
}

func maxL1Distance(s Space) (m int) {
	m = 0
	for i, p1 := range s {
		for j, p2 := range s {
			if i >= j {
				continue
			}
			t := AH.AbsInt(p1.x - p2.x) + AH.AbsInt(p1.y - p2.y) + AH.AbsInt(p1.z - p2.z)
			if t > m {
				m = t
			}
		}
	}
	return
}

func parseLine(s string, sep string) (poss Space) {
	poss = []Pos{}
	ps := strings.Split(s, sep)

	for _, p := range ps[1:] {
		ps := strings.Split(p, ",")
		x, _ := strconv.Atoi(ps[0])
		y, _ := strconv.Atoi(ps[1])
		z, _ := strconv.Atoi(ps[2])
		pos := Pos {x:x, y:y, z:z}
		poss = append(poss, pos)
	}

	return
}

func main() {
	sep := "|"
	ss, _ := AH.ParseLineGroups("../input/input19.txt", sep)
	Scanners := []Space{}
	for _, s := range ss {
		Scanners = append(Scanners, parseLine(s, sep))
	}

	stars, scanners := reduce(Scanners)

	AH.PrintSoln(19, len(stars), maxL1Distance(scanners))

	return
}

