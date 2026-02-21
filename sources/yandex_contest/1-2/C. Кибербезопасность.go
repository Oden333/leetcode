package main

import (
	"bufio"
	"fmt"
)

func C() {
	reader := bufio.NewReaderSize(Input, 1<<20)
	m := make([]int, 26)
	for {
		b, err := reader.ReadByte()
		if err != nil {
			break
		}
		if b > 122 || b < 97 {
			continue
		}
		m[int(b)%26]++
	}

	res := 1
	tempSum := 0
	for _, v := range m {
		if tempSum == 0 {
			tempSum = v
			continue
		}

		res += v * tempSum
		tempSum += v
	}

	/* 	res := 1
	   	for i := 0; i < len(m); i++ {
	   		for j := i + 1; j < len(m); j++ {
	   			res += m[i] * m[j]
	   		}
	   	}
	*/

	writer := bufio.NewWriter(Output)
	defer writer.Flush()

	writer.WriteString(fmt.Sprintf("%d", res))
	writer.WriteByte('\n')
}
