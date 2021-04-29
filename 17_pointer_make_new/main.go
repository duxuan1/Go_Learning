package main

import "fmt"

// func modify1(x int) {
// 	x = 100
// }

func modify2(x *int) {
	*x = 100
}

func main() {
	n := 18
	p := &n // & -> get address
	fmt.Println(p)
	fmt.Printf("%T\n", p)

	m := *p // * -> get value from address
	fmt.Println(m)

	// pass variable to function by address
	x := 10
	// modify1(x) no effect here
	fmt.Println(x)
	modify2(&x)
	fmt.Println(x)

	// assignment to pointer
	var a1 *int // memory not assigned, a == nil
	// can not assign any thing to a1
	fmt.Println(a1)

	// use new
	var a2 = new(int)
	*a2 = 100
	fmt.Println(*a2)

	// make can also allocate memory, but only for
	// map, slice and channel
	// new used for int, string, float
}
