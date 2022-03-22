package main

import (
	"fmt"
)

func main() {

	s := []int{1, 2, 3, 4, 5, 6}
	fmt.Printf("s = %v\n", s)
	fmt.Printf("length = %d\n", len(s))
	fmt.Printf("capacity = %d\n", cap(s))

	s = append(s, 20, 21, 22, 23, 24, 25)
	fmt.Printf("s = %v\n", s)
	fmt.Printf("length = %d\n", len(s))
	fmt.Printf("capacity =%d\n", cap(s))

	s = append(s, 30, 36)
	fmt.Printf("s = %v\n", s)
	fmt.Printf("length = %d\n", len(s))
	fmt.Printf("capacity =%d\n", cap(s))

}
