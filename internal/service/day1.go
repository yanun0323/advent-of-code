package service

import (
	"container/heap"
	"strconv"
)

func (svc Service) Day1A(calories []string) any {
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
			svc.l.Errorf("convert string %s failed, %+v", calory, err)
			return nil
		}
		temp += num
	}

	return max
}

func (svc Service) Day1B(calories []string) any {
	h := &Heap2022Day1{}
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
			svc.l.Errorf("convert string %s failed, %+v", calory, err)
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

type Heap2022Day1 struct {
	calory []int
}

func (h *Heap2022Day1) Len() int {
	return len(h.calory)
}
func (h *Heap2022Day1) Less(i, j int) bool {
	return h.calory[i] > h.calory[j]
}
func (h *Heap2022Day1) Swap(i, j int) {
	h.calory[i], h.calory[j] = h.calory[j], h.calory[i]
}
func (h *Heap2022Day1) Push(x interface{}) {
	h.calory = append(h.calory, x.(int))
}
func (h *Heap2022Day1) Pop() interface{} {
	x := h.calory[h.Len()-1]
	h.calory = h.calory[:h.Len()-1]
	return x
}
