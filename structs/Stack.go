package structs

// Stack — простая реализация стека на основе слайса.
// LIFO: Last In, First Out.
type Stack []int

// Push добавляет элемент на вершину стека. O(1) амортизированно.
func (s *Stack) Push(val int) {
	*s = append(*s, val)
}

// Pop удаляет и возвращает элемент с вершины стека
func (s *Stack) Pop() int {
	idx := len(*s) - 1
	val := (*s)[idx]
	*s = (*s)[:idx]
	return val
}

// Peek возвращает верхний элемент без удаления.
func (s *Stack) Peek() int {
	return (*s)[len(*s)-1]
}

// IsEmpty проверяет, пуст ли стек.
func (s *Stack) IsEmpty() bool {
	return s == nil || len(*s) == 0
}

// Size возвращает количество элементов.
func (s *Stack) Size() int {
	if s == nil {
		return 0
	}
	return len(*s)
}
