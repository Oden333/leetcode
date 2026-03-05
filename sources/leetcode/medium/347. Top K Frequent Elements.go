package medium

import "container/heap"

func topKFrequent(n []int, k int) []int {
	nums := make(map[int]int)
	for _, v := range n {
		nums[v]++
	}

	h := tkf{}
	heap.Init(&h)

	for k, v := range nums {
		heap.Push(&h, tfkel{k, v})
	}

	res := make([]int, 0)
	for k > 0 {
		res = append(res, heap.Pop(&h).(tfkel).val)
		k--
	}
	return res
}

//	type Interface interface {
//		sort.Interface
//		Push(x any) // add x as element Len()
//		Pop() any   // remove and return element Len() - 1.
//	}

type tfkel struct {
	val, count int
}

type tkf []tfkel

func (h tkf) Len() int {
	return len(h)
}

func (h tkf) Less(i, j int) bool {
	return h[i].count > h[j].count
}

func (h *tkf) Swap(i, j int) {
	(*h)[i], (*h)[j] = (*h)[j], (*h)[i]
}

func (h *tkf) Push(x any) {
	(*h) = append((*h), x.(tfkel))
}

func (h *tkf) Pop() any {
	res := (*h)[h.Len()-1]
	(*h) = (*h)[:h.Len()-1]
	return res
}
