package main

import AH "./adventhelper"

import (
	"strconv"
)

type Pos struct {
	x, y int
}

const Infinity = 1000000

func parseInput(ss []string) map[Pos]int {
	cave := make(map[Pos]int)

	for i, s := range ss {
		for j, r := range s {
			v, _ := strconv.Atoi(string(r))
			cave[Pos{x:i, y:j}] = v
		}
	}

	return cave
}

func expandCave(cave map[Pos]int, n int, size int) map[Pos]int {
	newCave := make(map[Pos]int)
	for i := 0; i < n; i++ {
		for j := 0; j < n; j++ {
			for pos, val := range cave {
				newVal := (val + i + j) 
				if newVal > 9 {
					newVal -= 9
				}
				newCave[Pos{x:pos.x + (i*size), y:pos.y + (j*size)}] = newVal
			}
		}
	}

	return newCave
}

func qCount(Q map[Pos]bool) (count int) {
	count = 0
	for _,v := range Q {
		if v {
			count += 1
		}
	}

	return
}

func initBoolGrid(n int) [][]bool {
	a := make([][]bool, n)
	for i := 0; i < n; i++ {
		a[i] = make([]bool, n)
	}

	return a
}

func initIntGrid(n int) [][]int {
	a := make([][]int, n)
	for i := 0; i < n; i++ {
		a[i] = make([]int, n)
	}

	return a
}

func minInMap(m [][]int, Q map[Pos]bool) (q Pos) {
	min := Infinity
	for p, _ := range Q {
		v := m[p.x][p.y]
		if v < min {
			min = v
			q = p
		}
	}

	return
}

func nbrs(p Pos, visited[][]bool, dim int) []Pos {
	ns := []Pos{}

	if p.y > 0 {
		np := Pos{p.x, p.y - 1}
		if !visited[p.x][p.y - 1] {
			ns = append(ns, np)
		}
	}

	if p.x > 0 {
		np := Pos{p.x - 1, p.y}
		if !visited[p.x - 1][p.y] {
			ns = append(ns, np)
		}
	}

	if p.y < (dim - 1) {
		np := Pos{p.x, p.y + 1}
		if !visited[p.x][p.y + 1] {
			ns = append(ns, np)
		}
	}

	if p.x < (dim - 1) {
		np := Pos{p.x + 1, p.y}
		if !visited[p.x + 1][p.y] {
			ns = append(ns, np)
		}
	}

	return ns
}

func dij(g map[Pos]int, source Pos, target Pos, dim int) int {
	dist := initIntGrid(dim)
	Flagged := make(map[Pos]bool)
	visited := initBoolGrid(dim)

	for vert, _ := range g {
		dist[vert.x][vert.y] = Infinity
	}
	dist[source.x][source.y] = 0
	Flagged[source] = true

	for ; true; {

		u := minInMap(dist, Flagged)
		delete(Flagged, u)
		visited[u.x][u.y] = true

		if (u == target) {
			return dist[u.x][u.y]
		}

		ns := nbrs(u, visited, dim)
		for _, n := range ns {
			alt := dist[u.x][u.y] + g[n]
			if alt < dist[n.x][n.y] {
				dist[n.x][n.y] = alt
			}

			Flagged[n] = true
		}
	}
	return -1
}


func main() {
	js, _ := AH.ReadStrFile("../input/input15.txt")
	cave := parseInput(js)

	target1 := Pos{x:99, y:99}
	dist1 := dij(cave, Pos{x:0, y:0}, target1, 100)

	target2 := Pos{x:499, y:499}
	bigCave := expandCave(cave, 5, 100)
	dist2 := dij(bigCave, Pos{x:0, y:0}, target2, 500)

	AH.PrintSoln(15, dist1, dist2)

	return
}
