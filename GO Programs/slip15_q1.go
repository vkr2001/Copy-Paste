package main
import "fmt"

func addSub(a, b int) (int, int, int) {
    sum := a + b
    diff := a - b
    prod := a * b
    return sum, diff, prod
}

func main() {
    var n1, n2 int
    fmt.Print("Enter the two numbers to add, subtract and multiply:")
    fmt.Scanln(&n1, &n2)
    sum, diff, prod := addSub(n1, n2)
    fmt.Printf("Sum: %d\n", sum)
    fmt.Printf("Difference: %d\n", diff)
    fmt.Printf("Product: %d\n", prod)
}
