package medium

import "slices"

func permute(nums []int) [][]int {
	// All the integers of nums are unique.
	res := make([][]int, 0, len(nums))

	var btr func(cur []int, set []int)
	btr = func(cur, set []int) {
		if len(set) == 0 {
			res = append(res, cur)
			return
		}

		for i, num := range set {
			/* append(slices.Clone(set[:i]), set[i+1:]...) */

			// сделать выбор
			set[0], set[i] = set[i], set[0]
			// backtrack(...)
			btr(append(slices.Clone(cur), num), set[1:])
			// отменить выбор
			set[i], set[0] = set[0], set[i]

		}

	}

	btr([]int{}, nums)
	return res
}
