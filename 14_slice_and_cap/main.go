package main

import "fmt"

func main() {
	// define slice, slice is similar to array
	var s1 []int // do not put length in []
	var s2 []string
	fmt.Println(s1)
	fmt.Println(s2)
	fmt.Println(s1 == nil) // nil meaning empty
	fmt.Println(s2 == nil)
	fmt.Println(len(s1))
	fmt.Println(len(s2))

	// initialize slice
	s1 = []int{1, 2, 3}
	s2 = []string{"1", "2", "3"}
	fmt.Println(s1)
	fmt.Println(s2)
	fmt.Println(len(s1))
	fmt.Println(len(s2))
	fmt.Println(cap(s1))
	fmt.Println(cap(s2))

	// slicing an array (similar to python), sliced array is referenced to origin
	a1 := [...]int{1, 3, 5, 7, 9, 11, 13}
	a2 := a1[1:4]
	a3 := a1[3:4]
	fmt.Println(a2)
	fmt.Printf("%T\n", a2)
	// cap indicates length of the bottom layer array
	fmt.Println(cap(a1)) // explanation: | 1 3 5 7 9 11 13 | -> 7
	fmt.Println(cap(a2)) // explanation: 1 | 3 5 7 9 11 13 | -> 6
	fmt.Println(cap(a3)) // explanation: 1 3 5 | 7 9 11 13 | -> 4
	fmt.Println(a1)

	// slicing in sliced slice
	a4 := a2[1:] // explanation: 3 | 5 7 9 11 13 | -> 5
	a4[0] = 100
	fmt.Println(a1)
	fmt.Println(cap(a4))
}
