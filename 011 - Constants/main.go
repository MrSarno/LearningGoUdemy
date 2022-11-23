package main

import "fmt"

// typed constants
const a int = 42
const b float64 = 78.87
const c string = "this is text"

// untyped constants
const (
	d = "hello, world"
	e = 15
	f = true
)

func main() {
	fmt.Println("a")
	fmt.Println("b")
	fmt.Println("c")
	fmt.Println("d")
	fmt.Println("e")
	fmt.Println("f")
}
