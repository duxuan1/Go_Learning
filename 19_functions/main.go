package main

import "fmt"

func f1() {
	fmt.Println("function f1")
}

func f2(name string) {
	fmt.Println("Hello", name)
}

func f3(x int, y int) int {
	sum := x + y
	return sum
	// or we can return x + y
}

// y is a slice
func f4(x int, y ...int) int {
	fmt.Println(x, y)
	return 0
}

// name return variable
func f5(x, y int) (sum int) {
	sum = x + y
	return // no need to return sum since sum is decleared
}

// multiple return variables
func f6(x, y int) (sum, sub int) {
	sum = x + y
	sub = x - y
	return
}

// defer function
func deferDemo() {
	fmt.Println("start")
	defer fmt.Println("1") // defer in reverse order
	defer fmt.Println("2")
	defer fmt.Println("3")
	defer fmt.Println("4")
	fmt.Println("end")
}

// for go language, return is not atomic operation
// step 1: assign value to return variable
// step 2: return true RET
// defer happen in between them
func defer1() int {
	x := 5
	defer func() {
		x++ // it modified x but not return value
	}()
	return x
}

func defer2() (x int) {
	defer func() {
		x++
	}()
	return x // x already assigned to 5, defer -> x++, return 6
}

func defer3() (y int) {
	x := 5
	defer func() {
		x++ // it modified x but not y
	}()
	return x
}

func defer4() (x int) {
	defer func(x int) {
		x++ // changed x in this func but not the x in defer4
	}(x)
	return 5
}

func main() {
	// call function
	f1()
	f2("dudu")
	s := f3(1, 2)
	fmt.Println(s)
	// or we can do this
	fmt.Println(f3(1, 2))

	f4(1, 1, 2, 3, 4, 5, 6, 7)
	fmt.Println(f5(1, 3))
	fmt.Println(f6(1, 3))
	// can not create func with name in func

	deferDemo()
	fmt.Println(defer1())
	fmt.Println(defer2())
	fmt.Println(defer3())
	fmt.Println(defer4())
}
