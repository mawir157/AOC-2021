package main

import AH "./adventhelper"

import (
	"sort"
	"strings"
)

func sortString(w string) string {
	s := strings.Split(w, "")
	sort.Strings(s)
	return strings.Join(s, "")
}

func parseLine(s string) (all []string, code []string) {
	ps := strings.Split(s, " | ")
	all = strings.Split(ps[0], " ")
	code = strings.Split(ps[1], " ")

	for i, v := range all {
		all[i] = sortString(v)
	}

	for i, v := range code {
		code[i] = sortString(v)
	}

	return
}

func stringDiff(a, b string) (diff string) {
	m := make(map[rune]bool)

	for _, item := range b {
		m[item] = true
	}

	for _, item := range a {
		if _, ok := m[item]; !ok {
			diff = diff + string(item)
		}
	}
	return
}

func filter(ss []string, s string) ([]string) {
	f := []string{}
	for _, v := range ss {
		if v != s {
			f = append(f, v)
		}
	}
	return f
}

// i don't know i there is a clever way to do this...
func unscramble(ss []string) ([]string) {
	// get the obvious words
	Word1 := "[1]"
	Word7 := "[7]"
	Word4 := "[4]"
	Word8 := "[8]"
	for _, v := range ss {
		if len(v) == 2 {
			Word1 = v
		} else if len(v) == 3 {
			Word7 = v
		} else if len(v) == 4 {
			Word4 = v
		} else if len(v) == 7 {
			Word8 = v
		}
	}

	ss = filter(ss, Word1)
	ss = filter(ss, Word7)
	ss = filter(ss, Word4)
	ss = filter(ss, Word8)

	// find the length 6 containing Word4
	Word9 := "[9]"
	for _, v := range ss {
		if len(v) != 6 {
			continue
		}

		q := stringDiff(Word4, v)
		if len(q) == 0 {
			Word9 = v
			break
		}
	}
	ss = filter(ss, Word9)

	// find the length 6 containing Word1
	Word0 := "[0]"
	Word6 := "[6]"
	for _, v := range ss {
		if len(v) != 6 {
			continue
		}

		q := stringDiff(Word1, v)
		if len(q) == 0 {
			Word0 = v
		} else {
			Word6 = v
		}
	}
	ss = filter(ss, Word6)
	ss = filter(ss, Word0)

	// find the length 5 word conaining Word1
	Word3 := "[3]"
	for _, v := range ss {
		if len(v) != 5 {
			continue
		}

		q := stringDiff(Word1, v)
		if len(q) == 0 {
			Word3 = v
			break
		}
	}
	ss = filter(ss, Word3)

	// find the length 5 not contained in Word9
	Word2 := "[2]"
	Word5 := "[5]"
	for _, v := range ss {
		if len(v) != 5 {
			continue
		}

		q := stringDiff(v, Word9)
		if len(q) == 0 {
			Word5 = v
		} else {
			Word2 = v
		}
	}
	ss = filter(ss, Word2)
	ss = filter(ss, Word5)

	decoded := []string{Word0, Word1, Word2, Word3, Word4,
                      Word5, Word6, Word7, Word8, Word9}

	return decoded
}

func match(ds []string, cs []string) (decode int) {

	for _, c := range cs {
		for i, d := range ds {
			if d == c {
				decode *= 10
				decode += i
			}
		}
	}

	return
}

func part1(ss []string) (total int) {
	for _, s := range ss {
		if (len(s) == 2) || (len(s) == 3) || (len(s) == 4) || (len(s) == 7) {
			total += 1
		}
	}

	return
}

func main() {
	ss, _ := AH.ReadStrFile("../input/input08.txt")

	total1 := 0
	total2 := 0
	for _, s := range ss {
		p, q := parseLine(s)
		total1 += part1(q)

		p = unscramble(p)
		total2 += match(p, q)
	}

	AH.PrintSoln(8, total1, total2)

	return
}
