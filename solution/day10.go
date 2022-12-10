package solution

import (
	"strconv"
	"strings"
)

type Day10 struct {
	Solution
}

func (s Solution) Day10(inputs []string) (any, any) {
	d := Day10{s}
	return d.PuzzleA(inputs), d.PuzzleB(inputs)
}

func (d Day10) PuzzleA(inputs []string) any {
	commands := d.parseCommand(inputs)
	sum, value := 0, 1
	for i, num := range commands {
		cycle := i + 1
		if d.needGetSignal(cycle) {
			sum += cycle * value
		}
		value += num
	}

	return sum
}

func (d Day10) PuzzleB(inputs []string) any {
	crt := make([][]byte, 6)
	for i := range crt {
		crt[i] = make([]byte, 40)
	}
	commands := d.parseCommand(inputs)
	value := 1

	for i, num := range commands {
		cycle := (i % 40) + 1
		if cycle >= value && cycle <= value+2 {
			crt[i/40][i%40] = '#'
		} else {
			crt[i/40][i%40] = '.'
		}

		value += num
	}

	ans := make([]string, 0, 6)
	for i := range crt {
		ans = append(ans, string(crt[i]))
	}

	return strings.Join(ans, "\n")
}

func (d Day10) needGetSignal(cycle int) bool {
	switch cycle {
	case 20, 60, 100, 140, 180, 220:
		return true
	}
	return false
}

func (d Day10) parseCommand(inputs []string) []int {
	add := make([]int, 0, 2*len(inputs))
	for _, input := range inputs {
		if len(input) == 0 {
			d.l.Warn("empty input row")
			continue
		}
		ss := strings.Split(input, " ")
		if len(ss) > 0 && ss[0] == "noop" {
			add = append(add, 0)
			continue
		}

		if len(ss) != 2 || ss[0] != "addx" {
			d.l.Warnf("wrong input row format: %s", input)
			continue
		}

		num, err := strconv.Atoi(ss[1])
		if err != nil {
			d.l.Errorf("convert row '%s' number failed, %+v", input, err)
		}
		add = append(add, 0, num)
	}
	return add
}
