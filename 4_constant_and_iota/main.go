package main

import "fmt"

// value of constant type won't change
const pi float32 = 3.1415

// group creation
const (
	STATUS_OK = 200
	STATUS_NOT_FOUND = 404
)

// special operation
const (
	n1 = 100
	n2
	n3 // n2 and n3 have the same value with n1
)

// iota is a timer, when const appear, iota is assigned to 0
// when more const creation going on, iota keep increasing
const (
	a1 = iota // 0
	a2 = iota // 1
	a3        // 2
	a4        // 4
)


// interview question

// const (
// 	b1 = iota // 0
// 	b2        // 1
// 	_         // 2, anonymous const still count for iota 
// 	b3        // 3
// )


// const (
// 	c1 = iota // 0
// 	c2 = 100  // 100
// 	c3        // 2
// 	c4        // 3
// )

// const (
// 	d1, d2 = iota + 1, iota + 1 // d1 = 1, d2 = 2
//  d3, d4 = iota + 1, iota + 2 // d3 = 2, d3 = 3
// )

// // define the magnitude
// const (
// 	_ = iota
// 	kb = 1 << (10 + iota) // iota = 1
// 	bb = 1 << (10 + iota) // iota = 2
// 	gb = 1 << (10 + iota)
// 	tb = 1 << (10 + iota)
// 	pb = 1 << (10 + iota)
// )

func main() {
	fmt.Println(pi)
	fmt.Println("n1: ", n1)
	fmt.Println("n2: ", n2)
	fmt.Println("n3: ", n3)
	fmt.Println("a1: ", a1)
	fmt.Println("a2: ", a2)
	fmt.Println("a3: ", a3)
	fmt.Println("a4: ", a4)
}
