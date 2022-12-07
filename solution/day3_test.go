package solution

var (
	_inputDay3 = `vJrwpWtwJgWrhcsFMMfFFhFp
jqHRNqRjqzjGDLGLrsFMfFZSrLrFZsSL
PmmdzqPrVvPwwTWBwg
wMqvLMZHhHMvwLHjbvcjnnSBnvTQFn
ttgJtRGJQctTZtZT
CrZsJsPPZsGzwwsLwLmpwMDw`
)

func (su SolutionSuite) TestDay3() {
	inputs := getInputs(_inputDay3)
	d := Day3{su.Solution}
	su.Equal(157, d.PuzzleA(inputs))
	su.Equal(70, d.PuzzleB(inputs))
}
