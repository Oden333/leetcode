package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"strings"
)

var (
	Input  io.Reader = os.Stdin
	Output io.Writer = os.Stdout
)

func main() {
	B()
}
func B() {
	reader := bufio.NewReaderSize(Input, 405)
	a, _ := reader.ReadString('\n')
	a = strings.TrimSpace(a)
	var (
		pos byte = 'L'
		// 0 = L
		// 1 = R

		sum int = 0
	)
	switch {
	case len(a) == 1:
		sum = 1
		if a[0] == 'B' {
			sum = 2
		}
	case len(a) == 2:
		sum += last[string(pos)+a[len(a)-2:]]
	default:
		for i := 0; i < len(a)-2; i++ {
			if a[i]&a[i+1]&pos == pos {
				cross(&pos)
				sum++
				continue
			}
			if a[i] == pos || a[i] == 'B' {
				sum++
				continue
			}
		}

		sum += last[string(pos)+a[len(a)-2:]]
	}

	writer := bufio.NewWriterSize(Output, 5)
	defer writer.Flush()

	writer.WriteString(fmt.Sprint(sum))
	writer.WriteByte('\n')
}

func cross(pos *byte) {
	if *pos == 'L' {
		*pos = 'R'
		return
	}
	*pos = 'L'
}

var (
	last = map[string]int{
		"LLL": 1,
		"LRL": 1,
		"LLR": 2,
		"LRR": 1,
		"LBL": 2,
		"LBR": 2,
		"LLB": 2,
		"LRB": 2,
		"LBB": 3,

		"RLL": 0,
		"RRL": 1,
		"RLR": 1,
		"RRR": 2,
		"RBL": 1,
		"RBR": 2,
		"RLB": 1,
		"RRB": 2,
		"RBB": 2,
	}
)
