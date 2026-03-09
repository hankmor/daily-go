package main

import (
	"fmt"
	"time"
)

func main() {
	m := make(map[int]int)
	
	go func() {
		for i := 0; i < 1000; i++ {
			m[i] = i
		}
	}()
	
	for i := 0; i < 1000; i++ {
		fmt.Println(m[i])
	}
	
	time.Sleep(time.Second)
}
