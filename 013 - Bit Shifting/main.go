package main

import "fmt"

const (
	_ = iota

	kb = 1024
	mb = kb * 1024
	gb = 1024 * mb

	kb2 = 1 << (iota * 10)
	mb2 = 1 << (iota * 10)
	gb2 = 1 << (iota * 10)
)

func main() {
	x := 4
	fmt.Printf("%d\t\t%b\n", x, x)

	y := x << 1
	fmt.Printf("%d\t\t%b\n", y, y)

	fmt.Printf("\n%d\t\t%b\n", kb, kb)
	fmt.Printf("%d\t\t%b\n", mb, mb)
	fmt.Printf("%d\t%b\n\n", gb, gb)

	fmt.Println(kb2)
	fmt.Println(mb2)
	fmt.Println(gb2)
}
