package main

import AH "./adventhelper"

import (
	"strings"
)

func parseInput(ss []string) (map[string]int, map[string]string, string) {
	rules := make(map[string]string)
	pairs := make(map[string]int)
	final := string(AH.FinalRune(ss[0]))

	for i := 2; i < len(ss); i++ {
		ps := strings.Split(ss[i], " -> ")
		rules[ps[0]] = ps[1]
	}

	for i := 0; i < len(ss[0]) - 1; i++ {
		str := ss[0][i:i+2]
		pairs[str] += 1
	}

	return pairs, rules, final
}

func applyRules(rs map[string]string, ps map[string]int) map[string]int {
	pairs := make(map[string]int)

	for pair, count := range ps {
		r := rs[pair]

		pairs[pair[:1] + r] += count
		pairs[r + pair[1:]] += count
	}

	return pairs
}

func repApplyRules(n int, rs map[string]string, ps map[string]int) map[string]int {
	for i := 0; i < n; i++ {
		ps = applyRules(rs, ps)
	}

	return ps
}

func score(ps map[string]int, final string) int {
	max := 0
	min := 1
	min <<= 62

	charCount := make(map[string]int)
	for pair, count := range ps {
		charCount[pair[:1]] += count
	}
	charCount[final] += 1

	for _, count := range charCount {
		if count > max {
			max = count
		}
		if count < min {
			min = count
		}
	}

	return max - min
}

func main() {
	js, _ := AH.ReadStrFile("../input/input14.txt")
	pairs, rules, final := parseInput(js)

	part1 := score(repApplyRules(10, rules, pairs), final)
	part2 := score(repApplyRules(40, rules, pairs), final)

	AH.PrintSoln(14, part1, part2)

	return
}
