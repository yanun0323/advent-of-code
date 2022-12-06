package service

func (s Solution) Day6A(inputs []string) any {
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

func (s Solution) Day6B(inputs []string) any {
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
