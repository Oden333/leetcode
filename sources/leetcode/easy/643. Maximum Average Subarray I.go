package easy

func findMaxAverage(n []int, k int) float64 {
	var max float64 = -(1 << 63) - 1
	if len(n) <= k {
		return avg(n)
	}

	for l := 0; l <= len(n)-k; l++ {
		if t := avg(n[l : l+k]); t > max {
			max = t
		}
	}

	return max
}

func avg(n []int) float64 {
	sum := 0
	for _, v := range n {
		sum += v
	}
	return float64(sum) / float64(len(n))
}
