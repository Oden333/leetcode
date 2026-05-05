package structs

import (
	"fmt"
	"testing"
)

func TestMinHeap_1(t *testing.T) {
	// input := []int{5, 3, 1, 7, 2, 9}
	// expected := []int{1, 2, 3, 5, 7, 9}

	h := &MinHeap{}
	for _, v := range []int{5, 3, 1, 7, 2} {
		h.Push(v)
	}

	for i := 0; i < 5; i++ {
		fmt.Println(h.Pop()) // Должно вывести: 1 2 3 5 7
	}
}
