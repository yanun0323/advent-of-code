package service

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
	su.Equal(157, su.Day3A(inputs))
	su.Equal(70, su.Day3B(inputs))
}
