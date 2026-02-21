package main

import (
	"bufio"
	"strconv"
	"strings"
)

func A() {
	//   Пример ввода и вывода числа n, где -10^9 < n < 10^9:
	reader := bufio.NewReaderSize(Input, 1<<20)
	num, _ := reader.ReadString('\n')
	n, _ := strconv.Atoi(strings.TrimSpace(num))
	var (
		minOdd  = 1001
		maxEven = 0
		counter = 1
		sum     = 0
		q       int
	)
	for range n {
		num, _ = reader.ReadString(' ')
		q, _ = strconv.Atoi(strings.TrimSpace(num))
		if counter%2 == 0 {
			if q > maxEven {
				maxEven = q
			}
			sum -= q
		} else {
			if q < minOdd {
				minOdd = q
			}
			sum += q
		}
		counter++
	}
	if maxEven > minOdd {
		sum = sum + 2*(maxEven-minOdd)
	}

	writer := bufio.NewWriterSize(Output, 1<<20)
	defer writer.Flush()

	writer.WriteString(strconv.Itoa(sum))
	writer.WriteByte('\n')
}
