package medium

func searchMatrix(matrix [][]int, target int) bool {
	var (
		// 1 <= n, m <= 300
		n, m = len(matrix) /* rows */, len(matrix[0]) - 1 /* columns */
		val  int
	)

	for i, j := 0, m; i < n && j >= 0; {
		val = matrix[i][j]

		switch {
		case val > target:
			j--
		case val < target:
			i++
		case val == target:
			return true
		}
	}

	return false

}
