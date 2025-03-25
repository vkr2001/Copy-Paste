package main
import "fmt"

func calc(a, b int) (int, int, int, float32) {
    sum := a + b
    diff := a - b
    prod := a * b
    div := float32(a) / float32(b)
    return sum, diff, prod, div
}

func main() {
    var n1, n2 int
    fmt.Print("Enter the two numbers to add, subtract, multiply and divide:")
    fmt.Scanln(&n1, &n2)
    sum, diff, prod, div := calc(n1, n2)
    fmt.Printf("Sum: %d\n", sum)
    fmt.Printf("Difference: %d\n", diff)
    fmt.Printf("Product: %d\n", prod)
    fmt.Printf("Division: %.2f\n", div)
}