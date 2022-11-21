package main

import "fmt"

func main() {
	s := "Hello, World!"
	fmt.Println(s)
	fmt.Printf("The variable 's' is of type: %T\n", s)
	fmt.Println("A string value is a - possibly empty - sequence of bytes. Strings are immutable.\n")

	// decimal
	bs := []byte(s)
	fmt.Println(bs)
	fmt.Printf("The variable 'bs' is of type: %T\n\n", bs)

	// utf-8
	for i := 0; i < len(s); i++ {
		fmt.Printf("%#U ", s[i])
	}

	// hexadecimal
	fmt.Println("")
	for i, v := range s {
		fmt.Printf("At index position %d we have hex %#x\n", i, v)
	}
}
