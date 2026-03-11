package medium

func search(in []int, target int) int {
	if len(in) == 0 {
		return -1
	}

	var (
		left, right = 0, len(in) - 1
		mid         int
	)

	for left <= right {
		mid = (right-left)/2 + left

		if in[mid] == target {
			return mid
		}

		// проверяем, какая часть отсортирована (переход может быть только в одной половине относительно mid)
		if in[left] <= in[mid] {
			// левая [left...mid] отсортирована

			// если таргет в отсортированной половине [ in[left], in[mid] ]
			if target >= in[left] && target < in[mid] {
				// сужаем поиск
				right = mid - 1
			} else {
				left = mid + 1
			}

		} else {
			// правая отсортирована [mid...right] отсортирована

			// если таргет в отсортированной половине
			if target > in[mid] && target <= in[right] {
				// сужаем поиск
				left = mid + 1
			} else { // иначе переходим на другую половину
				right = mid - 1
			}

		}
	}

	return -1
}
