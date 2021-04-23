package main

import "fmt"

// variable declaration and assignment
var s1 string = "xuan"
var s2 = "no type" // declaration can have no type
var a int = 1
var b bool = true

func foo() (int, string) {
	return 10, "lol"
}

func main() {
	fmt.Println(s1)
	fmt.Println(s2)
	fmt.Println(a)
	fmt.Println(b)

	s3 := "haha" // short declaration, can only use in function
	fmt.Println(s3)

	// Anonymous variable is special
	x, _ := foo()
	_, y := foo()
	fmt.Println("x: ", x)
	fmt.Println("y: ", y)
}
