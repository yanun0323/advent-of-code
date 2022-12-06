package service

var (
	_inputDay2 = `A Y
B X
C Z`
)

func (su SolutionSuite) TestDay2() {
	inputs := getInputs(_inputDay2)
	su.Equal(15, su.Day2A(inputs))
	su.Equal(12, su.Day2B(inputs))
}
