package main

import AH "./adventhelper"

func bitCount(m map[string]bool, size int) []int {
	counts := make([]int, size)

	for s, _ := range(m) {
		for i, r := range s {
			if r == '1' {
				counts[i] += 1
			}
  	}
	}

	return counts
}

func GammaEpsilon(m map[string]bool, size int) (gamma int, epsilon int) {
	total := len(m)
	mostCommon := bitCount(m, size)

	for _, b := range mostCommon {
		gamma <<= 1
		epsilon <<= 1

		if 2*b >= total {
			gamma += 1
		} else {
			epsilon += 1
		}
	}

	return
}

func stringToBin(s string) (bin int) {
	for _, r := range s {
		bin <<= 1

		if r == '1' {
			bin += 1
		}
	}	

	return
}

func setReduce(m map[string]bool, size int, good bool) int {
	for i := 0; i < size; i++ {
		total := len(m)
		mostCommon := bitCount(m, size)
		v := mostCommon[i]

		r := 'X'
		if 2*v >= total {
			r = '1'
		} else {
			r = '0'
		}

		for k, _ := range(m) {
			if ((rune(k[i]) != r) == good) { 
				delete(m, k)
			}
		}

		if len(m) == 1 {
			for k, _ := range(m) {
				return stringToBin(k)
			}
		}
	}

	return -1
}

func Oxygen(m map[string]bool, size int) (o2 int, co2 int) {
	o2Good  := make(map[string]bool)
	co2Good := make(map[string]bool)

	for k, v :=  range m {
		o2Good[k] = v
		co2Good[k] = v
	}

	// do these sperately
	o2  = setReduce(o2Good, size, true)
	co2 = setReduce(co2Good, size, false)

	return
}

func main() {
	bins, _ := AH.ReadStrFile("../input/input03.txt")
	binSet := make(map[string]bool)
	for _, v := range bins {
		binSet[v] = true
	}

	g, e := GammaEpsilon(binSet, len(bins[0]))
	o2, co2 := Oxygen(binSet, len(bins[0]))

	AH.PrintSoln(3, g * e, o2 * co2)

	return
}
