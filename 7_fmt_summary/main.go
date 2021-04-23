package main

import "fmt"

func main() {
	var n = 100
	fmt.Printf("%T\n", n) // check type
	fmt.Printf("%v\n", n) // print any type
	fmt.Printf("%b\n", n) // print binary
	fmt.Printf("%d\n", n) // print decimal
	fmt.Printf("%o\n", n) // print octal
	fmt.Printf("%x\n", n) // print hexadecimal

	var s = "Hello"
	fmt.Printf("%v\n", s)  // print string
	fmt.Printf("%s\n", s)  // print any type
	fmt.Printf("%#v\n", s) // print exact any type
}
