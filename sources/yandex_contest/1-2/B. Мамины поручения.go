package main

import (
	"bufio"
	"fmt"
	"math"
	"strconv"
	"strings"
)

func B() {
	reader := bufio.NewReader(Input)
	nums, _ := reader.ReadString('\n')
	ns := strings.Split(nums, " ")

	inta, _ := strconv.Atoi(ns[0])
	a := float64(inta)
	intb, _ := strconv.Atoi(ns[1])
	b := float64(intb)
	intc, _ := strconv.Atoi(ns[2])
	c := float64(intc)

	intv0, _ := strconv.Atoi(ns[3])
	v0 := float64(intv0)
	intv1, _ := strconv.Atoi(ns[4])
	v1 := float64(intv1)
	intv2, _ := strconv.Atoi(ns[5])
	v2 := float64(intv2)

	min4 := a/v0 + c/v1 + b/v2
	if x := (a/v0 + c/v0 + c/v1 + a/v2); min4 > x {
		min4 = x
	}
	if x := (b/v0 + c/v1 + a/v2); min4 > x {
		min4 = x
	}
	if x := (b/v0 + c/v0 + c/v1 + b/v2); min4 > x {
		min4 = x
	}

	minM := math.Min((b/v0 + b/v1), (a/v0 + c/v0 + c/v1 + a/v1))
	minP := math.Min((a/v0 + a/v1), (b/v0 + c/v0 + c/v1 + b/v1))

	res := math.Min(min4, minM+minP)

	writer := bufio.NewWriter(Output)
	defer writer.Flush()

	writer.WriteString(fmt.Sprintf("%.10f", res))
	writer.WriteByte('\n')
}
