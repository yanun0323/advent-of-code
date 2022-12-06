package service

var _inputDay5 = `    [D]    
[N] [C]    
[Z] [M] [P]
 1   2   3 

move 1 from 2 to 1
move 3 from 1 to 3
move 2 from 2 to 1
move 1 from 1 to 2`

func (su SolutionSuite) TestDay5() {
	inputs := getInputs(_inputDay5)
	su.Equal("CMZ", su.Day5A(inputs))
	su.Equal("MCD", su.Day5B(inputs))
}
