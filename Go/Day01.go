package main

import AH "./adventhelper"

func part1(vec []int) (acc int) {
	for i := 0; i < len(vec) - 1; i++ {
		if vec[i] < vec[i+1] {
			acc += 1
		}
	}

	return
}

func part2(vec []int) (acc int) {
	for i := 0; i < len(vec) - 3; i++ {
		if vec[i] < vec[i+3] {
			acc += 1
		}		
	}

	return
}

func main() {
	ints, _ := AH.ReadIntFile("../input/input01.txt")

	AH.PrintSoln(1, part1(ints), part2(ints))

	return
}
