package main

import AH "./adventhelper"

import (
	"strconv"
	"strings"
)

func crabEnergy(cs []int, p1 bool) (energy int) {
	energy = (1 << 31) // a big number 
	hi, lo := AH.MaxAndMin(cs)

	for i := lo; i <= hi; i++ {
		temp_energy := 0
		for _, v := range cs {
			n := AH.AbsInt(v - i)
			if p1 {
				temp_energy += n
			} else {
				temp_energy += (n * (n + 1)) / 2
			}
		}

		if temp_energy < energy {
			energy = temp_energy
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

	AH.PrintSoln(7, crabEnergy(crabs, true), crabEnergy(crabs, false))

	return
}
