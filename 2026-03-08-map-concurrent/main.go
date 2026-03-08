package main

import "fmt"

func main() {
	m := make(map[string]int)
	m["a"] = 1
	
	go func() {
		m["b"] = 2
	}()
	
	fmt.Println(m)
}
