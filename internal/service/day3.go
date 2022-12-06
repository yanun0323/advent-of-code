package service

func (s Solution) Day3A(inputs []string) any {
	sum := 0
	for _, input := range inputs {
		count := len(input)
		m := map[byte]bool{}
		for i := 0; i < count/2; i++ {
			m[input[i]] = true
		}
		for i := count / 2; i < count; i++ {
			if m[input[i]] {
				sum += day3GetScore(input[i])
				break
			}
		}
	}
	return sum
}

func day3GetScore(char byte) int {
	if char >= 'a' && char <= 'z' {
		return int(char - 'a' + 1)
	}
	return int(char - 'A' + 27)
}

func (s Solution) Day3B(inputs []string) any {
	m := map[byte]int{}
	count := len(inputs)
	if count > 3 {
		count = 3
	}
	last := count - 1
	sum := 0
	for i := 0; i < count; i++ {
		for j := range inputs[i] {
			if i == last && m[inputs[i][j]] == i {
				sum += day3GetScore(inputs[i][j])
				break
			}
			if m[inputs[i][j]] == i {
				m[inputs[i][j]]++
			}
		}
	}
	if len(inputs) > 3 {
		return sum + s.Day3B(inputs[3:]).(int)
	}
	return sum
}
