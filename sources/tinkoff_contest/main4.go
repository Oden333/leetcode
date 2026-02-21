package main

import (
	"fmt"
)

func main4() {

	var n int
	fmt.Scan(&n)

	numbers := make([]uint64, n)
	for i := 0; i < n; i++ {
		fmt.Scan(&numbers[i])
	}

	// numbers := []uint64{30, 30, 20, 11, 12, 1}

	lomutoQuickSort(numbers)
	uniqueNumbers := removeDuplicates(numbers)

	/* 	// Выводим отсортированный и уникальный список
	   	fmt.Println("Отсортированный и уникальный список:")

	   	for _, num := range uniqueNumbers {
	   		fmt.Printf("%d ", num)
	   	}
	   	fmt.Println() */

	/* 	// Выводим с условием
	   	fmt.Println("Вывод с условием:") */
	printWithCondition(uniqueNumbers)
}

func removeDuplicates(nums []uint64) []uint64 {
	if len(nums) == 0 {
		return nums
	}
	i := 0
	for j := 1; j < len(nums); j++ {
		if nums[j] != nums[i] {
			i++
			nums[i] = nums[j]
		}
	}
	return nums[:i+1]
}

func printWithCondition(nums []uint64) {

	fmt.Print(nums[0])

	buf := [2]uint64{nums[0], nums[0]}
	lastprinted := nums[0]
	for i := 1; i < len(nums); i++ {

		if nums[i] == (buf[1] + 1) {
			buf[0] = buf[1]
			buf[1] = nums[i]
			if i+1 == len(nums) {
				if buf[1] == (buf[0]+1) && buf[0] != lastprinted {
					fmt.Print(" ... ", buf[1])

				} else if buf[1] == (buf[0]+1) && buf[0] == lastprinted {
					fmt.Print(" ", buf[1])
				}
			}
		} else {
			if buf[1] == (buf[0]+1) && buf[0] != lastprinted {
				fmt.Print(" ... ", buf[1])

			} else if buf[1] == (buf[0]+1) && buf[0] == lastprinted {
				fmt.Print(" ", buf[1])
			}
			fmt.Print(" ", nums[i])
			lastprinted = nums[i]
			buf[0] = nums[i]
			buf[1] = nums[i]
		}
	}
}

func lomutoQuickSort(arr []uint64) {
	n := len(arr)
	for i := 0; i < n-1; i++ {
		minIdx := i
		for j := i + 1; j < n; j++ {
			if arr[j] < arr[minIdx] {
				minIdx = j
			}
		}
		arr[i], arr[minIdx] = arr[minIdx], arr[i]

		lomutoQuickSort(arr[i+1:])
	}
}

func mergeSort(arr []uint64) []uint64 {
	if len(arr) <= 1 {
		return arr
	}

	mid := len(arr) / 2
	left := mergeSort(arr[:mid])
	right := mergeSort(arr[mid:])

	return merge(left, right)
}

func merge(left, right []uint64) []uint64 {
	result := make([]uint64, len(left)+len(right))
	i := 0
	j := 0

	for i < len(left) && j < len(right) {
		if left[i] < right[j] {
			result[i+j] = left[i]
			i++
		} else {
			result[i+j] = right[j]
			j++
		}
	}

	for i < len(left) {
		result[i+j] = left[i]
		i++
	}
	for j < len(right) {
		result[i+j] = right[j]
		j++
	}

	return result
}
