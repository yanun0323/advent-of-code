package solution

var (
	_inputDay2 = `A Y
B X
C Z`
)

func (su SolutionSuite) TestDay2() {
	inputs := getInputs(_inputDay2)
	d := Day2{su.Solution}
	su.Equal(15, d.PuzzleA(inputs))
	su.Equal(12, d.PuzzleB(inputs))
}
