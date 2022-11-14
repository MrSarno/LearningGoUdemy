package main

import "fmt"

var x = 28                       // I could specify this is an int, but that isn't necessary
var y string = "this is my text" // an example of specifying a type
var z = `he said "this is a

raw string     literal"` // a raw string literal (as it says itself)

// remember Go is statically typed
// variables cannot be changed from one type to another

func main() {
	fmt.Println(x)
	fmt.Printf("%T\n", x)
	fmt.Println("---")
	fmt.Println(y)
	fmt.Printf("%T\n", y)
	fmt.Println("---")
	x = 15
	fmt.Println(x)
	fmt.Println("---")
	fmt.Println(z)
}
