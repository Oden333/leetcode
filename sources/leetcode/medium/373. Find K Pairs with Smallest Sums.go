package medium

import "container/heap"

func kSmallestPairs(nums1 []int, nums2 []int, k int) [][]int {
	h1 := new(ksp)
	for _, v := range nums1 {
		heap.Push(h1, v)
	}

	h2 := new(ksp)
	for _, v := range nums2 {
		heap.Push(h2, v)
	}

	res := make([][]int, k)
	for i := range k {
		res[i] = []int{heap.Pop(h1).(int), heap.Pop(h2).(int)}
	}

	return res
}

/*
type Interface interface { // size=16 (0x10)

	    sort.Interface
	    Push(x any) // add x as element Len()
	    Pop() any   // remove and return element Len() - 1.
	}

func (sort.Interface) Len() int
func (sort.Interface) Less(i int, j int) bool
func (sort.Interface) Swap(i int, j int)
*/

// min heap
type ksp []int

func (h ksp) Len() int {
	return len(h)
}

func (h ksp) Less(i int, j int) bool {
	return h[i] <= h[j]
}

func (h ksp) Swap(i int, j int) {
	h[i], h[j] = h[j], h[i]
}

func (h *ksp) Push(x any) {
	(*h) = append((*h), x.(int))
}

func (h *ksp) Pop() any {
	val := (*h)[h.Len()-1]
	(*h) = (*h)[:h.Len()-1]
	return val
}
