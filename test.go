package main

import "fmt"

func main() {
	fmt.Println("vim-go")
	var h = [...]string{"test1", "test2"}
	s1 := 0
	s2 := 1

	fmt.Println(h[s1])
	fmt.Println(h[s2])
}
