// B. Приближенный двоичный поиск
// Ограничение времени	1 секунда
// Ограничение памяти	64Mb
// Ввод	стандартный ввод или input.txt
// Вывод	стандартный вывод или output.txt
// Для каждого из чисел второй последовательности найдите ближайшее к нему в первой.

// Формат ввода
// В первой строке входных данных содержатся числа N и K (). Во второй строке задаются N чисел первого массива,
// отсортированного по неубыванию, а в третьей строке – K чисел второго массива. Каждое число в обоих массивах по модулю не превосходит 2⋅109.

// Формат вывода
// Для каждого из K чисел выведите в отдельную строку число из первого массива, наиболее близкое к данному.
// Если таких несколько, выведите меньшее из них.

package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
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
	curStrippedString := strings.Split(curString, " ")

	firstInt, _ := strconv.Atoi(curStrippedString[0])
	secondInt, _ := strconv.Atoi(curStrippedString[1])

	s.Scan()
	curString = s.Text()
	curStrippedString = strings.Split(curString, " ")

	firstSlice := make([]int, 0, firstInt)
	for i := 0; i < firstInt; i++ {
		val, _ := strconv.Atoi(curStrippedString[i])
		firstSlice = append(firstSlice, val)
	}

	s.Scan()
	curString = s.Text()
	curStrippedString = strings.Split(curString, " ")

	secondSlice := make([]int, 0, secondInt)
	for i := 0; i < secondInt; i++ {
		val, _ := strconv.Atoi(curStrippedString[i])
		secondSlice = append(secondSlice, val)
	}

	final := strings.Builder{}
	for _, i := range secondSlice {
		final.WriteString(fmt.Sprint(closestNumberSearch(firstSlice, i)))
		final.WriteString("\n")
	}

	fmt.Println(final.String())
}

func closestNumberSearch(in []int, searchFor int) int {

	var first, last = 0, len(in) - 1
	var mid int
	for first <= last {
		mid = ((last - first) / 2) + first

		if in[mid] == searchFor {
			return in[mid]
		} else if in[mid] > searchFor {
			last = mid - 1
		} else if in[mid] < searchFor {
			first = mid + 1
		}
	}
	abs := func(i int) int {
		if i >= 0 {
			return i
		} else {
			return -i
		}
	}
	if last < 0 {
		last = 0
	}
	if first >= len(in) {
		first = len(in) - 1
	}
	if abs(in[first]-searchFor) < abs(in[last]-searchFor) {
		return in[first]
	} else {
		return in[last]
	}
}
