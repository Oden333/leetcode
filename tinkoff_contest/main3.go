package main

import "fmt"

func findHeight(nums []int64) float32 {
	var count int64
	count = 0
	var sum float64
	sum = 0

	for i := 0; i < len(nums)-1; i++ {
		sum += float64(nums[i]+nums[i+1]) / 2
		count++
	}

	if count == 0 {
		return 0
	}

	return float32(sum / float64(count))
}

func main() {
	var n int
	fmt.Scan(&n)

	numbers := make([]int64, n)
	for i := 0; i < n; i++ {
		fmt.Scan(&numbers[i])
	}

	result := findHeight(numbers)

	fmt.Printf("%f", result)
}
