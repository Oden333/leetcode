package medium

type dt struct {
	val, idx int
}

func dailyTemperatures(nums []int) []int {
	mStack := make([]dt, 0)
	var val int

	for j := len(nums) - 1; j >= 0; j-- {
		val = nums[j]
		for len(mStack) > 0 {
			if peek := mStack[len(mStack)-1]; peek.val > val {
				nums[j] = peek.idx - j
				break
			}

			mStack = mStack[:len(mStack)-1]
		}

		if len(mStack) == 0 {
			nums[j] = 0
		}

		mStack = append(mStack, dt{val, j})
	}

	return nums
}
