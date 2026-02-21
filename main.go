package main

import (
	"fmt"
	// . "leetcode/Structs"
	. "leetcode/medium"
)

func main() {

	nums := [][]int{
		{2, 3, 1, 1, 4},
		{2, 3, 0, 1, 4},
	}
	for _, das := range nums {
		fmt.Println(Jump(das))
	}

	/*
		nums := [][]int{
			{3, 2, 1, 0, 4},
			{2, 3, 1, 1, 4},
		}
		for _, das := range nums {
			fmt.Println(CanJump(das))
		}
	*/

	/*
		values := []int{5, 8, 9, 2, 1, 3, 7, 4, 6}
		   	r := Fill(values)
		   	//fmt.Println(LevelOrder(r))
		   	fmt.Println(KthLargestLevelSum(r, 2))
	*/

	/*
		fmt.Println(LongestCommonPrefix([]string{"flower", "flow", "flight"}))
		fmt.Println(LongestCommonPrefix([]string{"ab", "a"}))
		fmt.Println(CommonChars([]string{"acabcddd", "bcbdbcbd", "baddbadb", "cbdddcac", "aacbcccd", "ccccddda", "cababaab", "addcaccd"}))
	*/

	/*
		fmt.Println(FindClosestNumber([]int{-4, -2, 1, 4, 8}))
		   	fmt.Println(FindClosestNumber([]int{2, -1, 1}))
	*/

	/* 	fmt.Println(Intersection([][]int{{3, 1, 2, 4, 5}, {1, 2, 3, 4}, {3, 4, 5, 6}}))
	   	nums := []int64{3, 1, 2, 4, 5}
	   	InsertionSort(nums)
	   	fmt.Println(nums) */
	//set.Gay()

	// type insdatat struct {
	// 	ggg
	// 	is_test   *bool
	// 	is_msater *bool
	// 	lang      *string
	// }

	// var q insdatat

	// fmt.Println(q.gr())
}

// type ggg struct {
// 	id int
// }

// func (g *ggg) gr() int {
// 	fmt.Println("")
// 	return 33
// }
