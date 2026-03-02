package medium

import (
	"fmt"
	"testing"
)

func Test_dailyTemperatures(t *testing.T) {
	tests := []struct {
		name string // description of this test case
		// Named input parameters for target function.
		nums []int
		want []int
	}{
		{
			nums: []int{73, 74, 75, 71, 69, 72, 76, 73},
			want: []int{1, 1, 4, 2, 1, 1, 0, 0},
		},
		{
			nums: []int{30, 40, 50, 60},
			want: []int{1, 1, 1, 0},
		},
		{
			nums: []int{30, 60, 90},
			want: []int{1, 1, 0},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := dailyTemperatures(tt.nums)
			if fmt.Sprintf("%v", tt.want) != fmt.Sprintf("%v", got) {
				t.Errorf("dailyTemperatures() = %v, want %v", got, tt.want)
			}
		})
	}
}
