package main

import (
	"fmt"
	"math"
	"os"
)

func main() {
	var n, w, h int
	f, _ := os.Open("input.txt")
	defer f.Close()
	fmt.Fscan(f, &n, &w, &h)

	side := float64(max(w, h))
	diagonal := math.Sqrt(float64(w*w + h*h))

	var maxSide float64
	for i := 0; i < n; i++ {
		var diploma int
		fmt.Scan(&diploma)
		maxSide = math.Max(maxSide, float64(diploma))
	}

	minBoardSide := math.Max(side, math.Ceil(diagonal))
	fmt.Printf("Минимальный размер стороны доски: %.2f\n", math.Max(minBoardSide, maxSide))
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
