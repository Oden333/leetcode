package main

import (
	"fmt"
	"os"
)

func main() {
	fmt.Println("GOTESTS_TEMPLATE_DIR =", os.Getenv("GOTESTS_TEMPLATE_DIR"))

	for _, e := range os.Environ() {
		fmt.Println(e)
	}
}
