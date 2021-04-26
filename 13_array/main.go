package main

import "fmt"

func main() {
	// create an array with lengrh 3
	var a1 [3]bool // bool array default to be false, int -> 0, string -> ""
	fmt.Println(a1)

	var a2 [4]bool // a2 has different type with a1, because they have different length
	fmt.Println(a2)

	// array initialization
	// method 1
	a1 = [3]bool{true, true, true}
	fmt.Println(a1)
	// method 2, ... will count length for you
	a3 := [...]int{1, 2, 3, 4, 5, 6, 7, 8, 9}
	fmt.Println(a3)
	// methoud 3, assign value to specific index
	a4 := [5]int{0: 1, 4: 2}
	fmt.Println(a4)

	// array looping
	cities := [...]string{"Beijing", "Shanghai", "Guangzhou"}
	n := len(cities)
	for i := 0; i < n; i++ {
		fmt.Println(cities[i])
	}
	// for range, similar to python for i in cities
	for i, v := range cities { // i is index, v is value
		fmt.Println(i, v)
	}

	// multi dimension array
	var a5 [3][2]int = [3][2]int{
		{1, 2},
		{3, 4},
		{5, 6},
	}
	for i := 0; i < 3; i++ {
		for j := 0; j < 2; j++ {
			fmt.Println(a5[i][j])
		}
	}
	for _, v := range a5 {
		fmt.Println(v)
	}

	// array is value assigned but not address assigned, doesn't like python
	b1 := [3]int{1, 2, 3}
	b2 := b1
	b2[0] = 100
	fmt.Println(b1, b2)

	// find sum
	array := [5]int{1, 3, 5, 7, 8}
	sum := 0
	for i := 0; i < len(array); i++ {
		sum += array[i]
	}
	fmt.Println("sum is", sum)
}
