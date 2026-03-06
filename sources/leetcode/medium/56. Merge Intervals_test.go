package medium

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_merge(t *testing.T) {
	type args struct {
		intervals [][]int
	}
	tests := []struct {
		name string
		args args
		want [][]int
	}{
		{
			args: args{[][]int{{1, 3}, {2, 6}, {8, 10}, {15, 18}}},
			want: [][]int{{1, 6}, {8, 10}, {15, 18}},
		},
		{
			args: args{[][]int{{1, 4}, {4, 5}}},
			want: [][]int{{1, 5}},
		},
		{
			args: args{[][]int{{4, 7}, {1, 4}}},
			want: [][]int{{1, 7}},
		},
		{
			args: args{[][]int{{1, 4}, {2, 3}}},
			want: [][]int{{1, 4}},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			assert.Equalf(t, tt.want, merge2(tt.args.intervals), "merge(%v)", tt.args.intervals)
		})
	}
}
