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

func Factorial(n int) uint64 {
	if n < 0 {
		return 0
	}
	if n == 0 {
		return 1
	}

	result := uint64(1)
	for i := 1; i <= n; i++ {
		result *= uint64(i)
	}
	return result
}
func main() {
	reader := bufio.NewReaderSize(Input, 1<<20)
	num, _ := reader.ReadString('\n')
	num = strings.TrimSpace(num)

	m := make(map[rune]int, 25)
	for _, letter := range num {
		m[letter]++
	}
	var x uint64 = 1

	for _, v := range m {
		x *= Factorial(v)
	}

	res := Factorial(len(num)) / x

	writer := bufio.NewWriterSize(Output, 1<<20)
	defer writer.Flush()

	writer.WriteString(fmt.Sprintf("%d", res))
	writer.WriteByte('\n')

}
