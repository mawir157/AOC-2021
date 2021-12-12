package main

import AH "./adventhelper"

import (
	"strings"
)

type Edge struct {
	from, to string
}

type Graph struct {
	vertices map[string]bool
	edges    []Edge
	part2    bool
	visited  map[string]int
}

func (g *Graph) addEdge (e Edge) {
	g.visited[e.from] = 0;
	g.visited[e.to] = 0;

	if e.from == strings.ToLower(e.from) {
		g.vertices[e.from] = false
	} else {
		g.vertices[e.from] = true
	}

	if e.to == strings.ToLower(e.to) {
		g.vertices[e.to] = false
	} else {
		g.vertices[e.to] = true
	}

	g.edges = append(g.edges, e)

	return
}

func (g *Graph) visit (v string) {
	if (!g.vertices[v]) {
		g.visited[v]++
		if (g.visited[v] > 1) {
			g.part2 = false;
		}
	}

	return
}

func (g Graph) adj (v string) (vs []string){
	allowed := 0
	if g.part2 {
		allowed = 1
	}

	for _, e := range g.edges {
		if e.from == v {
			if g.visited[e.to] <= allowed {
				vs = append(vs, e.to)
			}
		} else if e.to == v {
			if g.visited[e.from] <= allowed {
				vs = append(vs, e.from)
			}
		}
	}

	return
}

func buildGraph(ss []string, flag bool) (g Graph) {
	g.vertices = make(map[string]bool)
	g.visited = make(map[string]int)
	for _, s := range ss {
		parts := strings.Split(s, "-")
		e := Edge{from:parts[0], to:parts[1]}
		g.addEdge(e)
	}
	g.visit("start")
	g.visit("start")
	g.part2 = flag

	return
}

func CopyGraph(g Graph) (Graph) {
	new_vertices := make(map[string]bool)
	for k,v := range g.vertices {
		new_vertices[k] = v
	}

	new_edges := []Edge{}
	for _, v:= range g.edges {
		new_edges = append(new_edges, v)
	}

	new_visited := make(map[string]int)
	for k,v := range g.visited {
		new_visited[k] = v
	}

	return Graph{vertices:new_vertices, edges:new_edges, part2:g.part2,
	             visited:new_visited}
}

func countPaths(g Graph, from string, to string) int {
	if (from == to) {
		return 1
	}

	routes := 0
	nbrs := g.adj(from)
	for _, n := range nbrs {
		g_new := CopyGraph(g)
		g_new.visit(n)
		routes += countPaths(g_new, n, to)
	}

	return routes
}

func main() {
	js, _ := AH.ReadStrFile("../input/input12.txt")
	g1 := buildGraph(js, false)
	g2 := buildGraph(js, true)

	AH.PrintSoln(12, countPaths(g1, "start", "end"),
	                 countPaths(g2, "start", "end"))

	return
}
