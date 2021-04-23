package main

import "fmt"

func main() {
	// float
	f1 := 1.23456
	fmt.Printf("%T\n", f1) // default as float64

	f2 := float32(1.23456)
	fmt.Printf("%T\n", f2)

	// f1 = f2 ? can not do that because they are different type

	// bool
	// in go, bool can not cast to int, init to false and can
	// not be used in any calculation

	b1 := true
	var b2 bool
	fmt.Printf("%T value: %v\n", b1, b1)
	fmt.Printf("%T value: %v\n", b2, b2)
}
