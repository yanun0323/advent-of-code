package solution

import (
	"math"
	"strconv"
	"strings"
)

type Day7 struct {
	Solution
}

func (s Solution) Day7(inputs []string) (any, any) {
	d := Day7{s}
	return d.PuzzleA(inputs), d.PuzzleB(inputs)
}

func (d Day7) PuzzleA(inputs []string) any {
	dirMap := d.getDirMap(inputs)
	sum := 0
	for _, size := range dirMap {
		if size <= 100_000 {
			sum += size
		}
	}
	return sum
}

func (d Day7) PuzzleB(inputs []string) any {
	dirMap := d.getDirMap(inputs)
	ans := math.MaxInt
	requireSize := 30_000_000 - (70_000_000 - dirMap["//"])
	if requireSize < 0 {
		requireSize = 0
	}

	for _, size := range dirMap {
		if size >= requireSize && size < ans {
			ans = size
		}
	}
	return ans
}

func (d Day7) getDirMap(inputs []string) map[string]int {
	dirMap := map[string]int{}
	currentPath := []string{}
	dirStack := []Day7Dir{}
	for _, input := range inputs {
		if len(input) == 0 {
			continue
		}
		ss := strings.Split(input, " ")
		if ss[0] == "dir" {
			dirStack[len(dirStack)-1].dir = append(dirStack[len(dirStack)-1].dir, ss[1])
			continue
		}

		if ss[0] != "$" {
			dirStack[len(dirStack)-1].size += d.calculateFileSize(ss)
			continue
		}

		if ss[1] != "cd" {
			continue
		}

		if ss[2] != ".." {
			currentPath = append(currentPath, ss[2]+"/")
			dirStack = append(dirStack, Day7Dir{
				path: strings.Join(currentPath, ""),
			})
			continue
		}

		path := strings.Join(currentPath, "")
		if dirStack[len(dirStack)-1].path == path {
			dir := dirStack[len(dirStack)-1]
			dirStack = dirStack[:len(dirStack)-1]

			for _, childDir := range dir.dir {
				childDirPath := dir.path + childDir + "/"
				dir.size += dirMap[childDirPath]
			}
			dirMap[dir.path] = dir.size
		}
		currentPath = currentPath[:len(currentPath)-1]
	}

	for len(dirStack) > 0 {
		dir := dirStack[len(dirStack)-1]
		dirStack = dirStack[:len(dirStack)-1]

		for _, childDir := range dir.dir {
			childDirPath := dir.path + childDir + "/"
			dir.size += dirMap[childDirPath]
		}
		dirMap[dir.path] = dir.size
	}

	return dirMap
}

func (d Day7) calculateFileSize(ss []string) int {
	num, err := strconv.Atoi(ss[0])
	if err != nil {
		d.l.Errorf("convert string '%s' to int failed, %+v", strings.Join(ss, " "), err)
		return 0
	}
	return num
}

type Day7Dir struct {
	path string
	dir  []string
	size int
}
