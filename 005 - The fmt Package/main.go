package main

import "fmt"

var x = 86

func main() {
	fmt.Println("The variable x is set to:", x)
	fmt.Printf("The variable x is of variable type %T", x)
	fmt.Printf("\nThe variable x is equivalent to %b in base 2", x)
	fmt.Printf("\nThe variable x is equivalent to %o in base 8", x)
	fmt.Printf("\nThe variable x is equivalent to %x in base 16 (lower case letters)", x)
	fmt.Printf("\nThe variable x is equivalent to %X in base 16 (upper case letters)", x)
	fmt.Println("\n---")
	s := fmt.Sprintf("%#x\t%b\t%x/n", x, x, x) // Sprint = String PRINT (or 'print to string')
	fmt.Println(s)
}
