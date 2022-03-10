package main
import ("fmt")
 
func main() {
  var x uint = 500
  var y uint = 4500
  fmt.Printf("Type: %T, value: %v\n", x, x)
  fmt.Printf("Type: %T, value: %v", y, y)
}
