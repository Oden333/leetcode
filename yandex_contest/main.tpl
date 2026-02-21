package main

import (
	"bufio"
	"io"
	"os"
	"strconv"
	"strings"
)

var (
	Input  io.Reader = os.Stdin
	Output io.Writer = os.Stdout
)

func T() {
	reader := bufio.NewReaderSize(Input, 1<<20)
	num, _ := reader.ReadString('\n')
	n, _ := strconv.Atoi(strings.TrimSpace(num))

	writer := bufio.NewWriterSize(Output, 1<<20)
	defer writer.Flush()

	writer.WriteString(strconv.Itoa(n))
	writer.WriteByte('\n')

}
