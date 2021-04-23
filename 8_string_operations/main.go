package main

import (
	"fmt"
	"strings"
)

func main() {
	// \ has special meaning, \\ meaning single \
	path := "\"C:\\Users\\xuand\\go\\src\\Github\\Go_Learning\""
	fmt.Println("Current path is: ", path)

	var s string = "I'm OK"
	fmt.Println(s)

	// multi lines string
	var s2 string = `
	first line
	second line
	thrid line
	`
	fmt.Println(s2)

	var path2 string = `C:\Users\xuand\go\src\Github\Go_Learning\8_string_operations`
	fmt.Println(path2)

	// get length
	fmt.Println(len(path2))

	// string concatenation
	var name string = "this"
	var world string = "dream"
	fmt.Println(name + world)
	fmt.Printf("%s%s\n", name, world)
	var sum string = name + world
	fmt.Println(sum)

	// string spilt
	ret := strings.Split(path, "\\")
	fmt.Println(ret)

	// string prefix Suffix
	fmt.Println(strings.HasPrefix(path, "\"C:\\"))
	fmt.Println(strings.HasSuffix(path, "Go_Learning\""))

	// substring
	var s3 string = "abcd"
	fmt.Println(strings.Index(s3, "bc"))
	fmt.Println(strings.LastIndex(s3, "d"))

	// string join
	fmt.Println(strings.Join(ret, "\\"))
}
