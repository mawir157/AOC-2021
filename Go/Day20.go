package main

import AH "./adventhelper"

type Pos struct {
	X, Y int
}

type LIMITS struct {
	x_lo, x_hi, y_lo, y_hi int
}

type IMG map[Pos]bool
type LUT map[int]bool

func BuildLookupTable (s string) (table LUT) {
	table = make(LUT)
	for i, c := range s {
		table[i] = (c == '#')
	}

	return
}

func BuildImage (ss []string) (image IMG, lims LIMITS) {
	image = make(IMG)
	for i, s := range ss {
		for j, c := range s {
			if c == '#' {
				image[Pos{X:i, Y:j}] = true
			}
		}
	}

	lims = LIMITS{0, len(ss), 0, len(ss[0])}
	return
}

func NeighbourCode(image IMG, lut LUT, x int, y int, lims LIMITS, infinity int) (bool) {
	index := 0
	for dx := -1; dx < 2; dx++ {
		xx := x+dx
		for dy := -1; dy < 2; dy++ {
			yy := y+dy
			index *= 2
			if _, ok := image[Pos{X:xx, Y:yy}]; ok {
				index += 1
			} else if (xx < lims.x_lo + 1 || xx >= lims.x_hi - 1 ||
				       yy < lims.y_lo + 1 || yy >= lims.y_hi - 1) {
				index += infinity
			}
		}		
	}

	return lut[index]
}

func Decompress (image IMG, lut LUT, lims LIMITS, inf int) (new_image IMG) {
	new_image = make(IMG)
	for x := lims.x_lo; x < lims.x_hi; x++ {
		for y := lims.y_lo; y < lims.y_hi; y++ {
			if NeighbourCode(image, lut, x, y, lims, inf) {
				p := Pos{X:x, Y:y}
				new_image[p] = true
			}
		}		
	}
	return
}

func DecompressLoop (image IMG, lut LUT, times int, lims LIMITS) (IMG) {
	inf := 0
	for i := 0; i < times; i++ {
		lims.x_lo -= 1
		lims.x_hi += 1
		lims.y_lo -= 1
		lims.y_hi += 1

		new_image := Decompress(image, lut, lims, inf)
		image = new_image
		if lut[0] {
			inf = 1 - inf
		}
	}
	return image
}

func main() {
	js, _ := AH.ReadStrFile("../input/input20.txt")
	lut := BuildLookupTable(js[0])

	image, lims := BuildImage(js[1:])

	AH.PrintSoln(20, len(DecompressLoop(image, lut, 2, lims)),
	              len(DecompressLoop(image, lut, 50, lims)))

	return
}
