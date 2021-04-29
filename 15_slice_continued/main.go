package main

import "fmt"

func main() {
	// make(type, length)
	s1 := make([]int, 5)
	fmt.Printf("s1 = %v, len(s1) = %v, cap(s1) = %v\n",
		s1, len(s1), cap(s1))

	s2 := make([]int, 0, 10)
	fmt.Printf("s2 = %v, len(s2) = %v, cap(s2) = %v\n",
		s2, len(s2), cap(s2))

	// slice is reference type, can not be compared with
	// other slice
	// can only compare with nil
	// if backend array is nil, then slice is nil
	// if len() == 0 then slice is empty

	s3 := []int{1, 2, 3}
	s4 := s3
	fmt.Println(s3, s4)
	s3[0] = 1000
	fmt.Println(s3, s4)

	// looping a slice
	for i := 0; i < len(s4); i++ {
		fmt.Println(s4[i])
	}
	// for range
	for i, v := range s4 {
		fmt.Println(i, v)
	}
}
