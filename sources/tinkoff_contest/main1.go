package main

import (
	"fmt"
	"math"
)

func Score(x, y float64) int {
	distance := math.Sqrt(x*x + y*y)

	switch {
	case distance <= 0.1:
		return 3
	case distance > 0.1 && distance <= 0.8:
		return 2
	case distance > 0.8 && distance <= 1.0:
		return 1
	default:
		return 0
	}
}

func main1() {

	coordinates := make([][]float64, 3)

	for i := 0; i < 3; i++ {
		var x, y float64
		fmt.Scan(&x, &y)
		coordinates[i] = []float64{x, y}
	}
	totalScore := 0

	for _, coord := range coordinates {
		score := Score(coord[0], coord[1])
		totalScore += score
	}

	fmt.Println(totalScore)
}
