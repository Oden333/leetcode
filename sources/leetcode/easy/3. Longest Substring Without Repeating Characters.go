package easy

func lengthOfLongestSubstring(s string) int {
	tmp := []rune(s)

	maxM := make(map[rune]struct{})
	max := 0
	// for i := 0; i < len(tmp); i++ {
	for i := range tmp {
		for _, r := range tmp[i:] {
			if _, ok := maxM[r]; ok {
				maxM = make(map[rune]struct{})
				break
			}
			maxM[r] = struct{}{}
			if len(maxM) > max {
				max = len(maxM)
			}
		}
	}

	return max
}
