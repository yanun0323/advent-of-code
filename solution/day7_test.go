package solution

var _inputDay7 = `$ cd /
$ ls
dir a
14848514 b.txt
8504156 c.dat
dir d
$ cd a
$ ls
dir e
29116 f
2557 g
62596 h.lst
$ cd e
$ ls
584 i
$ cd ..
$ cd ..
$ cd d
$ ls
4060174 j
8033020 d.log
5626152 d.ext
7214296 k`

func (su SolutionSuite) TestDay7() {
	inputs := getInputs(_inputDay7)
	d := Day7{su.Solution}
	su.Equal(95437, d.PuzzleA(inputs))
	su.Equal(24933642, d.PuzzleB(inputs))
}
