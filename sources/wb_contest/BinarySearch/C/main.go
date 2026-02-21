package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
	"sync"
)

func main() {

	f, _ := os.Open("input.txt")
	defer f.Close()

	s := bufio.NewScanner(f)
	s.Split(bufio.ScanLines)

	const maxCapacity int = 10000000 // your required line length
	buf := make([]byte, maxCapacity)
	s.Buffer(buf, maxCapacity)

	s.Scan()
	curString := s.Text()
	curInt, _ := strconv.Atoi(curString)

	s.Scan()
	curString = s.Text()
	curStrippedString := strings.Split(curString, " ")

	firstSlice := make([]int, 0, curInt)
	for i := 0; i < curInt; i++ {
		val, _ := strconv.Atoi(curStrippedString[i])
		firstSlice = append(firstSlice, val)
	}

	s.Scan()
	curString = s.Text()
	curInt, _ = strconv.Atoi(curString)

	quickSort(firstSlice, 0, len(firstSlice)-1)

	final := strings.Builder{}

	ranges := make([][2]int, curInt)
	for i := 0; i < curInt; i++ {
		s.Scan()
		curString = s.Text()
		curStrippedString := strings.Split(curString, " ")
		left, _ := strconv.Atoi(curStrippedString[0])
		right, _ := strconv.Atoi(curStrippedString[1])
		ranges[i][0] = left
		ranges[i][1] = right
	}

	wg := sync.WaitGroup{}
	mu := sync.Mutex{}
	//fmt.Println(firstSlice)
	ansMap := make(map[int]int)
	batch := 0
	for i := 0; i < len(ranges); i += 2 {

		wg.Add(1)
		if i+2 < len(ranges) {
			batch = i + 2

		} else {
			batch = len(ranges)
		}

		go func(w *sync.WaitGroup, ranges [][2]int, idx int) {
			counter := 0
			for i := 0; i < len(ranges); i++ {
				leftIdx := closestLeftIdxSearch(firstSlice, ranges[i][0])
				rightIdx := closestRightIdxSearch(firstSlice, ranges[i][1])
				//fmt.Println(ranges[i], " indexes are ", leftIdx, rightIdx, "ans", rightIdx-leftIdx+1)

				if leftIdx == -1 || rightIdx == -1 {
					counter = 0
				} else {
					counter = rightIdx - leftIdx + 1
				}
				mu.Lock()
				//final.WriteString(fmt.Sprint(counter))
				//final.WriteString("\n")
				ansMap[i+idx] = counter
				mu.Unlock()

			}
			w.Done()
		}(&wg, ranges[i:batch], i)
	}
	wg.Wait()
	for i := 0; i < len(ranges); i++ {
		value := ansMap[i]
		final.WriteString(strconv.Itoa(value) + "\n")
	}

	fmt.Println(final.String())

}

func closestLeftIdxSearch(in []int, searchFor int) int {

	if searchFor > in[len(in)-1] {
		return -1
	}
	var left, right = 0, len(in) - 1
	var mid int
	for left <= right {
		mid = ((right - left) / 2) + left

		if in[mid] == searchFor {
			for (mid > 0) && (in[mid-1] == searchFor) {
				mid--
			}
			return mid
		} else if in[mid] > searchFor {
			right = mid - 1
		} else if in[mid] < searchFor {
			left = mid + 1
		}
	}

	if left >= len(in) {
		left = len(in) - 1
	}
	return left

}
func closestRightIdxSearch(in []int, searchFor int) int {

	if searchFor < in[0] {
		return -1
	}
	var left, right = 0, len(in) - 1
	var mid int
	for left <= right {
		mid = ((right - left) / 2) + left

		if in[mid] == searchFor {
			for (mid < len(in)-1) && (in[mid+1] == searchFor) {
				mid++
			}
			return mid
		} else if in[mid] > searchFor {
			right = mid - 1
		} else if in[mid] < searchFor {
			left = mid + 1
		}
	}

	if right < 0 {
		right = 0
	}

	return right
}

func quickSort(arr []int, low, high int) {
	if low < high {
		pivotIndex := partition(arr, low, high)

		quickSort(arr, low, pivotIndex-1)
		quickSort(arr, pivotIndex+1, high)
	}
}

func swap(arr []int, i, j int) {
	temp := arr[i]
	arr[i] = arr[j]
	arr[j] = temp
}

func partition(arr []int, low, high int) int {
	pivot := arr[high]

	i := low - 1

	for j := low; j < high; j++ {
		if arr[j] <= pivot {
			i++
			arr[i], arr[j] = arr[j], arr[i]
		}
	}

	swap(arr, i+1, high)

	return i + 1
}
