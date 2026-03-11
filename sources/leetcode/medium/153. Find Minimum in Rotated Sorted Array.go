package medium

func findMin(in []int) int {
	// 1 <= n <= 5000
	if len(in) == 1 {
		return in[0]
	}

	var (
		l, r = 0, len(in) - 1
		mid  int
		// -5000 <= nums[i] <= 5000
	)

	for l <= r {
		mid = (r-l)/2 + l

		// если мы дошли до начала разрыва, то возвращаем первый индекс
		if in[l] <= in[r] {
			return in[l]
		}

		if in[l] <= in[mid] {
			// [l...mid] - отсортирована, а значит разрыв справа, в [mid+1...r]

			// смещаемся направо чтобы найти разрыв
			l = mid + 1
		} else {
			// [mid...r] - отсортирована, а значит разрыв слева, в [l...mid]

			// смещаемся влево чтобы найти разрыв
			r = mid
		}
	}

	return in[l]
}
