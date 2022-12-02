package service

import "strings"

var (
	_inputDay2 = `A Y
B X
C Z`
)

func (su ServiceSuite) TestDay2() {
	inputs := strings.Split(_inputDay2, "\n")
	ansA := su.Day2A(inputs).(int)
	su.Equal(15, ansA)

	ansB := su.Day2B(inputs).(int)
	su.Equal(12, ansB)
}
