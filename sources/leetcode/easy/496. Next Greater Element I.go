package easy

//? For each 0 <= i < nums1.length, find the index j such that nums1[i] == nums2[j]
//? and determine the next greater element of nums2[j] in nums2.
//
//? If there is no next greater element, then the answer for this query is -1.

func nextGreaterElement(nums1 []int, nums2 []int) []int {
	res := make(map[int]int, len(nums2))

	mStack := make([]int, 0)

	for j := len(nums2) - 1; j >= 0; j-- {
		val := nums2[j]

		for len(mStack) > 0 {
			if svo := peek(mStack); svo > val {
				res[val] = svo
				break
			} else {
				mStack = pop(mStack)
			}
		}

		if len(mStack) == 0 {
			res[val] = -1
		}

		mStack = append(mStack, val)
	}

	for i := 0; i < len(nums1); i++ {
		if k, ok := res[nums1[i]]; ok {
			nums1[i] = k
		} else {
			nums1[i] = -1
		}
	}

	return nums1
}

func peek(n []int) int {
	return n[len(n)-1]
}

func pop(n []int) []int {
	return n[:len(n)-1]
}
