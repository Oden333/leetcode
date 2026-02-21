package set

import (
	"bufio"
	"fmt"
	"os"
)

func Gay() {

	f, _ := os.Open("input.txt")
	defer f.Close()

	s := bufio.NewScanner(f)
	s.Split(bufio.ScanWords)

	s.Scan()
	gen1 := s.Text()
	s.Scan()
	gen2 := s.Text()

	// gen1 := "ABBACAB"
	// gen2 := "BCABB"

	gen1Map := make(map[string]int)

	curGen := make([]byte, 2)
	for i := 0; i < len(gen1)-1; i++ {
		curGen[0] = gen1[i]
		curGen[1] = gen1[i+1]
		gen1Map[string(curGen)]++
	}

	counter := 0
	for i := 0; i < len(gen2)-1; i++ {
		curGen[0] = gen2[i]
		curGen[1] = gen2[i+1]
		_, ok := gen1Map[string(curGen)]
		if ok {
			counter += gen1Map[string(curGen)]
		}
	}

	fmt.Println(counter)

}
