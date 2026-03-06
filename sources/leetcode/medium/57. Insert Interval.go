package medium

import (
	"cmp"
	"slices"
)

func insert(intervals [][]int, newInterval []int) [][]int {
	if len(intervals) == 0 {
		return append(intervals, newInterval)
	}

	slices.SortFunc(intervals, func(a, b []int) int {
		return cmp.Compare(a[0], b[0])
	})

	var i int = len(intervals)
	if tmp := slices.IndexFunc(intervals, func(a []int) bool {
		if newInterval[0] <= a[0] {
			return true
		}
		return false
	}); tmp >= 0 {
		i = tmp
	}
	intervals = slices.Insert(intervals, i, newInterval)

	var (
		res  = make([][]int, 0)
		last = slices.Clone(intervals[0])
	)

	for _, b := range intervals[1:] {
		// набираем интервал
		if last[1] >= b[0] {
			last[1] = max(last[1], b[1])
			continue
		}

		// дропаем и идём дальше
		res = append(res, last)
		last = b
	}

	// последний интервал докинем
	res = append(res, last)

	return res
}
