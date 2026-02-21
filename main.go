package main

import (
	"fmt"
	"os"
)

// "fmt"
// . "leetcode/Structs"
// . "leetcode/medium"

func main() {
	fmt.Println("GOTESTS_TEMPLATE_DIR =", os.Getenv("GOTESTS_TEMPLATE_DIR"))

	// Посмотреть все переменные окружения
	for _, e := range os.Environ() {
		// если хотите увидеть все переменные, раскомментируйте:
		fmt.Println(e)
	}
}
