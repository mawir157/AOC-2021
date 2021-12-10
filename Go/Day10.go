package main

import AH "./adventhelper"

import (
	"sort"
)

func score1(s string) int {
	if s == ")" {
		return 3
	}
	if s == "]" {
		return 57
	}
	if s == "}" {
		return 1197
	}
	if s == ">" {
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
	if (l == '(' && r == ')') || (l == '{' && r == '}') ||
	   (l == '[' && r == ']') || (l == '<' && r == '>') {
		return true
	}

	return false
}

func illegal(s string) (string, string) {
	stack := ""
	for _, r := range s {
		if AH.ContainsChar("({<[", r) {// bra
			stack = string(r) + stack
		} else if (closing(AH.FirstRune(stack), r)) {
			stack = AH.TrimFirstRune(stack)
		} else {
			return string(r), "" // found bad character
		}
	}
	// reached the end of the string
	return "", stack
}

func main() {
	inputLines, _ := AH.ReadStrFile("../input/input10.txt")

	part1 := 0
	part2arr := []int{}
	goodCount := 0

	for _, s := range inputLines {
		a,b := illegal(s)

		part1 += score1(a)
		if len(b) > 0 {
			goodCount += 1
			part2arr = append(part2arr, score2(b))
		}
	}
	sort.Sort(sort.IntSlice(part2arr))

	AH.PrintSoln(10, part1, part2arr[goodCount / 2])

	return
}
