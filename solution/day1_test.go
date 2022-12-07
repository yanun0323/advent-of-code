package solution

import "strings"

var _inputDay1 = `1000
2000
3000

4000

5000
6000

7000
8000
9000

10000`

func (su SolutionSuite) TestDay1() {
	inputs := getInputs(_inputDay1)
	d := Day1{su.Solution}
	su.Equal(24000, d.PuzzleA(inputs))
	su.Equal(45000, d.PuzzleB(inputs))
}

func getInputs(input string) []string {
	return strings.Split(input, "\n")
}
