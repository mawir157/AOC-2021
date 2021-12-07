package main

import AH "./adventhelper"

import (
	"strconv"
	"strings"
)

func crabEnergy(cs []int) (energy1 int, energy2 int) {
	energy1 = (1 << 31)
	energy2 = (1 << 31)
	hi, lo := AH.MaxAndMin(cs)

	for i := lo; i <= hi; i++ {
		tempEnergy1 := 0
		tempEnergy2 := 0
		for _, v := range cs {
			n := AH.AbsInt(v - i)
			tempEnergy1 += n
			tempEnergy2 += (n * (n + 1)) / 2
		}

		if tempEnergy1 < energy1 {
			energy1 = tempEnergy1
		}

		if tempEnergy2 < energy2 {
			energy2 = tempEnergy2
		}
	}

	return
}

func main() {
	ss, _ := AH.ReadStrFile("../input/input07.txt")
	ps := strings.Split(ss[0], ",")
	crabs := []int{}
	for _, s := range ps {
		c, _ := strconv.Atoi(s)
		crabs = append(crabs, c)
	}

	p1, p2 := crabEnergy(crabs)

	AH.PrintSoln(7, p1, p2)

	return
}
