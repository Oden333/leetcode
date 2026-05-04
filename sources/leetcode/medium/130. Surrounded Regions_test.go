package medium

import (
	"runtime"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

func fmtBoard(b [][]byte) []string {
	rows := make([]string, len(b))
	for i, row := range b {
		rows[i] = string(row)
	}
	return rows
}

func Test_solve(t *testing.T) {
	t.Parallel()

	tests := []struct {
		name  string
		board [][]byte
		want  [][]byte
	}{
		{
			// X X X X      X X X X
			// X O O X  ->  X X X X
			// X X O X      X X X X
			// X O X X      X O X X
			name:  "surrounded region flipped, border O preserved",
			board: [][]byte{{'X', 'X', 'X', 'X'}, {'X', 'O', 'O', 'X'}, {'X', 'X', 'O', 'X'}, {'X', 'O', 'X', 'X'}},
			want:  [][]byte{{'X', 'X', 'X', 'X'}, {'X', 'X', 'X', 'X'}, {'X', 'X', 'X', 'X'}, {'X', 'O', 'X', 'X'}},
		},
		{
			// O connected to border through a chain — nothing flipped
			// X X X      X X X
			// O O O  ->  O O O
			// X X X      X X X
			name:  "O chain connected to left border",
			board: [][]byte{{'X', 'X', 'X'}, {'O', 'O', 'O'}, {'X', 'X', 'X'}},
			want:  [][]byte{{'X', 'X', 'X'}, {'O', 'O', 'O'}, {'X', 'X', 'X'}},
		},
		{
			// isolated O surrounded on all sides
			// X X X      X X X
			// X O X  ->  X X X
			// X X X      X X X
			name:  "single surrounded O",
			board: [][]byte{{'X', 'X', 'X'}, {'X', 'O', 'X'}, {'X', 'X', 'X'}},
			want:  [][]byte{{'X', 'X', 'X'}, {'X', 'X', 'X'}, {'X', 'X', 'X'}},
		},
		{
			name:  "no O cells",
			board: [][]byte{{'X', 'X'}, {'X', 'X'}},
			want:  [][]byte{{'X', 'X'}, {'X', 'X'}},
		},
		{
			name:  "1x1 O stays",
			board: [][]byte{{'O'}},
			want:  [][]byte{{'O'}},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()

			runtime.GC()
			var m1, m2 runtime.MemStats
			runtime.ReadMemStats(&m1)

			done := make(chan struct{}, 1)
			go func() { solve(tt.board); close(done) }()

			select {
			case <-done:
			case <-time.After(5 * time.Second):
				t.Fatal("TLE: exceeded 5s")
			}

			runtime.ReadMemStats(&m2)
			assert.LessOrEqual(t, m2.TotalAlloc-m1.TotalAlloc, uint64(50<<20), "MLE")
			assert.Equal(t, fmtBoard(tt.want), fmtBoard(tt.board))
		})
	}
}
