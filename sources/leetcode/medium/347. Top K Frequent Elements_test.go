package medium

import (
	"fmt"
	"testing"
)

func Test_topKFrequent(t *testing.T) {
	tests := []struct {
		name string // description of this test case
		// Named input parameters for target function.
		n    []int
		k    int
		want []int
	}{
		{
			n:    []int{1, 1, 1, 2, 2, 3},
			k:    2,
			want: []int{1, 2},
		},
		{
			n:    []int{1, 2, 1, 2, 1, 2, 3, 1, 3, 2},
			k:    2,
			want: []int{1, 2},
		},
		{
			n:    []int{1},
			k:    1,
			want: []int{1},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := topKFrequent(tt.n, tt.k)
			// TODO: update the condition below to compare got with tt.want.
			if fmt.Sprintf("%v", tt.want) != fmt.Sprintf("%v", got) {
				t.Errorf("topKFrequent() = %v, want %v", got, tt.want)
			}
		})
	}
}
