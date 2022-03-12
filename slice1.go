package main

import (
	"fmt"
)

func main() {
	myslice1 := []int{}
	fmt.Println(cap(myslice1))
	fmt.Println(len(myslice1))
	fmt.Println(myslice1)

	myslice2 := []string{"Go", "slices", "are", "powerful"}
	fmt.Println(cap(myslice2))
	fmt.Println(len(myslice2))
	fmt.Println(myslice2)

}
