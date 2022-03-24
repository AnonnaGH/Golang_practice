package main
import ("fmt")

func main() {
  fruits := [3]string{"apple", "orange", "banana"}

  for idx, _ := range fruits {
     fmt.Printf("%v\n", idx)
  }
}
