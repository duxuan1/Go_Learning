package main

import "fmt"

var x = 100 // define a global variable

func f1() {
	fmt.Println(x) // x not found in f1, search global
}

func main() {
	f1()
}
