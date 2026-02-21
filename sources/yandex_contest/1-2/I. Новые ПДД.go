package main

import (
	"bufio"
	"fmt"
	"math"
	"strconv"
	"strings"
)

/* var (
	Input  io.Reader = os.Stdin
	Output io.Writer = os.Stdout
)

func main() {
	I()
} */

func I() {
	reader := bufio.NewReader(Input)
	nums, _ := reader.ReadString('\n')
	ns := strings.Split(nums, " ")
	x1, _ := strconv.Atoi(ns[0])
	y1, _ := strconv.Atoi(strings.TrimSpace(ns[1]))

	nums, _ = reader.ReadString('\n')
	ns = strings.Split(nums, " ")
	x2, _ := strconv.Atoi(ns[0])
	y2, _ := strconv.Atoi(strings.TrimSpace(ns[1]))

	res := (int(math.Abs(float64(x1-x2)))-boolInt(x1 != x2))*3 +
		(int(math.Abs(float64(y1-y2)))-boolInt(y1 != y2))*3 +
		boolInt((y1 != y2) && (x1 != x2))

	writer := bufio.NewWriterSize(Output, 1<<20)
	defer writer.Flush()
	writer.WriteString(fmt.Sprintf("%d", res))
	writer.WriteByte('\n')
}

func boolInt(q bool) int {
	if q {
		return 1
	}
	return 0
}
