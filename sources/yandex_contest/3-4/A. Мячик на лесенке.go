package main

import (
	"bufio"
	"strconv"
	"strings"
)

/*
var (

	Input  io.Reader = os.Stdin
	Output io.Writer = os.Stdout

)

	func main() {
		A()
	}
*/
func A() {
	reader := bufio.NewReaderSize(Input, 8)
	num, _ := reader.ReadString('\n')
	n, _ := strconv.Atoi(strings.TrimSpace(num))

	stairs := make([]int, n+1)
	stairs[0] = 1
	for i := range n + 1 {
		if i-1 >= 0 {
			stairs[i] += stairs[i-1]
		}
		if i-2 >= 0 {
			stairs[i] += stairs[i-2]
		}
		if i-3 >= 0 {
			stairs[i] += stairs[i-3]
		}
	}

	writer := bufio.NewWriterSize(Output, 1<<20)
	defer writer.Flush()

	writer.WriteString(strconv.Itoa(stairs[n]))
	writer.WriteByte('\n')
}
