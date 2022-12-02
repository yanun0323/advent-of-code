package service

import (
	"strings"
)

var (
	_lose  = map[string]string{"A": "Z", "B": "X", "C": "Y"}
	_draw  = map[string]string{"A": "X", "B": "Y", "C": "Z"}
	_win   = map[string]string{"A": "Y", "B": "Z", "C": "X"}
	_score = map[string]int{"X": 1, "Y": 2, "Z": 3}

	_strategy = map[string]map[string]string{"X": _lose, "Y": _draw, "Z": _win}
)

func (svc Service) Day2A(inputs []string) any {
	score := 0
	for _, s := range inputs {
		chars := strings.Split(s, " ")
		if len(chars) != 2 {
			continue
		}
		score += getPairScore(chars[0], chars[1]) + _score[chars[1]]
	}
	return score
}

func getPairScore(opponent, player string) int {
	if _lose[opponent] == player {
		return 0
	}
	if _draw[opponent] == player {
		return 3
	}
	if _win[opponent] == player {
		return 6
	}
	return 0
}

func getChosenScore(char string) int {
	switch char {
	case "X":
		return 1
	case "Y":
		return 2
	case "Z":
		return 3
	}
	return 0
}

func (svc Service) Day2B(inputs []string) any {
	score := 0
	for _, s := range inputs {
		chars := strings.Split(s, " ")
		if len(chars) != 2 {
			continue
		}
		myShape := _strategy[chars[1]][chars[0]]
		score += getPairScore(chars[0], myShape) + _score[myShape]
	}
	return score
}
