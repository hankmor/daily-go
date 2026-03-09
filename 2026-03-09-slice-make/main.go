package main

import "fmt"

func main() {
	s1 := make([]int, 0, 5)
	s2 := []int{}

	fmt.Println(s1, s2)
	fmt.Println(len(s1), cap(s1))
	fmt.Println(len(s2), cap(s2))
}
