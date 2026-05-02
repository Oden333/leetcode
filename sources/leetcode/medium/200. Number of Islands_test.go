package medium

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestNumIslands(t *testing.T) {
	t.Parallel()

	tests := []struct {
		name string
		grid [][]byte
		want int
	}{
		{
			name: "empty grid",
			grid: [][]byte{},
			want: 0,
		},
		{
			name: "single water cell",
			grid: [][]byte{{'0'}},
			want: 0,
		},
		{
			name: "single land cell",
			grid: [][]byte{{'1'}},
			want: 1,
		},
		{
			name: "all water",
			grid: [][]byte{
				{'0', '0', '0'},
				{'0', '0', '0'},
				{'0', '0', '0'},
			},
			want: 0,
		},
		{
			name: "all land",
			grid: [][]byte{
				{'1', '1', '1'},
				{'1', '1', '1'},
				{'1', '1', '1'},
			},
			want: 1,
		},
		{
			// 1 1 0
			// 1 1 0
			// 0 0 1
			name: "two islands",
			grid: [][]byte{
				{'1', '1', '0'},
				{'1', '1', '0'},
				{'0', '0', '1'},
			},
			want: 2,
		},
		{
			// diagonal cells are NOT connected (only 4-directional adjacency)
			// 1 0
			// 0 1
			name: "diagonal cells not connected",
			grid: [][]byte{
				{'1', '0'},
				{'0', '1'},
			},
			want: 2,
		},
		{
			// leetcode example 1
			// 1 1 1 1 0
			// 1 1 0 1 0
			// 1 1 0 0 0
			// 0 0 0 0 0
			name: "one large island",
			grid: [][]byte{
				[]byte("11110"),
				[]byte("11010"),
				[]byte("11000"),
				[]byte("00000"),
			},
			want: 1,
		},
		{
			// leetcode example 2
			// 1 1 0 0 0
			// 1 1 0 0 0
			// 0 0 1 0 0
			// 0 0 0 1 1
			name: "three distinct islands",
			grid: [][]byte{
				[]byte("11000"),
				[]byte("11000"),
				[]byte("00100"),
				[]byte("00011"),
			},
			want: 3,
		},
		{
			// L-shaped island surrounded by water
			// 1 0 0
			// 1 0 0
			// 1 1 0
			name: "L-shaped island",
			grid: [][]byte{
				{'1', '0', '0'},
				{'1', '0', '0'},
				{'1', '1', '0'},
			},
			want: 1,
		},
		{
			// single row, alternating land and water
			// 1 0 1 0 1
			name: "single row alternating",
			grid: [][]byte{
				{'1', '0', '1', '0', '1'},
			},
			want: 3,
		},
		{
			// single column, alternating land and water
			// 1
			// 0
			// 1
			// 0
			// 1
			name: "single column alternating",
			grid: [][]byte{
				{'1'},
				{'0'},
				{'1'},
				{'0'},
				{'1'},
			},
			want: 3,
		},
		{
			// island forms a ring — interior water cell is surrounded but is water
			// 1 1 1
			// 1 0 1
			// 1 1 1
			name: "ring shaped island with water interior",
			grid: [][]byte{
				{'1', '1', '1'},
				{'1', '0', '1'},
				{'1', '1', '1'},
			},
			want: 1,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()
			got := numIslands(tt.grid)
			assert.Equal(t, tt.want, got)
		})
	}
}
