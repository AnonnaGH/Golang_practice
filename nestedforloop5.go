package main

import "fmt"

func main() {
	a := [2]string{"big", "testy"}
	b := [3]string{"apple", "orange", "banana"}

	for i := 0; i < len(a); i++ {
		for j := 0; j < len(b); j++ {
			fmt.Println(a[i], b[j])
		}
	}

}
