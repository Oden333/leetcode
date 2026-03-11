package medium

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_findMin(t *testing.T) {
	type args struct {
		in []int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			args: args{in: []int{3, 4, 5, 1, 2}},
			want: 1,
		},
		{
			args: args{in: []int{4, 5, 6, 7, 0, 1, 2}},
			want: 0,
		},
		{
			args: args{in: []int{11, 13, 15, 17}},
			want: 11,
		}, {
			args: args{in: []int{2, 1}},
			want: 1,
		}, {
			args: args{in: []int{3, 1, 2}},
			want: 1,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			assert.Equalf(t, tt.want, findMin(tt.args.in), "findMin(%v)", tt.args.in)
		})
	}
}
