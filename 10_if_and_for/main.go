package main

import "fmt"

func main() {
	// age := 19
	// if age >= 18 {
	// 	fmt.Println("valid")
	// } else if age >= 35 {
	// 	fmt.Println("invalid")
	// } else {
	// 	fmt.Println("other")
	// }

	// special case, age only created within if statements
	if age := 19; age >= 18 {
		fmt.Println("valid")
	} else if age >= 35 {
		fmt.Println("invalid")
	} else {
		fmt.Println("other")
	}

	// for loop
	for i := 0; i < 10; i++ {
		fmt.Println(i)
	}

	// for loop without creating variable in loop
	var i int = 5
	for ; i < 10; i++ {
		fmt.Println(i)
	}

	// i++ can be removed and becoming while loop
	var j int = 5
	for j < 10 {
		fmt.Println(j)
		j++
	}

	// infinite loop
	// for {
	// 	fmt.Println("lol")
	// }

	// for range loop
	s := "Hello"
	for i, v := range s {
		fmt.Printf("%d %c\n", i, v)
	}

	// print Multiplication table
	for i := 1; i < 10; i++ {
		for j := 1; j < 10; j++ {
			fmt.Println(i, " * ", j, " = ", i*j)
		}
	}
}
