package easy

type NumArray_ struct {
	arr []int
}

func Constructor(nums []int) NumArray_ {
	return NumArray_{arr: nums}
}

func (this *NumArray_) SumRange(l int, r int) int {
	var res int
	for l < r {
		res += this.arr[l] + this.arr[r]
		l++
		r--
	}
	if l == r {
		res += this.arr[l]
	}
	return res
}
