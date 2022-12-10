package solution

var _inputDay6_1 = map[int]string{
	5:  "bvwbjplbgvbhsrlpgdmjqwftvncz",
	6:  "nppdvjthqldpwncqszvftbrmjlhg",
	10: "nznrnfrfntjfmvfwmzdfjlvtqnbhcprsg",
	11: "zcfzfwzzqfrljwzlrfnpqdbhtmscgvjw",
}
var _inputDay6_2 = map[int]string{
	19: "mjqjpqmgbljsphdztnvjfqwrcgsmlb",
	23: "bvwbjplbgvbhsrlpgdmjqwftvncz",
	29: "nznrnfrfntjfmvfwmzdfjlvtqnbhcprsg",
	26: "zcfzfwzzqfrljwzlrfnpqdbhtmscgvjw",
}

func (su SolutionSuite) TestDay6() {
	d := Day6{su.Solution}
	for ans, input := range _inputDay6_1 {
		su.Equal(ans, d.PuzzleA([]string{input}))
	}
	for ans, input := range _inputDay6_2 {
		su.Equal(ans, d.PuzzleB([]string{input}))
	}
}
