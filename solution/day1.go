package solution

import (
	"container/heap"
	"strconv"
)

type Day1 struct {
	Solution
}

func (s Solution) Day1(inputs []string) (any, any) {
	d := Day1{s}
	return d.PuzzleA(inputs), d.PuzzleB(inputs)
}

func (d Day1) PuzzleA(calories []string) any {
	max := 0
	temp := 0
	for _, calory := range calories {
		if len(calory) == 0 {
			if temp > max {
				max = temp
			}
			temp = 0
			continue
		}
		num, err := strconv.Atoi(calory)
		if err != nil {
			d.l.Errorf("convert string %s failed, %+v", calory, err)
			return nil
		}
		temp += num
	}

	return max
}

func (d Day1) PuzzleB(calories []string) any {
	h := &Day1Heap{}
	heap.Init(h)

	carried := 0
	for _, calory := range calories {
		if len(calory) == 0 {
			heap.Push(h, carried)
			carried = 0
			continue
		}
		num, err := strconv.Atoi(calory)
		if err != nil {
			d.l.Errorf("convert string %s failed, %+v", calory, err)
			return nil
		}
		carried += num
	}
	if carried > 0 {
		heap.Push(h, carried)
	}

	sum := 0
	for i := 0; i < 3; i++ {
		sum += heap.Pop(h).(int)
	}

	return sum
}

type Day1Heap struct {
	calory []int
}

func (h *Day1Heap) Len() int {
	return len(h.calory)
}
func (h *Day1Heap) Less(i, j int) bool {
	return h.calory[i] > h.calory[j]
}
func (h *Day1Heap) Swap(i, j int) {
	h.calory[i], h.calory[j] = h.calory[j], h.calory[i]
}
func (h *Day1Heap) Push(x interface{}) {
	h.calory = append(h.calory, x.(int))
}
func (h *Day1Heap) Pop() interface{} {
	x := h.calory[h.Len()-1]
	h.calory = h.calory[:h.Len()-1]
	return x
}
