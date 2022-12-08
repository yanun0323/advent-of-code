package solution

type Day8 struct {
	Solution
}

func (s Solution) Day8(inputs []string) (any, any) {
	inputs = inputs[:len(inputs)-1]
	d := Day8{s}
	return d.PuzzleA(inputs), d.PuzzleB(inputs)
}

func (d Day8) PuzzleA(inputs []string) any {
	visibleGrid := make([][]bool, len(inputs))
	for i := range visibleGrid {
		visibleGrid[i] = make([]bool, len(inputs[0]))
	}

	var iterate func(highest byte, r, c, dR, dC int)
	iterate = func(highest byte, r, c, dR, dC int) {
		if r < 0 || c < 0 || r >= len(inputs) || c >= len(inputs[0]) {
			return
		}
		current := inputs[r][c]
		if current > highest {
			visibleGrid[r][c] = true
			iterate(current, r+dR, c+dC, dR, dC)
			return
		}

		iterate(highest, r+dR, c+dC, dR, dC)
	}

	ground := byte('0' - 1)

	for i := 0; i < len(inputs); i++ {
		iterate(ground, i, 0, 0, 1)                 /* from left to right */
		iterate(ground, i, len(inputs[0])-1, 0, -1) /* from right to left */
	}

	for i := 0; i < len(inputs[0]); i++ {
		iterate(ground, 0, i, 1, 0)              /* from top to bottom */
		iterate(ground, len(inputs)-1, i, -1, 0) /* from bottom to top */
	}

	visibleCount := 0
	for i := range visibleGrid {
		for j := range visibleGrid[i] {
			if visibleGrid[i][j] {
				visibleCount++
			}
		}
	}

	return visibleCount
}

func (d Day8) PuzzleB(inputs []string) any {
	scoreGrid := make([][]int, len(inputs))
	for i := range scoreGrid {
		scoreGrid[i] = make([]int, len(inputs[0]))
		for j := 0; j < len(inputs[0]); j++ {
			scoreGrid[i][j] = 1
		}
	}
	viewStack := Day8TreeStack{}
	var iterate func(r, c, dR, dC, depth int)
	iterate = func(r, c, dR, dC, depth int) {
		if r < 0 || c < 0 || r >= len(inputs) || c >= len(inputs[0]) {
			return
		}

		current := inputs[r][c]
		for viewStack.Len() > 0 && viewStack.Top().height < current {
			_ = viewStack.Pop()
		}

		if viewStack.Len() == 0 {
			scoreGrid[r][c] *= depth
		} else {
			scoreGrid[r][c] *= depth - viewStack.Top().depth 
		}

		viewStack.Push(Day8Tree{
			height: current,
			depth:  depth,
		})

		iterate(r+dR, c+dC, dR, dC, depth+1)
	}

	for i := 0; i < len(inputs); i++ {
		viewStack.Clear()
		iterate(i, 0, 0, 1, 0) /* from left to right */
		viewStack.Clear()
		iterate(i, len(inputs[0])-1, 0, -1, 0) /* from right to left */
	}

	for i := 0; i < len(inputs[0]); i++ {
		viewStack.Clear()
		iterate(0, i, 1, 0, 0) /* from top to bottom */
		viewStack.Clear()
		iterate(len(inputs)-1, i, -1, 0, 0) /* from bottom to top */
	}

	max := 0
	for i := range scoreGrid {
		for j := range scoreGrid[i] {
			if scoreGrid[i][j] > max {
				max = scoreGrid[i][j]
			}
		}
	}

	return max
}

type Day8Tree struct {
	height byte
	depth  int
}

type Day8TreeStack struct {
	stack []Day8Tree
}

func (s Day8TreeStack) Len() int {
	return len(s.stack)
}

func (s Day8TreeStack) Top() Day8Tree {
	return s.stack[s.Len()-1]
}

func (s *Day8TreeStack) Pop() Day8Tree {
	t := s.Top()
	s.stack = s.stack[:s.Len()-1]
	return t
}

func (s *Day8TreeStack) Push(t Day8Tree) {
	s.stack = append(s.stack, t)
}

func (s *Day8TreeStack) Clear() {
	s.stack = []Day8Tree{}
}
