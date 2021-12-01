package main

import AH "./adventhelper"

func diff_offset(vec []int, offset int) (acc int) {
	for i := 0; i < len(vec) - offset; i++ {
		if vec[i] < vec[i + offset] {
			acc += 1
		}
	}

	return
}

func main() {
	ints, _ := AH.ReadIntFile("../input/input01.txt")

	AH.PrintSoln(1, diff_offset(ints, 1), diff_offset(ints, 3))

	return
}
