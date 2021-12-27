package main

import AH "./adventhelper"

import (
	"strconv"
)

type Game struct {
	pos1   int
	score1 int
	pos2   int
	score2 int
}

var binomial = map[int]int{3:1, 4:3, 5:6, 6:7, 7:6, 8:3, 9:1}

func playGame1(p1 int, p2 int) int {
	n := 1
	rolls := 0
	score1 := 0
	score2 := 0
	for ; true ; {
		rolls += 3
		temp := p1 + (3*n + 3)
		p1 = ((temp - 1) % 10) + 1
		n += 3
		score1 += p1
		if (score1 >= 1000) {
			return rolls * score2
		}

		rolls += 3
		temp = p2 + (3*n + 3)
		p2 = ((temp - 1) % 10) + 1
		n += 3
		score2 += p2
		if (score2 >= 1000) {
			return rolls * score1
		}
 	}

	return -1
}

func sumValues(m map[Game]int) (total int) {
	for _, v := range m {
		total += v
	}

	return
}

func playGame2(p1 int, p2 int) int {
	win1 := 0
	win2 := 0
	initialState := Game{pos1: p1, score1: 0, pos2: p2, score2: 0}
	gameStates := make(map[Game]int)
	gameStates[initialState] = 1
	player1 := true
	for temp := 0; temp < 20; temp++ {
		nextStates := make(map[Game]int)
		for k, v := range gameStates {
			// player 1 turn
			if player1 {
				for s, t := range binomial {
					knew := k
					knew.pos1 = (((k.pos1 + s) - 1) % 10) + 1
					knew.score1 += knew.pos1
					nextStates[knew] += v * t
				}
			} else { // player2
				for s, t := range binomial {
					knew := k
					knew.pos2 = (((k.pos2 + s) - 1) % 10) + 1
					knew.score2 += knew.pos2
					nextStates[knew] += v	 * t
				}			
			}
		}
		player1 = !player1

		nextNextStates := make(map[Game]int)
		for k, v := range nextStates {
			if k.score1 >= 21 {
				win1 += v
			} else if k.score2 >= 21 {
				win2 += v
			} else {
				nextNextStates[k] = v
			}
		}

		gameStates = nextNextStates
	}

	if win1 > win2 {
		return win1
	} else {
		return win2
	}
}

func main() {
	ss, _ := AH.ReadStrFile("../input/input21.txt")
	pos1, _ := strconv.Atoi(AH.Drop(ss[0],28))
	pos2, _ := strconv.Atoi(AH.Drop(ss[1],28))

	AH.PrintSoln(21, playGame1(pos1, pos2), playGame2(pos1, pos2))

	return
}
