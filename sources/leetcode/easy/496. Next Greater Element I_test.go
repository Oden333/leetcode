package easy

import (
	"fmt"
	"testing"
)

func Test_nextGreaterElement(t *testing.T) {
	tests := []struct {
		name string // description of this test case
		// Named input parameters for target function.
		nums1 []int
		nums2 []int
		want  []int
	}{
		{
			// Input: nums1 = [4,1,2], nums2 = [1,3,4,2]п
			// Output: [-1,3,-1]

			nums1: []int{4, 1, 2},
			nums2: []int{1, 3, 4, 2},
			want:  []int{-1, 3, -1},
		},

		{
			// Input: nums1 = [2,4], nums2 = [1,2,3,4]
			// Output: [3,-1]

			nums1: []int{2, 4},
			nums2: []int{1, 2, 3, 4},
			want:  []int{3, -1},
		},
		{
			nums1: []int{1, 3, 5, 2, 4},
			nums2: []int{6, 5, 4, 3, 2, 1, 7},
			want:  []int{7, 7, 7, 7, 7},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := nextGreaterElement(tt.nums1, tt.nums2)
			if fmt.Sprintf("%v", tt.want) != fmt.Sprintf("%v", got) {
				t.Errorf("nextGreaterElement() = %v, want %v", got, tt.want)
			}
		})
	}
}
