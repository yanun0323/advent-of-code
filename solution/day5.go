package solution

import (
	"main/util"
	"strconv"
	"strings"
)

type Day5 struct {
	Solution
}

func (s Solution) Day5(inputs []string) (any, any) {
	d := Day5{s}
	return d.PuzzleA(inputs), d.PuzzleB(inputs)
}

func (d Day5) PuzzleA(inputs []string) any {
	cratesStr, commands := d.separateInputs(inputs)
	crates, lastNum := d.parseCrates(cratesStr)

	for _, cmd := range commands {
		crates.ExecuteCommandA(cmd)
	}

	ans := []byte{}
	for num := byte('1'); num <= lastNum; num++ {
		if len(crates[num]) == 0 {
			continue
		}
		ans = append(ans, crates[num][len(crates[num])-1])
	}
	return string(ans)
}

func (d Day5) PuzzleB(inputs []string) any {
	cratesStr, commands := d.separateInputs(inputs)
	crates, lastNum := d.parseCrates(cratesStr)

	for _, cmd := range commands {
		crates.ExecuteCommandB(cmd)
	}

	ans := []byte{}
	for num := byte('1'); num <= lastNum; num++ {
		if len(crates[num]) == 0 {
			continue
		}
		ans = append(ans, crates[num][len(crates[num])-1])
	}
	return string(ans)
}

func (d Day5) separateInputs(inputs []string) ([]string, []string) {
	for i := range inputs {
		if len(inputs[i]) == 0 {
			return inputs[:i], inputs[i+1:]
		}
	}
	return inputs, nil
}

func (d Day5) parseCrates(ss []string) (Day5Crates, byte) {
	crates := Day5Crates{}
	num := map[int]byte{}
	numRow := ss[len(ss)-1]
	lastNum := byte('0')
	for i := range numRow {
		if util.IsByteNumber(numRow[i]) {
			num[i] = numRow[i]
			if numRow[i] > lastNum {
				lastNum = numRow[i]
			}
		}
	}

	for i := len(ss) - 2; i >= 0; i-- {
		for k, v := range num {
			if len(ss[i]) > k && ss[i][k] != ' ' {
				crates[v] = append(crates[v], ss[i][k])
			}
		}
	}
	return crates, lastNum
}

type Day5Crates map[byte][]byte

func (c Day5Crates) ExecuteCommandA(cmd string) {
	if len(cmd) == 0 {
		return
	}
	count, from, to := c.ParseCommand(cmd)
	for count > 0 && len(c[from]) > 0 {
		c[to] = append(c[to], c[from][len(c[from])-1])
		c[from] = c[from][:len(c[from])-1]
		count--
	}
}
func (c Day5Crates) ExecuteCommandB(cmd string) {
	if len(cmd) == 0 {
		return
	}
	count, from, to := c.ParseCommand(cmd)
	fromCount := len(c[from])
	if count > fromCount {
		count = fromCount
	}
	c[to] = append(c[to], c[from][fromCount-count:]...)
	c[from] = c[from][:fromCount-count]
}

func (c Day5Crates) ParseCommand(cmd string) (int, byte, byte) {
	cmd = strings.ReplaceAll(cmd, "move ", "")
	cmd = strings.ReplaceAll(cmd, " from ", ",")
	cmd = strings.ReplaceAll(cmd, " to ", ",")
	str := strings.Split(cmd, ",")
	i, _ := strconv.Atoi(str[0])
	return i, str[1][0], str[2][0]
}
