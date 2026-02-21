package main

import (
	"bufio"
	"fmt"
	"strconv"
)

func E() {
	reader := bufio.NewReaderSize(Input, 20)

	s := bufio.NewScanner(reader)
	const maxCapacity int = 20
	buf := make([]byte, maxCapacity)
	s.Buffer(buf, maxCapacity)
	s.Split(bufio.ScanWords)

	s.Scan()
	n, _ := strconv.Atoi(s.Text())

	s.Scan()
	k, _ := strconv.Atoi(s.Text())
	writer := bufio.NewWriter(Output)
	defer writer.Flush()
	if k == 0 {
		writer.WriteString(fmt.Sprintf("%d", n))
		return
	}
	if n%5 == 0 {
		writer.WriteString(fmt.Sprintf("%d", n+n%10))
		return
	}
	if (n%2 != 0) && (k%4 == 0) {
		n += n % 10
		k--
	}
	for range k % 4 {
		n += n % 10
	}
	n += (k / 4) * 20

	writer.WriteString(fmt.Sprintf("%d", n))
	writer.WriteByte('\n')
}
