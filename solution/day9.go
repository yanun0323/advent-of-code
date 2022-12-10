package solution

import (
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
	var head, tail model.Coord
	tailBeen := map[model.Coord]bool{_startPoint: true}
	movements := d.parseCommand(inputs)

	for _, mv := range movements {
		head.X += mv.X
		head.Y += mv.Y
		if d.isNear(head, tail) {
			continue
		}
		tail = d.moveRope(head, tail)
		tailBeen[tail] = true
	}

	return len(tailBeen)
}

// BUG: Unsolved
func (d Day9) PuzzleB(inputs []string) any {
	var ropes [10]model.Coord
	tail := len(ropes) - 1
	tailBeen := map[model.Coord]bool{_startPoint: true}
	movements := d.parseCommand(inputs)

	for _, mv := range movements {
		ropes[0].X += mv.X
		ropes[0].Y += mv.Y
		for i := 1; i <= tail; i++ {
			if d.isNear(ropes[i-1], ropes[i]) {
				continue
			}
			ropes[i] = d.moveRope(ropes[i-1], ropes[i])
		}
		tailBeen[ropes[tail]] = true
	}

	return len(tailBeen)
}

func (d Day9) isNear(head, tail model.Coord) bool {
	return util.Abs(head.X-tail.X) <= 1 && util.Abs(head.Y-tail.Y) <= 1
}

func (d Day9) moveRope(target, current model.Coord) model.Coord {
	if util.Abs(target.X-current.X) >= 2 {
		current.X = (target.X + current.X) / 2
		current.Y = target.Y
		return current
	}

	if util.Abs(target.Y-current.Y) >= 2 {
		current.X = target.X
		current.Y = (target.Y + current.Y) / 2
		return current
	}

	current.X = (target.X + current.X) / 2
	current.Y = (target.Y + current.Y) / 2
	return current
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
