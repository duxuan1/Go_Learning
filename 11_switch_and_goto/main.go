package main

import "fmt"

// switch can simpilify many if statements
func switch_demo() {
	var n int = 5
	switch n {
	case 1:
		fmt.Println("1")
	case 2:
		fmt.Println("2")
	case 3:
		fmt.Println("3")
	case 4:
		fmt.Println("4")
	case 5:
		fmt.Println("5")
	default:
		fmt.Println("invalid input")
	}
}

func main() {
	// break
	for i := 0; i < 10; i++ {
		if i == 5 {
			break
		}
		fmt.Println(i)
	}
	fmt.Println("over")

	// continue
	for i := 0; i < 10; i++ {
		if i == 5 {
			continue
		}
		fmt.Println(i)
	}
	fmt.Println("over")

	// switch
	switch_demo()

	// goto and label
	for i := 0; i < 10; i++ {
		if i == 5 {
			goto end
		}
		fmt.Println(i)
	}
	fmt.Println("over")

end: // this is a label
	fmt.Println("end")
}
