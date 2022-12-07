package solution

import (
	"strconv"
	"strings"
)

type Day4 struct {
	Solution
}

func (s Solution) Day4(inputs []string) (any, any) {
	d := Day4{s}
	return d.PuzzleA(inputs), d.PuzzleB(inputs)
}

func (d Day4) PuzzleA(inputs []string) any {
	ans := 0
	for _, input := range inputs {
		str := strings.Split(input, ",")
		num := make([]int, 0, 4)
		for i := range str {
			for _, ss := range strings.Split(string(str[i]), "-") {
				n, _ := strconv.Atoi(ss)
				num = append(num, n)
			}
		}
		if d.isFullyContain(num) {
			ans++
		}
	}
	return ans
}

func (d Day4) PuzzleB(inputs []string) any {
	ans := 0
	for _, input := range inputs {
		str := strings.Split(input, ",")
		num := make([]int, 0, 4)
		for i := range str {
			for _, ss := range strings.Split(string(str[i]), "-") {
				n, _ := strconv.Atoi(ss)
				num = append(num, n)
			}
		}
		if d.isOverlapping(num) {
			ans++
		}
	}
	return ans
}

func (d Day4) isFullyContain(num []int) bool {
	if len(num) < 4 {
		return false
	}
	return (num[0] >= num[2] && num[1] <= num[3]) || (num[0] <= num[2] && num[1] >= num[3])
}

func (d Day4) isOverlapping(num []int) bool {
	if len(num) < 4 {
		return false
	}

	if num[0] > num[3] || num[1] < num[2] {
		return false
	}

	return true
}
