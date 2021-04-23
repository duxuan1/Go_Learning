package main

import (
	"fmt"
)

func main() {
	// string is inmutable
	// but we can modify string by some other ways
	s := "lol"
	s1 := []rune(s)
	s1[0] = 'o'
	fmt.Println(string(s1))

	// single quotation and double q are different
	c1 := "a" // string
	c2 := 'a' // int32
	fmt.Printf("c1: %T c2: %T\n", c1, c2)
	c3 := "H"       // string
	c4 := byte('H') // uint8
	fmt.Printf("c3: %T c4: %T\n", c3, c4)

	// type cast
	n := 10
	var f float64 = float64(n)
	fmt.Println(f)
}
