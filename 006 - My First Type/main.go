package main

import "fmt"

var a int

type johnsnumber int

var b johnsnumber

func main() {
	a = 1
	b = 2
	fmt.Println(a)
	fmt.Printf("%T\n\n", a)
	fmt.Println(b)
	fmt.Printf("%T\n", b)
}
