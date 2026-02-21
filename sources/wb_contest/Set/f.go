package set

import (
	"fmt"
)

func ContestF() {

	/* 	var firstNumber int
	   	fmt.Scan(&firstNumber)

	   	strings1 := make([]string, firstNumber)
	   	for i := 0; i < firstNumber; i++ {
	   		fmt.Scan(&strings1[i])
	   	}

	   	var secondNumber int
	   	fmt.Scan(&secondNumber)

	   	strings2 := make([]string, secondNumber)
	   	for i := 0; i < secondNumber; i++ {
	   		fmt.Scan(&strings2[i])
	   	} */

	// strings1 := []string{"ABC", "A37", "BCDA"}
	// strings2 := []string{"A317BD", "B137AC"}

	strings1 := []string{"1ABC", "3A4B"}
	strings2 := []string{"A143BC", "C143AB", "AAABC1"}

	str1Map := make(map[int](map[string]int))
	for i := 0; i < len(strings1); i++ {
		str1Map[i] = make(map[string]int)
		for _, j := range strings1[i] {
			str1Map[i][string(j)]++
		}

	}

	max := 0
	final := make(map[string]int)
	for _, value := range strings2 {
		count := 0

		str2Map := make(map[string]int)
		for _, j := range value {
			str2Map[string(j)]++
		}

		for key, _ := range str2Map {
			for i := 0; i < len(strings1); i++ {

				_, ok := str1Map[i][key]
				if ok {
					count++
				}

			}
		}
		final[value] = count
		if count >= max {
			max = count

			count = 0
		}
	}
	for key, value := range final {
		if value == max {
			fmt.Println(key)
		}
	}
}
