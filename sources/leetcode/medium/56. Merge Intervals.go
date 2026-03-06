package medium

import (
	"cmp"
	"slices"
)

func merge(intervals [][]int) [][]int {
	if len(intervals) <= 1 {
		return intervals
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
		res  = make([][]int, 0)
		last = intervals[0]
		j    = 1
	)

	for j < len(intervals) {
		b := intervals[j]
		if last[1] >= b[0] {
			if last[1] < b[1] {
				last[1] = b[1]
			}
		} else {
			res = append(res, last)
			last = b
		}
		j++
	}
	res = append(res, last)

	return res
}

func merge2(intervals [][]int) [][]int {
	if len(intervals) <= 1 {
		return intervals
	}

	slices.SortFunc(intervals, func(a, b []int) int {
		return cmp.Compare(a[0], b[0])
	})

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
