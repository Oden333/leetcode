package medium

import (
	"runtime"
	"sort"
	"time"
	"testing"

	"github.com/stretchr/testify/assert"
)

func sortPerms(perms [][]int) {
	for _, p := range perms {
		sort.Ints(p)
	}
	sort.Slice(perms, func(i, j int) bool {
		a, b := perms[i], perms[j]
		for k := range a {
			if k >= len(b) {
				return false
			}
			if a[k] != b[k] {
				return a[k] < b[k]
			}
		}
		return len(a) < len(b)
	})
}

func TestPermute(t *testing.T) {
	t.Parallel()

	tests := []struct {
		name string
		nums []int
		want [][]int
	}{
		{
			name: "single element",
			nums: []int{1},
			want: [][]int{{1}},
		},
		{
			name: "two elements",
			nums: []int{1, 2},
			want: [][]int{
				{1, 2},
				{2, 1},
			},
		},
		{
			name: "three elements — all 6 permutations",
			nums: []int{1, 2, 3},
			want: [][]int{
				{1, 2, 3},
				{1, 3, 2},
				{2, 1, 3},
				{2, 3, 1},
				{3, 1, 2},
				{3, 2, 1},
			},
		},
		{
			name: "negative numbers",
			nums: []int{-1, 2, -3},
			want: [][]int{
				{-3, -1, 2},
				{-3, 2, -1},
				{-1, -3, 2},
				{-1, 2, -3},
				{2, -3, -1},
				{2, -1, -3},
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()

			runtime.GC()
			var m1, m2 runtime.MemStats
			runtime.ReadMemStats(&m1)

			type result struct{ val [][]int }
			ch := make(chan result, 1)
			go func() {
				ch <- result{permute(tt.nums)}
			}()

			var got [][]int
			select {
			case r := <-ch:
				got = r.val
			case <-time.After(5 * time.Second):
				t.Fatal("TLE: exceeded 5s")
			}

			runtime.ReadMemStats(&m2)
			assert.LessOrEqual(t, m2.TotalAlloc-m1.TotalAlloc, uint64(50<<20), "MLE")

			sortPerms(got)
			sortPerms(tt.want)
			assert.Equal(t, tt.want, got)
		})
	}
}
