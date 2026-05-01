package easy

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestFloodFill(t *testing.T) {
	t.Parallel()

	tests := []struct {
		name  string
		setup func() [][]int
		sr    int
		sc    int
		color int
		want  [][]int
	}{
		{
			// before:        after:
			// 1 1 1          2 2 2
			// 1 1 0    →    2 2 0
			// 1 0 1          2 0 1
			name: "basic fill",
			setup: func() [][]int {
				return [][]int{{1, 1, 1}, {1, 1, 0}, {1, 0, 1}}
			},
			sr:    1,
			sc:    1,
			color: 2,
			want:  [][]int{{2, 2, 2}, {2, 2, 0}, {2, 0, 1}},
		},
		{
			// 0 0 0   color=0, sr=0 sc=0 — цвет уже совпадает, ничего не меняется
			// 0 0 0
			name: "same color as target, no change",
			setup: func() [][]int {
				return [][]int{{0, 0, 0}, {0, 0, 0}}
			},
			sr:    0,
			sc:    0,
			color: 0,
			want:  [][]int{{0, 0, 0}, {0, 0, 0}},
		},
		{
			// before: 5   after: 3
			name: "single pixel image",
			setup: func() [][]int {
				return [][]int{{5}}
			},
			sr:    0,
			sc:    0,
			color: 3,
			want:  [][]int{{3}},
		},
		{
			// before:        after:
			// 2 1 2          3 1 2
			// 1 1 1     →    1 1 1
			// 2 1 2          2 1 2
			// старт [0][0]=2, окружён цветом 1 — заливается только [0][0]
			name: "isolated start pixel",
			setup: func() [][]int {
				return [][]int{{2, 1, 2}, {1, 1, 1}, {2, 1, 2}}
			},
			sr:    0,
			sc:    0,
			color: 3,
			want:  [][]int{{3, 1, 2}, {1, 1, 1}, {2, 1, 2}},
		},
		{
			// before:     after:
			// 1 1    →   5 5
			// 1 1         5 5
			name: "full image fill",
			setup: func() [][]int {
				return [][]int{{1, 1}, {1, 1}}
			},
			sr:    0,
			sc:    0,
			color: 5,
			want:  [][]int{{5, 5}, {5, 5}},
		},
		{
			// before:        after:
			// 1 1 0          2 2 0
			// 1 0 0    →    2 0 0
			// старт из угла [0][0], заливка останавливается на нулях
			name: "start from corner",
			setup: func() [][]int {
				return [][]int{{1, 1, 0}, {1, 0, 0}}
			},
			sr:    0,
			sc:    0,
			color: 2,
			want:  [][]int{{2, 2, 0}, {2, 0, 0}},
		},
		{
			// before: 1 1 1 0 1
			// after:  4 4 4 0 1
			// одна строка, заливка слева до разрыва на нуле
			name: "single row",
			setup: func() [][]int {
				return [][]int{{1, 1, 1, 0, 1}}
			},
			sr:    0,
			sc:    0,
			color: 4,
			want:  [][]int{{4, 4, 4, 0, 1}},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()
			got := floodFill(tt.setup(), tt.sr, tt.sc, tt.color)
			assert.Equal(t, tt.want, got)
		})
	}
}
