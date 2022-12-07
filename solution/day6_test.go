package solution

var _inputDay6A = map[int]string{
	5:  "bvwbjplbgvbhsrlpgdmjqwftvncz",
	6:  "nppdvjthqldpwncqszvftbrmjlhg",
	10: "nznrnfrfntjfmvfwmzdfjlvtqnbhcprsg",
	11: "zcfzfwzzqfrljwzlrfnpqdbhtmscgvjw",
}
var _inputDay6B = map[int]string{
	19: "mjqjpqmgbljsphdztnvjfqwrcgsmlb",
	23: "bvwbjplbgvbhsrlpgdmjqwftvncz",
	29: "nznrnfrfntjfmvfwmzdfjlvtqnbhcprsg",
	26: "zcfzfwzzqfrljwzlrfnpqdbhtmscgvjw",
}

func (su SolutionSuite) TestDay6() {
	d := Day6{su.Solution}
	for ans, input := range _inputDay6A {
		su.Equal(ans, d.PuzzleA([]string{input}))
	}
	for ans, input := range _inputDay6B {
		su.Equal(ans, d.PuzzleB([]string{input}))
	}
}
