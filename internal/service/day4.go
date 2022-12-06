package service

import (
	"strconv"
	"strings"
)

func (s Solution) Day4A(inputs []string) any {
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
		if day4IsFullyContain(num) {
			ans++
		}
	}
	return ans
}

func (s Solution) Day4B(inputs []string) any {
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
		if day4IsOverlapping(num) {
			ans++
		}
	}
	return ans
}

func day4IsFullyContain(num []int) bool {
	if len(num) < 4 {
		return false
	}
	return (num[0] >= num[2] && num[1] <= num[3]) || (num[0] <= num[2] && num[1] >= num[3])
}

func day4IsOverlapping(num []int) bool {
	if len(num) < 4 {
		return false
	}

	if num[0] > num[3] || num[1] < num[2] {
		return false
	}

	return true
}
