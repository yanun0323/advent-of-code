package solution

var _inputDay4 = `2-4,6-8
2-3,4-5
5-7,7-9
2-8,3-7
6-6,4-6
2-6,4-8`

func (su SolutionSuite) TestDay4() {
	inputs := getInputs(_inputDay4)
	d := Day4{su.Solution}
	su.Equal(2, d.PuzzleA(inputs))
	su.Equal(4, d.PuzzleB(inputs))
}