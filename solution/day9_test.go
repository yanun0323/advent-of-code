package solution

var (
	_inputDay9_1 = `R 4
U 4
L 3
D 1
R 4
D 1
L 5
R 2`

	_inputDay9_2 = `R 5
U 8
L 8
D 3
R 17
D 10
L 25
U 20`
)

func (su SolutionSuite) TestDay9() {
	inputs := getInputs(_inputDay9_1)
	inputs2 := getInputs(_inputDay9_2)
	d := Day9{su.Solution}
	su.Equal(13, d.PuzzleA(inputs))
	su.Equal(1, d.PuzzleB(inputs))
	su.Equal(36, d.PuzzleB(inputs2))
}
