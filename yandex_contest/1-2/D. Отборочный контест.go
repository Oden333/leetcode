package main

import (
	"bufio"
	"strconv"
)

func D() {
	reader := bufio.NewReader(Input)

	s := bufio.NewScanner(reader)
	s.Split(bufio.ScanWords)
	s.Scan()
	arrsize, _ := strconv.Atoi(s.Text())

	s.Scan()
	len, _ := strconv.Atoi(s.Text())

	m := map[string]int{}
	for range arrsize {
		s.Scan()
		str := s.Text()

		m[str]++
	}

	writer := bufio.NewWriter(Output)
	defer writer.Flush()

	for len > 0 {
		for k, v := range m {
			if v == 0 {
				continue
			}

			writer.WriteString(k)
			len--
			m[k]--

			if len == 0 {
				break
			}

			writer.WriteString(" ")
		}
	}
}
