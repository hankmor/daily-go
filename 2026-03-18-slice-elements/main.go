package main

import "fmt"

func modify(s []int) {
	s[0] = 100
}

func main() {
	arr := []int{1, 2, 3}
	modify(arr)
	fmt.Println(arr)
}
