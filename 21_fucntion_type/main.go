package main

import "fmt"

func f1() string {
	return "hello"
}

func f2() int {
	return 1
}

// function as argument
func f3(x func() int) {
	ret := x()
	fmt.Println(ret)
}

func f4(x int, y int) int {
	return x + y
}

// function as return value
func f5(x func() int) func(int, int) int {
	return f4
}

func main() {
	a := f1()
	fmt.Printf("%T\n", a)
	b := f2()
	fmt.Printf("%T\n", b)
	f3(f2)
	fmt.Printf("%T\n", f4)
	f5(f2)
}
