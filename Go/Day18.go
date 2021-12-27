package main

import AH "./adventhelper"

import (
	// "fmt"
	"strconv"
)

type Vertex struct {
	value      int
	parent     *Vertex
	lhs        *Vertex
	rhs        *Vertex
}

func (v Vertex) toString() string {
	if (v.value >= 0) {
		s_v := strconv.Itoa(v.value)
		return s_v
	} else {
		return ("[" + v.lhs.toString() + "," + v.rhs.toString() + "]")
	}
}

func combineTrees(l string, r string) (int, string) {
	comb_string := ("[" + l + "," + r + "]")
	tree := parseToBintree(comb_string, nil)
	tree.fullReduce()
	return tree.magnitude(), tree.toString()
}

func (v Vertex) magnitude() int {
	if (v.value >= 0) {
		return v.value
	} else {
		return (3 * v.lhs.magnitude() + 2 * v.rhs.magnitude())
	}	
}

func findExplodable(v *Vertex, depth int) *Vertex {
	// we can't do any further down
	if (v.lhs == nil) || (v.rhs == nil) {
		return nil
	}
	// if both children are value vertices we are done
	if (v.lhs.value >= 0) && (v.rhs.value >= 0) && (depth > 3) {
		return v
	}

	// try to go left
	q := findExplodable(v.lhs, depth + 1)
	if (q != nil) {
		return q
	}

	// try to go right
	q = findExplodable(v.rhs, depth + 1)
	if (q != nil) {
		return q
	}

	// we have got to here without finding an explodable vertex
	return nil
}

func findL(v * Vertex) (* Vertex) {
	// go to parent
	par := v.parent
	if par == nil {
		return nil
	}

	// we have come from the right...
	if (*v != *par.lhs) {
		// ...so go down the left...
		v = par.lhs
		// ...and keep to the right until we hit the end
		for ; v.rhs != nil ; {
			v = v.rhs
		}
		return v
	} else {
		return findL(par)
	}

	
}

func findR(v * Vertex) (* Vertex) {
	// go to parent
	par := v.parent
	if par == nil {
		return nil
	}

	// we have come from the left...
	if (*v != *par.rhs) {
		// ...so go down the right...
		v = par.rhs
		// ...and keep to the left until we hit the end
		for ; v.lhs != nil ; {
			v = v.lhs
		}
		return v
	} else {
		return findR(par)
	}

}

func (tree *Vertex) explode() (changed bool) {
	expVertex := findExplodable(tree, 0)
	changed = false
	for ; expVertex != nil; {
		changed = true

		l_val := expVertex.lhs.value
		r_val := expVertex.rhs.value

		left  := findL(expVertex)
		right := findR(expVertex)

		// find expVertex's parent

		expVertex.lhs = nil
		expVertex.rhs = nil
		expVertex.value = 0

		if (left == nil) {
			expVertex.value = 0
		} else {
			left.value += l_val
		}

		if (right == nil) {
			expVertex.value = 0
		} else {
			right.value += r_val
		}
		
		expVertex = findExplodable(tree, 0)
	}

	return
}

func div(i int) (l, r int) {
	if i % 2 == 0 {
		return (i / 2), (i / 2)
	} else {
		return (i / 2), ((i / 2) + 1)
	}
}

func (v *Vertex)split() bool {
	// this vertex is a value
	// fmt.Println("Split?", v.value)
	if (v.lhs == nil) && (v.rhs == nil) {
		if v.value > 9 {
			l_val, r_val := div(v.value)
			v_left  := Vertex{value:l_val, lhs:nil, rhs:nil, parent:v}
			v_right := Vertex{value:r_val, lhs:nil, rhs:nil, parent:v}

			v.value = -1
			v.lhs = & v_left
			v.rhs = & v_right

			return true
		} else {
			return false
		}
	}

	// this vertex is not a value so check if we can split on the left
	// b := v.lhs.split()
	if v.lhs.split() {
		return true
	}

	// this vertex is not a value so check if we can split on the right
	// b = v.rhs.split()
	if v.rhs.split() {
		return true
	}

	// we haven't found and split a vertex
	return false
}

func (tree *Vertex)fullReduce() {
	changed := true

	for ; changed ; {
		changed = tree.explode()
		changed = changed || tree.split()
	}
}

func parseToBintree(s string, parent *Vertex) Vertex {
	t := AH.TrimFirstRune(AH.TrimLastRune(s)) // remove outer [ ]
	l_string, r_string := splitOnTopComma(t)
	l_final := !AH.ContainsChar(l_string, ',')
	var l_val int
	if l_final {
		l_val, _ = strconv.Atoi(l_string)
	} else {
		l_val = 0
	}

	r_final := !AH.ContainsChar(r_string, ',')
	var r_val int
	if r_final {
		r_val, _ = strconv.Atoi(r_string)
	} else {
		r_val = 0
	}

	var lv Vertex
	var rv Vertex
	current := Vertex{value: -1, parent: parent}
	if (l_final && r_final) {
		lv = Vertex{value:l_val, lhs:nil, rhs:nil, parent:&current}
		rv = Vertex{value:r_val, lhs:nil, rhs:nil, parent:&current}
	} else if (l_final) {
		lv = Vertex{value:l_val, lhs:nil, rhs:nil, parent:&current}
		rv = parseToBintree(r_string, &current)
	} else if (r_final) {
		lv = parseToBintree(l_string, &current)
		rv = Vertex{value:r_val, lhs:nil, rhs:nil, parent:&current}
	} else {
		lv = parseToBintree(l_string, &current)
		rv = parseToBintree(r_string, &current)
	}
	current.lhs = &lv
	current.rhs = &rv
	return current
}

func splitOnTopComma(s string) (lhs string, rhs string) {
	depth := 0
	for i, r := range s {
		switch r {
			case '[':
				depth += 1
				lhs = lhs + string(r)
			case ']':
				depth -= 1
			case ',': // minimum packet
				if depth == 0 {
					lhs = AH.Take(s, i)
					rhs = AH.Drop(s, i + 1)

					return
				}
		}
	}
	return
}


func main() {
	ss, _ := AH.ReadStrFile("../input/input18.txt")

	// part 1
	sum := ""
	for i := 0; i < len(ss); i++ {
		_, sum = combineTrees(sum, ss[i])
	}
	part1 := parseToBintree(sum, nil)

	// part 2
	part2 := 0
	for i := 0; i < len(ss); i++ {
		for j := 1; j < len(ss); j++ {
			m0, _ := combineTrees(ss[i], ss[j])
			if (m0 > part2) {
				part2 = m0
			}

			m1, _ := combineTrees(ss[j], ss[i])
			if (m1 > part2) {
				part2 = m1
			}
		}
	}

	AH.PrintSoln(18, part1.magnitude(), part2)

	return
}
