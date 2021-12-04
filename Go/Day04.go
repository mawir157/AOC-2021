package main

import AH "./adventhelper"

import (
	"strconv"
	"strings"
)

type Board struct {
	Values []int
}

func parseNumbers(s string) (nums []int) {
	lines := strings.Split(s, ",")

	for _, l := range lines {
		if len(l) == 0 { continue }
		n, _ := strconv.Atoi(l)
		nums = append(nums, n)
	}
	return
}

func buildBoard(ss []string) (b Board) {
	for _, s := range ss {
		lines := strings.Split(s, " ")
		temp := []int{}
		for _, l := range lines {
			if len(l) == 0 { continue }
			n, _ := strconv.Atoi(l)
			temp = append(temp, n)
		}
		b.Values = append(b.Values, temp...)
	}
	return
}

func parseBoards(ss []string) (bs []Board) {
	for i := 0; i < len(ss); i += 5 {
		b := buildBoard(ss[i:i+5])
		bs = append(bs, b)
	}

	return
}

func (b *Board) callNumber(n int) () {
	for i, v := range b.Values {
		if v == n {
			b.Values[i] = -1
		}
	}
}

func (b Board) winning() bool {
	for i := 0; i < len(b.Values); i += 5 {
		if sumSlice(b.Values[i:i+5]) == -5 {
			return true
		}
	}

	for i := 0; i < 5; i += 1 {
		if b.Values[i] + b.Values[i+5] + b.Values[i+10] + b.Values[i+15] +
		   b.Values[i+20] == -5 {
			return true
		}
	}
	return false
}

func (b Board) score() (score int) {
	for _, v := range b.Values {
		if v >= 0 {
			score += v
		}
	}

	return
}

func sumSlice(is []int) (t int) {
	for _, v := range is { t += v }
	return
}

func part1(bs []Board, ns []int) int {
	for _, n := range ns {
		for _, b := range bs {
			b.callNumber(n)
			if b.winning() {
				return b.score() * n
			}
		}
	}

	return 0
}

func part2(bs []Board, ns []int) int {
	slowest_n := 0
	slowest_b := -1
	for ib, b :=  range bs {
		for in, n := range ns {
			b.callNumber(n)
			if b.winning() {
				if in > slowest_n {
					slowest_n = in
					slowest_b = ib
				}
				break
			}
		}
	}

	return bs[slowest_b].score() * ns[slowest_n]
}

func main() {
	ss, _ := AH.ReadStrFile("../input/input04.txt")

	calls := parseNumbers(ss[0])
	boards := parseBoards(ss[1:])

	p1 := part1(boards, calls)
	p2 := part2(boards, calls)

	AH.PrintSoln(4, p1, p2)

	return
}
