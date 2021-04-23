package main

import "fmt"

// variables must be declared then being used
// var name type
var name string
var age int
var isOK bool

// group declaration
var (
	s string
	a int
	b bool
)

// variable initialization
func main() {
	name = "xuan"
	age = 22
	isOK = true

	s = "1"
	a = 1
	b = true
	// declared variable must be used in go
	fmt.Printf("name: %s", name)
	fmt.Println(age)
	fmt.Println(isOK)

	fmt.Println(s)
	fmt.Println(a)
	fmt.Println(b)
}
