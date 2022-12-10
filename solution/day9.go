package solution

import (
	"fmt"
	"main/model"
	"main/util"
	"strconv"
	"strings"
)

type Day9 struct {
	Solution
}

var _startPoint = model.Coord{X: 0, Y: 0}

func (s Solution) Day9(inputs []string) (any, any) {
	d := Day9{s}
	return d.PuzzleA(inputs), d.PuzzleB(inputs)
}

func (d Day9) PuzzleA(inputs []string) any {
	var head, tail, tailNext model.Coord
	tailBeen := map[model.Coord]bool{_startPoint: true}
	commands := d.parseCommand(inputs)

	isHeadNear := func() bool {
		return util.Abs(head.X-tail.X) <= 1 && util.Abs(head.Y-tail.Y) <= 1
	}

	for _, cmd := range commands {
		head.X += cmd.X
		head.Y += cmd.Y
		if isHeadNear() {
			tailNext = head
			continue
		}
		tail = tailNext
		tailBeen[tail] = true
		tailNext = head
	}

	return len(tailBeen)
}

// TODO: Unsolved
func (d Day9) PuzzleB(inputs []string) any {
	var ropes, ropesNext [10]model.Coord
	tailBeen := map[model.Coord]bool{_startPoint: true}
	commands := d.parseCommand(inputs)

	isPrevNear := func(i int) bool {
		return util.Abs(ropes[i-1].X-ropes[i].X) <= 1 && util.Abs(ropes[i-1].Y-ropes[i].Y) <= 1
	}
	max := _startPoint
	min := _startPoint

	for _, cmd := range commands {
		ropes[0].X += cmd.X
		ropes[0].Y += cmd.Y
		max.X, max.Y = util.Max(ropes[0].X, max.X), util.Max(ropes[0].Y, max.Y)
		min.X, min.Y = util.Min(ropes[0].X, min.X), util.Min(ropes[0].Y, min.Y)
		for i := 1; i < len(ropes); i++ {
			if isPrevNear(i) {
				ropesNext[i] = ropes[i-1]
				break
			}
			ropes[i] = ropesNext[i]
			if i == len(ropes)-1 {
				tailBeen[ropes[i]] = true
			}
			ropesNext[i] = ropes[i-1]
		}
	}
	c := 0
	fmt.Println("max:", max.X, max.Y)
	fmt.Println("min:", min.X, min.Y)
	for y := max.Y; y >= min.Y; y-- {
		for x := min.X; x < max.X; x++ {
			char := "."
			if tailBeen[model.Coord{X: x, Y: y}] {
				char = "#"
				c++
			}
			fmt.Print(string(char))
		}
		fmt.Println("")
	}
	fmt.Println("count:", c)

	return len(tailBeen)
}

func (d Day9) parseCommand(inputs []string) []model.Coord {
	coords := []model.Coord{}
	for _, input := range inputs {
		if len(input) == 0 {
			d.l.Warn("empty input row")
			continue
		}
		ss := strings.Split(input, " ")
		if len(ss) != 2 {
			d.l.Warn("wrong input row format")
			continue
		}
		num, err := strconv.Atoi(ss[1])
		if err != nil {
			d.l.Errorf("parse command '%s' failed, %+v", input, err)
			continue
		}
		coord := model.Coord{}
		switch ss[0] {
		case "R":
			coord.X = 1
		case "U":
			coord.Y = 1
		case "L":
			coord.X = -1
		case "D":
			coord.Y = -1
		}
		for i := 0; i < num; i++ {
			coords = append(coords, coord)
		}
	}
	return coords
}
