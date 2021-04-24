package main

import "fmt"

func main() {
	var (
		a = 5
		b = 5
	)
	fmt.Println(a + b)
	fmt.Println(a - b)
	fmt.Println(a * b)
	fmt.Println(a / b)
	fmt.Println(a % b)
	a++
	b--
	fmt.Println(a)

	// relational operators
	fmt.Println(a == b) // compare type and value!
	fmt.Println(a != b)
	fmt.Println(a > b)
	fmt.Println(a >= b)
	fmt.Println(a < b)
	fmt.Println(a <= b)

	// logical operators
	age := 22
	// and
	if age > 18 && age < 60 {
		fmt.Println("working labour")
	} else {
		fmt.Println("students")
	}
	// or
	if age > 18 || age < 60 {
		fmt.Println("working labour")
	} else {
		fmt.Println("students")
	}
	// priority similar to other programming language

	// not -> !
	isMarried := false
	fmt.Println(isMarried)  // false
	fmt.Println(!isMarried) // true

	// bit operators
	// binary rep of 5 is 101 and
	//               2 is  10
	fmt.Println(5 & 2)   // bit and 010
	fmt.Println(5 | 2)   // bit or  111
	fmt.Println(5 ^ 2)   // bit xor 111
	fmt.Println(5 << 1)  // bit left shift 101 -> 10
	fmt.Println(1 << 10) // bit left shift 10000000000 = 1024
	fmt.Println(5 >> 2)  // bit right shift 101 -> 1

	// assignment operator
	var x int = 10
	x += 1
	x -= 1
	x *= 1
	x /= 1
	x %= 1
	x <<= 1
	x >>= 1
	x &= 1
	x |= 1
	x ^= 1
	fmt.Println(x)
}
