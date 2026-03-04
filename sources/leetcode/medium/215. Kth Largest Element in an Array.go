package medium

import "container/heap"

func findKthLargest(nums []int, k int) int {
	f := new(fkl)

	for _, v := range nums {
		heap.Push(f, v)
	}

	var res int
	for k > 0 {
		res = heap.Pop(f).(int)
		k--
	}

	return res
}

// maxheap
type fkl []int

func (h fkl) Len() int {
	return len(h)
}
func (h fkl) Less(i, j int) bool {
	return h[i] > h[j]
}

func (h fkl) Swap(i, j int) {
	h[i], h[j] = h[j], h[i]
}

func (h *fkl) Push(x any) {
	(*h) = append((*h), x.(int))
}

func (h *fkl) Pop() any {

	res := (*h)[h.Len()-1]
	(*h) = (*h)[:h.Len()-1]
	return res
}
