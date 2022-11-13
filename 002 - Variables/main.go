package main

import "fmt"

func main() {
	// short declaration operator
	x := 37

	// y = 41
	// won't work due to scope - y would be undefined outside of the main function

	fmt.Println("The variable x is set to", x, ".")

	foo()
}

func foo() {
	y := 53

	fmt.Println("The variable y is set to", y, ".")
}
