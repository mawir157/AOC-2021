package main

import AH "./adventhelper"

import (
	"sort"
)

func score1(r rune) int {
	if r == ')' {
		return 3
	}
	if r == ']' {
		return 57
	}
	if r == '}' {
		return 1197
	}
	if r == '>' {
		return 25137
	}

	return 0
}

func score2(s string) (score int) {
	score = 0
	for _, v := range s {
		score *= 5
		if v == '(' {
			score += 1
		}
		if v == '[' {
			score += 2
		}
		if v == '{' {
			score += 3
		}
		if v == '<' {
			score += 4
		}
	}

	return
}

func closing(l rune, r rune) bool {
	return (l == '(' && r == ')') || (l == '{' && r == '}') ||
	       (l == '[' && r == ']') || (l == '<' && r == '>')
}

func illegal(s string) (rune, string) {
	stack := ""
	for _, r := range s {
		if AH.ContainsChar("({<[", r) {// bra
			stack = string(r) + stack
		} else if (closing(AH.FirstRune(stack), r)) {
			stack = AH.TrimFirstRune(stack)
		} else {
			return r, "" // found bad character
		}
	}
	// reached the end of the string
	return 'X', stack
}

func main() {
	inputLines, _ := AH.ReadStrFile("../input/input10.txt")

	part1 := 0
	part2arr := []int{}
	goodCount := 0

	for _, s := range inputLines {
		a,b := illegal(s)

		part1 += score1(a)
		if a == 'X' {
			goodCount += 1
			part2arr = append(part2arr, score2(b))
		}
	}
	sort.Sort(sort.IntSlice(part2arr))

	AH.PrintSoln(10, part1, part2arr[goodCount / 2])

	return
}
