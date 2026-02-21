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
		// fmt.Printf("%d\n", -(1<<63 - 1))
		// fmt.Printf("%b", 1<<63-1)
		F()
	}
*/
func F() {
	reader := bufio.NewReaderSize(Input, 1<<20)
	in, _ := reader.ReadString('\n')
	nums := strings.Split(in, " ")
	strs, _ := strconv.Atoi(nums[0])
	cols, _ := strconv.Atoi(strings.TrimSpace(nums[1]))

	var rowCount = make(map[int][]bool, strs)

	var maxRow int = -(1<<63 - 1)
	var maxIdx int

	var minCol = make([]int, cols)

	var curRow int
	var idx int

	for i := 0; i < strs; {
		b, err := reader.ReadByte()
		if b == '\n' || err != nil {
			if maxRow < curRow {
				maxRow = curRow
				maxIdx = i
			}

			curRow = 0
			idx = 0
			i++
			continue
		}

		switch b {
		case '+':
			curRow++
			minCol[idx]++
		case '-':
			minCol[idx]--
			curRow--
		case '?':
			minCol[idx]--
			curRow++
			if rowCount[i] == nil {
				rowCount[i] = make([]bool, cols)
			}
			rowCount[i][idx] = true
		}
		idx++
	}

	var minC = minCol[0]
	if z, ok := rowCount[maxIdx]; ok {
		if len(z) > 0 && z[0] {
			minC += 2
		}
	}

	for i, x := range minCol {
		if z, ok := rowCount[maxIdx]; ok {
			if len(z) >= i && z[i] {
				x += 2
			}
		}

		if minC > x {
			minC = x
		}
	}

	writer := bufio.NewWriterSize(Output, 10)
	defer writer.Flush()

	// if maxRow > minC {
	writer.WriteString(strconv.Itoa(maxRow - minC))
	// } else {
	// writer.WriteString(strconv.Itoa(minC - maxRow))
	// }
	writer.WriteByte('\n')
}

/*
package main

import (
  "bufio"
  "fmt"
  "io"
  "os"
)

var (
  Input  io.Reader = os.Stdin
  Output io.Writer = os.Stdout
)

func main() {
  in := bufio.NewReader(Input)
  out := bufio.NewWriter(Output)
  defer out.Flush()

  var n, m int
  fmt.Fscan(in, &n, &m)

  grid := make([]string, n)
  for i := 0; i < n; i++ {
    fmt.Fscan(in, &grid[i])
  }

  fixedRow := make([]int, n)
  fixedCol := make([]int, m)
  qRow := make([]int, n)
  qCol := make([]int, m)

  for i := 0; i < n; i++ {
    for j := 0; j < m; j++ {
      switch grid[i][j] {
      case '+':
        fixedRow[i]++
        fixedCol[j]++
      case '-':
        fixedRow[i]--
        fixedCol[j]--
      case '?':
        qRow[i]++
        qCol[j]++
      }
    }
  }

  ans := -1000000000
  for i := 0; i < n; i++ {
    for j := 0; j < m; j++ {
      maxR := fixedRow[i] + qRow[i]
      minC := fixedCol[j] - qCol[j]

      var diff int
      if grid[i][j] == '?' {
        // Мы не можем одновременно получить +2 для строки и +2 для разности столбца
        // Фактически мы теряем 2 по сравнению с "идеальной" разностью
        diff = maxR - minC - 2
      } else {
        diff = maxR - minC
      }

      if diff > ans {
        ans = diff
      }
    }
  }

  fmt.Fprintln(out, ans)
}
*/
