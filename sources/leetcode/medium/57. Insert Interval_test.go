package medium

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_insert(t *testing.T) {
	type args struct {
		intervals [][]int
	}
	tests := []struct {
		name        string
		args        args
		want        [][]int
		newInterval []int
	}{
		{
			args:        args{[][]int{{1, 3}, {6, 9}}},
			newInterval: []int{2, 5},
			want:        [][]int{{1, 5}, {6, 9}},
		},
		{
			args:        args{[][]int{{1, 2}, {3, 5}, {6, 7}, {8, 10}, {12, 16}}},
			newInterval: []int{4, 8},
			want:        [][]int{{1, 2}, {3, 10}, {12, 16}},
		},
		{
			args:        args{[][]int{{1, 5}}},
			newInterval: []int{2, 3},
			want:        [][]int{{1, 5}},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			assert.Equalf(t, tt.want, insert(tt.args.intervals, tt.newInterval), "merge(%v)", tt.args.intervals)
		})
	}
}
