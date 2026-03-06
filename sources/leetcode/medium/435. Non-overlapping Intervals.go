package medium

import (
	"slices"
)

func eraseOverlapIntervals(intervals [][]int) int {
	if len(intervals) <= 1 {
		return 0
	}

	slices.SortFunc(intervals, func(a, b []int) int {
		switch {
		case a[0] < b[0]:
			return -1
		case a[0] == b[0]:
			switch {
			case a[1] < b[1]:
				return -1
			case a[1] == b[1]:
				return 0
			case a[1] > b[1]:
				return 1
			}
		}
		return 1
	})

	var (
		res  int
		last = slices.Clone(intervals[0])
	)

	for _, b := range intervals[1:] {
		if last[1] > b[0] {
			last[1] = min(last[1], b[1])
			res++
			continue
		}

		last = b
	}

	return res
}
