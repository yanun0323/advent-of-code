package solution

var (
	_inputDay8 = `30373
25512
65332
33549
35390`
	_inputDay8_2 = `99999
99999
56789
98765
99999`
	_inputDay8_3 = `9999999999
8889699999
0123456789
9999799999
9999999999`
)

func (su SolutionSuite) TestDay8() {
	inputs := getInputs(_inputDay8)
	inputs2 := getInputs(_inputDay8_2)
	inputs3 := getInputs(_inputDay8_3)
	d := Day8{su.Solution}
	su.Equal(21, d.PuzzleA(inputs))
	su.Equal(22, d.PuzzleA(inputs2))
	su.Equal(35, d.PuzzleA(inputs3))

	su.Equal(8, d.PuzzleB(inputs))
}
