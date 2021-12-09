package main

import AH "./adventhelper"

import (
	"sort"
	"strconv"
)

type Pos struct {
	r int
	c int
}

func parseInput(ss []string) [][]int {
	grid := [][]int{}

	for _, s := range ss {
		temp := []int{}
		for _, r := range s {
			v, _ := strconv.Atoi(string(r))
			temp = append(temp, v)
		}
		grid = append(grid, temp)
	}

	return grid
}

func part1(grid [][]int) (d int, ps []Pos) {
	ps = []Pos{}
	for i, row := range grid {
		for j, value := range row {
			low := true
			if (i > 0) {
				low = low && (grid[i - 1][j] > value)
			}
			if (i < len(grid) - 1) {
				low = low && (grid[i + 1][j] > value)
			}
			if (j != 0) {
				low = low && (grid[i][j - 1] > value)
			}
			if (j < len(row) - 1) {
				low = low && (grid[i][j + 1] > value)
			}

			if low {
				d += (value + 1)
				ps = append(ps, Pos{i,j})
			}
		}
	}

	return
}

func floodFill(grid [][]int, p Pos) (size int) {
	if (p.r < 0) || (p.r >= 100) || (p.c < 0) || (p.c >= 100) {
		return 0
	}

	if grid[p.r][p.c] == 9 {
		return 0
	}

	size = 1
	grid[p.r][p.c] = 9
	size += floodFill(grid, Pos{p.r - 1, p.c})
	size += floodFill(grid, Pos{p.r + 1, p.c})
	size += floodFill(grid, Pos{p.r, p.c - 1})
	size += floodFill(grid, Pos{p.r, p.c + 1})

	return size
}

func part2(grid [][]int, lows []Pos) (int) {
	basins := []int{}
	for _, p := range lows {
		basins = append(basins, floodFill(grid, p))
	}

	sort.Sort(sort.Reverse(sort.IntSlice(basins)))
	return basins[0] * basins[1] * basins[2]
}

func main() {
	ss, _ := AH.ReadStrFile("../input/input09.txt")
	oceanFloor := parseInput(ss)

	p1, lows:= part1(oceanFloor)
	AH.PrintSoln(9, p1, part2(oceanFloor, lows))

	return
}
