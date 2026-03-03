package structs

// Heap — это бинарное дерево, но хранится в слайсе:
type MinHeap []int

// Len, Less, Swap, Push, Pop

func (h MinHeap) Len() int {
	return len(h)
}

func (h MinHeap) Less(i, j int) bool {
	return h[i] < h[j]
}

func (h MinHeap) Swap(i, j int) {
	h[i], h[j] = h[j], h[i]
}

// Push — сравнивает новый элемент с родителем (решает, всплывать ли)
func (h *MinHeap) Push(x any) {
	(*h) = append((*h), x.(int))
}

// Pop — сравнивает детей (решает, кто станет новым корнем)
func (h *MinHeap) Pop() any {
	val := (*h)[h.Len()-1]
	(*h) = (*h)[:h.Len()-1]
	return val
}
