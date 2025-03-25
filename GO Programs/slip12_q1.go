package main
import "fmt"

func swap(x *int, y *int) {
  temp := *x
  *x = *y
  *y = temp
}

func main() {
  var p, q int
  fmt.Print("Enter the two numbers to swap:")
  fmt.Scan(&p, &q)
  fmt.Println("Before swap:", p, q)
  swap(&p, &q)
  fmt.Println("After swap:", p, q)
}