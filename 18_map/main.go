package main

import (
	"fmt"
	"math/rand"
	"sort"
)

func main() {
	// map[key]value
	var m1 = make(map[string]int, 10)
	m1["a"] = 18
	m1["b"] = 19
	fmt.Println(m1)

	// get access to key
	value, found := m1["c"]

	if !found {
		fmt.Println("not such key")
	} else {
		fmt.Println(value)
	}

	// looping a map
	for k, v := range m1 {
		fmt.Println(k, v)
	}

	// only key
	for k := range m1 {
		fmt.Println(k)
	}

	// only value
	for _, value := range m1 {
		fmt.Println(value)
	}

	// delete key
	delete(m1, "a")
	fmt.Println(m1)

	// if we delete a key that is not exist?
	// it will do nothing and won't return error

	// make a map with
	// student number as key
	// score as value
	var scoreMap = make(map[string]int, 100)

	for i := 0; i < 100; i++ {
		student_number := fmt.Sprintf("student%02d", i)
		score := rand.Intn(100)
		scoreMap[student_number] = score
	}
	fmt.Println(scoreMap)

	// access key in order
	var keys = make([]string, 100)
	for key := range scoreMap {
		keys = append(keys, key)
	}

	// sort key slice
	sort.Strings(keys)
	for _, key := range keys {
		fmt.Println(key, scoreMap[key])
	}

	// map in slice
	var s1 = make([]map[int]int, 10)
	// each map in slice not init yet
	// init mao in slice

	s1[0] = make(map[int]int, 1)
	s1[0][10] = 1
	fmt.Println(s1)

	// slice in map
	var m2 = make(map[string][]int, 10)
	m2["a1"] = []int{10, 20, 30}
	fmt.Println(m2)
}
