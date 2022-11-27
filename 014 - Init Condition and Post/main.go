package main

import (
	"fmt"
	"strconv"
)

func main() {
	// for init; condition; post { ... }
	for i := 0; i <= 100; i++ {
		fmt.Println("Hello, loop! Iteration: " + strconv.Itoa(i))
	}
}
