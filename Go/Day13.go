package main

import AH "./adventhelper"

import (
	"fmt"
	"strconv"
	"strings"
)

type Point struct {
	x, y int
}

type Fold struct {
	axis  bool
	value int
}

func uniqueCount(intSlice []Point) (count int) {
	keys := make(map[Point]bool)
		for _, entry := range intSlice {
			if _, value := keys[entry]; !value {
				keys[entry] = true
				count += 1
			}
		}
	return count
}

func parseInput(ss []string) (map[Point]bool, []Fold) {
	points := make(map[Point]bool)
	folds := []Fold{}

	group := 0;
	for i, l := range ss {
		if (len(l) > 10) {
			break
		}
		s := strings.Split(l, ",")
		l, _ := strconv.Atoi(s[0])
		r, _ := strconv.Atoi(s[1])

		points[Point{x:l, y:r}] = true
		group = i
	}

	for i := group + 1; i < len(ss); i++ {
		str := ss[i]
		b := str[11] == 'y'
		tail := str[13:]
		v, _ := strconv.Atoi(tail)
		folds = append(folds, Fold{axis:b, value:v})
	}	
	return points, folds
}

func doFold(ps map[Point]bool, f Fold) map[Point]bool {
	newPoints := make(map[Point]bool)
	for p, _ := range ps {
		var pNew Point
		if (f.axis) {
			if (p.y > f.value) {
				pNew = Point{p.x, 2*f.value - p.y}
			} else {
			pNew = Point{p.x, p.y}
			}
		} else {
			if (p.x > f.value) {
				pNew = Point{2*f.value - p.x, p.y}
			} else {
			pNew = Point{p.x, p.y}
			}			
		}
		newPoints[pNew] = true
	}

	return newPoints
}

func printGrid(ps map[Point]bool) {
	for i := 0; i < 6; i++ {
		for j := 0; j < 40; j++ {
			p := Point{j,i}
			if _, value := ps[p]; value {
				fmt.Print("#")
			} else {
				fmt.Print(" ")
			}
		}
		fmt.Println("")
	}
}

func main() {
	js, _ := AH.ReadStrFile("../input/input13.txt")

	points, folds := parseInput(js)

	part2 := points
	for _, f := range folds {
		part2 = doFold(part2, f)
	}

	AH.PrintSoln(13, len(doFold(points, folds[0])), 0)
	printGrid(part2)

	return
}
