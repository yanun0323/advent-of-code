package solution

type Day6 struct {
	Solution
}

func (s Solution) Day6(inputs []string) (any, any) {
	d := Day6{s}
	return d.PuzzleA(inputs), d.PuzzleB(inputs)
}

func (d Day6) PuzzleA(inputs []string) any {
	input := inputs[0]
	m := map[byte]int{}
	q := make([]byte, 0, 4)
	for i := range input {
		if len(m) == 4 {
			return i
		}

		q = append(q, input[i])
		m[input[i]]++

		if len(q) <= 4 {
			continue
		}

		m[q[0]]--
		if m[q[0]] == 0 {
			delete(m, q[0])
		}
		q = q[1:]
	}
	return 0
}

func (d Day6) PuzzleB(inputs []string) any {
	input := inputs[0]
	m := map[byte]int{}
	q := make([]byte, 0, 14)
	for i := range input {
		if len(m) == 14 {
			return i
		}

		q = append(q, input[i])
		m[input[i]]++

		if len(q) <= 14 {
			continue
		}

		m[q[0]]--
		if m[q[0]] == 0 {
			delete(m, q[0])
		}
		q = q[1:]
	}
	return 0
}
