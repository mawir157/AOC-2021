package main

import AH "./adventhelper"

import (
	"strconv"
	"strings"
)

func parseInput (s string) (fs map[int]int) {
	ps := strings.Split(s, ",")
	temp := []int{}
	for _, p := range ps {
		f, _ := strconv.Atoi(p)
		temp = append(temp, f)
	}

	fs = make(map[int]int)
	for _, v := range temp {
		fs[v] +=1
	}

	return
}

func oneDay(fs map[int]int) (gs map[int]int) {
	gs = make(map[int]int)
	for k, v := range fs {
		if k == 0 {
			gs[6] += v
			gs[8] += v
		} else {
			gs[k-1] += v
		}
	}

	return
}

func life(fs map[int]int, n int) (map[int]int) {
	for i := 0; i < n; i++ {
		fs = oneDay(fs)
	}

	return fs
}

func sum(fs map[int]int) (t int) {
	for _, v := range fs {
		t += v
	}

	return
}

func main() {
	ss, _ := AH.ReadStrFile("../input/input06.txt")
	is1 := parseInput(ss[0])
	is2 := parseInput(ss[0])

	AH.PrintSoln(6, sum(life(is1, 80)), sum(life(is2, 256)))

	return
}
