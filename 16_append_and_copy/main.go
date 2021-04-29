package main

import "fmt"

func delete(slice []int, index int) []int {
	return append(slice[:index], slice[index+1:]...)
}

func main() {
	s1 := []int{1, 2, 3}
	// s1[3] = 4, no enough space, error
	fmt.Printf("s1 = %v, len(s1) = %v, cap(s1) = %v\n",
		s1, len(s1), cap(s1))
	// append(target, value)
	s1 = append(s1, 4) // backend array will expand
	fmt.Printf("s1 = %v, len(s1) = %v, cap(s1) = %v\n",
		s1, len(s1), cap(s1)) // then cap doubled here

	// when slice cap < 1024, then double expand
	// when slice cap >= 1024, then slice expand bt 1/4

	numSlice := []int{}
	for i := 0; i < 10; i++ {
		numSlice = append(numSlice, i)
		fmt.Printf("numSlice = %v, len = %v, cap = %v\n",
			numSlice, len(numSlice), cap(numSlice))
	}

	// can append multiple elements
	s1 = append(s1, 5, 6, 7)
	fmt.Printf("s1 = %v, len(s1) = %v, cap(s1) = %v\n",
		s1, len(s1), cap(s1))

	s2 := []int{100}
	s2 = append(s2, s1...) // ...meaning split
	fmt.Printf("s2 = %v, len(s2) = %v, cap(s2) = %v\n",
		s2, len(s2), cap(s2))

	// copy
	a1 := []int{1, 2, 3}
	a2 := a1
	a3 := make([]int, 3)
	// copy(dest, sour), copy is value type
	copy(a3, a1)
	a1[0] = 100
	fmt.Println(a1, a2, a3)

	// delete by index 1
	a1 = append(a1[:1], a1[2:]...)
	fmt.Println(a1, a2, a3)

	// delete by function
	a2 = delete(a2, 1)
	fmt.Println(a1, a2, a3)
}
