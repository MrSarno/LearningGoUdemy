package main

import "fmt"

func main() {
	a := 13
	b := 44
	c := 8.354792
	d := 454545.45454545
	e := 987654321012345678
	f := -209
	var g int8 = -4
	var h uint8 = 2
	var i byte = 6 // byte is just an alias
	fmt.Printf("%T\n", a)
	fmt.Printf("%T\n", b)
	fmt.Printf("%T\n", c)
	fmt.Printf("%T\n", d)
	fmt.Printf("%T\n", e)
	fmt.Printf("%T\n", f)
	fmt.Printf("%T\n", g)
	fmt.Printf("%T\n", h)
	fmt.Printf("%T\n", i)
}
