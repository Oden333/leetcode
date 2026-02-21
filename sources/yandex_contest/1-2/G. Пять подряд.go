package main

import (
	"bufio"
	"strconv"
	"strings"
)

/* var (
  Input  io.Reader = os.Stdin
  Output io.Writer = os.Stdout
)
	func main() {
		G()
	}

*/

func G() {
	reader := bufio.NewReaderSize(Input, 1<<20)
	in, _ := reader.ReadString('\n')
	in = strings.Trim(in, "\n")

	nums := strings.Split(in, " ")
	rows, _ := strconv.Atoi(nums[0])
	cols, _ := strconv.Atoi(nums[1])

	writer := bufio.NewWriterSize(Output, 1<<20)
	defer writer.Flush()

	if rows < 5 && cols < 5 {
		writer.WriteString("No")
		return
	}

	matrix := make([][]string, rows)
	idx := 0
	numX := 0
	numO := 0

	for i := range rows {
		matrix[i] = make([]string, cols)
	}

	for i := 0; i < rows; {
		b, err := reader.ReadByte()
		if b == '\n' || err != nil {
			i++
			idx = 0

			continue
		}

		switch b {
		case 'X':
			matrix[i][idx] = "X"
			idx++

			numX++
			if numX == 5 {
				writer.WriteString("Yes")
				return
			}
			numO = 0
		case 'O':
			matrix[i][idx] = "O"
			idx++

			numO++
			if numO == 5 {
				writer.WriteString("Yes")
				return
			}
			numX = 0
		case '.':
			matrix[i][idx] = "."
			idx++

			numX = 0
			numO = 0
		}
	}

	for i := range rows {
		for j := range cols {
			if matrix[i][j] == "." {
				continue
			}

			if i-4 >= 0 {
				if matrix[i][j] == matrix[i-1][j] &&
					matrix[i][j] == matrix[i-2][j] &&
					matrix[i][j] == matrix[i-3][j] &&
					matrix[i][j] == matrix[i-4][j] {
					writer.WriteString("Yes")
					return
				}
			}

			if i+4 < rows {
				if matrix[i][j] == matrix[i+1][j] &&
					matrix[i][j] == matrix[i+2][j] &&
					matrix[i][j] == matrix[i+3][j] &&
					matrix[i][j] == matrix[i+4][j] {
					writer.WriteString("Yes")
					return
				}
			}

			if i+4 < rows && j-4 >= 0 {
				if matrix[i][j] == matrix[i+1][j-1] &&
					matrix[i][j] == matrix[i+2][j-2] &&
					matrix[i][j] == matrix[i+3][j-3] &&
					matrix[i][j] == matrix[i+4][j-4] {
					writer.WriteString("Yes")
					return
				}
			}

			if i+4 < rows && j+4 < cols {
				if matrix[i][j] == matrix[i+1][j+1] &&
					matrix[i][j] == matrix[i+2][j+2] &&
					matrix[i][j] == matrix[i+3][j+3] &&
					matrix[i][j] == matrix[i+4][j+4] {
					writer.WriteString("Yes")
					return
				}
			}

			if i-4 >= 0 && j+4 < cols {
				if matrix[i][j] == matrix[i-1][j+1] &&
					matrix[i][j] == matrix[i-2][j+2] &&
					matrix[i][j] == matrix[i-3][j+3] &&
					matrix[i][j] == matrix[i-4][j+4] {
					writer.WriteString("Yes")
					return
				}
			}

			if i-4 >= 0 && j-4 >= 0 {
				if matrix[i][j] == matrix[i-1][j-1] &&
					matrix[i][j] == matrix[i-2][j-2] &&
					matrix[i][j] == matrix[i-3][j-3] &&
					matrix[i][j] == matrix[i-4][j-4] {
					writer.WriteString("Yes")
					return
				}
			}

		}
	}

	writer.WriteString("No")
	writer.WriteByte('\n')
}
