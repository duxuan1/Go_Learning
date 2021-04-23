package main

import "fmt"

func main() {
	// integer type
	// decimal
	var i1 = 101
	fmt.Printf("%d\n", i1) // print as decimal
	fmt.Printf("%o\n", i1) // print as octal
	fmt.Printf("%x\n", i1) // print as hexadecimal
	fmt.Printf("%b\n", i1) // print as binary
	fmt.Printf("%T\n", i1) // check variable type

	// octal number, start with 0 
	i2 := 077
	fmt.Printf("%d\n", i2)

	// Hexadecimal number, start with 0x, range 0 ~ f
	i3 := 0x1234567
	fmt.Printf("%d\n", i3)

	// int8 declaration
	i4 := int8(9)
	fmt.Printf("%T\n", i4)


	// float type
	
}
